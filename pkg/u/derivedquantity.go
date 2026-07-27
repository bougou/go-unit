package u

import (
	"fmt"
	"math"
)

// DerivedQuantity (导出量) is a numeric value with a compound derived unit.
//
// Example:
//
//	unit := NewDerivedUnit().Length(Meter.Prefix(Kilo), 1).Time(Hour, -1)
//	speed := NewDerivedQuantity(60, unit) // 60 km/h
type DerivedQuantity struct {
	// Value is the numeric magnitude in Unit.
	Value float64
	// Unit is the compound derived unit. Nil means dimensionless display.
	Unit *DerivedUnit
}

// NewDerivedQuantity creates a derived quantity, e.g. 10 km·h⁻¹.
// For a fluent call on a known unit, use DerivedUnit.Of.
//
// Example: NewDerivedQuantity(60, speedUnit) // 60 km/h
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

// String formats q as "value unit" using superscript exponents.
// Prefers the SI special name when one is set (e.g. "N").
//
// Example: NewDerivedQuantity(10, Newton).String() // "10 N"
func (q DerivedQuantity) String() string {
	opt := defaultFormatOption
	opt.expSign = ExpSignSup
	return q.formatWith(opt)
}

// Format formats q with optional FormatOption values.
// Defaults match Symbol (special name when set, caret exponents for compound form).
// Use WithCompoundSymbol(true) for the base-unit expression, and
// WithExpSign(ExpSignSup) for typographic exponents.
//
// Example:
//
//	NewDerivedQuantity(1234.5, Newton).Format(
//		WithPrecision(1),
//		WithNumberDelimiter(NumberDelimiterComma),
//	) // "1,234.5 N"
func (q DerivedQuantity) Format(options ...FormatOption) string {
	return q.formatWith(applyFormatOptions(options...))
}

func (q DerivedQuantity) formatWith(opt formatOption) string {
	value := formatQuantityValue(q.Value, opt)
	if q.Unit == nil {
		return value
	}
	return joinQuantityString(value, q.Unit.symbolWith(opt))
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

// Prefix converts q to the same unit scaled by an SI decimal prefix.
//
// Named derived units (e.g. Ohm) use DerivedUnit.Prefix. Unnamed units with
// exactly one active base dimension prefix that base unit (e.g. A → mA after
// Sqrt of ampere-squared). Multi-term unnamed units and nil are unchanged.
//
// Example: Ohm.Of(2e6).Prefix(Mega) // 2 MΩ
// Example: Watt.Of(P).Div(Ohm.Of(R)).Sqrt().Prefix(Milli) // current in mA
func (q DerivedQuantity) Prefix(factor SIPrefix) DerivedQuantity {
	if q.Unit == nil {
		return q
	}
	if q.Unit.specialSymbol != "" {
		return q.By(q.Unit.Prefix(factor))
	}
	terms := q.Unit.terms()
	if len(terms) != 1 {
		return q
	}
	def, ok := terms[0].unit.Def()
	if !ok {
		return q
	}
	prefixed, ok := prefixBaseDimensionUnit(terms[0].unit, def.Dimension, factor)
	if !ok {
		return q
	}
	return q.byDimension(def.Dimension, prefixed)
}

// prefixBaseDimensionUnit applies an SI prefix to a single base-dimension unit.
func prefixBaseDimensionUnit(unit Unit, dim Dimension, factor SIPrefix) (Unit, bool) {
	switch dim {
	case DimLength:
		return Unit(LengthUnit(unit).Prefix(factor)), true
	case DimMass:
		return Unit(MassUnit(unit).Prefix(factor)), true
	case DimTime:
		return Unit(TimeUnit(unit).Prefix(factor)), true
	case DimCurrent:
		return Unit(CurrentUnit(unit).Prefix(factor)), true
	case DimTemperature:
		return Unit(TemperatureUnit(unit).Prefix(factor)), true
	case DimAmount:
		return Unit(AmountUnit(unit).Prefix(factor)), true
	case DimLuminous:
		return Unit(LuminousUnit(unit).Prefix(factor)), true
	default:
		return "", false
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
// Same dimension alone is not enough for arithmetic: each shared base dimension
// must also use proportionally related units (see Add/Sub/Mul/Div).
func (q DerivedQuantity) Compatible(other DerivedQuantity) bool {
	if q.Unit == nil || other.Unit == nil {
		return false
	}
	return mustDerivedDim(q.Unit).Equal(mustDerivedDim(other.Unit))
}

// Add returns q plus other in q's unit.
// Requires the same derived dimension and proportionally related units on every
// base dimension (affine pairs such as °C/K are rejected). Otherwise returns q unchanged.
func (q DerivedQuantity) Add(other DerivedQuantity) DerivedQuantity {
	if !q.Compatible(other) || !derivedUnitsConvertible(q.Unit, other.Unit) {
		return q
	}
	converted := other.By(q.Unit)
	return DerivedQuantity{Value: q.Value + converted.Value, Unit: q.Unit}
}

// Sub returns q minus other in q's unit.
// Same requirements as Add: equal derived dimension and proportional units per
// base dimension. Otherwise returns q unchanged.
func (q DerivedQuantity) Sub(other DerivedQuantity) DerivedQuantity {
	if !q.Compatible(other) || !derivedUnitsConvertible(q.Unit, other.Unit) {
		return q
	}
	converted := other.By(q.Unit)
	return DerivedQuantity{Value: q.Value - converted.Value, Unit: q.Unit}
}

// Mul returns the product of q and other. The result unit prefers q's unit per dimension.
// Every base dimension present in both operands must use proportionally related units;
// otherwise returns q unchanged.
func (q DerivedQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// Div returns q divided by other. Same-dimension division yields a dimensionless quantity.
// Overlapping base dimensions must use proportionally related units. Division by zero
// or non-proportional overlap returns q unchanged.
func (q DerivedQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// Sqrt returns the square root of q (Root(2)).
// The result is unchanged when any base exponent is odd, the value is negative,
// or the unit is nil.
//
// Example: Watt.Of(20).Div(Ohm.Of(5000)).Sqrt() // √(P/R) → current
func (q DerivedQuantity) Sqrt() DerivedQuantity {
	return q.Root(2)
}

// Root returns the nth root of q: value and every base-dimension exponent are scaled by 1/n.
// n must be >= 2. Even roots of negative values, and exponents not divisible by n,
// return q unchanged (same silent-failure style as Add/Div).
//
// Example: area.Root(2) // length when area has dimension L²
func (q DerivedQuantity) Root(n int) DerivedQuantity {
	if q.Unit == nil || n < 2 {
		return q
	}
	if q.Value < 0 && n%2 == 0 {
		return q
	}
	unit, ok := rootDerivedUnit(q.Unit, n)
	if !ok {
		return q
	}
	base := q.Value * mustFactorToBase(q.Unit)
	rootBase := nthRoot(base, n)
	factor := mustFactorToBase(unit)
	if factor == 0 {
		return q
	}
	return DerivedQuantity{Value: rootBase / factor, Unit: unit}
}

func nthRoot(v float64, n int) float64 {
	if n == 2 {
		return math.Sqrt(v)
	}
	if v < 0 {
		return -math.Pow(-v, 1/float64(n))
	}
	return math.Pow(v, 1/float64(n))
}

func mulDerivedQuantities(q, other DerivedQuantity) DerivedQuantity {
	if q.Unit == nil || other.Unit == nil {
		return q
	}
	if !derivedUnitsOverlapProportional(q.Unit, other.Unit) {
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
	if !derivedUnitsOverlapProportional(q.Unit, other.Unit) {
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
