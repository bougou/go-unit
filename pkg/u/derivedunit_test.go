package u

import (
	"math"
	"testing"
)

func TestDerivedUnitEmpty(t *testing.T) {
	u := NewDerivedUnit()

	if u.Symbol(WithExpSign(ExpSignSup)) != "" {
		t.Fatalf("Symbol() = %q, want empty", u.Symbol(WithExpSign(ExpSignSup)))
	}
	if u.Symbol() != "" {
		t.Fatalf("Symbol() = %q, want empty", u.Symbol())
	}
	if !u.Dim().Equal(DerivedDimension{}) {
		t.Fatalf("Dim() = %+v, want zero", u.Dim())
	}
	if u.Key() != "" {
		t.Fatalf("Key() = %q, want empty", u.Key())
	}
	if u.FactorToBase() != 1 {
		t.Fatalf("FactorToBase() = %v, want 1", u.FactorToBase())
	}
}

func TestDerivedUnitSymbol(t *testing.T) {
	speedUnit := NewDerivedUnit().Length(LengthUnit(Kilometer), 1).Time(TimeUnit(Hour), -1)
	forceUnit := NewDerivedUnit().Mass(MassUnit(Kilogram), 1).
		Length(LengthUnit(Meter), 1).
		Time(TimeUnit(Second), -2)

	tests := []struct {
		name    string
		unit    *DerivedUnit
		options []SymbolOption
		want    string
	}{
		{
			name: "default carat negative",
			unit: speedUnit,
			want: "km·h^-1",
		},
		{
			name: "slash division",
			unit: speedUnit,
			options: []SymbolOption{
				WithDivSign(DivSignSlash),
			},
			want: "km/h",
		},
		{
			name: "superscript",
			unit: speedUnit,
			options: []SymbolOption{
				WithExpSign(ExpSignSup),
			},
			want: "km·h⁻¹",
		},
		{
			name: "star multiply",
			unit: speedUnit,
			options: []SymbolOption{
				WithMulSign(MulSignStar),
			},
			want: "km*h^-1",
		},
		{
			name: "force slash",
			unit: forceUnit,
			options: []SymbolOption{
				WithDivSign(DivSignSlash),
			},
			want: "kg·m/s^2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.unit.Symbol(tt.options...)
			if got != tt.want {
				t.Fatalf("Symbol() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDerivedUnitAreaPerTime(t *testing.T) {
	u := NewDerivedUnit().Length(LengthUnit(Centimeter), 2).Time(TimeUnit(Hour), -1)

	sign := u.Symbol(WithExpSign(ExpSignSup))
	if sign != "cm²·h⁻¹" {
		t.Fatalf("sign = %q, want cm²·h⁻¹", sign)
	}
	dim := u.Dim()
	wantDim := DerivedDimension{L: 2, T: -1}
	if !dim.Equal(wantDim) {
		t.Fatalf("dim = %+v, want %+v", dim, wantDim)
	}
}

func TestNewDerivedQuantity(t *testing.T) {
	unit := NewDerivedUnit().Length(LengthUnit(Centimeter), 2).Time(TimeUnit(Hour), -1)
	speed := NewDerivedQuantity(25, unit)

	if speed.Value != 25 {
		t.Fatalf("value = %v, want 25", speed.Value)
	}
	if speed.String() != "25 cm²·h⁻¹" {
		t.Fatalf("String() = %q, want 25 cm²·h⁻¹", speed.String())
	}
}

func TestDerivedQuantityForceUnit(t *testing.T) {
	u := NewDerivedUnit().Mass(MassUnit(Kilogram), 1).
		Length(LengthUnit(Meter), 1).
		Time(TimeUnit(Second), -2)

	if u.Symbol(WithExpSign(ExpSignSup)) != "kg·m·s⁻²" {
		t.Fatalf("sign = %q, want kg·m·s⁻²", u.Symbol(WithExpSign(ExpSignSup)))
	}
}

func TestDerivedUnitSpecialSymbol(t *testing.T) {
	force := ForceUnit

	if force.SpecialSymbol() != "N" {
		t.Fatalf("SpecialSymbol() = %q, want N", force.SpecialSymbol())
	}
	if force.Symbol() != "kg·m·s^-2" {
		t.Fatalf("Symbol() = %q, want kg·m·s^-2", force.Symbol())
	}
	if force.Symbol(WithExpSign(ExpSignSup)) != "kg·m·s⁻²" {
		t.Fatalf("compound Symbol() = %q, want kg·m·s⁻²", force.Symbol(WithExpSign(ExpSignSup)))
	}
	if force.Symbol(WithNamedSymbol(true)) != "N" {
		t.Fatalf("WithNamedSymbol(true) = %q, want N", force.Symbol(WithNamedSymbol(true)))
	}

	q := NewDerivedQuantity(10, force)
	if q.String() != "10 kg·m·s⁻²" {
		t.Fatalf("String() = %q, want 10 kg·m·s⁻²", q.String())
	}
	if Unit(force.Key()).Symbol() != "kg·m·s^-2" {
		t.Fatalf("Unit.Symbol() = %q, want kg·m·s^-2", Unit(force.Key()).Symbol())
	}
	if Unit(force.Key()).Symbol(WithNamedSymbol(true)) != "N" {
		t.Fatalf("Unit.Symbol(WithNamedSymbol(true)) = %q, want N", Unit(force.Key()).Symbol(WithNamedSymbol(true)))
	}
}

func TestDerivedQuantitySI(t *testing.T) {
	speedUnit := NewDerivedUnit().Length(LengthUnit(Kilometer), 1).Time(TimeUnit(Hour), -1)
	speed := NewDerivedQuantity(50, speedUnit)

	si := speed.SI()
	want := 50.0 * 1000.0 / 3600.0
	if math.Abs(si.Value-want) > 1e-12 {
		t.Fatalf("SI value = %v, want %v", si.Value, want)
	}
	if si.Unit.Symbol(WithExpSign(ExpSignSup)) != "m·s⁻¹" {
		t.Fatalf("SI sign = %q, want m·s⁻¹", si.Unit.Symbol(WithExpSign(ExpSignSup)))
	}
}

func TestDerivedUnitIntern(t *testing.T) {
	u := NewDerivedUnit().Mass(MassUnit(Kilogram), 1).
		Length(LengthUnit(Meter), 1).
		Time(TimeUnit(Second), -2)

	canonical, err := u.Intern()
	if err != nil {
		t.Fatal(err)
	}
	again, err := u.Intern()
	if err != nil {
		t.Fatal(err)
	}
	if canonical != again {
		t.Fatal("Intern should return the same canonical instance")
	}
	if canonical.Symbol(WithExpSign(ExpSignSup)) != "kg·m·s⁻²" {
		t.Fatalf("sign = %q, want kg·m·s⁻²", canonical.Symbol(WithExpSign(ExpSignSup)))
	}
}
