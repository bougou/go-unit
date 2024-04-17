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
//   - plain number, like: "1024", or "0.5"
//   - number and SI unit prefix, like: "1024 G"
//   - number and IEC unit prefix, like: "1024Gi"
//
// Note, the strings of the following formats are NOT valid:
//   - "1024 MiB"        -> Valid: "1024 Mi"
//   - "1024 Bytes"      -> Valid: "1024"
//   - "1024 Bytes/s"    -> Valid: "1024"
//   - "1024 Gb/s"       -> Valid: "1024 G"
//   - "1024 Kbit"       -> Valid: "1024 K"
//   - "1024 bps"        -> Valid: "1024"
//
// The "MiB", 'Bytes, 'Bytes/s, 'Gb/s', 'Kbit', "bps" are NOT "unix prefix", they are real "unit".
// You should just keep the "unit prefix" part of the "unit" before calling 'PrefixParse'.
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
