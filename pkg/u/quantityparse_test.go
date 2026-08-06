package u

import (
	"math"
	"testing"

	"github.com/bougou/go-unit/pkg/prefix"
)

func TestQuantityParse(t *testing.T) {
	tests := []struct {
		input string
		value float64
		unit  Unit
	}{
		{"10 m", 10, Unit(Meter)},
		{"10m", 10, Unit(Meter)},
		{"  1.5  km  ", 1.5, Unit(Meter.Prefix(prefix.Kilo))},
		{"100 kg", 100, Unit(Kilogram)},
		{"-5 °C", -5, Unit(Celsius)},
		{"+3.2 s", 3.2, Unit(Second)},
		{"1,024 mm", 1024, Unit(Meter.Prefix(prefix.Milli))},
		{"1_024 mm", 1024, Unit(Meter.Prefix(prefix.Milli))},
		{"10 oz t", 10, Unit(TroyOunce)},
		{"5 N", 5, Unit(Newton.Key())},
		{"5 kg·m·s^-2", 5, Unit(Newton.Key())},
		{"5 kg·m·s⁻²", 5, Unit(Newton.Key())},
		{"60 m/s", 60, Unit(Speed.Key())},
		{"1.5e3 m", 1500, Unit(Meter)},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := QuantityParse(tt.input)
			if err != nil {
				t.Fatalf("QuantityParse(%q) error: %v", tt.input, err)
			}
			if got.Unit != tt.unit {
				t.Fatalf("unit = %q, want %q", got.Unit, tt.unit)
			}
			if math.Abs(got.Value-tt.value) > 1e-12 {
				t.Fatalf("value = %v, want %v", got.Value, tt.value)
			}
		})
	}
}

func TestQuantityParseInvalid(t *testing.T) {
	tests := []string{
		"",
		"   ",
		"abc m",
		"10",
		"10 unknown",
		"1 024 m",
		"1 024 000 m",
	}

	for _, input := range tests {
		if _, err := QuantityParse(input); err == nil {
			t.Fatalf("QuantityParse(%q) should fail", input)
		}
	}
}

func TestQuantityMustParse(t *testing.T) {
	got := QuantityMustParse("10 m")
	if got.Value != 10 || got.Unit != Unit(Meter) {
		t.Fatalf("QuantityMustParse = %+v, want 10 m", got)
	}
}
