package u

import (
	"math"
	"testing"
)

func TestElectricalDerivedDimensions(t *testing.T) {
	cases := []struct {
		name string
		unit *DerivedUnit
		dim  DerivedDimension
	}{
		{"Volt", Volt, DimElectricPotential},
		{"Ohm", Ohm, DimResistance},
		{"Impedance", Impedance, DimResistance},
		{"Reactance", Reactance, DimResistance},
		{"Siemens", Siemens, DimConductance},
		{"Admittance", Admittance, DimConductance},
		{"Susceptance", Susceptance, DimConductance},
		{"Coulomb", Coulomb, DimElectricCharge},
		{"AmpereHour", AmpereHour, DimElectricCharge},
		{"Farad", Farad, DimCapacitance},
		{"Henry", Henry, DimInductance},
		{"Watt", Watt, DimPower},
		{"VoltAmpere", VoltAmpere, DimPower},
		{"Var", Var, DimPower},
		{"ElectricField", ElectricField, DimElectricField},
		{"ElectricDisplacement", ElectricDisplacement, DimElectricDisplacement},
		{"ElectricFlux", ElectricFlux, DimElectricFlux},
		{"CurrentDensity", CurrentDensity, DimCurrentDensity},
		{"Resistivity", Resistivity, DimResistivity},
		{"Conductivity", Conductivity, DimConductivity},
		{"Permittivity", Permittivity, DimPermittivity},
		{"Permeability", Permeability, DimPermeability},
		{"MagneticFieldStrength", MagneticFieldStrength, DimMagneticFieldStrength},
		{"Reluctance", Reluctance, DimReluctance},
		{"Weber", Weber, DimMagneticFlux},
		{"Tesla", Tesla, DimMagneticFluxDensity},
	}
	for _, tc := range cases {
		if !tc.unit.Dim().Equal(tc.dim) {
			t.Fatalf("%s dim = %+v, want %+v", tc.name, tc.unit.Dim(), tc.dim)
		}
	}
}

func TestElectricalNamedSymbols(t *testing.T) {
	cases := []struct {
		unit *DerivedUnit
		want string
	}{
		{VoltAmpere, "VA"},
		{Var, "var"},
		{AmpereHour, "A·h"},
		{ElectricField, "V/m"},
		{ElectricDisplacement, "C/m²"},
		{ElectricFlux, "V·m"},
		{CurrentDensity, "A/m²"},
		{Resistivity, "Ω·m"},
		{Conductivity, "S/m"},
		{Permittivity, "F/m"},
		{Permeability, "H/m"},
		{MagneticFieldStrength, "A/m"},
		{Reluctance, "H^-1"},
		{ElectricField.Prefix(Kilo), "kV/m"},
		{Resistivity.Prefix(Milli), "mΩ·m"},
		{AmpereHour.Prefix(Kilo), "kA·h"},
	}
	for _, tc := range cases {
		if got := tc.unit.Symbol(); got != tc.want {
			t.Fatalf("Symbol() = %q, want %q", got, tc.want)
		}
	}
}

func TestElectricalCircuitAliases(t *testing.T) {
	if Impedance != Ohm || Reactance != Ohm {
		t.Fatal("Impedance/Reactance should alias Ohm")
	}
	if Admittance != Siemens || Susceptance != Siemens {
		t.Fatal("Admittance/Susceptance should alias Siemens")
	}
}

func TestAmpereHourScale(t *testing.T) {
	if AmpereHour.FactorToBase() != 3600 {
		t.Fatalf("AmpereHour FactorToBase = %g, want 3600", AmpereHour.FactorToBase())
	}
	q := AmpereHour.Of(2).By(Coulomb)
	if q.Value != 7200 || q.Unit != Coulomb {
		t.Fatalf("2 A·h → By(Coulomb) = %v %q, want 7200 C", q.Value, q.Unit.Symbol())
	}
}

