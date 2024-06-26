package unit

import (
	"testing"
)

func TestDelimitInt(t *testing.T) {
	var tests = []struct {
		val       int64
		delimiter NumberDelimiter
		exp       string
	}{
		{123456, NumberDelimiterNone, "123456"},
		{123456, NumberDelimiterComma, "123,456"},
		{123456, NumberDelimiterSpace, "123 456"},
		{1000000000, NumberDelimiterDot, "1.000.000.000"},
		{-123456789, NumberDelimiterUnderscore, "-123_456_789"},
		{-123456789, NumberDelimiterThinSpace, "-123 456 789"},
		{-9223372036854775808, NumberDelimiterComma, "-9,223,372,036,854,775,808"},
		{-9223372036854775807, NumberDelimiterComma, "-9,223,372,036,854,775,807"},
		{-9223372036854775808, NumberDelimiterUnderscore, "-9_223_372_036_854_775_808"},
	}

	for _, tt := range tests {
		got := DelimitInt(tt.val, tt.delimiter)
		if got != tt.exp {
			t.Errorf("DelimitInt val (%v) not expected, got: (%s), expected: (%s)\n", tt.val, got, tt.exp)
		}
	}
}

func TestTrimDelimiter(t *testing.T) {
	var tests = []struct {
		val string
		exp string
	}{
		{"123456", "123456"},
		{"123,456", "123456"},
		{"123 456", "123456"},
		{"1.000.000.000", "1.000.000.000"}, // dot delimiter not trimmed
		{"-123_456_789", "-123456789"},
		{"-123 456 789", "-123456789"},
		{"-9,223,372,036,854,775,808", "-9223372036854775808"},
		{"-9,223,372,036,854,775,807", "-9223372036854775807"},
		{"-9_223_372_036_854_775_808", "-9223372036854775808"},
	}

	for _, tt := range tests {
		got := TrimDelimiter(tt.val)
		if got != tt.exp {
			t.Errorf("TrimDelimiter val (%v) not expected, got: (%s), expected: (%s)\n", tt.val, got, tt.exp)
		}
	}
}

func TestCommaInt(t *testing.T) {
	var tests = []struct {
		val int64
		exp string
	}{
		{123456, "123,456"},
		{1000000000, "1,000,000,000"},
		{-123456789, "-123,456,789"},
		{-9223372036854775808, "-9,223,372,036,854,775,808"},
		{-9223372036854775807, "-9,223,372,036,854,775,807"},
	}

	for _, tt := range tests {
		got := CommaInt(tt.val)
		if got != tt.exp {
			t.Errorf("CommaInt val (%v) not expected, got: (%s), expected: (%s)\n", tt.val, got, tt.exp)
		}
	}
}

func TestCommaFloat64(t *testing.T) {
	var tests = []struct {
		val       float64
		precision int
		exp       string
	}{
		{123456, 0, "123,456"},
		{123456, 3, "123,456.000"},
		{123456.1234, 3, "123,456.123"},
		{1000000000, 0, "1,000,000,000"},
		{-123456789, 0, "-123,456,789"},
		{-9223372036854775808, 0, "-9,223,372,036,854,775,808"},
		{-9223372036854775000, 0, "-9,223,372,036,854,774,784"},
		{-9223372036854775806, 3, "-9,223,372,036,854,775,808.000"}, // !!! Why
		{-3.4e+38, 0, "-339,999,999,999,999,996,123,846,586,046,231,871,488"},
		{3.4e+38, 0, "339,999,999,999,999,996,123,846,586,046,231,871,488"},
	}

	for _, tt := range tests {
		got := CommaFloat(tt.val, tt.precision)
		if got != tt.exp {
			t.Errorf("CommaFloat val (%v) with precision (%v) not expected, got: (%s), expected: (%s)\n", tt.val, tt.precision, got, tt.exp)
		}
	}

}
