package u

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
// other is converted to q's unit via the base unit, then values are added.
func (q Quantity) Add(other Quantity) Quantity {
	if !q.Compatible(other) {
		return q
	}
	if _, ok := q.Unit.Def(); !ok {
		return q
	}
	converted := other.By(q.Unit)
	return Quantity{Value: q.Value + converted.Value, Unit: q.Unit}
}

// Sub returns q minus other in q's unit. Incompatible or invalid operands return q unchanged.
// other is converted to q's unit via the base unit, then values are subtracted.
func (q Quantity) Sub(other Quantity) Quantity {
	if !q.Compatible(other) {
		return q
	}
	if _, ok := q.Unit.Def(); !ok {
		return q
	}
	converted := other.By(q.Unit)
	return Quantity{Value: q.Value - converted.Value, Unit: q.Unit}
}

// Mul returns the product of q and other as a derived quantity.
// Multiplying quantities adds their derived dimensions.
//
// Example: Length(2, Kilometer).Mul(Length(3, Meter)) // 6 km·m → area unit km²
func (q Quantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// Div returns q divided by other as a derived quantity.
// Same-dimension division yields a dimensionless quantity (NoneUnit).
//
// Example: Length(10, Kilometer).Div(Time(2, Hour)) // 5 km/h
func (q Quantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q Quantity) MulV(v float64) Quantity {
	return Quantity{Value: q.Value * v, Unit: q.Unit}
}

// DivV divides q by v in q's unit. Division by zero returns q unchanged.
func (q Quantity) DivV(v float64) Quantity {
	if v == 0 {
		return q
	}
	return Quantity{Value: q.Value / v, Unit: q.Unit}
}
