package u

import (
	"math"
	"testing"
)

func TestQuantityAddSub(t *testing.T) {
	sum := Length(5, Kilometer).Add(Length(500, Meter))
	if sum.Value != 5.5 || sum.Unit != Unit(Kilometer) {
		t.Fatalf("Add = %+v, want 5.5 km", sum)
	}

	diff := Length(5, Kilometer).Sub(Length(500, Meter))
	if diff.Value != 4.5 || diff.Unit != Unit(Kilometer) {
		t.Fatalf("Sub = %+v, want 4.5 km", diff)
	}

	unchanged := Quantity(Length(5, Meter)).Add(Quantity(Mass(10, Kilogram)))
	if unchanged.Value != 5 || unchanged.Unit != Unit(Meter) {
		t.Fatalf("incompatible Add = %+v, want 5 m", unchanged)
	}
}

func TestQuantityAddAffineTemperature(t *testing.T) {
	sum := Temperature(30, Celsius).Add(Temperature(10, Celsius))
	if sum.Value != 40 || sum.Unit != Unit(Celsius) {
		t.Fatalf("Add = %+v, want 40 °C", sum)
	}

	crossSum := Temperature(30, Celsius).Add(Temperature(10, Kelvin))
	// 10 K = −263.15 °C as an absolute reading; 30 + (−263.15) = −233.15 °C
	if math.Abs(crossSum.Value-(-233.15)) > 1e-9 || crossSum.Unit != Unit(Celsius) {
		t.Fatalf("cross affine Add = %+v, want −233.15 °C", crossSum)
	}

	diff := Temperature(100, Celsius).Sub(Temperature(273.15, Kelvin))
	if diff.Value != 100 || diff.Unit != Unit(Celsius) {
		t.Fatalf("cross affine Sub = %+v, want 100 °C", diff)
	}
}

func TestQuantityMulDiv(t *testing.T) {
	area := Length(2, Kilometer).Mul(Length(3, Meter))
	if math.Abs(area.Value-0.006) > 1e-12 {
		t.Fatalf("Mul value = %v, want 0.006", area.Value)
	}
	if area.Unit.Symbol(WithExpSign(ExpSignSup)) != "km²" {
		t.Fatalf("Mul unit = %q, want km²", area.Unit.Symbol(WithExpSign(ExpSignSup)))
	}

	speed := Length(10, Kilometer).Div(Time(2, Hour))
	if speed.Value != 5 {
		t.Fatalf("Div value = %v, want 5", speed.Value)
	}
	if speed.Unit.Symbol(WithExpSign(ExpSignSup)) != "km·h⁻¹" {
		t.Fatalf("Div unit = %q, want km·h⁻¹", speed.Unit.Symbol(WithExpSign(ExpSignSup)))
	}

	ratio := Mass(10, Kilogram).Div(Mass(2, Kilogram))
	if ratio.Value != 5 {
		t.Fatalf("ratio value = %v, want 5", ratio.Value)
	}
	if ratio.Unit != NoneUnit {
		t.Fatalf("ratio unit = %v, want NoneUnit", ratio.Unit)
	}
	if ratio.String() != "5" {
		t.Fatalf("ratio String() = %q, want 5", ratio.String())
	}

	momentum := Mass(2, Kilogram).Mul(Length(3, Meter))
	if momentum.Value != 6 {
		t.Fatalf("cross Mul value = %v, want 6", momentum.Value)
	}
	if momentum.Unit.Symbol(WithExpSign(ExpSignSup)) != "kg·m" {
		t.Fatalf("cross Mul unit = %q, want kg·m", momentum.Unit.Symbol(WithExpSign(ExpSignSup)))
	}
}

func TestQuantityMulVDivV(t *testing.T) {
	doubled := Length(3, Meter).MulV(2)
	if doubled.Value != 6 || doubled.Unit != Unit(Meter) {
		t.Fatalf("MulV = %+v, want 6 m", doubled)
	}

	half := Length(3, Meter).DivV(2)
	if half.Value != 1.5 || half.Unit != Unit(Meter) {
		t.Fatalf("DivV = %+v, want 1.5 m", half)
	}

	unchanged := Length(3, Meter).DivV(0)
	if unchanged.Value != 3 || unchanged.Unit != Unit(Meter) {
		t.Fatalf("DivV(0) = %+v, want 3 m", unchanged)
	}
}

func TestDerivedQuantityAddMul(t *testing.T) {
	speedUnit := NewDerivedUnit().Length(LengthUnit(Kilometer), 1).Time(TimeUnit(Hour), -1)
	a := NewDerivedQuantity(60, speedUnit)
	b := NewDerivedQuantity(30, speedUnit)

	sum := a.Add(b)
	if sum.Value != 90 || sum.Unit != speedUnit {
		t.Fatalf("Add = %+v, want 90 km·h^-1", sum)
	}

	forceUnit := NewDerivedUnit().Mass(MassUnit(Kilogram), 1).
		Length(LengthUnit(Meter), 1).
		Time(TimeUnit(Second), -2)
	force := NewDerivedQuantity(10, forceUnit)
	meterUnit, err := NewDerivedUnit().Length(LengthUnit(Meter), 1).Intern()
	if err != nil {
		t.Fatal(err)
	}
	distance := NewDerivedQuantity(5, meterUnit)

	work := force.Mul(distance)
	if work.Value != 50 {
		t.Fatalf("Mul value = %v, want 50", work.Value)
	}
	if work.Unit.Symbol(WithExpSign(ExpSignSup)) != "kg·m²·s⁻²" {
		t.Fatalf("Mul unit = %q, want kg·m²·s⁻²", work.Unit.Symbol(WithExpSign(ExpSignSup)))
	}
}

func TestDerivedQuantityMulDivBaseQuantity(t *testing.T) {
	forceUnit := NewDerivedUnit().Mass(MassUnit(Kilogram), 1).
		Length(LengthUnit(Meter), 1).
		Time(TimeUnit(Second), -2)
	force := NewDerivedQuantity(10, forceUnit)

	work := force.Mul(Length(5, Meter))
	if work.Value != 50 {
		t.Fatalf("Derived Mul base value = %v, want 50", work.Value)
	}
	if work.Unit.Symbol(WithExpSign(ExpSignSup)) != "kg·m²·s⁻²" {
		t.Fatalf("Derived Mul base unit = %q, want kg·m²·s⁻²", work.Unit.Symbol(WithExpSign(ExpSignSup)))
	}

	spring := force.Div(Length(2, Meter))
	if spring.Value != 5 {
		t.Fatalf("Derived Div base value = %v, want 5", spring.Value)
	}
	if spring.Unit.Symbol(WithExpSign(ExpSignSup)) != "kg·s⁻²" {
		t.Fatalf("Derived Div base unit = %q, want kg·s⁻²", spring.Unit.Symbol(WithExpSign(ExpSignSup)))
	}

	scaled := Length(3, Meter).Mul(force)
	if scaled.Value != 30 {
		t.Fatalf("base Mul Derived value = %v, want 30", scaled.Value)
	}
	if scaled.Unit.Symbol(WithExpSign(ExpSignSup)) != "kg·m²·s⁻²" {
		t.Fatalf("base Mul Derived unit = %q, want kg·m²·s⁻²", scaled.Unit.Symbol(WithExpSign(ExpSignSup)))
	}
}
