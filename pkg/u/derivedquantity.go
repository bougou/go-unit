package u

import (
	"fmt"
	"math"
)

// DerivedQuantity (导出量) is a numeric value with a compound derived unit.
//
// Example:
//
//	unit := NewDerivedUnit().Length(Kilometer, 1).Time(Hour, -1)
//	speed := NewDerivedQuantity(60, unit) // 60 km/h
type DerivedQuantity struct {
	// Value is the numeric magnitude in Unit.
	Value float64
	// Unit is the compound derived unit. Nil means dimensionless display.
	Unit *DerivedUnit
}

// NewDerivedQuantity creates a derived quantity, e.g. 10 km·h⁻¹.
func NewDerivedQuantity(value float64, unit *DerivedUnit) DerivedQuantity {
	return DerivedQuantity{Value: value, Unit: unit}
}

// DerivedQuantityParse parses a derived quantity from text.
// Base-dimension inputs such as "10 m" are rejected; use typed parsers instead.
//
// Example: DerivedQuantityParse("5 N") // 5 N
func DerivedQuantityParse(s string) (DerivedQuantity, error) {
	q, err := QuantityParse(s)
	if err != nil {
		return DerivedQuantity{}, err
	}
	du, ok := q.Unit.DerivedUnit()
	if !ok {
		return DerivedQuantity{}, fmt.Errorf("unit %q is not a derived unit: %w", q.Unit.Symbol(), ErrDimension)
	}
	return DerivedQuantity{Value: q.Value, Unit: du}, nil
}

// DerivedQuantityMustParse parses a derived quantity and panics on error.
func DerivedQuantityMustParse(s string) DerivedQuantity {
	q, err := DerivedQuantityParse(s)
	if err != nil {
		panic(err)
	}
	return q
}

// String formats q as "value unit" using the compound unit symbol.
//
// Example: NewDerivedQuantity(10, ForceUnit).String() // "10 kg·m·s⁻²"
func (q DerivedQuantity) String() string {
	if q.Unit == nil {
		return fmt.Sprintf("%g", q.Value)
	}
	sign := q.Unit.Symbol(WithExpSign(ExpSignSup))
	if sign == "" {
		return fmt.Sprintf("%g", q.Value)
	}
	return fmt.Sprintf("%g %s", q.Value, sign)
}

// SI converts the quantity to SI base units for each dimension.
func (q DerivedQuantity) SI() DerivedQuantity {
	return DerivedQuantity{
		Value: q.Value * mustFactorToBase(q.Unit),
		Unit:  q.Unit.SI(),
	}
}

// By converts q to target when both share the same derived dimension and every
// dimension whose unit changes converts proportionally via the dimension base unit.
// Otherwise q is returned unchanged.
//
// Example: speed in km/h converted to m/s via By on the SI unit
func (q DerivedQuantity) By(target *DerivedUnit) DerivedQuantity {
	if q.Unit == nil || target == nil {
		return q
	}
	if !mustDerivedDim(q.Unit).Equal(mustDerivedDim(target)) {
		return q
	}
	if !derivedUnitsConvertible(q.Unit, target) {
		return q
	}
	return DerivedQuantity{
		Value: q.Value * derivedUnitConversionFactor(q.Unit, target),
		Unit:  target,
	}
}

// ByLength converts only the length term of q's unit to u.
// Other dimensions are unchanged. Non-proportional conversion returns q unchanged.
func (q DerivedQuantity) ByLength(u LengthUnit) DerivedQuantity {
	return q.byDimension(DimLength, Unit(u))
}

// ByMass converts only the mass term of q's unit to u.
func (q DerivedQuantity) ByMass(u MassUnit) DerivedQuantity {
	return q.byDimension(DimMass, Unit(u))
}

// ByTime converts only the time term of q's unit to u.
func (q DerivedQuantity) ByTime(u TimeUnit) DerivedQuantity {
	return q.byDimension(DimTime, Unit(u))
}

// ByCurrent converts only the current term of q's unit to u.
func (q DerivedQuantity) ByCurrent(u CurrentUnit) DerivedQuantity {
	return q.byDimension(DimCurrent, Unit(u))
}

// ByTemperature converts only the temperature term of q's unit to u.
func (q DerivedQuantity) ByTemperature(u TemperatureUnit) DerivedQuantity {
	return q.byDimension(DimTemperature, Unit(u))
}

// ByAmount converts only the amount-of-substance term of q's unit to u.
func (q DerivedQuantity) ByAmount(u AmountUnit) DerivedQuantity {
	return q.byDimension(DimAmount, Unit(u))
}

