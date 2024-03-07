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

// PrefixParse converts the string s to a float64 value.
// A valid input string s could be:
//   - plain number, like: "1024"
//   - number and SI unit prefix, like: "1024 G"
//   - number and IEC unit prefix, like: "1024Gi"
//
// Note, the strings of the following formats are NOT valid:
//   - "1024 Bytes"
//   - "1024 Gb/s"
//   - "1024 Kbit"
//   - "1024 Bytes/s"
//
// The 'Bytes, 'Bytes/s, 'b/s', 'bit' are NOT "unix prefix", they are real "unit".
// You should cut them out from the string before 'PrefixParse'.
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
