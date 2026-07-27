package u

// Dimensionless (无量纲) and other derived units may share the same base-unit composition but
// carry different SI special names (专用名称). Intern (内化注册) distinguishes named variants by appending
// "#<symbol>" to the composition key (dimensionless named units use "#<symbol>" only).
// The first registrant for a composition also claims the bare composition key so
// unnamed Intern calls can share that canonical instance.
//
// Radian and Steradian are registered in unit_si_derived.go.

// NoneUnit is the generic unnamed dimensionless unit, e.g. from same-dimension division.
// Compare with == after Intern; e.g. Mass(10, Kilogram).Div(Mass(2, Kilogram)).Unit == NoneUnit.
var NoneUnit *DerivedUnit

func init() {
	NoneUnit = NewDerivedUnit().MustIntern()
}

// Dimensionless creates a dimensionless derived quantity with the generic unnamed unit.
//
// Example: Dimensionless(3.5) // "3.5"
func Dimensionless(value float64) DerivedQuantity {
	return DerivedQuantity{Value: value, Unit: NoneUnit}
}
