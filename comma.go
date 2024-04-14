package unit

import (
	"bytes"
	"strconv"
	"strings"
)

type NumberDelimiter rune

const (
	NumberDelimiterNone       = NumberDelimiter(0)   // "1234"
	NumberDelimiterComma      = NumberDelimiter(',') // "1,234"
	NumberDelimiterUnderscore = NumberDelimiter('_') // "1_234"
	NumberDelimiterSpace      = NumberDelimiter(' ') // "1 234"
	NumberDelimiterThinSpace  = NumberDelimiter(' ') // "1 234", U+2009
	NumberDelimiterDot        = NumberDelimiter('.') // "1.234"
)

func DelimitInt(n int64, delimiter NumberDelimiter) string {
	if delimiter == NumberDelimiterNone {
		return strconv.FormatInt(n, 10)
	}

	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}

	numOfDelimiter := (numOfDigits - 1) / 3

	out := make([]rune, len(in)+numOfDelimiter)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = rune(in[i])
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = rune(delimiter)
		}
	}
}

func TrimDelimiter(s string) string {
	out := strings.ReplaceAll(s, string(NumberDelimiterComma), "")
	out = strings.ReplaceAll(out, string(NumberDelimiterUnderscore), "")
	out = strings.ReplaceAll(out, string(NumberDelimiterSpace), "")
	out = strings.ReplaceAll(out, string(NumberDelimiterThinSpace), "")

	// out = strings.ReplaceAll(out, string(NumberDelimiterDot), "")
	return out
}

// ref: https://stackoverflow.com/questions/13020308/how-to-fmt-printf-an-integer-with-thousands-comma
// ref: https://gosamples.dev/print-number-thousands-separator/
func CommaInt(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

func CommaFloat(val float64, precision int) string {
	buf := &bytes.Buffer{}
	if val < 0 {
		buf.Write([]byte{'-'})
		val = 0 - val
	}

	comma := []byte{','}

	parts := strings.Split(strconv.FormatFloat(val, 'f', precision, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}
	return buf.String()

}
