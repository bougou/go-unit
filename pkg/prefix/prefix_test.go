package prefix

import (
	"fmt"
	"testing"
)

func Test_exponentOfValue(t *testing.T) {
	var tests = []struct {
		val float64
		exp int
	}{
		{1023, 0},
		{1024, 1},
		{1025, 1},
		{1048575, 1},
		{Mebi, 2},
		{1048577, 2},
		{Gibi - 1, 2},
		{Gibi, 3},
		{Gibi + 1, 3},
	}

	fmt.Println(formatScales(scalesIEC))

	for _, tt := range tests {
		got := exponentOfValue(tt.val, scalesIEC)
		if got != tt.exp {
			t.Errorf("exponentOfValue(%f, %v) not expect, got: %v, expect: %v", tt.val, scalesIEC, got, tt.exp)
			continue
		}
	}
}
