package u

import (
	"math"
	"testing"

	"github.com/bougou/go-unit/pkg/prefix"
)

func TestBaseUnitP(t *testing.T) {
	km := Meter.Prefix(prefix.Kilo)
	if Unit(km).Symbol() != "km" {
		t.Fatalf("Meter.Prefix(prefix.Kilo) symbol = %q, want km", Unit(km).Symbol())
	}
	def, ok := Unit(km).Def()
	if !ok || def.Scale != float64(prefix.Kilo) {
		t.Fatalf("Meter.Prefix(prefix.Kilo) scale = %v ok=%v, want %g", def.Scale, ok, float64(prefix.Kilo))
	}
	if Meter.Prefix(prefix.Kilo) != km {
		t.Fatal("Meter.Prefix(prefix.Kilo) should be stable")
	}

	if Unit(Second.Prefix(prefix.Milli)).Symbol() != "ms" {
		t.Fatalf("Second.Prefix(prefix.Milli) symbol = %q, want ms", Unit(Second.Prefix(prefix.Milli)).Symbol())
	}
	if Unit(Ampere.Prefix(prefix.Micro)).Symbol() != "μA" {
		t.Fatalf("Ampere.Prefix(prefix.Micro) symbol = %q, want μA", Unit(Ampere.Prefix(prefix.Micro)).Symbol())
	}

	if Kilogram.Prefix(prefix.Milli) != Gram {
		t.Fatalf("Kilogram.Prefix(prefix.Milli) = %q, want %q", Kilogram.Prefix(prefix.Milli), Gram)
	}
	if Gram.Prefix(prefix.Kilo) != Kilogram {
		t.Fatalf("Gram.Prefix(prefix.Kilo) = %q, want %q", Gram.Prefix(prefix.Kilo), Kilogram)
	}
	if Unit(Gram.Prefix(prefix.Milli)).Symbol() != "mg" {
		t.Fatalf("Gram.Prefix(prefix.Milli) symbol = %q, want mg", Unit(Gram.Prefix(prefix.Milli)).Symbol())
	}
	if Unit(Gram.Prefix(prefix.Mega)).Symbol() != "Mg" {
		t.Fatalf("Gram.Prefix(prefix.Mega) symbol = %q, want Mg", Unit(Gram.Prefix(prefix.Mega)).Symbol())
	}

	q := Length(1000, Meter).By(Meter.Prefix(prefix.Kilo))
	if q.Value != 1 || q.Unit != Unit(Meter.Prefix(prefix.Kilo)) {
		t.Fatalf("1000 m.By(km) = %v %q, want 1 km", q.Value, q.Unit)
	}
}

func TestPrefix(t *testing.T) {
	r := Ohm.Of(2e6).Prefix(prefix.Mega)
	if r.Value != 2 || r.Unit.Symbol() != "MΩ" {
		t.Fatalf("Prefix = %v %q, want 2 MΩ", r.Value, r.Unit.Symbol())
	}
	d := Length(1000, Meter).Prefix(prefix.Kilo)
	if d.Value != 1 || d.Unit != Unit(Meter.Prefix(prefix.Kilo)) {
		t.Fatalf("Length.Prefix = %v %q, want 1 km", d.Value, d.Unit)
	}
	m := Mass(1, Kilogram).Prefix(prefix.Milli)
	if m.Value != 1000 || m.Unit != Unit(Gram) {
		t.Fatalf("Mass.Prefix = %v %q, want 1000 g", m.Value, m.Unit)
	}

	// Sqrt(P/R) yields unnamed A; Prefix should still convert to mA.
	i := Watt.Prefix(prefix.Milli).Of(20).Div(Ohm.Of(5000).Prefix(prefix.Kilo)).Sqrt().Prefix(prefix.Milli)
	if math.Abs(i.Value-2) > 1e-12 || i.Unit.Symbol() != "mA" {
		t.Fatalf("√(P/R).Prefix(prefix.Milli) = %v %q, want 2 mA", i.Value, i.Unit.Symbol())
	}
}

func TestKilogramPOnlyMilli(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("Kilogram.Prefix(prefix.Kilo) should panic")
		}
	}()
	_ = Kilogram.Prefix(prefix.Kilo)
}

func TestBaseUnitPAffineRejected(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("Celsius.P should panic")
		}
	}()
	_ = Celsius.Prefix(prefix.Kilo)
}

func TestParsePrefixedBaseUnit(t *testing.T) {
	q, err := LengthQuantityParse("10 km")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if q.Unit != Unit(Meter.Prefix(prefix.Kilo)) || q.Value != 10 {
		t.Fatalf("got %v %q", q.Value, q.Unit)
	}
	q2, err := MassQuantityParse("5 mg")
	if err != nil {
		t.Fatalf("parse mg: %v", err)
	}
	if q2.Unit != Unit(Gram.Prefix(prefix.Milli)) {
		t.Fatalf("mg unit = %q, want Gram.Prefix(prefix.Milli)", q2.Unit)
	}
}
