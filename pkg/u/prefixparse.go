package u

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	// ErrSyntax indicates the input string could not be parsed.
	ErrSyntax = fmt.Errorf("syntax error")
	// ErrInvalidMode indicates an unsupported PrefixMode.
	ErrInvalidMode = fmt.Errorf("invalid prefix mode")
	// ErrInvalidSymbol indicates an unknown prefix symbol.
	ErrInvalidSymbol = fmt.Errorf("invalid prefix symbol")
)

// PrefixParse converts the string s to a float64 value.
// Valid inputs:
//   - plain number: "1024", "0.5"
//   - number with SI prefix: "1024 G"
//   - number with IEC prefix: "1024Gi"
//
// The following are not valid prefix inputs (strip the unit suffix first):
//   - "1024 MiB"     → use "1024 Mi"
//   - "1024 Bytes"   → use "1024"
//   - "1024 Bytes/s" → use "1024"
//   - "1024 Gb/s"    → use "1024 G"
//   - "1024 Kbit"    → use "1024 K"
//   - "1024 bps"     → use "1024"
//
// "MiB", "Bytes", "Gb/s", "Kbit", and "bps" are physical units, not numeric prefixes.
// Keep only the prefix portion before calling PrefixParse.
func PrefixParse(s string, mode PrefixMode) (val float64, err error) {
	s = strings.TrimSpace(s)

	if mode == Auto {
		if len(s) > 0 && s[len(s)-1] == 'i' {
			mode = IEC
		} else {
			mode = SI
		}
	}

	var number, prefix string
	if i := strings.IndexAny(s, AllValidSymbols); i >= 0 {
		number = s[:i]
		prefix = s[i:]
	} else {
		number = s[:]
		prefix = ""
	}
	number = strings.TrimSpace(number)
	number = TrimDelimiter(number)

	var symbol rune

	if prefix == "" {
		symbol = rune(fakeSymbol)
	} else {
		for i, c := range prefix {
			switch i {
			case 0:
				symbol = c

			case 1:
				if !(mode == IEC || mode == ForceSI || mode == ForceIEC) {
					return val, fmt.Errorf("'i' suffix is invalid for mode %s: %w", mode, ErrSyntax)
				}
				if c != 'i' {
					return val, ErrSyntax
				}

			default:
				return val, ErrSyntax
			}
		}
	}

	val, err = strconv.ParseFloat(number, 64)
	if err != nil {
		return val, fmt.Errorf("parse number failed %s: %w", number, ErrSyntax)
	}

	scale, _, err := getScaleOfSymbol(symbol, mode)
	if err != nil {
		return val, fmt.Errorf("get scale for symbol in %q: %w", s, ErrSyntax)
	}

	return val * scale, nil

}
