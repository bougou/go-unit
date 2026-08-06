package prefix

import (
	"testing"
)

func Test_PrefixParseInvalid(t *testing.T) {
	tests := []struct {
		str  string
		mode PrefixMode
	}{
		{"1024 MiB", IEC},
		{"1024 Mx", IEC},
		{"1024 Mx", ForceIEC},
		{"1024 Mx", ForceSI},
		{"1024 Bytes", SI},
		{"1024 Mi", SI},
	}

	for _, tt := range tests {
		_, err := PrefixParse(tt.str, tt.mode)
		if err == nil {
			t.Errorf("parse str (%s) with mode (%s) should fail", tt.str, tt.mode)
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
		{"1_024_000Mi", Auto, 1024 * 1024 * 1024 * 1000},
		{"1 024 000 Mi", Auto, 1024 * 1024 * 1024 * 1000},
		{"1 024 000 Mi", Auto, 1024 * 1024 * 1024 * 1000},
		{"1 024 000 Mi", IEC, 1024 * 1024 * 1024 * 1000},
		{"1 024 000 Mi", ForceSI, 1024 * 1000 * 1000 * 1000},
		{"1 024 000 M", ForceSI, 1024 * 1000 * 1000 * 1000},
		{"1 024 000 Mi", ForceIEC, 1024 * 1024 * 1024 * 1000},
		{"1 024 000 M", ForceIEC, 1024 * 1024 * 1024 * 1000},
		{"1024Ki", Auto, 1048576},
		{"1024 Ki", Auto, 1024 * 1024},
		{"1024   Ki", Auto, 1024 * 1024},
		{"1.2K", Auto, 1200},
		{"1.2k", Auto, 1200},
		{"2.0Ti", Auto, 2 * 1024 * 1024 * 1024 * 1024},
		{"9M", Auto, 9 * 1000 * 1000},
		{"9", Auto, 9},
		{"9M", SI1024, 9 * 1024 * 1024},
		{"9.412 μ", Auto, 9.412 * float64(Micro)},
		{"100m", SI, 0.1},
		{"0.1", SI, 0.1},
		{"100 μ", Auto, 100 * float64(Micro)},
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
