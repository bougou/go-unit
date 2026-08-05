package u

// Typed unit and quantity wrappers provide compile-time dimension safety (编译期量纲安全).
// Each SI (国际单位制) base dimension has a *Unit alias and a *Quantity struct that delegate
// to Unit and Quantity. Methods mirror Quantity: Add/Sub require the same dimension;
// Mul/Div can cross dimensions and return DerivedQuantity (导出量).
// Try* counterparts return errors when the silent chainable methods would leave q unchanged.

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
// Example: Length(42, Meter.Prefix(Kilo)) // 42 km
func Length(value float64, unit LengthUnit) LengthQuantity {
	return LengthQuantity{Value: value, Unit: Unit(unit)}
}

// Of creates a length quantity with value in u.
//
// Example: Meter.Prefix(Kilo).Of(42) // 42 km
func (u LengthUnit) Of(value float64) LengthQuantity {
	return Length(value, u)
}

// String formats q as "value symbol".
func (q LengthQuantity) String() string {
	return Quantity(q).String()
}

// Format formats q with optional value and symbol options.
func (q LengthQuantity) Format(options ...FormatOption) string {
	return Quantity(q).Format(options...)
}

// Base converts q to the SI base unit of length (meter).
func (q LengthQuantity) Base() LengthQuantity {
	return LengthQuantity(Quantity(q).Base())
}

// TryBase converts q to the SI base unit of length (meter).
func (q LengthQuantity) TryBase() (LengthQuantity, error) {
	r, err := Quantity(q).TryBase()
	return LengthQuantity(r), err
}

// By converts q to another length unit.
//
// Example: Length(1000, Meter).By(Meter.Prefix(Kilo)) // 1 km
func (q LengthQuantity) By(u LengthUnit) LengthQuantity {
	return LengthQuantity(Quantity(q).By(Unit(u)))
}

// TryBy converts q to another length unit.
func (q LengthQuantity) TryBy(u LengthUnit) (LengthQuantity, error) {
	r, err := Quantity(q).TryBy(Unit(u))
	return LengthQuantity(r), err
}

// Prefix converts q to LengthUnit(q.Unit).Prefix(factor).
//
// Example: Length(1000, Meter).Prefix(Kilo) // 1 km
func (q LengthQuantity) Prefix(factor SIPrefix) LengthQuantity {
	return q.By(LengthUnit(q.Unit).Prefix(factor))
}

