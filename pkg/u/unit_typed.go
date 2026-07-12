package u

// Typed unit and quantity wrappers provide compile-time dimension safety (编译期量纲安全).
// Each SI (国际单位制) base dimension has a *Unit alias and a *Quantity struct that delegate
// to Unit and Quantity. Methods mirror Quantity: Add/Sub require the same dimension;
// Mul/Div can cross dimensions and return DerivedQuantity (导出量).

// LengthUnit is a unit of length (dimension L). See unit_si.go for constants.
type LengthUnit Unit

// LengthQuantity is a length value. Use Length() to construct.
//
// Example:
//
//	d := Length(5, Meter)
//	area := d.Mul(Length(2, Meter)) // DerivedQuantity in m²
type LengthQuantity Quantity

// Length creates a length quantity.
//
// Example: Length(42, Kilometer) // 42 km
func Length(value float64, unit LengthUnit) LengthQuantity {
	return LengthQuantity{Value: value, Unit: Unit(unit)}
}

// String formats q as "value symbol".
func (q LengthQuantity) String() string {
	return Quantity(q).String()
}

// Base converts q to the SI base unit of length (meter).
func (q LengthQuantity) Base() LengthQuantity {
	return LengthQuantity(Quantity(q).Base())
}

// By converts q to another length unit.
//
// Example: Length(1000, Meter).By(Kilometer) // 1 km
func (q LengthQuantity) By(u LengthUnit) LengthQuantity {
	return LengthQuantity(Quantity(q).By(Unit(u)))
}

