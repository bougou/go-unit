package u

import (
	"errors"
	"math"
	"testing"
)

func TestTryAddSubDimension(t *testing.T) {
	sum, err := Length(5, Meter.Prefix(Kilo)).TryAdd(Length(500, Meter))
	if err != nil {
		t.Fatalf("TryAdd: %v", err)
	}
	if sum.Value != 5.5 || sum.Unit != Unit(Meter.Prefix(Kilo)) {
		t.Fatalf("TryAdd = %+v, want 5.5 km", sum)
	}

	_, err = Quantity(Length(5, Meter)).TryAdd(Quantity(Mass(10, Kilogram)))
	if !errors.Is(err, ErrDimension) {
		t.Fatalf("TryAdd incompatible err = %v, want ErrDimension", err)
	}
}

func TestTryByAndDivV(t *testing.T) {
	got, err := Length(1000, Meter).TryBy(Meter.Prefix(Kilo))
	if err != nil {
		t.Fatalf("TryBy: %v", err)
	}
	if got.Value != 1 || got.Unit != Unit(Meter.Prefix(Kilo)) {
		t.Fatalf("TryBy = %+v, want 1 km", got)
	}

	_, err = Quantity(Length(1, Meter)).TryBy(Unit(Second))
	if !errors.Is(err, ErrDimension) {
		t.Fatalf("TryBy cross-dimension err = %v, want ErrDimension", err)
	}

	_, err = Length(3, Meter).TryDivV(0)
	if !errors.Is(err, ErrDivByZero) {
		t.Fatalf("TryDivV(0) err = %v, want ErrDivByZero", err)
	}
}

func TestTryDerivedAddIncompatible(t *testing.T) {
	celsiusLen := NewDerivedUnit().
		Temperature(TemperatureUnit(Celsius), 1).
		Length(LengthUnit(Meter), 1).MustIntern()
	kelvinLen := NewDerivedUnit().
		Temperature(TemperatureUnit(Kelvin), 1).
		Length(LengthUnit(Meter), 1).MustIntern()
	cq := NewDerivedQuantity(20, celsiusLen)
	kq := NewDerivedQuantity(293.15, kelvinLen)

	_, err := cq.TryAdd(kq)
	if !errors.Is(err, ErrIncompatible) {
		t.Fatalf("TryAdd affine overlap err = %v, want ErrIncompatible", err)
	}
}

func TestTryRoot(t *testing.T) {
	area := Length(2, Meter.Prefix(Kilo)).Mul(Length(8, Meter.Prefix(Kilo)))
	side, err := area.TrySqrt()
	if err != nil {
		t.Fatalf("TrySqrt: %v", err)
	}
	if math.Abs(side.Value-4) > 1e-12 {
		t.Fatalf("TrySqrt value = %v, want 4", side.Value)
	}

	odd := Length(4, Meter).Mul(Length(1, Meter)).Div(Time(1, Second))
	_, err = odd.TrySqrt()
	if !errors.Is(err, ErrRoot) {
		t.Fatalf("TrySqrt odd exponents err = %v, want ErrRoot", err)
	}

	_, err = area.TryRoot(1)
	if !errors.Is(err, ErrRoot) {
		t.Fatalf("TryRoot(1) err = %v, want ErrRoot", err)
	}
}

func TestTryMulDiv(t *testing.T) {
	speed, err := Length(10, Meter.Prefix(Kilo)).TryDiv(Time(2, Hour))
	if err != nil {
		t.Fatalf("TryDiv: %v", err)
	}
	if speed.Value != 5 {
		t.Fatalf("TryDiv value = %v, want 5", speed.Value)
	}

	_, err = Newton.Of(10).TryDiv(Newton.Of(0))
	if !errors.Is(err, ErrDivByZero) {
		t.Fatalf("TryDiv by zero err = %v, want ErrDivByZero", err)
	}
}
