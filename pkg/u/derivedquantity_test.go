package u

import (
	"math"
	"testing"

	"github.com/bougou/go-unit/pkg/prefix"
)

func TestUnitsProportional(t *testing.T) {
	if !unitsProportional(Unit(Meter.Prefix(prefix.Kilo)), Unit(Meter)) {
		t.Fatal("km and m should be proportional")
	}
	if unitsProportional(Unit(Celsius), Unit(Kelvin)) {
		t.Fatal("°C and K should not be proportional")
	}
	if !unitsProportional(Unit(Meter), Unit(Meter)) {
		t.Fatal("same unit should be proportional")
	}
}

func TestDerivedQuantityBy(t *testing.T) {
	speedUnit := NewDerivedUnit().Length(LengthUnit(Meter.Prefix(prefix.Kilo)), 1).Time(TimeUnit(Hour), -1)
	speed := NewDerivedQuantity(60, speedUnit)

	siUnit := speedUnit.Base()
	got := speed.By(siUnit)
	want := 60.0 * 1000.0 / 3600.0
	if math.Abs(got.Value-want) > 1e-12 {
		t.Fatalf("By SI value = %v, want %v", got.Value, want)
	}
	if got.Unit.Symbol(WithExpSign(ExpSignSup)) != "m·s⁻¹" {
		t.Fatalf("By SI unit = %q, want m·s⁻¹", got.Unit.Symbol(WithExpSign(ExpSignSup)))
	}

	otherDim := NewDerivedUnit().Length(LengthUnit(Meter), 2)
	if changed := speed.By(otherDim); changed.Value != speed.Value || changed.Unit != speed.Unit {
		t.Fatalf("incompatible By should return unchanged, got %+v", changed)
	}
}

func TestDerivedQuantityByAffineTemperature(t *testing.T) {
	unit := NewDerivedUnit().
		Temperature(TemperatureUnit(Celsius), 1).
		Length(LengthUnit(Meter), 1)
	q := NewDerivedQuantity(100, unit)

	target := NewDerivedUnit().
		Temperature(TemperatureUnit(Kelvin), 1).
		Length(LengthUnit(Meter), 1)
	got := q.By(target)
	if got.Value != q.Value || got.Unit != q.Unit {
		t.Fatalf("affine temperature By should return unchanged, got %+v", got)
	}
}

func TestDerivedQuantityByLength(t *testing.T) {
	speedUnit := NewDerivedUnit().Length(LengthUnit(Meter.Prefix(prefix.Kilo)), 1).Time(TimeUnit(Hour), -1)
	speed := NewDerivedQuantity(60, speedUnit)

	got := speed.ByLength(Meter)
	if math.Abs(got.Value-60000) > 1e-12 {
		t.Fatalf("ByLength value = %v, want 60000", got.Value)
	}
	if got.Unit.Symbol(WithExpSign(ExpSignSup)) != "m·h⁻¹" {
		t.Fatalf("ByLength unit = %q, want m·h⁻¹", got.Unit.Symbol(WithExpSign(ExpSignSup)))
	}
	if got.Unit.t.unit != Unit(Hour) {
		t.Fatalf("time unit should stay hour, got %v", got.Unit.t.unit)
	}
}

func TestDerivedQuantityByTime(t *testing.T) {
	speedUnit := NewDerivedUnit().Length(LengthUnit(Meter.Prefix(prefix.Kilo)), 1).Time(TimeUnit(Hour), -1)
	speed := NewDerivedQuantity(60, speedUnit)

	got := speed.ByTime(Second)
	want := 60.0 / 3600.0
	if math.Abs(got.Value-want) > 1e-12 {
		t.Fatalf("ByTime value = %v, want %v", got.Value, want)
	}
	if got.Unit.Symbol(WithExpSign(ExpSignSup)) != "km·s⁻¹" {
		t.Fatalf("ByTime unit = %q, want km·s⁻¹", got.Unit.Symbol(WithExpSign(ExpSignSup)))
	}
}

func TestDerivedQuantityByNoOpChaining(t *testing.T) {
	speedUnit := NewDerivedUnit().Length(LengthUnit(Meter.Prefix(prefix.Kilo)), 1).Time(TimeUnit(Hour), -1)
	speed := NewDerivedQuantity(60, speedUnit)

	affineUnit := NewDerivedUnit().
		Temperature(TemperatureUnit(Celsius), 1).
		Length(LengthUnit(Meter), 1)
	q := NewDerivedQuantity(10, affineUnit)

	got := q.ByTemperature(Kelvin).ByLength(Meter).MulV(2)
	if got.Value != 20 || got.Unit != q.Unit {
		t.Fatalf("chaining after no-op ByTemperature = %+v, want value 20 unchanged unit", got)
	}

	got = speed.ByLength(Meter).ByTime(Second).Base()
	if math.Abs(got.Value-16.666666666666668) > 1e-12 {
		t.Fatalf("chaining proportional conversions then SI = %v", got.Value)
	}
}
