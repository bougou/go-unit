package u

import (
	"math"
	"testing"
)

func TestAreaDerivedDimensions(t *testing.T) {
	cases := []struct {
		name string
		unit *DerivedUnit
	}{
		{"SquareMeter", SquareMeter},
		{"SquareKilometer", SquareKilometer},
		{"SquareCentimeter", SquareCentimeter},
		{"SquareMillimeter", SquareMillimeter},
		{"Are", Are},
		{"Hectare", Hectare},
		{"Acre", Acre},
		{"SquareInch", SquareInch},
		{"SquareFoot", SquareFoot},
		{"SquareYard", SquareYard},
		{"SquareMile", SquareMile},
		{"Mu", Mu},
		{"Qing", Qing},
		{"Barn", Barn},
	}
	for _, tc := range cases {
		if !tc.unit.Dim().Equal(DimArea) {
			t.Fatalf("%s dim = %+v, want DimArea", tc.name, tc.unit.Dim())
		}
	}
}

func TestAreaNamedSymbols(t *testing.T) {
	cases := []struct {
		unit *DerivedUnit
		want string
	}{
		{SquareMeter, "m²"},
		{SquareKilometer, "km²"},
		{SquareCentimeter, "cm²"},
		{SquareMillimeter, "mm²"},
		{Are, "a"},
		{Hectare, "ha"},
		{Acre, "ac"},
		{SquareInch, "in²"},
		{SquareFoot, "ft²"},
		{SquareYard, "yd²"},
		{SquareMile, "mi²"},
		{Mu, "mu"},
		{Qing, "qing"},
		{Barn, "b"},
	}
	for _, tc := range cases {
		if got := tc.unit.Symbol(); got != tc.want {
			t.Fatalf("Symbol() = %q, want %q", got, tc.want)
		}
	}
}

func TestAreaFactorToBase(t *testing.T) {
	cases := []struct {
		unit *DerivedUnit
		want float64
	}{
		{SquareMeter, 1},
		{SquareKilometer, 1e6},
		{SquareCentimeter, 1e-4},
		{SquareMillimeter, 1e-6},
		{Are, 100},
		{Hectare, 1e4},
		{Acre, 4046.8564224},
		{SquareFoot, 0.3048 * 0.3048},
		{SquareInch, 0.0254 * 0.0254},
		{SquareYard, 0.9144 * 0.9144},
		{SquareMile, 1609.344 * 1609.344},
		{Mu, 2000.0 / 3.0},
		{Qing, 200000.0 / 3.0},
		{Barn, 1e-28},
	}
	for _, tc := range cases {
		if math.Abs(tc.unit.FactorToBase()-tc.want) > 1e-12*math.Max(1, math.Abs(tc.want)) {
			t.Fatalf("%s FactorToBase = %g, want %g", tc.unit.Symbol(), tc.unit.FactorToBase(), tc.want)
		}
	}
}

func TestAreaConversions(t *testing.T) {
	ha := Hectare.Of(1).By(SquareMeter)
	if ha.Value != 1e4 || ha.Unit != SquareMeter {
		t.Fatalf("1 ha → m² = %v %q, want 10000 m²", ha.Value, ha.Unit.Symbol())
	}

	km2 := SquareKilometer.Of(2).By(SquareMeter)
	if km2.Value != 2e6 || km2.Unit != SquareMeter {
		t.Fatalf("2 km² → m² = %v %q, want 2e6 m²", km2.Value, km2.Unit.Symbol())
	}

	cm2 := SquareMeter.Of(1).By(SquareCentimeter)
	if cm2.Value != 1e4 || cm2.Unit != SquareCentimeter {
		t.Fatalf("1 m² → cm² = %v %q, want 10000 cm²", cm2.Value, cm2.Unit.Symbol())
	}

	mu := Mu.Of(1).By(SquareMeter)
	if math.Abs(mu.Value-2000.0/3.0) > 1e-12 || mu.Unit != SquareMeter {
		t.Fatalf("1 mu → m² = %v %q, want %g m²", mu.Value, mu.Unit.Symbol(), 2000.0/3.0)
	}

	qing := Qing.Of(1).By(Mu)
	if math.Abs(qing.Value-100) > 1e-12 || qing.Unit != Mu {
		t.Fatalf("1 qing → mu = %v %q, want 100 mu", qing.Value, qing.Unit.Symbol())
	}

	// Prefix on m² is not km² (would be 10³, not 10⁶).
	if SquareMeter.Prefix(Kilo).FactorToBase() == SquareKilometer.FactorToBase() {
		t.Fatal("SquareMeter.Prefix(Kilo) must not equal SquareKilometer")
	}
}

func TestAreaParse(t *testing.T) {
	cases := []struct {
		input string
		unit  *DerivedUnit
		value float64
	}{
		{"2 ha", Hectare, 2},
		{"1.5 km²", SquareKilometer, 1.5},
		{"100 mu", Mu, 100},
		{"3 ac", Acre, 3},
	}
	for _, tc := range cases {
		q, err := DerivedQuantityParse(tc.input)
		if err != nil {
			t.Fatalf("parse %q: %v", tc.input, err)
		}
		if q.Unit != tc.unit || q.Value != tc.value {
			t.Fatalf("parse %q = %v %q, want %v %q", tc.input, q.Value, q.Unit.Symbol(), tc.value, tc.unit.Symbol())
		}
	}
}
