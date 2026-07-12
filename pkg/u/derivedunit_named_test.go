package u

import "testing"

func TestNamedUnitsSameCompositionDistinct(t *testing.T) {
	if HertzUnit == BecquerelUnit {
		t.Fatal("Hz and Bq should be distinct canonical instances")
	}
	if GrayUnit == SievertUnit {
		t.Fatal("Gy and Sv should be distinct canonical instances")
	}
	if HertzUnit.Key() != "second^-1#Hz" || BecquerelUnit.Key() != "second^-1#Bq" {
		t.Fatalf("Hz key %q Bq key %q", HertzUnit.Key(), BecquerelUnit.Key())
	}
	if GrayUnit.Key() != "meter^2*second^-2#Gy" || SievertUnit.Key() != "meter^2*second^-2#Sv" {
		t.Fatalf("Gy key %q Sv key %q", GrayUnit.Key(), SievertUnit.Key())
	}
}

func TestNamedUnitsSameCompositionSameDimension(t *testing.T) {
	if !HertzUnit.Dim().Equal(BecquerelUnit.Dim()) || !HertzUnit.Dim().Equal(DimFrequency) {
		t.Fatalf("Hz dim %+v Bq dim %+v", HertzUnit.Dim(), BecquerelUnit.Dim())
	}
	if !GrayUnit.Dim().Equal(SievertUnit.Dim()) || !GrayUnit.Dim().Equal(DimAbsorbedDose) {
		t.Fatalf("Gy dim %+v Sv dim %+v", GrayUnit.Dim(), SievertUnit.Dim())
	}
}

func TestNamedUnitsSameCompositionConvertible(t *testing.T) {
	fiftyHz := NewDerivedQuantity(50, HertzUnit)
	asBq := fiftyHz.By(BecquerelUnit)
	if asBq.Value != 50 || asBq.Unit != BecquerelUnit {
		t.Fatalf("By(BecquerelUnit) = %+v, want 50 Bq", asBq)
	}
}

func TestNamedUnitAliasesUnnamedIntern(t *testing.T) {
	unnamed, err := NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 1).Time(Second, -2).Intern()
	if err != nil {
		t.Fatal(err)
	}
	if unnamed != ForceUnit {
		t.Fatal("unnamed force intern should alias to ForceUnit registered with N")
	}

	generic, err := NewDerivedUnit().Time(Second, -1).Intern()
	if err != nil {
		t.Fatal(err)
	}
	if generic != HertzUnit {
		t.Fatal("unnamed s^-1 intern should alias to first named registrant HertzUnit")
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
