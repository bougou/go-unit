package unit

import (
	"testing"
)

func Test_Parse(t *testing.T) {
	var tests = []struct {
		str    string
		mode   Mode
		expect float64
	}{
		{"1024Mi", Auto, 1024 * 1024 * 1024},
		{"1,024 Mi", Auto, 1024 * 1024 * 1024},
		{"1024Ki", Auto, 1048576},
		{"1024 Ki", Auto, 1024 * 1024},
		{"1024   Ki", Auto, 1024 * 1024},
		{"1.2K", Auto, 1200},
		{"1.2k", Auto, 1200},
		{"2.0Ti", Auto, 2 * 1024 * 1024 * 1024 * 1024},
		{"9M", Auto, 9 * 1000 * 1000},
		{"9", Auto, 9},
		{"9M", SI1024, 9 * 1024 * 1024},
		{"9.412 μ", Auto, 9.412 * Micro},
		{"100m", SI, 0.1},
		{"100 μ", Auto, 100 * Micro},
	}

	for _, tt := range tests {
		got, err := Parse(tt.str, tt.mode)
		if err != nil {
			t.Error(err)
			continue
		}

		if got != tt.expect {
			t.Errorf("not expect, got: %v, expect: %v\n", got, tt.expect)
			continue
		}
	}
}
