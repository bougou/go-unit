package u

import "testing"

func TestRegisterDerivedUnitForce(t *testing.T) {
	want := DerivedDimension{M: 1, L: 1, T: -2}
	dim := derivedDimensionFromTerms([]unitTerm{
		{unit: Unit(Kilogram), exp: 1},
		{unit: Unit(Meter), exp: 1},
		{unit: Unit(Second), exp: -2},
	})
	if !dim.Equal(want) {
		t.Fatalf("dim = %+v, want %+v", dim, want)
	}

	def, err := derivedUnitFromTerms([]unitTerm{
		{unit: Unit(Kilogram), exp: 1},
		{unit: Unit(Meter), exp: 1},
		{unit: Unit(Second), exp: -2},
	})
	if err != nil {
		t.Fatal(err)
	}
	if def.Symbol(WithExpSign(ExpSignSup)) != "kg·m·s⁻²" {
		t.Fatalf("sign = %q, want kg·m·s⁻²", def.Symbol(WithExpSign(ExpSignSup)))
	}
}
