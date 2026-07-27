package u

import (
	"math"
	"testing"
)

func TestDerivedUnitParse(t *testing.T) {
	tests := []struct {
		symbol string
		check  func(t *testing.T, du *DerivedUnit)
	}{
		{
			symbol: "km/h",
			check: func(t *testing.T, du *DerivedUnit) {
				if du.Dim() != DimSpeed {
					t.Fatalf("dim = %+v, want speed", du.Dim())
				}
				if du.l.unit != Unit(Meter.Prefix(Kilo)) || du.l.exp != 1 {
					t.Fatalf("length term = %+v", du.l)
				}
				if du.t.unit != Unit(Hour) || du.t.exp != -1 {
					t.Fatalf("time term = %+v", du.t)
				}
			},
		},
		{
			symbol: "kg·m·s^-2",
			check: func(t *testing.T, du *DerivedUnit) {
				if du.Dim() != DimForce {
					t.Fatalf("dim = %+v, want force", du.Dim())
				}
			},
		},
		{
			symbol: "kg·m/s^2",
			check: func(t *testing.T, du *DerivedUnit) {
				if du.Dim() != DimForce {
					t.Fatalf("dim = %+v, want force", du.Dim())
				}
			},
		},
		{
			symbol: "m·s⁻¹",
			check: func(t *testing.T, du *DerivedUnit) {
				if du.Dim() != DimSpeed {
					t.Fatalf("dim = %+v, want speed", du.Dim())
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.symbol, func(t *testing.T) {
			du, err := DerivedUnitParse(tt.symbol)
			if err != nil {
				t.Fatal(err)
			}
			tt.check(t, du)
		})
	}
}

func TestDerivedUnitParseSpecialNameRejected(t *testing.T) {
	if _, err := DerivedUnitParse("N"); err == nil {
		t.Fatal("DerivedUnitParse(N) should fail")
	}
}

func TestDerivedUnitParseSeparatorRules(t *testing.T) {
	if _, err := DerivedUnitParse("kmh"); err == nil {
		t.Fatal("DerivedUnitParse(kmh) should fail without separator")
	}

	du, err := DerivedUnitParse("ms")
	if err != nil {
		t.Fatal(err)
	}
	if du.t.unit != Unit(Second.Prefix(Milli)) || du.t.exp != 1 {
		t.Fatalf("ms = %+v, want millisecond", du.t)
	}

	force, err := DerivedUnitParse("kg m s^-2")
	if err != nil {
		t.Fatal(err)
	}
	if force.Dim() != DimForce {
		t.Fatalf("kg m s^-2 dim = %+v, want force", force.Dim())
	}

	speed, err := DerivedUnitParse("kg m/s^2")
	if err != nil {
		t.Fatal(err)
	}
	if speed.Dim() != DimForce {
		t.Fatalf("kg m/s^2 dim = %+v, want force", speed.Dim())
	}

	mps, err := DerivedUnitParse("m s^-1")
	if err != nil {
		t.Fatal(err)
	}
	if mps.Dim() != DimSpeed {
		t.Fatalf("m s^-1 dim = %+v, want speed", mps.Dim())
	}
}

func TestQuantityParseUnregisteredDerived(t *testing.T) {
	q, err := QuantityParse("60 km/h")
	if err != nil {
		t.Fatal(err)
	}
	du, ok := q.Unit.DerivedUnit()
	if !ok {
		t.Fatal("expected derived unit")
	}
	if du.Dim() != DimSpeed {
		t.Fatalf("dim = %+v, want speed", du.Dim())
	}
	if math.Abs(q.Value-60) > 1e-12 {
		t.Fatalf("value = %v, want 60", q.Value)
	}

	q2, err := QuantityParse("60km/h")
	if err != nil {
		t.Fatal(err)
	}
	if q2.Value != 60 {
		t.Fatalf("value = %v, want 60", q2.Value)
	}
}

func TestDerivedQuantityParseKmPerHour(t *testing.T) {
	q, err := DerivedQuantityParse("60 km/h")
	if err != nil {
		t.Fatal(err)
	}
	if q.Unit.Dim() != DimSpeed || q.Value != 60 {
		t.Fatalf("got %+v", q)
	}
}
