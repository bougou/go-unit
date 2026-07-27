package u

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
	left, ok := receiver.asDerivedQuantity()
	if !ok {
		return DerivedQuantity{}
	}
	right, ok := other.asDerivedQuantity()
	if !ok {
		return left
	}
	return mulDerivedQuantities(left, right)
}

func divQuantities(receiver derivedQuantity, other derivedQuantity) DerivedQuantity {
	left, ok := receiver.asDerivedQuantity()
	if !ok {
		return DerivedQuantity{}
	}
	right, ok := other.asDerivedQuantity()
	if !ok {
		return left
	}
	return divDerivedQuantities(left, right)
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
//
// Example: Quantity{Value: 1, Unit: Unit(Meter.Prefix(Kilo))}.Base() // 1000 m
func (q Quantity) Base() Quantity {
	if _, ok := q.Unit.DerivedUnit(); ok {
		return q
	}

	def, ok := q.Unit.Def()
	if !ok {
		return q
	}
	baseUnit := def.Dimension.Base()
	if baseUnit == "" {
		return q
	}
	return Quantity{Value: def.ToBase(q.Value), Unit: baseUnit}
}

// By converts q to another unit within the same dimension.
//
// Example: Length(1000, Meter).By(Meter.Prefix(Kilo)) // 1 km
func (q Quantity) By(u Unit) Quantity {
	targetDef, ok := u.Def()
	if !ok {
		return q
	}
	sourceDef, ok := q.Unit.Def()
	if !ok || sourceDef.Dimension != targetDef.Dimension {
		return q
	}
	baseVal := sourceDef.ToBase(q.Value)
	return Quantity{Value: targetDef.FromBase(baseVal), Unit: u}
}
