package u

import "testing"

func TestDerivedDimensionMulDiv(t *testing.T) {
	// F × L → energy (M L² T⁻²)
	got := DimForce.Mul(DerivedDimension{L: 1})
	if !got.Equal(DimEnergy) {
		t.Fatalf("DimForce.Mul(L): got %#v, want %#v", got, DimEnergy)
	}

	// energy / time → power
	got = DimEnergy.Div(DerivedDimension{T: 1})
	if !got.Equal(DimPower) {
		t.Fatalf("DimEnergy.Div(T): got %#v, want %#v", got, DimPower)
	}

	// length / length → dimensionless
	got = DerivedDimension{L: 1}.Div(DerivedDimension{L: 1})
	if !got.Equal(DerivedDimension{}) {
		t.Fatalf("L.Div(L): got %#v, want dimensionless", got)
	}
}

func TestDerivedDimensionRoot(t *testing.T) {
	got, ok := DerivedDimension{I: 2}.Root(2)
	if !ok || !got.Equal(DerivedDimension{I: 1}) {
		t.Fatalf("I².Root(2): got (%#v, %v), want ({I:1}, true)", got, ok)
	}

	got, ok = DerivedDimension{L: 4, T: -2}.Root(2)
	if !ok || !got.Equal(DerivedDimension{L: 2, T: -1}) {
		t.Fatalf("L⁴T⁻².Root(2): got (%#v, %v), want ({L:2,T:-1}, true)", got, ok)
	}

	_, ok = DerivedDimension{L: 1}.Root(2)
	if ok {
		t.Fatal("L¹.Root(2): want ok=false (odd exponent)")
	}

	_, ok = DerivedDimension{I: 2}.Root(1)
	if ok {
		t.Fatal("Root(1): want ok=false (n < 2)")
	}

	got, ok = DimPower.Div(DimResistance).Root(2)
	if !ok || !got.Equal(DerivedDimension{I: 1}) {
		t.Fatalf("Power/Resistance Root(2): got (%#v, %v), want ({I:1}, true)", got, ok)
	}
}

func TestDerivedDimensionAddSubHomogeneous(t *testing.T) {
	sum, ok := DimEnergy.Add(DimEnergy)
	if !ok || !sum.Equal(DimEnergy) {
		t.Fatalf("DimEnergy.Add(DimEnergy): got (%#v, %v), want (DimEnergy, true)", sum, ok)
	}

	diff, ok := DimEnergy.Sub(DimEnergy)
	if !ok || !diff.Equal(DimEnergy) {
		t.Fatalf("DimEnergy.Sub(DimEnergy): got (%#v, %v), want (DimEnergy, true)", diff, ok)
	}

	_, ok = DimEnergy.Add(DimForce)
	if ok {
		t.Fatal("DimEnergy.Add(DimForce): want ok=false (dimensional inconsistency)")
	}

	_, ok = DimEnergy.Sub(DimSpeed)
	if ok {
		t.Fatal("DimEnergy.Sub(DimSpeed): want ok=false (dimensional inconsistency)")
	}
}