// Compatible reports whether other has the same dimension (length).
func (q LengthQuantity) Compatible(other LengthQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit. Incompatible operands return q unchanged.
func (q LengthQuantity) Add(other LengthQuantity) LengthQuantity {
	return LengthQuantity(Quantity(q).Add(Quantity(other)))
}

// Sub returns q minus other in q's unit. Incompatible operands return q unchanged.
func (q LengthQuantity) Sub(other LengthQuantity) LengthQuantity {
	return LengthQuantity(Quantity(q).Sub(Quantity(other)))
}

// Mul returns the product as a derived quantity.
//
// Example: Length(10, Kilometer).Mul(Length(3, Meter)) // 30 km·m
func (q LengthQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
//
// Example: Length(10, Kilometer).Div(Time(2, Hour)) // 5 km/h
func (q LengthQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q LengthQuantity) MulV(v float64) LengthQuantity {
	return LengthQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit. Division by zero returns q unchanged.
func (q LengthQuantity) DivV(v float64) LengthQuantity {
	return LengthQuantity(Quantity(q).DivV(v))
}

// TimeUnit is a unit of time (dimension T).
type TimeUnit Unit

// TimeQuantity is a time value. Use Time() to construct.
type TimeQuantity Quantity

// Time creates a time quantity.
//
// Example: Time(2, Hour) // 2 h
func Time(value float64, unit TimeUnit) TimeQuantity {
	return TimeQuantity{Value: value, Unit: Unit(unit)}
}

// String formats q as "value symbol".
func (q TimeQuantity) String() string {
	return Quantity(q).String()
}

// Base converts q to the SI base unit of time (second).
func (q TimeQuantity) Base() TimeQuantity {
	return TimeQuantity(Quantity(q).Base())
}

// By converts q to another time unit.
func (q TimeQuantity) By(u TimeUnit) TimeQuantity {
	return TimeQuantity(Quantity(q).By(Unit(u)))
}

// Compatible reports whether other has the same dimension (time).
func (q TimeQuantity) Compatible(other TimeQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit.
func (q TimeQuantity) Add(other TimeQuantity) TimeQuantity {
	return TimeQuantity(Quantity(q).Add(Quantity(other)))
}

// Sub returns q minus other in q's unit.
func (q TimeQuantity) Sub(other TimeQuantity) TimeQuantity {
	return TimeQuantity(Quantity(q).Sub(Quantity(other)))
}

// Mul returns the product as a derived quantity.
func (q TimeQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q TimeQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q TimeQuantity) MulV(v float64) TimeQuantity {
	return TimeQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q TimeQuantity) DivV(v float64) TimeQuantity {
	return TimeQuantity(Quantity(q).DivV(v))
}

// MassUnit is a unit of mass (dimension M).
type MassUnit Unit

// MassQuantity is a mass value. Use Mass() to construct.
type MassQuantity Quantity

// Mass creates a mass quantity.
//
// Example: Mass(2, Kilogram) // 2 kg
func Mass(value float64, unit MassUnit) MassQuantity {
	return MassQuantity{Value: value, Unit: Unit(unit)}
}

// String formats q as "value symbol".
func (q MassQuantity) String() string {
	return Quantity(q).String()
}

// Base converts q to the SI base unit of mass (kilogram).
func (q MassQuantity) Base() MassQuantity {
	return MassQuantity(Quantity(q).Base())
}

// By converts q to another mass unit.
func (q MassQuantity) By(u MassUnit) MassQuantity {
	return MassQuantity(Quantity(q).By(Unit(u)))
}

// Compatible reports whether other has the same dimension (mass).
func (q MassQuantity) Compatible(other MassQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit.
func (q MassQuantity) Add(other MassQuantity) MassQuantity {
	return MassQuantity(Quantity(q).Add(Quantity(other)))
}

// Sub returns q minus other in q's unit.
func (q MassQuantity) Sub(other MassQuantity) MassQuantity {
	return MassQuantity(Quantity(q).Sub(Quantity(other)))
}

// Mul returns the product as a derived quantity.
//
// Example: Mass(2, Kilogram).Mul(Length(3, Meter)) // 6 kg·m
func (q MassQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q MassQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q MassQuantity) MulV(v float64) MassQuantity {
	return MassQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q MassQuantity) DivV(v float64) MassQuantity {
	return MassQuantity(Quantity(q).DivV(v))
}

// CurrentUnit is a unit of electric current (dimension I).
type CurrentUnit Unit

// CurrentQuantity is a current value. Use Current() to construct.
type CurrentQuantity Quantity

// Current creates a current quantity.
//
// Example: Current(5, Ampere) // 5 A
func Current(value float64, unit CurrentUnit) CurrentQuantity {
	return CurrentQuantity{Value: value, Unit: Unit(unit)}
}

// String formats q as "value symbol".
func (q CurrentQuantity) String() string {
	return Quantity(q).String()
}

// Base converts q to the SI base unit of current (ampere).
func (q CurrentQuantity) Base() CurrentQuantity {
	return CurrentQuantity(Quantity(q).Base())
}

// By converts q to another current unit.
func (q CurrentQuantity) By(u CurrentUnit) CurrentQuantity {
	return CurrentQuantity(Quantity(q).By(Unit(u)))
}

// Compatible reports whether other has the same dimension (current).
func (q CurrentQuantity) Compatible(other CurrentQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit.
func (q CurrentQuantity) Add(other CurrentQuantity) CurrentQuantity {
	return CurrentQuantity(Quantity(q).Add(Quantity(other)))
}

// Sub returns q minus other in q's unit.
func (q CurrentQuantity) Sub(other CurrentQuantity) CurrentQuantity {
	return CurrentQuantity(Quantity(q).Sub(Quantity(other)))
}

// Mul returns the product as a derived quantity.
func (q CurrentQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q CurrentQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q CurrentQuantity) MulV(v float64) CurrentQuantity {
	return CurrentQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q CurrentQuantity) DivV(v float64) CurrentQuantity {
	return CurrentQuantity(Quantity(q).DivV(v))
}

// TemperatureUnit is a unit of thermodynamic temperature (dimension Θ).
type TemperatureUnit Unit

// TemperatureQuantity is a temperature value. Use Temperature() to construct.
type TemperatureQuantity Quantity

// Temperature creates a temperature quantity.
//
// Example: Temperature(25, Celsius) // 25 °C
func Temperature(value float64, unit TemperatureUnit) TemperatureQuantity {
	return TemperatureQuantity{Value: value, Unit: Unit(unit)}
}

// String formats q as "value symbol".
func (q TemperatureQuantity) String() string {
	return Quantity(q).String()
}

// Base converts q to the SI base unit of temperature (kelvin).
func (q TemperatureQuantity) Base() TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).Base())
}

// By converts q to another temperature unit.
func (q TemperatureQuantity) By(u TemperatureUnit) TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).By(Unit(u)))
}

