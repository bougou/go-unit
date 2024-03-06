package unit

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrSyntax        = fmt.Errorf("error syntax")
	ErrInvalidMode   = fmt.Errorf("invalid unit mode")
	ErrInvalidSymbol = fmt.Errorf("invalid unit symbol")
)

// Parse converts the string s to a float64 value.
func Parse(s string, mode Mode) (val float64, err error) {
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
	number = strings.ReplaceAll(number, ",", "")

	var symbol rune

	if prefix == "" {
		symbol = rune(fakeSymbol)
	} else {
		for i, c := range prefix {
			switch i {
			case 0:
				symbol = c

			case 1:
				if mode != IEC {
					return val, ErrSyntax
				}

			default:
				return val, ErrSyntax
			}
		}
	}

	val, err = strconv.ParseFloat(number, 64)
	if err != nil {
		return val, ErrSyntax
	}

	scale, _, err := getScaleOfSymbol(symbol, mode)
	if err != nil {
		return val, ErrSyntax
	}

	return val * scale, nil

}
