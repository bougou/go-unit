package u

import "testing"

func TestDimensionlessUnitsDistinct(t *testing.T) {
	if Radian == Steradian || Radian == NoneUnit || Steradian == NoneUnit {
		t.Fatal("dimensionless SI units should be distinct canonical instances")
	}
	if Radian.Key() != "#rad" || Steradian.Key() != "#sr" || NoneUnit.Key() != "" {
		t.Fatalf("keys = rad %q sr %q none %q", Radian.Key(), Steradian.Key(), NoneUnit.Key())
	}
}

func TestDimensionlessUnitsSameDimension(t *testing.T) {
	want := DerivedDimension{}
	for _, u := range []*DerivedUnit{NoneUnit, Radian, Steradian} {
		if !u.Dim().Equal(want) {
			t.Fatalf("%v Dim() = %+v, want zero", u, u.Dim())
		}
	}
}

func TestDimensionlessUnitsConvertible(t *testing.T) {
	pi := NewDerivedQuantity(3.14, Radian)
	asSteradian := pi.By(Steradian)
	if asSteradian.Value != 3.14 || asSteradian.Unit != Steradian {
		t.Fatalf("By(Steradian) = %+v, want 3.14 sr", asSteradian)
	}
}

func TestDimensionlessUnitsNamedSymbol(t *testing.T) {
	if Radian.Symbol() != "rad" {
		t.Fatalf("Radian named symbol = %q, want rad", Radian.Symbol())
	}
	if Steradian.Symbol() != "sr" {
		t.Fatalf("Steradian named symbol = %q, want sr", Steradian.Symbol())
	}
}

func TestDimensionlessUnitsIntern(t *testing.T) {
	again, err := NewDerivedUnit().Named("rad").Intern()
	if err != nil {
		t.Fatal(err)
	}
	if again != Radian {
		t.Fatal("Intern should return canonical Radian")
	}
}

func TestDimensionless(t *testing.T) {
	q := Dimensionless(3.5)
	if q.Value != 3.5 || q.Unit != NoneUnit {
		t.Fatalf("Dimensionless = %+v", q)
	}
}
