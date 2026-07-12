package u

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"unicode"
)

var (
	baseUnitBySymbol    map[string]Unit
	baseSymbolsByLength []string
	baseSymbolIndexOnce sync.Once
)

func ensureBaseUnitSymbolIndex() {
	baseSymbolIndexOnce.Do(func() {
		baseUnitBySymbol = make(map[string]Unit)
		for unit, def := range unitRegistry {
			if def.Symbol != "" {
				baseUnitBySymbol[def.Symbol] = unit
			}
		}
		baseSymbolsByLength = make([]string, 0, len(baseUnitBySymbol))
		for sym := range baseUnitBySymbol {
			baseSymbolsByLength = append(baseSymbolsByLength, sym)
		}
		sort.Slice(baseSymbolsByLength, func(i, j int) bool {
			li, lj := len(baseSymbolsByLength[i]), len(baseSymbolsByLength[j])
			if li != lj {
				return li > lj
			}
			return baseSymbolsByLength[i] < baseSymbolsByLength[j]
		})
	})
}

func matchBaseUnitSymbol(s string) (Unit, int, bool) {
	ensureBaseUnitSymbolIndex()
	for _, sym := range baseSymbolsByLength {
		if strings.HasPrefix(s, sym) {
			return baseUnitBySymbol[sym], len(sym), true
		}
	}
	return "", 0, false
}

// DerivedUnitParse parses a derived unit symbol composed of registered base-unit
// symbols (e.g. "km/h", "kg·m·s^-2", "m·s⁻¹"). Adjacent units must be separated
// by ·, *, /, or whitespace (MulSignSpace); longest-match keeps "ms" as millisecond.
// SI special names such as "N" are not expanded.
//
// The returned unit is interned when possible.
func DerivedUnitParse(symbol string) (*DerivedUnit, error) {
	symbol = strings.TrimSpace(symbol)
	if symbol == "" {
		return nil, ErrSyntax
	}

	terms, err := parseDerivedSymbolTerms(symbol)
	if err != nil {
		return nil, err
	}
	if len(terms) == 0 {
		return nil, fmt.Errorf("empty derived unit %q: %w", symbol, ErrSyntax)
	}

	u, err := buildDerivedUnitFromTerms(terms)
	if err != nil {
		return nil, err
	}
	interned, err := u.Intern()
	if err != nil {
		return nil, err
	}
	return interned, nil
}

func parseDerivedSymbolTerms(symbol string) ([]unitTerm, error) {
	if num, den, ok := splitDerivedSlash(symbol); ok {
		numTerms, err := tokenizeDerivedProduct(num, 1)
		if err != nil {
			return nil, err
		}
		denTerms, err := tokenizeDerivedProduct(den, -1)
		if err != nil {
			return nil, err
		}
		return mergeUnitTerms(numTerms, denTerms)
	}
	return tokenizeDerivedProduct(symbol, 1)
}

func splitDerivedSlash(s string) (num, den string, ok bool) {
	depth := 0
	for i, r := range s {
		switch r {
		case '(':
			depth++
		case ')':
			if depth > 0 {
				depth--
			}
		case '/':
			if depth == 0 {
				left := strings.TrimSpace(s[:i])
				right := strings.TrimSpace(s[i+1:])
				if right == "" {
					return "", "", false
				}
				if left == "" || left == "1" {
					return "", unwrapDerivedParens(right), true
				}
				return left, unwrapDerivedParens(right), true
			}
		}
	}
	return "", "", false
}

func unwrapDerivedParens(s string) string {
	s = strings.TrimSpace(s)
	for {
		if !strings.HasPrefix(s, "(") || !strings.HasSuffix(s, ")") {
			return s
		}
		inner := strings.TrimSpace(s[1 : len(s)-1])
		if strings.ContainsRune(inner, '(') || strings.ContainsRune(inner, ')') {
			return s
		}
		s = inner
	}
}

func tokenizeDerivedProduct(s string, sideSign int8) ([]unitTerm, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}

	var terms []unitTerm
	first := true
	for len(s) > 0 {
		if !first {
			consumed := skipDerivedMulSign(s)
			if consumed == 0 {
				return nil, fmt.Errorf("missing unit separator before %q: %w", s, ErrSyntax)
			}
			s = s[consumed:]
		}
		first = false

		unit, symLen, ok := matchBaseUnitSymbol(s)
		if !ok {
			return nil, fmt.Errorf("unknown unit factor in %q: %w", s, ErrSyntax)
		}
		exp, expLen, err := parseDerivedExponent(s[symLen:])
		if err != nil {
			return nil, err
		}
		if exp == 0 {
			exp = 1
		}
		if sideSign < 0 {
			exp = -absInt8(exp)
		}
		terms = append(terms, unitTerm{unit: unit, exp: exp})
		s = s[symLen+expLen:]
	}
	return terms, nil
}

func skipDerivedMulSign(s string) int {
	switch {
	case strings.HasPrefix(s, string(MulSignDot)):
		return len(string(MulSignDot))
	case strings.HasPrefix(s, string(MulSignStar)):
		return len(string(MulSignStar))
	default:
		i := 0
		for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
			i++
		}
		return i
	}
}

func parseDerivedExponent(s string) (exp int8, consumed int, err error) {
	if s == "" {
		return 0, 0, nil
	}
	if s[0] == '^' {
		exp, rest, err := parseSignedIntegerPrefix(s[1:])
		if err != nil {
			return 0, 0, err
		}
		return exp, 1 + len(s[1:]) - len(rest), nil
	}
	if exp, width, err := utf8SuperscriptRunes(s); width > 0 || err != nil {
		return exp, width, err
	}
	return 0, 0, nil
}

