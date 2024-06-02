package unit

import "testing"

func Test_UnitIsBase(t *testing.T) {
	tests := []struct {
		i      int
		u      Unit
		isBase bool
	}{
		{0, TimeSecond, true},
		{1, MassGram, true},
		{2, MassKilogram, false},
		{3, LengthMeter, true},
		{4, TimeHour, false},
	}

	for _, tt := range tests {
		if reason, isValid := tt.u.Valid(); !isValid {
			t.Errorf("test [%d] unit (%s), not valid unit, %s", tt.i, tt.u, reason)
		}
		_, actual := tt.u.IsBase()
		if tt.isBase != actual {
			t.Errorf("test [%d] unit (%s), isBase not matched, expect (%v), actual (%v)", tt.i, tt.u, tt.isBase, actual)
		}
	}
}