// Compatible reports whether other has the same dimension (length).
func (q LengthQuantity) Compatible(other LengthQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit. Incompatible operands return q unchanged.
func (q LengthQuantity) Add(other LengthQuantity) LengthQuantity {
	return LengthQuantity(Quantity(q).Add(Quantity(other)))
}

// TryAdd returns q plus other in q's unit.
func (q LengthQuantity) TryAdd(other LengthQuantity) (LengthQuantity, error) {
	r, err := Quantity(q).TryAdd(Quantity(other))
	return LengthQuantity(r), err
}

// Sub returns q minus other in q's unit. Incompatible operands return q unchanged.
func (q LengthQuantity) Sub(other LengthQuantity) LengthQuantity {
	return LengthQuantity(Quantity(q).Sub(Quantity(other)))
}

// TrySub returns q minus other in q's unit.
func (q LengthQuantity) TrySub(other LengthQuantity) (LengthQuantity, error) {
	r, err := Quantity(q).TrySub(Quantity(other))
	return LengthQuantity(r), err
}

// Mul returns the product as a derived quantity.
//
// Example: Length(10, Meter.Prefix(Kilo)).Mul(Length(3, Meter)) // 30 km·m
func (q LengthQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// TryMul returns the product as a derived quantity.
func (q LengthQuantity) TryMul(other derivedQuantity) (DerivedQuantity, error) {
	return tryMulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
//
// Example: Length(10, Meter.Prefix(Kilo)).Div(Time(2, Hour)) // 5 km/h
func (q LengthQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// TryDiv returns the quotient as a derived quantity.
func (q LengthQuantity) TryDiv(other derivedQuantity) (DerivedQuantity, error) {
	return tryDivQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q LengthQuantity) MulV(v float64) LengthQuantity {
	return LengthQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit. Division by zero returns q unchanged.
func (q LengthQuantity) DivV(v float64) LengthQuantity {
	return LengthQuantity(Quantity(q).DivV(v))
}

// TryDivV divides q by v in q's unit.
func (q LengthQuantity) TryDivV(v float64) (LengthQuantity, error) {
	r, err := Quantity(q).TryDivV(v)
	return LengthQuantity(r), err
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

// Of creates a time quantity with value in u.
//
// Example: Hour.Of(2) // 2 h
func (u TimeUnit) Of(value float64) TimeQuantity {
	return Time(value, u)
}

// String formats q as "value symbol".
func (q TimeQuantity) String() string {
	return Quantity(q).String()
}

// Format formats q with optional value and symbol options.
func (q TimeQuantity) Format(options ...FormatOption) string {
	return Quantity(q).Format(options...)
}

// Base converts q to the SI base unit of time (second).
func (q TimeQuantity) Base() TimeQuantity {
	return TimeQuantity(Quantity(q).Base())
}

// TryBase converts q to the SI base unit of time (second).
func (q TimeQuantity) TryBase() (TimeQuantity, error) {
	r, err := Quantity(q).TryBase()
	return TimeQuantity(r), err
}

// By converts q to another time unit.
func (q TimeQuantity) By(u TimeUnit) TimeQuantity {
	return TimeQuantity(Quantity(q).By(Unit(u)))
}

// TryBy converts q to another time unit.
func (q TimeQuantity) TryBy(u TimeUnit) (TimeQuantity, error) {
	r, err := Quantity(q).TryBy(Unit(u))
	return TimeQuantity(r), err
}

// Prefix converts q to TimeUnit(q.Unit).Prefix(factor).
//
// Example: Time(1, Second).Prefix(Milli) // 1000 ms
func (q TimeQuantity) Prefix(factor SIPrefix) TimeQuantity {
	return q.By(TimeUnit(q.Unit).Prefix(factor))
}

// Compatible reports whether other has the same dimension (time).
func (q TimeQuantity) Compatible(other TimeQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit.
func (q TimeQuantity) Add(other TimeQuantity) TimeQuantity {
	return TimeQuantity(Quantity(q).Add(Quantity(other)))
}

// TryAdd returns q plus other in q's unit.
func (q TimeQuantity) TryAdd(other TimeQuantity) (TimeQuantity, error) {
	r, err := Quantity(q).TryAdd(Quantity(other))
	return TimeQuantity(r), err
}

// Sub returns q minus other in q's unit.
func (q TimeQuantity) Sub(other TimeQuantity) TimeQuantity {
	return TimeQuantity(Quantity(q).Sub(Quantity(other)))
}

// TrySub returns q minus other in q's unit.
func (q TimeQuantity) TrySub(other TimeQuantity) (TimeQuantity, error) {
	r, err := Quantity(q).TrySub(Quantity(other))
	return TimeQuantity(r), err
}

// Mul returns the product as a derived quantity.
func (q TimeQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// TryMul returns the product as a derived quantity.
func (q TimeQuantity) TryMul(other derivedQuantity) (DerivedQuantity, error) {
	return tryMulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q TimeQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// TryDiv returns the quotient as a derived quantity.
func (q TimeQuantity) TryDiv(other derivedQuantity) (DerivedQuantity, error) {
	return tryDivQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q TimeQuantity) MulV(v float64) TimeQuantity {
	return TimeQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q TimeQuantity) DivV(v float64) TimeQuantity {
	return TimeQuantity(Quantity(q).DivV(v))
}

// TryDivV divides q by v in q's unit.
func (q TimeQuantity) TryDivV(v float64) (TimeQuantity, error) {
	r, err := Quantity(q).TryDivV(v)
	return TimeQuantity(r), err
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

// Of creates a mass quantity with value in u.
//
// Example: Kilogram.Of(2) // 2 kg
func (u MassUnit) Of(value float64) MassQuantity {
	return Mass(value, u)
}

// String formats q as "value symbol".
func (q MassQuantity) String() string {
	return Quantity(q).String()
}

// Format formats q with optional value and symbol options.
func (q MassQuantity) Format(options ...FormatOption) string {
	return Quantity(q).Format(options...)
}

// Base converts q to the SI base unit of mass (kilogram).
func (q MassQuantity) Base() MassQuantity {
	return MassQuantity(Quantity(q).Base())
}

// TryBase converts q to the SI base unit of mass (kilogram).
func (q MassQuantity) TryBase() (MassQuantity, error) {
	r, err := Quantity(q).TryBase()
	return MassQuantity(r), err
}

// By converts q to another mass unit.
func (q MassQuantity) By(u MassUnit) MassQuantity {
	return MassQuantity(Quantity(q).By(Unit(u)))
}

// TryBy converts q to another mass unit.
func (q MassQuantity) TryBy(u MassUnit) (MassQuantity, error) {
	r, err := Quantity(q).TryBy(Unit(u))
	return MassQuantity(r), err
}

// Prefix converts q to MassUnit(q.Unit).Prefix(factor).
//
// Example: Mass(1, Kilogram).Prefix(Milli) // 1000 g
func (q MassQuantity) Prefix(factor SIPrefix) MassQuantity {
	return q.By(MassUnit(q.Unit).Prefix(factor))
}

// Compatible reports whether other has the same dimension (mass).
func (q MassQuantity) Compatible(other MassQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit.
func (q MassQuantity) Add(other MassQuantity) MassQuantity {
	return MassQuantity(Quantity(q).Add(Quantity(other)))
}

// TryAdd returns q plus other in q's unit.
func (q MassQuantity) TryAdd(other MassQuantity) (MassQuantity, error) {
	r, err := Quantity(q).TryAdd(Quantity(other))
	return MassQuantity(r), err
}

// Sub returns q minus other in q's unit.
func (q MassQuantity) Sub(other MassQuantity) MassQuantity {
	return MassQuantity(Quantity(q).Sub(Quantity(other)))
}

// TrySub returns q minus other in q's unit.
func (q MassQuantity) TrySub(other MassQuantity) (MassQuantity, error) {
	r, err := Quantity(q).TrySub(Quantity(other))
	return MassQuantity(r), err
}

// Mul returns the product as a derived quantity.
//
// Example: Mass(2, Kilogram).Mul(Length(3, Meter)) // 6 kg·m
func (q MassQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// TryMul returns the product as a derived quantity.
func (q MassQuantity) TryMul(other derivedQuantity) (DerivedQuantity, error) {
	return tryMulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q MassQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// TryDiv returns the quotient as a derived quantity.
func (q MassQuantity) TryDiv(other derivedQuantity) (DerivedQuantity, error) {
	return tryDivQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q MassQuantity) MulV(v float64) MassQuantity {
	return MassQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q MassQuantity) DivV(v float64) MassQuantity {
	return MassQuantity(Quantity(q).DivV(v))
}

// TryDivV divides q by v in q's unit.
func (q MassQuantity) TryDivV(v float64) (MassQuantity, error) {
	r, err := Quantity(q).TryDivV(v)
	return MassQuantity(r), err
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

// Of creates a current quantity with value in u.
//
// Example: Ampere.Of(5) // 5 A
func (u CurrentUnit) Of(value float64) CurrentQuantity {
	return Current(value, u)
}

// String formats q as "value symbol".
func (q CurrentQuantity) String() string {
	return Quantity(q).String()
}

// Format formats q with optional value and symbol options.
func (q CurrentQuantity) Format(options ...FormatOption) string {
	return Quantity(q).Format(options...)
}

// Base converts q to the SI base unit of current (ampere).
func (q CurrentQuantity) Base() CurrentQuantity {
	return CurrentQuantity(Quantity(q).Base())
}

// TryBase converts q to the SI base unit of current (ampere).
func (q CurrentQuantity) TryBase() (CurrentQuantity, error) {
	r, err := Quantity(q).TryBase()
	return CurrentQuantity(r), err
}

// By converts q to another current unit.
func (q CurrentQuantity) By(u CurrentUnit) CurrentQuantity {
	return CurrentQuantity(Quantity(q).By(Unit(u)))
}

// TryBy converts q to another current unit.
func (q CurrentQuantity) TryBy(u CurrentUnit) (CurrentQuantity, error) {
	r, err := Quantity(q).TryBy(Unit(u))
	return CurrentQuantity(r), err
}

// Prefix converts q to CurrentUnit(q.Unit).Prefix(factor).
//
// Example: Current(1, Ampere).Prefix(Milli) // 1000 mA
func (q CurrentQuantity) Prefix(factor SIPrefix) CurrentQuantity {
	return q.By(CurrentUnit(q.Unit).Prefix(factor))
}

// Compatible reports whether other has the same dimension (current).
func (q CurrentQuantity) Compatible(other CurrentQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit.
func (q CurrentQuantity) Add(other CurrentQuantity) CurrentQuantity {
	return CurrentQuantity(Quantity(q).Add(Quantity(other)))
}

// TryAdd returns q plus other in q's unit.
func (q CurrentQuantity) TryAdd(other CurrentQuantity) (CurrentQuantity, error) {
	r, err := Quantity(q).TryAdd(Quantity(other))
	return CurrentQuantity(r), err
}

// Sub returns q minus other in q's unit.
func (q CurrentQuantity) Sub(other CurrentQuantity) CurrentQuantity {
	return CurrentQuantity(Quantity(q).Sub(Quantity(other)))
}

// TrySub returns q minus other in q's unit.
func (q CurrentQuantity) TrySub(other CurrentQuantity) (CurrentQuantity, error) {
	r, err := Quantity(q).TrySub(Quantity(other))
	return CurrentQuantity(r), err
}

// Mul returns the product as a derived quantity.
func (q CurrentQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// TryMul returns the product as a derived quantity.
func (q CurrentQuantity) TryMul(other derivedQuantity) (DerivedQuantity, error) {
	return tryMulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q CurrentQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// TryDiv returns the quotient as a derived quantity.
func (q CurrentQuantity) TryDiv(other derivedQuantity) (DerivedQuantity, error) {
	return tryDivQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q CurrentQuantity) MulV(v float64) CurrentQuantity {
	return CurrentQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q CurrentQuantity) DivV(v float64) CurrentQuantity {
	return CurrentQuantity(Quantity(q).DivV(v))
}

// TryDivV divides q by v in q's unit.
func (q CurrentQuantity) TryDivV(v float64) (CurrentQuantity, error) {
	r, err := Quantity(q).TryDivV(v)
	return CurrentQuantity(r), err
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

// Of creates a temperature quantity with value in u.
//
// Example: Celsius.Of(25) // 25 °C
func (u TemperatureUnit) Of(value float64) TemperatureQuantity {
	return Temperature(value, u)
}

// String formats q as "value symbol".
func (q TemperatureQuantity) String() string {
	return Quantity(q).String()
}

// Format formats q with optional value and symbol options.
func (q TemperatureQuantity) Format(options ...FormatOption) string {
	return Quantity(q).Format(options...)
}

// Base converts q to the SI base unit of temperature (kelvin).
func (q TemperatureQuantity) Base() TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).Base())
}

// TryBase converts q to the SI base unit of temperature (kelvin).
func (q TemperatureQuantity) TryBase() (TemperatureQuantity, error) {
	r, err := Quantity(q).TryBase()
	return TemperatureQuantity(r), err
}

// By converts q to another temperature unit.
func (q TemperatureQuantity) By(u TemperatureUnit) TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).By(Unit(u)))
}

// TryBy converts q to another temperature unit.
func (q TemperatureQuantity) TryBy(u TemperatureUnit) (TemperatureQuantity, error) {
	r, err := Quantity(q).TryBy(Unit(u))
	return TemperatureQuantity(r), err
}

// Prefix converts q to TemperatureUnit(q.Unit).Prefix(factor).
// Affine units reject Prefix and therefore panic.
//
// Example: Temperature(1, Kelvin).Prefix(Milli) // 1000 mK
func (q TemperatureQuantity) Prefix(factor SIPrefix) TemperatureQuantity {
	return q.By(TemperatureUnit(q.Unit).Prefix(factor))
}

// Compatible reports whether other has the same dimension (temperature).
func (q TemperatureQuantity) Compatible(other TemperatureQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit. Uses affine conversion for °C/°F.
func (q TemperatureQuantity) Add(other TemperatureQuantity) TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).Add(Quantity(other)))
}

// TryAdd returns q plus other in q's unit.
func (q TemperatureQuantity) TryAdd(other TemperatureQuantity) (TemperatureQuantity, error) {
	r, err := Quantity(q).TryAdd(Quantity(other))
	return TemperatureQuantity(r), err
}

// Sub returns q minus other in q's unit.
func (q TemperatureQuantity) Sub(other TemperatureQuantity) TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).Sub(Quantity(other)))
}

// TrySub returns q minus other in q's unit.
func (q TemperatureQuantity) TrySub(other TemperatureQuantity) (TemperatureQuantity, error) {
	r, err := Quantity(q).TrySub(Quantity(other))
	return TemperatureQuantity(r), err
}

// Mul returns the product as a derived quantity.
func (q TemperatureQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// TryMul returns the product as a derived quantity.
func (q TemperatureQuantity) TryMul(other derivedQuantity) (DerivedQuantity, error) {
	return tryMulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q TemperatureQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// TryDiv returns the quotient as a derived quantity.
func (q TemperatureQuantity) TryDiv(other derivedQuantity) (DerivedQuantity, error) {
	return tryDivQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q TemperatureQuantity) MulV(v float64) TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q TemperatureQuantity) DivV(v float64) TemperatureQuantity {
	return TemperatureQuantity(Quantity(q).DivV(v))
}

// TryDivV divides q by v in q's unit.
func (q TemperatureQuantity) TryDivV(v float64) (TemperatureQuantity, error) {
	r, err := Quantity(q).TryDivV(v)
	return TemperatureQuantity(r), err
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

// Of creates an amount quantity with value in u.
//
// Example: Mole.Of(1) // 1 mol
func (u AmountUnit) Of(value float64) AmountQuantity {
	return Amount(value, u)
}

// String formats q as "value symbol".
func (q AmountQuantity) String() string {
	return Quantity(q).String()
}

// Format formats q with optional value and symbol options.
func (q AmountQuantity) Format(options ...FormatOption) string {
	return Quantity(q).Format(options...)
}

// Base converts q to the SI base unit of amount (mole).
func (q AmountQuantity) Base() AmountQuantity {
	return AmountQuantity(Quantity(q).Base())
}

// TryBase converts q to the SI base unit of amount (mole).
func (q AmountQuantity) TryBase() (AmountQuantity, error) {
	r, err := Quantity(q).TryBase()
	return AmountQuantity(r), err
}

// By converts q to another amount unit.
func (q AmountQuantity) By(u AmountUnit) AmountQuantity {
	return AmountQuantity(Quantity(q).By(Unit(u)))
}

// TryBy converts q to another amount unit.
func (q AmountQuantity) TryBy(u AmountUnit) (AmountQuantity, error) {
	r, err := Quantity(q).TryBy(Unit(u))
	return AmountQuantity(r), err
}

// Prefix converts q to AmountUnit(q.Unit).Prefix(factor).
//
// Example: Amount(1, Mole).Prefix(Milli) // 1000 mmol
func (q AmountQuantity) Prefix(factor SIPrefix) AmountQuantity {
	return q.By(AmountUnit(q.Unit).Prefix(factor))
}

// Compatible reports whether other has the same dimension (amount).
func (q AmountQuantity) Compatible(other AmountQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit.
func (q AmountQuantity) Add(other AmountQuantity) AmountQuantity {
	return AmountQuantity(Quantity(q).Add(Quantity(other)))
}

// TryAdd returns q plus other in q's unit.
func (q AmountQuantity) TryAdd(other AmountQuantity) (AmountQuantity, error) {
	r, err := Quantity(q).TryAdd(Quantity(other))
	return AmountQuantity(r), err
}

// Sub returns q minus other in q's unit.
func (q AmountQuantity) Sub(other AmountQuantity) AmountQuantity {
	return AmountQuantity(Quantity(q).Sub(Quantity(other)))
}

// TrySub returns q minus other in q's unit.
func (q AmountQuantity) TrySub(other AmountQuantity) (AmountQuantity, error) {
	r, err := Quantity(q).TrySub(Quantity(other))
	return AmountQuantity(r), err
}

// Mul returns the product as a derived quantity.
func (q AmountQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// TryMul returns the product as a derived quantity.
func (q AmountQuantity) TryMul(other derivedQuantity) (DerivedQuantity, error) {
	return tryMulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q AmountQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// TryDiv returns the quotient as a derived quantity.
func (q AmountQuantity) TryDiv(other derivedQuantity) (DerivedQuantity, error) {
	return tryDivQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q AmountQuantity) MulV(v float64) AmountQuantity {
	return AmountQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q AmountQuantity) DivV(v float64) AmountQuantity {
	return AmountQuantity(Quantity(q).DivV(v))
}

// TryDivV divides q by v in q's unit.
func (q AmountQuantity) TryDivV(v float64) (AmountQuantity, error) {
	r, err := Quantity(q).TryDivV(v)
	return AmountQuantity(r), err
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

// Of creates a luminous-intensity quantity with value in u.
//
// Example: Candela.Of(100) // 100 cd
func (u LuminousUnit) Of(value float64) LuminousQuantity {
	return Luminous(value, u)
}

// String formats q as "value symbol".
func (q LuminousQuantity) String() string {
	return Quantity(q).String()
}

// Format formats q with optional value and symbol options.
func (q LuminousQuantity) Format(options ...FormatOption) string {
	return Quantity(q).Format(options...)
}

// Base converts q to the SI base unit of luminous intensity (candela).
func (q LuminousQuantity) Base() LuminousQuantity {
	return LuminousQuantity(Quantity(q).Base())
}

// TryBase converts q to the SI base unit of luminous intensity (candela).
func (q LuminousQuantity) TryBase() (LuminousQuantity, error) {
	r, err := Quantity(q).TryBase()
	return LuminousQuantity(r), err
}

// By converts q to another luminous unit.
func (q LuminousQuantity) By(u LuminousUnit) LuminousQuantity {
	return LuminousQuantity(Quantity(q).By(Unit(u)))
}

// TryBy converts q to another luminous unit.
func (q LuminousQuantity) TryBy(u LuminousUnit) (LuminousQuantity, error) {
	r, err := Quantity(q).TryBy(Unit(u))
	return LuminousQuantity(r), err
}

// Prefix converts q to LuminousUnit(q.Unit).Prefix(factor).
//
// Example: Luminous(1, Candela).Prefix(Milli) // 1000 mcd
func (q LuminousQuantity) Prefix(factor SIPrefix) LuminousQuantity {
	return q.By(LuminousUnit(q.Unit).Prefix(factor))
}

// Compatible reports whether other has the same dimension (luminous intensity).
func (q LuminousQuantity) Compatible(other LuminousQuantity) bool {
	return Quantity(q).Compatible(Quantity(other))
}

// Add returns q plus other in q's unit.
func (q LuminousQuantity) Add(other LuminousQuantity) LuminousQuantity {
	return LuminousQuantity(Quantity(q).Add(Quantity(other)))
}

// TryAdd returns q plus other in q's unit.
func (q LuminousQuantity) TryAdd(other LuminousQuantity) (LuminousQuantity, error) {
	r, err := Quantity(q).TryAdd(Quantity(other))
	return LuminousQuantity(r), err
}

// Sub returns q minus other in q's unit.
func (q LuminousQuantity) Sub(other LuminousQuantity) LuminousQuantity {
	return LuminousQuantity(Quantity(q).Sub(Quantity(other)))
}

// TrySub returns q minus other in q's unit.
func (q LuminousQuantity) TrySub(other LuminousQuantity) (LuminousQuantity, error) {
	r, err := Quantity(q).TrySub(Quantity(other))
	return LuminousQuantity(r), err
}

// Mul returns the product as a derived quantity.
func (q LuminousQuantity) Mul(other derivedQuantity) DerivedQuantity {
	return mulQuantities(q, other)
}

// TryMul returns the product as a derived quantity.
func (q LuminousQuantity) TryMul(other derivedQuantity) (DerivedQuantity, error) {
	return tryMulQuantities(q, other)
}

// Div returns the quotient as a derived quantity.
func (q LuminousQuantity) Div(other derivedQuantity) DerivedQuantity {
	return divQuantities(q, other)
}

// TryDiv returns the quotient as a derived quantity.
func (q LuminousQuantity) TryDiv(other derivedQuantity) (DerivedQuantity, error) {
	return tryDivQuantities(q, other)
}

// MulV scales q by v in q's unit.
func (q LuminousQuantity) MulV(v float64) LuminousQuantity {
	return LuminousQuantity(Quantity(q).MulV(v))
}

// DivV divides q by v in q's unit.
func (q LuminousQuantity) DivV(v float64) LuminousQuantity {
	return LuminousQuantity(Quantity(q).DivV(v))
}

// TryDivV divides q by v in q's unit.
func (q LuminousQuantity) TryDivV(v float64) (LuminousQuantity, error) {
	r, err := Quantity(q).TryDivV(v)
	return LuminousQuantity(r), err
}