func parseSignedIntegerPrefix(s string) (val int8, rest string, err error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, s, fmt.Errorf("missing exponent after ^: %w", ErrSyntax)
	}

	negative := false
	switch s[0] {
	case '-':
		negative = true
		s = s[1:]
	case '+':
		s = s[1:]
	}
	if s == "" {
		return 0, "", fmt.Errorf("missing exponent after sign: %w", ErrSyntax)
	}

	i := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		i++
	}
	if i == 0 {
		return 0, s, fmt.Errorf("invalid exponent %q: %w", s, ErrSyntax)
	}

	val, err = parseSmallInt8(s[:i])
	if err != nil {
		return 0, s, err
	}
	if negative {
		val = -val
	}
	return val, s[i:], nil
}

func utf8SuperscriptRunes(s string) (exp int8, width int, err error) {
	if s == "" {
		return 0, 0, nil
	}

	runes := []rune(s)
	if runes[0] == '⁻' {
		if len(runes) == 1 {
			return 0, 0, fmt.Errorf("incomplete superscript exponent: %w", ErrSyntax)
		}
		digit, ok := superscriptDigit(runes[1])
		if !ok {
			return 0, 0, fmt.Errorf("invalid superscript exponent: %w", ErrSyntax)
		}
		if len(runes) == 2 {
			return -digit, len(string(runes[:2])), nil
		}
		val, w, err := readSuperscriptDigits(runes[1:])
		if err != nil {
			return 0, 0, err
		}
		return -val, len(string(runes[:1+w])), nil
	}

	val, w, err := readSuperscriptDigits(runes)
	if err != nil {
		return 0, 0, err
	}
	if w == 0 {
		return 0, 0, nil
	}
	return val, len(string(runes[:w])), nil
}

func readSuperscriptDigits(runes []rune) (exp int8, count int, err error) {
	if len(runes) == 0 {
		return 0, 0, nil
	}
	digits := make([]rune, 0, len(runes))
	for _, r := range runes {
		d, ok := superscriptDigit(r)
		if !ok {
			break
		}
		digits = append(digits, rune('0'+d))
	}
	if len(digits) == 0 {
		return 0, 0, nil
	}
	val, err := parseSmallInt8(string(digits))
	if err != nil {
		return 0, 0, err
	}
	return val, len(digits), nil
}

func superscriptDigit(r rune) (int8, bool) {
	switch r {
	case '⁰':
		return 0, true
	case '¹':
		return 1, true
	case '²':
		return 2, true
	case '³':
		return 3, true
	case '⁴':
		return 4, true
	case '⁵':
		return 5, true
	case '⁶':
		return 6, true
	case '⁷':
		return 7, true
	case '⁸':
		return 8, true
	case '⁹':
		return 9, true
	default:
		return 0, false
	}
}

func parseSmallInt8(s string) (int8, error) {
	var n int64
	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			return 0, fmt.Errorf("invalid exponent %q: %w", s, ErrSyntax)
		}
		n = n*10 + int64(ch-'0')
		if n > 127 {
			return 0, fmt.Errorf("exponent out of range %q: %w", s, ErrSyntax)
		}
	}
	return int8(n), nil
}

func absInt8(v int8) int8 {
	if v < 0 {
		return -v
	}
	return v
}

func mergeUnitTerms(parts ...[]unitTerm) ([]unitTerm, error) {
	byDim := map[Dimension]unitTerm{}
	for _, terms := range parts {
		for _, t := range terms {
			if t.exp == 0 {
				continue
			}
			def, ok := t.unit.Def()
			if !ok {
				return nil, fmt.Errorf("unknown unit %q: %w", t.unit, ErrSyntax)
			}
			existing, exists := byDim[def.Dimension]
			if exists {
				if existing.unit != t.unit {
					return nil, fmt.Errorf("conflicting units for dimension %s: %w", def.Dimension, ErrSyntax)
				}
				sum := int(existing.exp) + int(t.exp)
				if sum < -127 || sum > 127 {
					return nil, fmt.Errorf("exponent out of range for dimension %s: %w", def.Dimension, ErrSyntax)
				}
				t.exp = int8(sum)
			}
			byDim[def.Dimension] = t
		}
	}
	out := make([]unitTerm, 0, len(byDim))
	for _, t := range byDim {
		out = append(out, t)
	}
	return out, nil
}

func buildDerivedUnitFromTerms(terms []unitTerm) (*DerivedUnit, error) {
	u := NewDerivedUnit()
	for _, t := range terms {
		if t.exp == 0 {
			continue
		}
		def, ok := t.unit.Def()
		if !ok {
			return nil, fmt.Errorf("unknown unit %q: %w", t.unit, ErrSyntax)
		}
		if !setDerivedUnitExp(u, def.Dimension, t.unit, t.exp) {
			return nil, fmt.Errorf("duplicate dimension %s in derived unit: %w", def.Dimension, ErrSyntax)
		}
	}
	return u, nil
}

func lookupUnitSymbolOrDerived(symbol string) (Unit, bool) {
	if unit, ok := lookupUnitSymbol(symbol); ok {
		return unit, true
	}
	du, err := DerivedUnitParse(symbol)
	if err != nil {
		return "", false
	}
	key := du.Key()
	registerUnitSymbol(symbol, key)
	return key, true
}
