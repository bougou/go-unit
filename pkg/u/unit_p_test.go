package u

import (
	"math"
	"testing"
)

func TestBaseUnitP(t *testing.T) {
	km := Meter.Prefix(Kilo)
	if Unit(km).Symbol() != "km" {
		t.Fatalf("Meter.Prefix(Kilo) symbol = %q, want km", Unit(km).Symbol())
	}
	def, ok := Unit(km).Def()
	if !ok || def.Scale != float64(Kilo) {
		t.Fatalf("Meter.Prefix(Kilo) scale = %v ok=%v, want %g", def.Scale, ok, float64(Kilo))
	}
	if Meter.Prefix(Kilo) != km {
		t.Fatal("Meter.Prefix(Kilo) should be stable")
	}

	if Unit(Second.Prefix(Milli)).Symbol() != "ms" {
		t.Fatalf("Second.Prefix(Milli) symbol = %q, want ms", Unit(Second.Prefix(Milli)).Symbol())
	}
	if Unit(Ampere.Prefix(Micro)).Symbol() != "μA" {
		t.Fatalf("Ampere.Prefix(Micro) symbol = %q, want μA", Unit(Ampere.Prefix(Micro)).Symbol())
	}

	if Kilogram.Prefix(Milli) != Gram {
		t.Fatalf("Kilogram.Prefix(Milli) = %q, want %q", Kilogram.Prefix(Milli), Gram)
	}
	if Gram.Prefix(Kilo) != Kilogram {
		t.Fatalf("Gram.Prefix(Kilo) = %q, want %q", Gram.Prefix(Kilo), Kilogram)
	}
	if Unit(Gram.Prefix(Milli)).Symbol() != "mg" {
		t.Fatalf("Gram.Prefix(Milli) symbol = %q, want mg", Unit(Gram.Prefix(Milli)).Symbol())
	}
	if Unit(Gram.Prefix(Mega)).Symbol() != "Mg" {
		t.Fatalf("Gram.Prefix(Mega) symbol = %q, want Mg", Unit(Gram.Prefix(Mega)).Symbol())
	}

	q := Length(1000, Meter).By(Meter.Prefix(Kilo))
	if q.Value != 1 || q.Unit != Unit(Meter.Prefix(Kilo)) {
		t.Fatalf("1000 m.By(km) = %v %q, want 1 km", q.Value, q.Unit)
	}
}

func TestPrefix(t *testing.T) {
	r := Ohm.Of(2e6).Prefix(Mega)
	if r.Value != 2 || r.Unit.Symbol() != "MΩ" {
		t.Fatalf("Prefix = %v %q, want 2 MΩ", r.Value, r.Unit.Symbol())
	}
	d := Length(1000, Meter).Prefix(Kilo)
	if d.Value != 1 || d.Unit != Unit(Meter.Prefix(Kilo)) {
		t.Fatalf("Length.Prefix = %v %q, want 1 km", d.Value, d.Unit)
	}
	m := Mass(1, Kilogram).Prefix(Milli)
	if m.Value != 1000 || m.Unit != Unit(Gram) {
		t.Fatalf("Mass.Prefix = %v %q, want 1000 g", m.Value, m.Unit)
	}

	// Sqrt(P/R) yields unnamed A; Prefix should still convert to mA.
	i := Watt.Prefix(Milli).Of(20).Div(Ohm.Of(5000).Prefix(Kilo)).Sqrt().Prefix(Milli)
	if math.Abs(i.Value-2) > 1e-12 || i.Unit.Symbol() != "mA" {
		t.Fatalf("√(P/R).Prefix(Milli) = %v %q, want 2 mA", i.Value, i.Unit.Symbol())
	}
}

func TestKilogramPOnlyMilli(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("Kilogram.Prefix(Kilo) should panic")
		}
	}()
	_ = Kilogram.Prefix(Kilo)
}

func TestBaseUnitPAffineRejected(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("Celsius.P should panic")
		}
	}()
	_ = Celsius.Prefix(Kilo)
}

func TestParsePrefixedBaseUnit(t *testing.T) {
	q, err := LengthQuantityParse("10 km")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if q.Unit != Unit(Meter.Prefix(Kilo)) || q.Value != 10 {
		t.Fatalf("got %v %q", q.Value, q.Unit)
	}
	q2, err := MassQuantityParse("5 mg")
	if err != nil {
		t.Fatalf("parse mg: %v", err)
	}
	if q2.Unit != Unit(Gram.Prefix(Milli)) {
		t.Fatalf("mg unit = %q, want Gram.Prefix(Milli)", q2.Unit)
	}
}