// ByLuminous converts only the luminous-intensity term of q's unit to u.
func (q DerivedQuantity) ByLuminous(u LuminousUnit) DerivedQuantity {
	return q.byDimension(DimLuminous, Unit(u))
}

func (q DerivedQuantity) byDimension(dim Dimension, newUnit Unit) DerivedQuantity {
	if q.Unit == nil {
		return q
	}
	term := q.Unit.term(dim)
	if term == nil || !term.Active() {
		return q
	}
	if term.unit == newUnit {
		return q
	}
	if !unitsProportional(term.unit, newUnit) {
		return q
	}
	sourceDef, ok := term.unit.Def()
	if !ok {
		return q
	}
	targetDef, ok := newUnit.Def()
	if !ok {
		return q
	}
	ratio := sourceDef.Scale / targetDef.Scale
	factor := math.Pow(ratio, float64(term.exp))

	unit := q.Unit.clone()
	termPtr := unit.term(dim)
	termPtr.unit = newUnit
	return DerivedQuantity{
		Value: q.Value * factor,
		Unit:  unit,
	}
}

// Compatible reports whether other has the same derived dimension as q.
func (q DerivedQuantity) Compatible(other DerivedQuantity) bool {
	if q.Unit == nil || other.Unit == nil {
		return false
	}
	return mustDerivedDim(q.Unit).Equal(mustDerivedDim(other.Unit))
}

// Add returns q plus other in q's unit. Incompatible operands return q unchanged.
func (q DerivedQuantity) Add(other DerivedQuantity) DerivedQuantity {
	if !q.Compatible(other) {
		return q
	}
	if q.Unit == nil || other.Unit == nil {
		return q
	}
	baseSum := q.Value*mustFactorToBase(q.Unit) + other.Value*mustFactorToBase(other.Unit)
	factor := mustFactorToBase(q.Unit)
	if factor == 0 {
		return q
	}
	return DerivedQuantity{
		Value: baseSum / factor,
		Unit:  q.Unit,
	}
}

// Sub returns q minus other in q's unit. Incompatible operands return q unchanged.
func (q DerivedQuantity) Sub(other DerivedQuantity) DerivedQuantity {
	if !q.Compatible(other) {
		return q
	}
	if q.Unit == nil || other.Unit == nil {
		return q
	}
	baseDiff := q.Value*mustFactorToBase(q.Unit) - other.Value*mustFactorToBase(other.Unit)
	factor := mustFactorToBase(q.Unit)
	if factor == 0 {
		return q
	}
	return DerivedQuantity{
		Value: baseDiff / factor,
		Unit:  q.Unit,
	}
}

// Mul returns the product of q and other. The result unit prefers q's unit per dimension.
func (q DerivedQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// Div returns q divided by other. Same-dimension division yields a dimensionless quantity.
// Division by zero returns q unchanged.
func (q DerivedQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

func mulDerivedQuantities(q, other DerivedQuantity) DerivedQuantity {
	if q.Unit == nil || other.Unit == nil {
		return q
	}
	unit := mulDerivedUnits(q.Unit, other.Unit)
	baseResult := q.Value * mustFactorToBase(q.Unit) * other.Value * mustFactorToBase(other.Unit)
	factor := mustFactorToBase(unit)
	if factor == 0 {
		return q
	}
	return DerivedQuantity{
		Value: baseResult / factor,
		Unit:  unit,
	}
}

func divDerivedQuantities(q, other DerivedQuantity) DerivedQuantity {
	if q.Unit == nil || other.Unit == nil {
		return q
	}
	divisor := other.Value * mustFactorToBase(other.Unit)
	if divisor == 0 {
		return q
	}
	unit := divDerivedUnits(q.Unit, other.Unit)
	baseResult := q.Value * mustFactorToBase(q.Unit) / divisor
	factor := mustFactorToBase(unit)
	if factor == 0 {
		return q
	}
	return DerivedQuantity{
		Value: baseResult / factor,
		Unit:  unit,
	}
}

// MulV scales q by v in q's unit.
func (q DerivedQuantity) MulV(v float64) DerivedQuantity {
	return DerivedQuantity{Value: q.Value * v, Unit: q.Unit}
}

// DivV divides q by v in q's unit. Division by zero returns q unchanged.
func (q DerivedQuantity) DivV(v float64) DerivedQuantity {
	if v == 0 {
		return q
	}
	return DerivedQuantity{Value: q.Value / v, Unit: q.Unit}
}