// Compatible reports whether other has the same dimension (temperature).
func (q TemperatureQuantity) Compatible(other TemperatureQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit. Uses affine conversion for °C/°F.
func (q TemperatureQuantity) Add(other TemperatureQuantity) TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).Add(Quantity(other)))
}

// Sub returns q minus other in q's unit.
func (q TemperatureQuantity) Sub(other TemperatureQuantity) TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).Sub(Quantity(other)))
}

// Mul returns the product as a derived quantity.
func (q TemperatureQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q TemperatureQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q TemperatureQuantity) MulV(v float64) TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q TemperatureQuantity) DivV(v float64) TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).DivV(v))
}

// AmountUnit is a unit of amount of substance (dimension N).
type AmountUnit Unit

// AmountQuantity is an amount-of-substance value. Use Amount() to construct.
type AmountQuantity Quantity

// Amount creates an amount quantity.
//
// Example: Amount(1, Mole) // 1 mol
func Amount(value float64, unit AmountUnit) AmountQuantity {
	return AmountQuantity{Value: value, Unit: Unit(unit)}
}

// String formats q as "value symbol".
func (q AmountQuantity) String() string {
	return Quantity(q).String()
}

// Base converts q to the SI base unit of amount (mole).
func (q AmountQuantity) Base() AmountQuantity {
	return AmountQuantity(Quantity(q).Base())
}

// By converts q to another amount unit.
func (q AmountQuantity) By(u AmountUnit) AmountQuantity {
	return AmountQuantity(Quantity(q).By(Unit(u)))
}

// Compatible reports whether other has the same dimension (amount).
func (q AmountQuantity) Compatible(other AmountQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit.
func (q AmountQuantity) Add(other AmountQuantity) AmountQuantity {
	return AmountQuantity(Quantity(q).Add(Quantity(other)))
}

// Sub returns q minus other in q's unit.
func (q AmountQuantity) Sub(other AmountQuantity) AmountQuantity {
	return AmountQuantity(Quantity(q).Sub(Quantity(other)))
}

// Mul returns the product as a derived quantity.
func (q AmountQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q AmountQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q AmountQuantity) MulV(v float64) AmountQuantity {
	return AmountQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q AmountQuantity) DivV(v float64) AmountQuantity {
	return AmountQuantity(Quantity(q).DivV(v))
}

// LuminousUnit is a unit of luminous intensity (dimension J).
type LuminousUnit Unit

// LuminousQuantity is a luminous-intensity value. Use Luminous() to construct.
type LuminousQuantity Quantity

// Luminous creates a luminous-intensity quantity.
//
// Example: Luminous(100, Candela) // 100 cd
func Luminous(value float64, unit LuminousUnit) LuminousQuantity {
	return LuminousQuantity{Value: value, Unit: Unit(unit)}
}

// String formats q as "value symbol".
func (q LuminousQuantity) String() string {
	return Quantity(q).String()
}

// Base converts q to the SI base unit of luminous intensity (candela).
func (q LuminousQuantity) Base() LuminousQuantity {
	return LuminousQuantity(Quantity(q).Base())
}

// By converts q to another luminous unit.
func (q LuminousQuantity) By(u LuminousUnit) LuminousQuantity {
	return LuminousQuantity(Quantity(q).By(Unit(u)))
}

// Compatible reports whether other has the same dimension (luminous intensity).
func (q LuminousQuantity) Compatible(other LuminousQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit.
func (q LuminousQuantity) Add(other LuminousQuantity) LuminousQuantity {
	return LuminousQuantity(Quantity(q).Add(Quantity(other)))
}

// Sub returns q minus other in q's unit.
func (q LuminousQuantity) Sub(other LuminousQuantity) LuminousQuantity {
	return LuminousQuantity(Quantity(q).Sub(Quantity(other)))
}

// Mul returns the product as a derived quantity.
func (q LuminousQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q LuminousQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q LuminousQuantity) MulV(v float64) LuminousQuantity {
	return LuminousQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q LuminousQuantity) DivV(v float64) LuminousQuantity {
	return LuminousQuantity(Quantity(q).DivV(v))
}
