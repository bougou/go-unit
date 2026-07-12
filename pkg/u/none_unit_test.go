package u

import "testing"

func TestDimensionlessUnitsDistinct(t *testing.T) {
	if RadianUnit == SteradianUnit || RadianUnit == NoneUnit || SteradianUnit == NoneUnit {
		t.Fatal("dimensionless SI units should be distinct canonical instances")
	}
	if RadianUnit.Key() != "#rad" || SteradianUnit.Key() != "#sr" || NoneUnit.Key() != "" {
		t.Fatalf("keys = rad %q sr %q none %q", RadianUnit.Key(), SteradianUnit.Key(), NoneUnit.Key())
	}
}

func TestDimensionlessUnitsSameDimension(t *testing.T) {
	want := DerivedDimension{}
	for _, u := range []*DerivedUnit{NoneUnit, RadianUnit, SteradianUnit} {
		if !u.Dim().Equal(want) {
			t.Fatalf("%v Dim() = %+v, want zero", u, u.Dim())
		}
	}
}

func TestDimensionlessUnitsConvertible(t *testing.T) {
	pi := NewDerivedQuantity(3.14, RadianUnit)
	asSteradian := pi.By(SteradianUnit)
	if asSteradian.Value != 3.14 || asSteradian.Unit != SteradianUnit {
		t.Fatalf("By(SteradianUnit) = %+v, want 3.14 sr", asSteradian)
	}
}

func TestDimensionlessUnitsNamedSymbol(t *testing.T) {
	if RadianUnit.Symbol(WithNamedSymbol(true)) != "rad" {
		t.Fatalf("RadianUnit named symbol = %q, want rad", RadianUnit.Symbol(WithNamedSymbol(true)))
	}
	if SteradianUnit.Symbol(WithNamedSymbol(true)) != "sr" {
		t.Fatalf("SteradianUnit named symbol = %q, want sr", SteradianUnit.Symbol(WithNamedSymbol(true)))
	}
}

func TestDimensionlessUnitsIntern(t *testing.T) {
	again, err := NewDerivedUnit().Named("rad").Intern()
	if err != nil {
		t.Fatal(err)
	}
	if again != RadianUnit {
		t.Fatal("Intern should return canonical RadianUnit")
	}
}

func TestDimensionless(t *testing.T) {
	q := Dimensionless(3.5)
	if q.Value != 3.5 || q.Unit != NoneUnit {
		t.Fatalf("Dimensionless = %+v", q)
	}
}
