package u

func quantityAsDerived(q Quantity) (DerivedQuantity, bool) {
	du, ok := derivedUnitFromBaseUnit(q.Unit)
	if !ok {
		return DerivedQuantity{}, false
	}
	return DerivedQuantity{Value: q.Value, Unit: du}, true
}
