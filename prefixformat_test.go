package unit

import (
	"testing"
)

func Test_PrefixFormat(t *testing.T) {
	var tests = []struct {
		val    float64
		mode   PrefixMode
		expect string
		option prefixFormatOption
	}{
		{0.1, SI, "100m", prefixFormatOption{}},
		{0.568, SI, "568m", prefixFormatOption{}},
		{568, SI, "568", prefixFormatOption{}},
		{0.000568, SI, "568u", prefixFormatOption{}},
		{0.000000568, SI, "568n", prefixFormatOption{}},
		{1048576, IEC, "1Mi", prefixFormatOption{}},
		{2048, IEC, "2Ki", prefixFormatOption{}},
		{2048, Auto, "2K", prefixFormatOption{}},
		{2048, SI, "2K", prefixFormatOption{}},
		{2048, SI1024, "2K", prefixFormatOption{}},
		{2048, IEC, "2.000Ki", prefixFormatOption{false, 3, RoundMethodFloor, 0}},
		{2048, Auto, "2.048K", prefixFormatOption{false, 3, RoundMethodFloor, 0}},
		{2048, SI, "2.048K", prefixFormatOption{false, 3, RoundMethodFloor, 0}},
		{2048, SI1024, "2.000K", prefixFormatOption{false, 3, RoundMethodFloor, 0}},
		{1025, IEC, "1Ki", prefixFormatOption{}},
		{9.412, Auto, "9.41", prefixFormatOption{false, 2, RoundMethodFloor, 0}},
		{1024 * 1024 * 1024, IEC, "1Gi", prefixFormatOption{}},
		{1024 * 1024 * 1024, IEC, "1.00Gi", prefixFormatOption{false, 2, RoundMethodFloor, 0}},
		{1023 * 1024 * 1024, IEC, "1023.00Mi", prefixFormatOption{false, 2, RoundMethodFloor, 0}},
		{1025 * 1024 * 1024, IEC, "1.00098 Gi", prefixFormatOption{true, 5, RoundMethodFloor, 0}},
		{2047 * 1024 * 1024, IEC, "1Gi", prefixFormatOption{}},
		{2047 * 1024 * 1024, IEC, "2Gi", prefixFormatOption{false, 0, RoundMethodCeil, 0}},
		{1023 * 1024 * 1024, IEC, "1023Mi", prefixFormatOption{false, 0, RoundMethodCeil, 0}},
		{1025 * 1024 * 1024, IEC, "2 Gi", prefixFormatOption{true, 0, RoundMethodCeil, 0}},
		{1010 * 1000 * 1000, SI, "1.01G", prefixFormatOption{false, 2, RoundMethodFloor, 0}},
		{1001 * 1000 * 1000, SI, "1G", prefixFormatOption{}},                              // 向下取整
		{1001 * 1000 * 1000, SI, "2G", prefixFormatOption{false, 0, RoundMethodCeil, 0}},  // 向上取整
		{1999 * 1000 * 1000, SI, "1G", prefixFormatOption{false, 0, RoundMethodFloor, 0}}, // 向下取整
		{1401 * 1000 * 1000, SI, "1G", prefixFormatOption{false, 0, RoundMethodRound, 0}}, // 四舍五入
		{1501 * 1000 * 1000, SI, "2G", prefixFormatOption{false, 0, RoundMethodRound, 0}}, // 四舍五入
		{1000 * 1000 * 1000, SI, "1.00G", prefixFormatOption{false, 2, RoundMethodFloor, 0}},
		{999 * 1000 * 1000, SI, "999.000 M", prefixFormatOption{true, 3, RoundMethodFloor, 0}},
		{999.99 * 1000 * 1000, SI, "999.990 M", prefixFormatOption{true, 3, RoundMethodFloor, 0}},
		{1024*1024*1024 + 1024*1024*1024*0.049, IEC, "1Gi", prefixFormatOption{false, 0, RoundMethodDifference, 0.05}},
		{1024*1024*1024 + 1024*1024*1024*0.05, IEC, "2Gi", prefixFormatOption{false, 0, RoundMethodDifference, 0.05}},
		{1024*1024*1024 + 1024*1024*1024*0.051, IEC, "2Gi", prefixFormatOption{false, 0, RoundMethodDifference, 0.05}},
		{1024*1024*1024 + 1024*1024*1024*0.029, IEC, "1Gi", prefixFormatOption{false, 0, RoundMethodDifference, 0.03}},
		{1024*1024*1024 + 1024*1024*1024*0.030, IEC, "2Gi", prefixFormatOption{false, 0, RoundMethodDifference, 0.03}},
		{1024*1024*1024 + 1024*1024*1024*0.031, IEC, "2Gi", prefixFormatOption{false, 0, RoundMethodDifference, 0.03}},
		{-(1024*1024*1024 + 1024*1024*1024*0.029), IEC, "-1Gi", prefixFormatOption{false, 0, RoundMethodDifference, 0.03}},
		{-(1024*1024*1024 + 1024*1024*1024*0.030), IEC, "-2Gi", prefixFormatOption{false, 0, RoundMethodDifference, 0.03}},
		{-(1024*1024*1024 + 1024*1024*1024*0.031), IEC, "-2Gi", prefixFormatOption{false, 0, RoundMethodDifference, 0.03}},
	}

	for _, tt := range tests {

		got := PrefixFormat(tt.val, tt.mode,
			WithSpace(tt.option.space),
			WithPrecision(tt.option.precision),
			WithRoundMethod(tt.option.roundMethod),
			WithRoundDifference(tt.option.roundDifference),
		)
		if got != tt.expect {
			t.Errorf("format val (%v) with mode (%s) failed, not expect, got: %v, expect: %v", tt.val, tt.mode, got, tt.expect)
			continue
		}
	}
}
