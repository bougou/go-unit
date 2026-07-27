package u

import (
	"math"
	"testing"
)

func TestQuantityAddSub(t *testing.T) {
	sum := Length(5, Meter.Prefix(Kilo)).Add(Length(500, Meter))
	if sum.Value != 5.5 || sum.Unit != Unit(Meter.Prefix(Kilo)) {
		t.Fatalf("Add = %+v, want 5.5 km", sum)
	}

	diff := Length(5, Meter.Prefix(Kilo)).Sub(Length(500, Meter))
	if diff.Value != 4.5 || diff.Unit != Unit(Meter.Prefix(Kilo)) {
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
	area := Length(2, Meter.Prefix(Kilo)).Mul(Length(3, Meter))
	if math.Abs(area.Value-0.006) > 1e-12 {
		t.Fatalf("Mul value = %v, want 0.006", area.Value)
	}
	if area.Unit.Symbol(WithExpSign(ExpSignSup)) != "km²" {
		t.Fatalf("Mul unit = %q, want km²", area.Unit.Symbol(WithExpSign(ExpSignSup)))
	}

	speed := Length(10, Meter.Prefix(Kilo)).Div(Time(2, Hour))
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
	speedUnit := NewDerivedUnit().Length(LengthUnit(Meter.Prefix(Kilo)), 1).Time(TimeUnit(Hour), -1)
	a := NewDerivedQuantity(60, speedUnit)
	b := NewDerivedQuantity(30, speedUnit)

	sum := a.Add(b)
	if sum.Value != 90 || sum.Unit != speedUnit {
		t.Fatalf("Add = %+v, want 90 km·h^-1", sum)
	}

	// Proportional cross-unit add: 36 km/h + 10 m/s = 72 km/h
	ms := speedUnit.SI()
	cross := NewDerivedQuantity(36, speedUnit).Add(NewDerivedQuantity(10, ms))
	if math.Abs(cross.Value-72) > 1e-9 || cross.Unit != speedUnit {
		t.Fatalf("cross-unit Add = %+v, want 72 km/h", cross)
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
	if work.Unit.Symbol() != "J" {
		t.Fatalf("Mul unit = %q, want J", work.Unit.Symbol())
	}
}

func TestDerivedQuantityArithRejectsAffineOverlap(t *testing.T) {
	celsiusLen := NewDerivedUnit().
		Temperature(TemperatureUnit(Celsius), 1).
		Length(LengthUnit(Meter), 1).MustIntern()
	kelvinLen := NewDerivedUnit().
		Temperature(TemperatureUnit(Kelvin), 1).
		Length(LengthUnit(Meter), 1).MustIntern()
	cq := NewDerivedQuantity(20, celsiusLen)
	kq := NewDerivedQuantity(293.15, kelvinLen)

	if !cq.Compatible(kq) {
		t.Fatal("same derived dimension should be Compatible")
	}
	if sum := cq.Add(kq); sum.Value != cq.Value || sum.Unit != cq.Unit {
		t.Fatalf("Add with affine °C/K should leave q unchanged, got %+v", sum)
	}
	if diff := cq.Sub(kq); diff.Value != cq.Value || diff.Unit != cq.Unit {
		t.Fatalf("Sub with affine °C/K should leave q unchanged, got %+v", diff)
	}

	cOnly := NewDerivedUnit().Temperature(TemperatureUnit(Celsius), 1).MustIntern()
	kOnly := NewDerivedUnit().Temperature(TemperatureUnit(Kelvin), 1).MustIntern()
	left := NewDerivedQuantity(2, celsiusLen)
	right := NewDerivedQuantity(3, kOnly)
	if prod := left.Mul(right); prod.Value != left.Value || prod.Unit != left.Unit {
		t.Fatalf("Mul with overlapping affine temperature should leave q unchanged, got %+v", prod)
	}
	if quot := left.Div(NewDerivedQuantity(4, kOnly)); quot.Value != left.Value || quot.Unit != left.Unit {
		t.Fatalf("Div with overlapping affine temperature should leave q unchanged, got %+v", quot)
	}

	// No overlapping temperature: °C × kg is fine (no shared base dimension).
	massUnit := NewDerivedUnit().Mass(MassUnit(Kilogram), 1).MustIntern()
	scaled := NewDerivedQuantity(2, cOnly).Mul(NewDerivedQuantity(3, massUnit))
	if scaled.Value != 6 {
		t.Fatalf("Mul without overlapping temperature = %v, want 6", scaled.Value)
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
	if work.Unit.Symbol() != "J" {
		t.Fatalf("Derived Mul base unit = %q, want J", work.Unit.Symbol())
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
	if scaled.Unit.Symbol() != "J" {
		t.Fatalf("base Mul Derived unit = %q, want J", scaled.Unit.Symbol())
	}
}

func TestDerivedQuantitySqrtPowerResistance(t *testing.T) {
	// I = √(P/R): 20 mW / 5 kΩ → 2 mA
	p := Watt.Prefix(Milli).Of(20)
	r := Ohm.Of(5000).Prefix(Kilo)
	i := p.Div(r).Sqrt()
	if math.Abs(i.Value-0.002) > 1e-12 {
		t.Fatalf("√(P/R) value = %v, want 0.002 A", i.Value)
	}
	if !i.Unit.Dim().Equal(DerivedDimension{I: 1}) {
		t.Fatalf("√(P/R) dim = %#v, want {I:1}", i.Unit.Dim())
	}
	if i.Unit.Symbol() != "A" {
		t.Fatalf("√(P/R) unit = %q, want A", i.Unit.Symbol())
	}
}

func TestDerivedQuantitySqrtArea(t *testing.T) {
	area := Length(2, Meter.Prefix(Kilo)).Mul(Length(8, Meter.Prefix(Kilo))) // 16 km²
	side := area.Sqrt()
	if math.Abs(side.Value-4) > 1e-12 {
		t.Fatalf("√(area) value = %v, want 4 km", side.Value)
	}
	if side.Unit.Symbol(WithExpSign(ExpSignSup)) != "km" {
		t.Fatalf("√(area) unit = %q, want km", side.Unit.Symbol(WithExpSign(ExpSignSup)))
	}
}

func TestDerivedQuantityRootRejection(t *testing.T) {
	odd := Length(4, Meter).Mul(Length(1, Meter)).Div(Time(1, Second)) // m²/s — T exp not divisible by 2
	if got := odd.Sqrt(); got.Value != odd.Value || got.Unit != odd.Unit {
		t.Fatalf("Sqrt with odd-ready-fail dim should leave q unchanged, got %+v", got)
	}

	neg := Length(4, Meter).Mul(Length(1, Meter)).MulV(-1)
	if got := neg.Sqrt(); got.Value != neg.Value || got.Unit != neg.Unit {
		t.Fatalf("Sqrt of negative should leave q unchanged, got %+v", got)
	}

	cube := Length(8, Meter).Mul(Length(1, Meter)).Mul(Length(1, Meter)) // 8 m³
	side := cube.Root(3)
	if math.Abs(side.Value-2) > 1e-12 {
		t.Fatalf("Root(3) of 8 m³ value = %v, want 2 m", side.Value)
	}

	negCube := cube.MulV(-1)
	sideNeg := negCube.Root(3)
	if math.Abs(sideNeg.Value-(-2)) > 1e-12 {
		t.Fatalf("Root(3) of −8 m³ value = %v, want −2 m", sideNeg.Value)
	}

	if got := cube.Root(1); got.Value != cube.Value || got.Unit != cube.Unit {
		t.Fatalf("Root(1) should leave q unchanged, got %+v", got)
	}

	none := Mass(9, Kilogram).Div(Mass(1, Kilogram))
	rooted := none.Sqrt()
	if math.Abs(rooted.Value-3) > 1e-12 || rooted.Unit != NoneUnit {
		t.Fatalf("Sqrt of dimensionless = %+v, want 3 NoneUnit", rooted)
	}
}
