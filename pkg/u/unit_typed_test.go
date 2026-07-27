package u

import "testing"

func TestLengthQuantityBaseAndBy(t *testing.T) {
	q := Length(1, LengthUnit(Meter.Prefix(Kilo)))

	base := q.Base()
	if base.Value != 1000 || base.Unit != Unit(Meter) {
		t.Fatalf("Base() = %+v, want 1000 m", base)
	}

	back := base.By(LengthUnit(Meter.Prefix(Kilo)))
	if back.Value != 1 || back.Unit != Unit(Meter.Prefix(Kilo)) {
		t.Fatalf("By(km) = %+v, want 1 kilometer", back)
	}
}

func TestX(t *testing.T) {
	Length(100, Meter.Prefix(Kilo))
}

func TestTypedUnitOf(t *testing.T) {
	tests := []struct {
		name string
		got  Quantity
		want Quantity
	}{
		{"Length", Quantity(Meter.Prefix(Kilo).Of(42)), Quantity(Length(42, Meter.Prefix(Kilo)))},
		{"Time", Quantity(Hour.Of(2)), Quantity(Time(2, Hour))},
		{"Mass", Quantity(Kilogram.Of(3)), Quantity(Mass(3, Kilogram))},
		{"Current", Quantity(Ampere.Of(10)), Quantity(Current(10, Ampere))},
		{"Temperature", Quantity(Celsius.Of(25)), Quantity(Temperature(25, Celsius))},
		{"Amount", Quantity(Mole.Of(1)), Quantity(Amount(1, Mole))},
		{"Luminous", Quantity(Candela.Of(100)), Quantity(Luminous(100, Candela))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.Value != tt.want.Value || tt.got.Unit != tt.want.Unit {
				t.Fatalf("Of() = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}
