package u

import (
	"strconv"
	"strings"
)

// formatQuantityValue formats the numeric part of a quantity.
func formatQuantityValue(value float64, opt formatOption) string {
	if opt.precision < 0 {
		s := strconv.FormatFloat(value, 'g', -1, 64)
		if opt.delimiter == NumberDelimiterNone {
			return s
		}
		return delimitNumberString(s, opt.delimiter)
	}
	return DelimitFloat(value, opt.precision, opt.delimiter)
}

// joinQuantityString joins a formatted value and unit symbol with a single space.
func joinQuantityString(value, unitSymbol string) string {
	if unitSymbol == "" {
		return value
	}
	return value + " " + unitSymbol
}

// delimitNumberString inserts a thousands separator into an already-formatted
// number string (supports leading sign, decimal point, and e/E exponent).
func delimitNumberString(s string, delimiter NumberDelimiter) string {
	if delimiter == NumberDelimiterNone || s == "" {
		return s
	}

	sign := ""
	rest := s
	if rest[0] == '+' || rest[0] == '-' {
		sign = rest[:1]
		rest = rest[1:]
	}

	exp := ""
	if i := strings.IndexAny(rest, "eE"); i >= 0 {
		exp = rest[i:]
		rest = rest[:i]
	}

	intPart, fracPart := rest, ""
	if i := strings.IndexByte(rest, '.'); i >= 0 {
		intPart = rest[:i]
		fracPart = rest[i:]
	}

	return sign + delimitDigitString(intPart, delimiter) + fracPart + exp
}

func delimitDigitString(digits string, delimiter NumberDelimiter) string {
	if delimiter == NumberDelimiterNone || len(digits) <= 3 {
		return digits
	}
	n := len(digits)
	numDelim := (n - 1) / 3
	out := make([]rune, n+numDelim)
	for i, j, k := n-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = rune(digits[i])
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = rune(delimiter)
		}
	}
}
