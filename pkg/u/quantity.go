package u

import "fmt"

// derivedQuantity is implemented by Quantity, the seven SI base-dimension typed Quantity types,
// and DerivedQuantity. It allows Mul and Div across base and derived quantities.
type derivedQuantity interface {
	asDerivedQuantity() (DerivedQuantity, bool)
}

func (q Quantity) asDerivedQuantity() (DerivedQuantity, bool) {
	return quantityAsDerived(q)
}

func (q LengthQuantity) asDerivedQuantity() (DerivedQuantity, bool) {
	return quantityAsDerived(Quantity(q))
}

func (q TimeQuantity) asDerivedQuantity() (DerivedQuantity, bool) {
	return quantityAsDerived(Quantity(q))
}

func (q MassQuantity) asDerivedQuantity() (DerivedQuantity, bool) {
	return quantityAsDerived(Quantity(q))
}

func (q CurrentQuantity) asDerivedQuantity() (DerivedQuantity, bool) {
	return quantityAsDerived(Quantity(q))
}

func (q TemperatureQuantity) asDerivedQuantity() (DerivedQuantity, bool) {
	return quantityAsDerived(Quantity(q))
}

func (q AmountQuantity) asDerivedQuantity() (DerivedQuantity, bool) {
	return quantityAsDerived(Quantity(q))
}

func (q LuminousQuantity) asDerivedQuantity() (DerivedQuantity, bool) {
	return quantityAsDerived(Quantity(q))
}

func (q DerivedQuantity) asDerivedQuantity() (DerivedQuantity, bool) {
	if q.Unit == nil {
		return DerivedQuantity{}, false
	}
	return q, true
}

func mulQuantities(receiver derivedQuantity, other derivedQuantity) DerivedQuantity {
	r, _ := tryMulQuantities(receiver, other)
	return r
}

func divQuantities(receiver derivedQuantity, other derivedQuantity) DerivedQuantity {
	r, _ := tryDivQuantities(receiver, other)
	return r
}

func tryMulQuantities(receiver derivedQuantity, other derivedQuantity) (DerivedQuantity, error) {
	left, ok := receiver.asDerivedQuantity()
	if !ok {
		return DerivedQuantity{}, fmt.Errorf("receiver unit: %w", ErrInvalidUnit)
	}
	right, ok := other.asDerivedQuantity()
	if !ok {
		return left, fmt.Errorf("operand unit: %w", ErrInvalidUnit)
	}
	return tryMulDerivedQuantities(left, right)
}

func tryDivQuantities(receiver derivedQuantity, other derivedQuantity) (DerivedQuantity, error) {
	left, ok := receiver.asDerivedQuantity()
	if !ok {
		return DerivedQuantity{}, fmt.Errorf("receiver unit: %w", ErrInvalidUnit)
	}
	right, ok := other.asDerivedQuantity()
	if !ok {
		return left, fmt.Errorf("operand unit: %w", ErrInvalidUnit)
	}
	return tryDivDerivedQuantities(left, right)
}

// Quantity is a numeric value with a single base-dimension unit.
// For compile-time dimension checks, prefer typed quantities such as LengthQuantity.
//
// Example: Quantity{Value: 3, Unit: Unit(Meter)}
type Quantity struct {
	// Value is the numeric magnitude in Unit.
	Value float64
	// Unit is the measurement unit identifier.
	Unit Unit
}

// String formats q as "value symbol", e.g. "3 m".
func (q Quantity) String() string {
	return q.Format()
}

// Format formats q with optional value and symbol options.
//
// Example:
//
//	Quantity{Value: 1234.5, Unit: Unit(Meter)}.Format(
//		WithPrecision(1),
//		WithNumberDelimiter(NumberDelimiterComma),
//	) // "1,234.5 m"
func (q Quantity) Format(options ...FormatOption) string {
	opt := applyFormatOptions(options...)
	value := formatQuantityValue(q.Value, opt)
	return joinQuantityString(value, q.Unit.symbolWith(opt))
}

// QuantityMustParse parses a quantity and panics on error.
func QuantityMustParse(s string) Quantity {
	q, err := QuantityParse(s)
	if err != nil {
		panic(err)
	}
	return q
}

// Base converts q to the SI (国际单位制) base unit of its dimension.
// On failure, returns q unchanged. Prefer TryBase when errors must be observed.
//
// Example: Quantity{Value: 1, Unit: Unit(Meter.Prefix(Kilo))}.Base() // 1000 m
func (q Quantity) Base() Quantity {
	r, _ := q.TryBase()
	return r
}

// TryBase converts q to the SI base unit of its dimension.
func (q Quantity) TryBase() (Quantity, error) {
	if _, ok := q.Unit.DerivedUnit(); ok {
		return q, fmt.Errorf("derived unit has no single base: %w", ErrInvalidUnit)
	}

	def, ok := q.Unit.Def()
	if !ok {
		return q, fmt.Errorf("unknown unit %q: %w", q.Unit, ErrInvalidUnit)
	}
	baseUnit := def.Dimension.Base()
	if baseUnit == "" {
		return q, fmt.Errorf("dimension %s has no base unit: %w", def.Dimension, ErrInvalidUnit)
	}
	return Quantity{Value: def.ToBase(q.Value), Unit: baseUnit}, nil
}

// By converts q to another unit within the same dimension.
// On failure, returns q unchanged. Prefer TryBy when errors must be observed.
//
// Example: Length(1000, Meter).By(Meter.Prefix(Kilo)) // 1 km
func (q Quantity) By(u Unit) Quantity {
	r, _ := q.TryBy(u)
	return r
}

// TryBy converts q to another unit within the same dimension.
func (q Quantity) TryBy(u Unit) (Quantity, error) {
	targetDef, ok := u.Def()
	if !ok {
		return q, fmt.Errorf("unknown target unit %q: %w", u, ErrInvalidUnit)
	}
	sourceDef, ok := q.Unit.Def()
	if !ok {
		return q, fmt.Errorf("unknown unit %q: %w", q.Unit, ErrInvalidUnit)
	}
	if sourceDef.Dimension != targetDef.Dimension {
		return q, fmt.Errorf("want %s, got %s: %w", sourceDef.Dimension, targetDef.Dimension, ErrDimension)
	}
	baseVal := sourceDef.ToBase(q.Value)
	return Quantity{Value: targetDef.FromBase(baseVal), Unit: u}, nil
}
