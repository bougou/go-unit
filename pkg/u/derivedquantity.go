package u

import (
	"fmt"
	"math"

	"github.com/bougou/go-unit/pkg/prefix"
)

// DerivedQuantity (导出量) is a numeric value with a compound derived unit.
//
// Example:
//
//	unit := NewDerivedUnit().Length(Meter.Prefix(prefix.Kilo), 1).Time(Hour, -1)
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

// Base converts the quantity to SI base units for each dimension.
func (q DerivedQuantity) Base() DerivedQuantity {
	return DerivedQuantity{
		Value: q.Value * mustFactorToBase(q.Unit),
		Unit:  q.Unit.Base(),
	}
}

// By converts q to target when both share the same derived dimension and every
// dimension whose unit changes converts proportionally via the dimension base unit.
// Otherwise q is returned unchanged. Prefer TryBy when errors must be observed.
//
// Example: speed in km/h converted to m/s via By on the SI unit
func (q DerivedQuantity) By(target *DerivedUnit) DerivedQuantity {
	r, _ := q.TryBy(target)
	return r
}

// TryBy converts q to target when both share the same derived dimension and every
// dimension whose unit changes converts proportionally via the dimension base unit.
func (q DerivedQuantity) TryBy(target *DerivedUnit) (DerivedQuantity, error) {
	if q.Unit == nil || target == nil {
		return q, fmt.Errorf("by: %w", ErrInvalidUnit)
	}
	if !mustDerivedDim(q.Unit).Equal(mustDerivedDim(target)) {
		return q, fmt.Errorf("by: %w", ErrDimension)
	}
	if !derivedUnitsConvertible(q.Unit, target) {
		return q, fmt.Errorf("by: %w", ErrIncompatible)
	}
	return DerivedQuantity{
		Value: q.Value * derivedUnitConversionFactor(q.Unit, target),
		Unit:  target,
	}, nil
}

// Prefix converts q to the same unit scaled by an SI decimal prefix.
//
// Named derived units (e.g. Ohm) use DerivedUnit.Prefix. Unnamed units with
// exactly one active base dimension prefix that base unit (e.g. A → mA after
// Sqrt of ampere-squared). Multi-term unnamed units and nil are unchanged.
// Prefer TryPrefix when errors must be observed.
//
// Example: Ohm.Of(2e6).Prefix(prefix.Mega) // 2 MΩ
// Example: Watt.Of(P).Div(Ohm.Of(R)).Sqrt().Prefix(prefix.Milli) // current in mA
func (q DerivedQuantity) Prefix(factor prefix.SIPrefix) DerivedQuantity {
	r, _ := q.TryPrefix(factor)
	return r
}

// TryPrefix converts q to the same unit scaled by an SI decimal prefix.
func (q DerivedQuantity) TryPrefix(factor prefix.SIPrefix) (DerivedQuantity, error) {
	if q.Unit == nil {
		return q, fmt.Errorf("prefix: %w", ErrInvalidUnit)
	}
	if q.Unit.namedSymbol != "" {
		return q.TryBy(q.Unit.Prefix(factor))
	}
	terms := q.Unit.terms()
	if len(terms) != 1 {
		return q, fmt.Errorf("prefix requires a named unit or single-dimension unit: %w", ErrInvalidUnit)
	}
	def, ok := terms[0].unit.Def()
	if !ok {
		return q, fmt.Errorf("unknown unit %q: %w", terms[0].unit, ErrInvalidUnit)
	}
	prefixed, ok := prefixBaseDimensionUnit(terms[0].unit, def.Dimension, factor)
	if !ok {
		return q, fmt.Errorf("prefix unsupported for dimension %s: %w", def.Dimension, ErrInvalidUnit)
	}
	return q.tryByDimension(def.Dimension, prefixed)
}

