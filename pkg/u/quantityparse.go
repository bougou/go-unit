package u

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var (
	unitBySymbol    map[string]Unit
	symbolsByLength []string
	symbolIndexOnce sync.Once
)

func registerUnitSymbol(symbol string, unit Unit) {
	if symbol == "" {
		return
	}
	if unitBySymbol == nil {
		unitBySymbol = make(map[string]Unit)
	}
	if existing, ok := unitBySymbol[symbol]; ok && existing != unit {
		return
	}
	unitBySymbol[symbol] = unit
}

func ensureSymbolIndex() {
	symbolIndexOnce.Do(func() {
		symbolsByLength = make([]string, 0, len(unitBySymbol))
		for sym := range unitBySymbol {
			symbolsByLength = append(symbolsByLength, sym)
		}
		sort.Slice(symbolsByLength, func(i, j int) bool {
			li, lj := len(symbolsByLength[i]), len(symbolsByLength[j])
			if li != lj {
				return li > lj
			}
			return symbolsByLength[i] < symbolsByLength[j]
		})
	})
}

func splitQuantityString(s string) (number, unitSymbol string, err error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", "", ErrSyntax
	}

	ensureSymbolIndex()
	for _, sym := range symbolsByLength {
		if !strings.HasSuffix(s, sym) {
			continue
		}
		prefix := strings.TrimSpace(s[:len(s)-len(sym)])
		if prefix == "" {
			continue
		}
		if strings.ContainsAny(prefix, " \t\n\r") {
			continue
		}
		if _, err := parseQuantityNumber(prefix); err != nil {
			continue
		}
		return prefix, sym, nil
	}

	if number, unitSymbol, ok := splitNumberAndUnitPrefix(s); ok {
		return number, unitSymbol, nil
	}
	return "", "", fmt.Errorf("unknown unit in %q: %w", s, ErrSyntax)
}

func splitNumberAndUnitPrefix(s string) (number, unitSymbol string, ok bool) {
	i := 0
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		i++
	}
	sawDigit := false
	for i < len(s) {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ',', '_':
			sawDigit = true
			i++
		case '.':
			sawDigit = true
			i++
		case 'e', 'E':
			i++
			if i < len(s) && (s[i] == '+' || s[i] == '-') {
				i++
			}
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				i++
			}
		default:
			goto done
		}
	}
done:
	if !sawDigit || i == 0 || i >= len(s) {
		return "", "", false
	}
	number = strings.TrimSpace(s[:i])
	unitSymbol = strings.TrimSpace(s[i:])
	if unitSymbol == "" {
		return "", "", false
	}
	if strings.ContainsAny(number, " \t\n\r") {
		return "", "", false
	}
	if _, err := parseQuantityNumber(number); err != nil {
		return "", "", false
	}
	if _, ok := lookupUnitSymbolOrDerived(unitSymbol); !ok {
		return "", "", false
	}
	return number, unitSymbol, true
}

func parseQuantityNumber(s string) (float64, error) {
	if strings.ContainsAny(s, " \t\n\r") {
		return 0, ErrSyntax
	}
	normalized := TrimDelimiter(s)
	val, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return 0, fmt.Errorf("parse number failed %s: %w", s, ErrSyntax)
	}
	return val, nil
}

func lookupUnitSymbol(symbol string) (Unit, bool) {
	ensureSymbolIndex()
	unit, ok := unitBySymbol[symbol]
	return unit, ok
}

// ErrDimension indicates the parsed unit belongs to a different dimension than required.
var ErrDimension = fmt.Errorf("dimension mismatch")

func parseBaseDimensionQuantity(s string, want Dimension) (Quantity, error) {
	q, err := QuantityParse(s)
	if err != nil {
		return Quantity{}, err
	}
	def, ok := q.Unit.Def()
	if !ok {
		return Quantity{}, fmt.Errorf("unit %q is not a base-dimension quantity: %w", q.Unit.Symbol(), ErrDimension)
	}
	if def.Dimension != want {
		return Quantity{}, fmt.Errorf("want %s unit, got %s: %w", want, def.Dimension, ErrDimension)
	}
	return q, nil
}

// QuantityParse parses text of the form "value unit".
// The numeric value must not contain spaces; optional spaces may appear between
// the value and the unit. The unit may be any registered base or derived unit
// symbol, including compound and SI special-name forms.
//
// Examples:
//
//	QuantityParse("10 m")        // 10 meters
//	QuantityParse("1.5km")       // 1.5 kilometers
//	QuantityParse("5 N")         // 5 newtons
//	QuantityParse("60 m/s")      // 60 m/s
func QuantityParse(s string) (Quantity, error) {
	number, unitSymbol, err := splitQuantityString(s)
	if err != nil {
		return Quantity{}, err
	}

	val, err := parseQuantityNumber(number)
	if err != nil {
		return Quantity{}, err
	}

	unit, ok := lookupUnitSymbolOrDerived(unitSymbol)
	if !ok {
		return Quantity{}, fmt.Errorf("unknown unit %q: %w", unitSymbol, ErrSyntax)
	}

	return Quantity{Value: val, Unit: unit}, nil
}
