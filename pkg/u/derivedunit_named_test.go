package u

import "testing"

func TestNamedUnitsSameCompositionDistinct(t *testing.T) {
	if Hertz == Becquerel {
		t.Fatal("Hz and Bq should be distinct canonical instances")
	}
	if Gray == Sievert {
		t.Fatal("Gy and Sv should be distinct canonical instances")
	}
	if Hertz.Key() != "second^-1#Hz" || Becquerel.Key() != "second^-1#Bq" {
		t.Fatalf("Hz key %q Bq key %q", Hertz.Key(), Becquerel.Key())
	}
	if Gray.Key() != "meter^2*second^-2#Gy" || Sievert.Key() != "meter^2*second^-2#Sv" {
		t.Fatalf("Gy key %q Sv key %q", Gray.Key(), Sievert.Key())
	}
}

func TestNamedUnitsSameCompositionSameDimension(t *testing.T) {
	if !Hertz.Dim().Equal(Becquerel.Dim()) || !Hertz.Dim().Equal(DimFrequency) {
		t.Fatalf("Hz dim %+v Bq dim %+v", Hertz.Dim(), Becquerel.Dim())
	}
	if !Gray.Dim().Equal(Sievert.Dim()) || !Gray.Dim().Equal(DimAbsorbedDose) {
		t.Fatalf("Gy dim %+v Sv dim %+v", Gray.Dim(), Sievert.Dim())
	}
}

func TestNamedUnitsSameCompositionConvertible(t *testing.T) {
	fiftyHz := NewDerivedQuantity(50, Hertz)
	asBq := fiftyHz.By(Becquerel)
	if asBq.Value != 50 || asBq.Unit != Becquerel {
		t.Fatalf("By(Becquerel) = %+v, want 50 Bq", asBq)
	}
}

func TestUnnamedFirstThenNamedSameComposition(t *testing.T) {
	generic, err := NewDerivedUnit().Length(Meter, 2).Time(Second, -2).Intern()
	if err != nil {
		t.Fatal(err)
	}
	named, err := NewDerivedUnit().Length(Meter, 2).Time(Second, -2).Named("foo").Intern()
	if err != nil {
		t.Fatal(err)
	}
	if generic == named {
		t.Fatal("generic and differently named units should not share identity")
	}
	if !generic.Dim().Equal(named.Dim()) {
		t.Fatal("dimensions should still match")
	}
}
