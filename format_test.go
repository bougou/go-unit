package unit

import (
	"testing"
)

func Test_Format(t *testing.T) {
	var tests = []struct {
		val    float64
		mode   Mode
		expect string
		option formatOption
	}{
		{0.1, SI, "100m", formatOption{}},
		{0.568, SI, "568m", formatOption{}},
		{568, SI, "568", formatOption{}},
		{0.000568, SI, "568u", formatOption{}},
		{0.000000568, SI, "568n", formatOption{}},
		{2048, IEC, "2Ki", formatOption{}},
		{1025, IEC, "1Ki", formatOption{}},
		{9.412, Auto, "9.41", formatOption{false, 2, RoundMethodFloor, false, 0}},
		{1024 * 1024 * 1024, IEC, "1Gi", formatOption{}},
		{1024 * 1024 * 1024, IEC, "1.00Gi", formatOption{false, 2, RoundMethodFloor, false, 0}},
		{1023 * 1024 * 1024, IEC, "1023.00Mi", formatOption{false, 2, RoundMethodFloor, false, 0}},
		{1025 * 1024 * 1024, IEC, "1.00098 Gi", formatOption{true, 5, RoundMethodFloor, false, 0}},
		{2047 * 1024 * 1024, IEC, "1Gi", formatOption{}},
		{2047 * 1024 * 1024, IEC, "2Gi", formatOption{false, 0, RoundMethodCeil, false, 0}},
		{1023 * 1024 * 1024, IEC, "1023Mi", formatOption{false, 0, RoundMethodCeil, false, 0}},
		{1025 * 1024 * 1024, IEC, "2 Gi", formatOption{true, 0, RoundMethodCeil, false, 0}},
		{1010 * 1000 * 1000, SI, "1.01G", formatOption{false, 2, RoundMethodFloor, false, 0}},
		{1001 * 1000 * 1000, SI, "1G", formatOption{}},                                     // 向下取整
		{1001 * 1000 * 1000, SI, "2G", formatOption{false, 0, RoundMethodCeil, false, 0}},  // 向上取整
		{1401 * 1000 * 1000, SI, "1G", formatOption{false, 0, RoundMethodRound, false, 0}}, // 四舍五入
		{1501 * 1000 * 1000, SI, "2G", formatOption{false, 0, RoundMethodRound, false, 0}}, // 四舍五入
		{1000 * 1000 * 1000, SI, "1.00G", formatOption{false, 2, RoundMethodFloor, false, 0}},
		{999 * 1000 * 1000, SI, "999.000 M", formatOption{true, 3, RoundMethodFloor, false, 0}},
		{999.99 * 1000 * 1000, SI, "999.990 M", formatOption{true, 3, RoundMethodFloor, false, 0}},
		{1024*1024*1024 + 1024*1024*1024*0.049, IEC, "1Gi", formatOption{false, 0, RoundMethodDifference, false, 0.05}},
		{1024*1024*1024 + 1024*1024*1024*0.05, IEC, "2Gi", formatOption{false, 0, RoundMethodDifference, false, 0.05}},
		{1024*1024*1024 + 1024*1024*1024*0.051, IEC, "2Gi", formatOption{false, 0, RoundMethodDifference, false, 0.05}},
		{1024*1024*1024 + 1024*1024*1024*0.029, IEC, "1Gi", formatOption{false, 0, RoundMethodDifference, false, 0.03}},
		{1024*1024*1024 + 1024*1024*1024*0.030, IEC, "2Gi", formatOption{false, 0, RoundMethodDifference, false, 0.03}},
		{1024*1024*1024 + 1024*1024*1024*0.031, IEC, "2Gi", formatOption{false, 0, RoundMethodDifference, false, 0.03}},
		{-(1024*1024*1024 + 1024*1024*1024*0.029), IEC, "-1Gi", formatOption{false, 0, RoundMethodDifference, false, 0.03}},
		{-(1024*1024*1024 + 1024*1024*1024*0.030), IEC, "-2Gi", formatOption{false, 0, RoundMethodDifference, false, 0.03}},
		{-(1024*1024*1024 + 1024*1024*1024*0.031), IEC, "-2Gi", formatOption{false, 0, RoundMethodDifference, false, 0.03}},
	}

	for _, tt := range tests {

		got := Format(tt.val, tt.mode,
			WithSpace(tt.option.space),
			WithPrecision(tt.option.precision),
			WithRoundMethod(tt.option.roundMethod),
			WithRoundDifference(tt.option.roundDifference),
		)
		if got != tt.expect {
			t.Errorf("not expect, got: %v, expect: %v", got, tt.expect)
			continue
		}
	}
}