// prefixBaseDimensionUnit applies an SI prefix to a single base-dimension unit.
func prefixBaseDimensionUnit(unit Unit, dim Dimension, factor prefix.SIPrefix) (Unit, bool) {
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
// Prefer TryByLength when errors must be observed.
func (q DerivedQuantity) ByLength(u LengthUnit) DerivedQuantity {
	r, _ := q.TryByLength(u)
	return r
}

// TryByLength converts only the length term of q's unit to u.
func (q DerivedQuantity) TryByLength(u LengthUnit) (DerivedQuantity, error) {
	return q.tryByDimension(DimLength, Unit(u))
}

// ByMass converts only the mass term of q's unit to u.
// Prefer TryByMass when errors must be observed.
func (q DerivedQuantity) ByMass(u MassUnit) DerivedQuantity {
	r, _ := q.TryByMass(u)
	return r
}

// TryByMass converts only the mass term of q's unit to u.
func (q DerivedQuantity) TryByMass(u MassUnit) (DerivedQuantity, error) {
	return q.tryByDimension(DimMass, Unit(u))
}

// ByTime converts only the time term of q's unit to u.
// Prefer TryByTime when errors must be observed.
func (q DerivedQuantity) ByTime(u TimeUnit) DerivedQuantity {
	r, _ := q.TryByTime(u)
	return r
}

// TryByTime converts only the time term of q's unit to u.
func (q DerivedQuantity) TryByTime(u TimeUnit) (DerivedQuantity, error) {
	return q.tryByDimension(DimTime, Unit(u))
}

// ByCurrent converts only the current term of q's unit to u.
// Prefer TryByCurrent when errors must be observed.
func (q DerivedQuantity) ByCurrent(u CurrentUnit) DerivedQuantity {
	r, _ := q.TryByCurrent(u)
	return r
}

// TryByCurrent converts only the current term of q's unit to u.
func (q DerivedQuantity) TryByCurrent(u CurrentUnit) (DerivedQuantity, error) {
	return q.tryByDimension(DimCurrent, Unit(u))
}

// ByTemperature converts only the temperature term of q's unit to u.
// Prefer TryByTemperature when errors must be observed.
func (q DerivedQuantity) ByTemperature(u TemperatureUnit) DerivedQuantity {
	r, _ := q.TryByTemperature(u)
	return r
}

// TryByTemperature converts only the temperature term of q's unit to u.
func (q DerivedQuantity) TryByTemperature(u TemperatureUnit) (DerivedQuantity, error) {
	return q.tryByDimension(DimTemperature, Unit(u))
}

// ByAmount converts only the amount-of-substance term of q's unit to u.
// Prefer TryByAmount when errors must be observed.
func (q DerivedQuantity) ByAmount(u AmountUnit) DerivedQuantity {
	r, _ := q.TryByAmount(u)
	return r
}

// TryByAmount converts only the amount-of-substance term of q's unit to u.
func (q DerivedQuantity) TryByAmount(u AmountUnit) (DerivedQuantity, error) {
	return q.tryByDimension(DimAmount, Unit(u))
}

// ByLuminous converts only the luminous-intensity term of q's unit to u.
// Prefer TryByLuminous when errors must be observed.
func (q DerivedQuantity) ByLuminous(u LuminousUnit) DerivedQuantity {
	r, _ := q.TryByLuminous(u)
	return r
}

// TryByLuminous converts only the luminous-intensity term of q's unit to u.
func (q DerivedQuantity) TryByLuminous(u LuminousUnit) (DerivedQuantity, error) {
	return q.tryByDimension(DimLuminous, Unit(u))
}

func (q DerivedQuantity) tryByDimension(dim Dimension, newUnit Unit) (DerivedQuantity, error) {
	if q.Unit == nil {
		return q, fmt.Errorf("by %s: %w", dim, ErrInvalidUnit)
	}
	term := q.Unit.term(dim)
	if term == nil || !term.Active() {
		return q, fmt.Errorf("unit has no %s term: %w", dim, ErrDimension)
	}
	if term.unit == newUnit {
		return q, nil
	}
	if !unitsProportional(term.unit, newUnit) {
		return q, fmt.Errorf("by %s: %w", dim, ErrIncompatible)
	}
	sourceDef, ok := term.unit.Def()
	if !ok {
		return q, fmt.Errorf("unknown unit %q: %w", term.unit, ErrInvalidUnit)
	}
	targetDef, ok := newUnit.Def()
	if !ok {
		return q, fmt.Errorf("unknown target unit %q: %w", newUnit, ErrInvalidUnit)
	}
	ratio := sourceDef.Scale / targetDef.Scale
	factor := math.Pow(ratio, float64(term.exp))

	unit := q.Unit.clone()
	termPtr := unit.term(dim)
	termPtr.unit = newUnit
	return DerivedQuantity{
		Value: q.Value * factor,
		Unit:  unit,
	}, nil
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
// Prefer TryAdd when errors must be observed.
func (q DerivedQuantity) Add(other DerivedQuantity) DerivedQuantity {
	r, _ := q.TryAdd(other)
	return r
}

// TryAdd returns q plus other in q's unit.
func (q DerivedQuantity) TryAdd(other DerivedQuantity) (DerivedQuantity, error) {
	if !q.Compatible(other) {
		return q, fmt.Errorf("add: %w", ErrDimension)
	}
	if !derivedUnitsConvertible(q.Unit, other.Unit) {
		return q, fmt.Errorf("add: %w", ErrIncompatible)
	}
	converted, err := other.TryBy(q.Unit)
	if err != nil {
		return q, err
	}
	return DerivedQuantity{Value: q.Value + converted.Value, Unit: q.Unit}, nil
}

// Sub returns q minus other in q's unit.
// Same requirements as Add: equal derived dimension and proportional units per
// base dimension. Otherwise returns q unchanged.
// Prefer TrySub when errors must be observed.
func (q DerivedQuantity) Sub(other DerivedQuantity) DerivedQuantity {
	r, _ := q.TrySub(other)
	return r
}

// TrySub returns q minus other in q's unit.
func (q DerivedQuantity) TrySub(other DerivedQuantity) (DerivedQuantity, error) {
	if !q.Compatible(other) {
		return q, fmt.Errorf("sub: %w", ErrDimension)
	}
	if !derivedUnitsConvertible(q.Unit, other.Unit) {
		return q, fmt.Errorf("sub: %w", ErrIncompatible)
	}
	converted, err := other.TryBy(q.Unit)
	if err != nil {
		return q, err
	}
	return DerivedQuantity{Value: q.Value - converted.Value, Unit: q.Unit}, nil
}

// Mul returns the product of q and other. The result unit prefers q's unit per dimension.
// Every base dimension present in both operands must use proportionally related units;
// otherwise returns q unchanged. Prefer TryMul when errors must be observed.
func (q DerivedQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// TryMul returns the product of q and other.
func (q DerivedQuantity) TryMul(other derivedQuantity) (DerivedQuantity, error) {
	return tryMulQuantities(q, other)
}

// Div returns q divided by other. Same-dimension division yields a dimensionless quantity.
// Overlapping base dimensions must use proportionally related units. Division by zero
// or non-proportional overlap returns q unchanged. Prefer TryDiv when errors must be observed.
func (q DerivedQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// TryDiv returns q divided by other.
func (q DerivedQuantity) TryDiv(other derivedQuantity) (DerivedQuantity, error) {
	return tryDivQuantities(q, other)
}

// Sqrt returns the square root of q (Root(2)).
// The result is unchanged when any base exponent is odd, the value is negative,
// or the unit is nil. Prefer TrySqrt when errors must be observed.
//
// Example: Watt.Of(20).Div(Ohm.Of(5000)).Sqrt() // √(P/R) → current
func (q DerivedQuantity) Sqrt() DerivedQuantity {
	return q.Root(2)
}

// TrySqrt returns the square root of q (TryRoot(2)).
func (q DerivedQuantity) TrySqrt() (DerivedQuantity, error) {
	return q.TryRoot(2)
}

// Root returns the nth root of q: value and every base-dimension exponent are scaled by 1/n.
// n must be >= 2. Even roots of negative values, and exponents not divisible by n,
// return q unchanged (same silent-failure style as Add/Div).
// Prefer TryRoot when errors must be observed.
//
// Example: area.Root(2) // length when area has dimension L²
func (q DerivedQuantity) Root(n int) DerivedQuantity {
	r, _ := q.TryRoot(n)
	return r
}

// TryRoot returns the nth root of q.
func (q DerivedQuantity) TryRoot(n int) (DerivedQuantity, error) {
	if q.Unit == nil {
		return q, fmt.Errorf("root: %w", ErrInvalidUnit)
	}
	if n < 2 {
		return q, fmt.Errorf("root degree %d: %w", n, ErrRoot)
	}
	if q.Value < 0 && n%2 == 0 {
		return q, fmt.Errorf("even root of negative value: %w", ErrRoot)
	}
	unit, ok := rootDerivedUnit(q.Unit, n)
	if !ok {
		return q, fmt.Errorf("exponents not divisible by %d: %w", n, ErrRoot)
	}
	base := q.Value * mustFactorToBase(q.Unit)
	rootBase := nthRoot(base, n)
	factor := mustFactorToBase(unit)
	if factor == 0 {
		return q, fmt.Errorf("root unit factor: %w", ErrInvalidUnit)
	}
	return DerivedQuantity{Value: rootBase / factor, Unit: unit}, nil
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

func tryMulDerivedQuantities(q, other DerivedQuantity) (DerivedQuantity, error) {
	if q.Unit == nil || other.Unit == nil {
		return q, fmt.Errorf("mul: %w", ErrInvalidUnit)
	}
	if !derivedUnitsOverlapProportional(q.Unit, other.Unit) {
		return q, fmt.Errorf("mul: %w", ErrIncompatible)
	}
	unit := mulDerivedUnits(q.Unit, other.Unit)
	baseResult := q.Value * mustFactorToBase(q.Unit) * other.Value * mustFactorToBase(other.Unit)
	factor := mustFactorToBase(unit)
	if factor == 0 {
		return q, fmt.Errorf("mul unit factor: %w", ErrInvalidUnit)
	}
	return DerivedQuantity{
		Value: baseResult / factor,
		Unit:  unit,
	}, nil
}

func tryDivDerivedQuantities(q, other DerivedQuantity) (DerivedQuantity, error) {
	if q.Unit == nil || other.Unit == nil {
		return q, fmt.Errorf("div: %w", ErrInvalidUnit)
	}
	if !derivedUnitsOverlapProportional(q.Unit, other.Unit) {
		return q, fmt.Errorf("div: %w", ErrIncompatible)
	}
	divisor := other.Value * mustFactorToBase(other.Unit)
	if divisor == 0 {
		return q, ErrDivByZero
	}
	unit := divDerivedUnits(q.Unit, other.Unit)
	baseResult := q.Value * mustFactorToBase(q.Unit) / divisor
	factor := mustFactorToBase(unit)
	if factor == 0 {
		return q, fmt.Errorf("div unit factor: %w", ErrInvalidUnit)
	}
	return DerivedQuantity{
		Value: baseResult / factor,
		Unit:  unit,
	}, nil
}

// MulV scales q by v in q's unit.
func (q DerivedQuantity) MulV(v float64) DerivedQuantity {
	return DerivedQuantity{Value: q.Value * v, Unit: q.Unit}
}

// DivV divides q by v in q's unit. Division by zero returns q unchanged.
// Prefer TryDivV when errors must be observed.
func (q DerivedQuantity) DivV(v float64) DerivedQuantity {
	r, _ := q.TryDivV(v)
	return r
}

// TryDivV divides q by v in q's unit.
func (q DerivedQuantity) TryDivV(v float64) (DerivedQuantity, error) {
	if v == 0 {
		return q, ErrDivByZero
	}
	return DerivedQuantity{Value: q.Value / v, Unit: q.Unit}, nil
}
