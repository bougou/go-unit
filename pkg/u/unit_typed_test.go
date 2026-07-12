package u

import "testing"

func TestLengthQuantityBaseAndBy(t *testing.T) {
	q := Length(1, LengthUnit(Kilometer))

	base := q.Base()
	if base.Value != 1000 || base.Unit != Unit(Meter) {
		t.Fatalf("Base() = %+v, want 1000 m", base)
	}

	back := base.By(LengthUnit(Kilometer))
	if back.Value != 1 || back.Unit != Unit(Kilometer) {
		t.Fatalf("By(km) = %+v, want 1 kilometer", back)
	}
}

func TestX(t *testing.T) {
	Length(100, Kilometer)
}
