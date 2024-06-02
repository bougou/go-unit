package unit

import (
	"testing"
)

func Test_PrefixParseNotValid(t *testing.T) {
	var tests = []struct {
		str   string
		mode  PrefixMode
		valid bool
	}{
		{"1024 MiB", IEC, false},
		{"1024 Bytes", SI, false},
		{"1024 Gb/s", SI, false},
		{"1024 Kbit", SI, false},
		{"1024 bps", SI, false},

		{"1024 Mi", IEC, true},
		{"1024 Mi", SI, false},

		{"1024 M", SI, true},
		{"1024 M", IEC, true}, // This is also supported.
	}
	for _, tt := range tests {
		_, err := PrefixParse(tt.str, tt.mode)
		if err != nil {
			if tt.valid {
				t.Errorf("parse str (%s) with mode (%s) should not failed, err: %s", tt.str, tt.mode, err)
			}
		} else {
			if !tt.valid {
				t.Errorf("parse str (%s) with mode (%s) should failed", tt.str, tt.mode)
			}
		}

	}
}

func Test_PrefixParse(t *testing.T) {
	var tests = []struct {
		str    string
		mode   PrefixMode
		expect float64
	}{
		{"1024Mi", Auto, 1024 * 1024 * 1024},
		{"1,024 Mi", Auto, 1024 * 1024 * 1024},
		{"1,024 M", Auto, 1024 * 1000 * 1000},
		{"1,024 M", IEC, 1024 * 1024 * 1024},
		{"1_024_000Mi", Auto, 1024 * 1024 * 1024 * 1000},
		{"1 024 000 Mi", Auto, 1024 * 1024 * 1024 * 1000},
		{"1 024 000 Mi", Auto, 1024 * 1024 * 1024 * 1000},
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
		{"0.1", SI, 0.1},
		{"100 μ", Auto, 100 * Micro},
		{"1024Yi", Auto, 1024 * Yobi},
	}

	for _, tt := range tests {
		got, err := PrefixParse(tt.str, tt.mode)
		if err != nil {
			t.Errorf("parse str (%s) with mode (%s) failed, err: %s\n", tt.str, tt.mode, err)
			continue
		}

		if got != tt.expect {
			t.Errorf("parse str (%s) with mode (%s) not expect, got: %v, expect: %v\n", tt.str, tt.mode, got, tt.expect)
			continue
		}
	}
}
