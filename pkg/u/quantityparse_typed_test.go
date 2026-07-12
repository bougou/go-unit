package u

import (
	"errors"
	"math"
	"testing"
)

func TestLengthQuantityParse(t *testing.T) {
	got, err := LengthQuantityParse("10 km")
	if err != nil {
		t.Fatal(err)
	}
	if got.Value != 10 || got.Unit != Unit(Kilometer) {
		t.Fatalf("got %+v, want 10 km", got)
	}

	if _, err := LengthQuantityParse("5 s"); !errors.Is(err, ErrDimension) {
		t.Fatalf("LengthQuantityParse(5 s) err = %v, want ErrDimension", err)
	}
	if _, err := LengthQuantityParse("5 N"); !errors.Is(err, ErrDimension) {
		t.Fatalf("LengthQuantityParse(5 N) err = %v, want ErrDimension", err)
	}
}

func TestTimeQuantityParse(t *testing.T) {
	got, err := TimeQuantityParse("2 h")
	if err != nil {
		t.Fatal(err)
	}
	if got.Value != 2 || got.Unit != Unit(Hour) {
		t.Fatalf("got %+v, want 2 h", got)
	}
}

func TestMassQuantityParse(t *testing.T) {
	got, err := MassQuantityParse("100 kg")
	if err != nil {
		t.Fatal(err)
	}
	if got.Value != 100 || got.Unit != Unit(Kilogram) {
		t.Fatalf("got %+v, want 100 kg", got)
	}
}

func TestTemperatureQuantityParse(t *testing.T) {
	got, err := TemperatureQuantityParse("-5 °C")
	if err != nil {
		t.Fatal(err)
	}
	if got.Value != -5 || got.Unit != Unit(Celsius) {
		t.Fatalf("got %+v, want -5 °C", got)
	}
}

func TestDerivedQuantityParse(t *testing.T) {
	tests := []struct {
		input string
		value float64
		unit  *DerivedUnit
	}{
		{"5 N", 5, ForceUnit},
		{"60 m/s", 60, SpeedUnit},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := DerivedQuantityParse(tt.input)
			if err != nil {
				t.Fatalf("DerivedQuantityParse(%q) error: %v", tt.input, err)
			}
			if got.Unit != tt.unit {
				t.Fatalf("unit = %v, want %v", got.Unit, tt.unit)
			}
			if math.Abs(got.Value-tt.value) > 1e-12 {
				t.Fatalf("value = %v, want %v", got.Value, tt.value)
			}
		})
	}

	if _, err := DerivedQuantityParse("10 m"); !errors.Is(err, ErrDimension) {
		t.Fatalf("DerivedQuantityParse(10 m) err = %v, want ErrDimension", err)
	}
}

func TestTypedQuantityMustParse(t *testing.T) {
	if got := LengthQuantityMustParse("10 m"); got.Value != 10 || got.Unit != Unit(Meter) {
		t.Fatalf("LengthQuantityMustParse = %+v, want 10 m", got)
	}
	if got := DerivedQuantityMustParse("5 N"); got.Value != 5 || got.Unit != ForceUnit {
		t.Fatalf("DerivedQuantityMustParse = %+v, want 5 N", got)
	}
}
