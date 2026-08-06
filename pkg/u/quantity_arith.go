package u

import "fmt"

// Compatible reports whether other has the same base dimension as q.
func (q Quantity) Compatible(other Quantity) bool {
	src, ok := q.Unit.Def()
	if !ok {
		return false
	}
	dst, ok := other.Unit.Def()
	if !ok {
		return false
	}
	return src.Dimension == dst.Dimension
}

// Add returns q plus other in q's unit. Incompatible or invalid operands return q unchanged.
// Prefer TryAdd when errors must be observed.
// other is converted to q's unit via the base unit, then values are added.
func (q Quantity) Add(other Quantity) Quantity {
	r, _ := q.TryAdd(other)
	return r
}

// TryAdd returns q plus other in q's unit.
func (q Quantity) TryAdd(other Quantity) (Quantity, error) {
	if !q.Compatible(other) {
		return q, fmt.Errorf("add: %w", ErrDimension)
	}
	if _, ok := q.Unit.Def(); !ok {
		return q, fmt.Errorf("unknown unit %q: %w", q.Unit, ErrInvalidUnit)
	}
	converted, err := other.TryBy(q.Unit)
	if err != nil {
		return q, err
	}
	return Quantity{Value: q.Value + converted.Value, Unit: q.Unit}, nil
}

// Sub returns q minus other in q's unit. Incompatible or invalid operands return q unchanged.
// Prefer TrySub when errors must be observed.
// other is converted to q's unit via the base unit, then values are subtracted.
func (q Quantity) Sub(other Quantity) Quantity {
	r, _ := q.TrySub(other)
	return r
}

// TrySub returns q minus other in q's unit.
func (q Quantity) TrySub(other Quantity) (Quantity, error) {
	if !q.Compatible(other) {
		return q, fmt.Errorf("sub: %w", ErrDimension)
	}
	if _, ok := q.Unit.Def(); !ok {
		return q, fmt.Errorf("unknown unit %q: %w", q.Unit, ErrInvalidUnit)
	}
	converted, err := other.TryBy(q.Unit)
	if err != nil {
		return q, err
	}
	return Quantity{Value: q.Value - converted.Value, Unit: q.Unit}, nil
}

// Mul returns the product of q and other as a derived quantity.
// Multiplying quantities adds their derived dimensions.
// On failure, returns a zero or unchanged-style result. Prefer TryMul when errors must be observed.
//
// Example: Length(2, Meter.Prefix(prefix.Kilo)).Mul(Length(3, Meter)) // 6 km·m → area unit km²
func (q Quantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// TryMul returns the product of q and other as a derived quantity.
func (q Quantity) TryMul(other derivedQuantity) (DerivedQuantity, error) {
	return tryMulQuantities(q, other)
}

// Div returns q divided by other as a derived quantity.
// Same-dimension division yields a dimensionless quantity (NoneUnit).
// On failure, returns q unchanged as a derived quantity. Prefer TryDiv when errors must be observed.
//
// Example: Length(10, Meter.Prefix(prefix.Kilo)).Div(Time(2, Hour)) // 5 km/h
func (q Quantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// TryDiv returns q divided by other as a derived quantity.
func (q Quantity) TryDiv(other derivedQuantity) (DerivedQuantity, error) {
	return tryDivQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q Quantity) MulV(v float64) Quantity {
	return Quantity{Value: q.Value * v, Unit: q.Unit}
}

// DivV divides q by v in q's unit. Division by zero returns q unchanged.
// Prefer TryDivV when errors must be observed.
func (q Quantity) DivV(v float64) Quantity {
	r, _ := q.TryDivV(v)
	return r
}

// TryDivV divides q by v in q's unit.
func (q Quantity) TryDivV(v float64) (Quantity, error) {
	if v == 0 {
		return q, ErrDivByZero
	}
	return Quantity{Value: q.Value / v, Unit: q.Unit}, nil
}