func TestApparentAndReactivePowerSameAsWatt(t *testing.T) {
	if Watt == VoltAmpere || Watt == Var || VoltAmpere == Var {
		t.Fatal("W, VA, var should be distinct named instances")
	}
	s := VoltAmpere.Of(100).By(Watt)
	if s.Value != 100 || s.Unit != Watt {
		t.Fatalf("100 VA → By(Watt) = %v %q", s.Value, s.Unit.Symbol())
	}
	q := Var.Of(50).By(Watt)
	if q.Value != 50 || q.Unit != Watt {
		t.Fatalf("50 var → By(Watt) = %v %q", q.Value, q.Unit.Symbol())
	}
}

func TestElectricalDimensionAlgebra(t *testing.T) {
	length := DerivedDimension{L: 1}
	area := DerivedDimension{L: 2}
	current := DerivedDimension{I: 1}

	// E = V / m
	if !DimElectricPotential.Div(length).Equal(DimElectricField) {
		t.Fatal("V/m should be DimElectricField")
	}
	// resistivity = ohm * m
	if !DimResistance.Mul(length).Equal(DimResistivity) {
		t.Fatal("ohm·m should be DimResistivity")
	}
	// conductivity = S / m = 1 / resistivity
	if !DimConductance.Div(length).Equal(DimConductivity) {
		t.Fatalf("S/m dim = %+v, want %+v", DimConductance.Div(length), DimConductivity)
	}
	if !DimAngle.Div(DimResistivity).Equal(DimConductivity) {
		t.Fatal("1/resistivity should be DimConductivity")
	}
	// permittivity = F / m, permeability = H / m
	if !DimCapacitance.Div(length).Equal(DimPermittivity) {
		t.Fatal("F/m should be DimPermittivity")
	}
	if !DimInductance.Div(length).Equal(DimPermeability) {
		t.Fatal("H/m should be DimPermeability")
	}
	// reluctance = 1 / H
	if !DimAngle.Div(DimInductance).Equal(DimReluctance) {
		t.Fatal("1/H should be DimReluctance")
	}
	// D = C / m^2, J = A / m^2, H = A / m
	if !DimElectricCharge.Div(area).Equal(DimElectricDisplacement) {
		t.Fatal("C/m^2 should be DimElectricDisplacement")
	}
	if !current.Div(area).Equal(DimCurrentDensity) {
		t.Fatal("A/m^2 should be DimCurrentDensity")
	}
	if !current.Div(length).Equal(DimMagneticFieldStrength) {
		t.Fatal("A/m should be DimMagneticFieldStrength")
	}
	// electric flux = E * m^2 = V·m
	if !DimElectricField.Mul(area).Equal(DimElectricFlux) {
		t.Fatal("E·m^2 should be DimElectricFlux")
	}
}

func TestElectricalQuantityRelations(t *testing.T) {
	// E = U / d
	e := Volt.Of(220).Div(Meter.Of(2))
	e = e.By(ElectricField)
	if math.Abs(e.Value-110) > 1e-12 || e.Unit != ElectricField {
		t.Fatalf("U/d = %v %q, want 110 V/m", e.Value, e.Unit.Symbol())
	}

	// ρ = R · A / L → here check Ω·m via Mul length
	rho := Ohm.Of(1.5).Mul(Meter.Of(1)).By(Resistivity)
	if math.Abs(rho.Value-1.5) > 1e-12 || rho.Unit != Resistivity {
		t.Fatalf("Ω·m = %v %q, want 1.5 Ω·m", rho.Value, rho.Unit.Symbol())
	}

	// σ = 1/ρ
	sigma := Dimensionless(1).Div(Resistivity.Of(2)).By(Conductivity)
	if math.Abs(sigma.Value-0.5) > 1e-12 || sigma.Unit != Conductivity {
		t.Fatalf("1/ρ = %v %q, want 0.5 S/m", sigma.Value, sigma.Unit.Symbol())
	}
}
