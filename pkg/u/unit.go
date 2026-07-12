package u

import "fmt"

// Unit is the stable internal identifier for a registered or derived unit.
// Display symbols come from UnitDef.Symbol or DerivedUnit.Symbol.
//
// Example: Unit(Kilometer).Symbol() // "km"
type Unit string

// unitDef describes a unit and how to convert it to its dimension's SI (国际单位制) base unit.
//
// Conversion uses an affine map (仿射映射) between unit values and base-unit values:
//
//	base = value × Scale + Offset
//	value = (base − Offset) / Scale
//
// Scale carries the proportional (ratio) part of the conversion; Offset carries the
// origin shift. When Offset is zero, conversion is a pure scale change (e.g. km → m).
// When both are non-trivial, conversion is affine (e.g. °C → K: Scale=1, Offset=273.15).
// The base unit of each dimension always has Scale=1 and Offset=0.
type unitDef struct {
	// Dimension is the SI (国际单位制) base dimension this unit belongs to.
	Dimension Dimension

	// Unit is the internal registry identifier (stable key for lookup and storage).
	Unit Unit

	// Symbol is the short display symbol shown to users (e.g. "km", "°C", "mol").
	Symbol string

	// Name is the full English name of the unit (e.g. "kilometer", "celsius").
	Name string

	// Scale is the multiplicative factor in the affine conversion to the base unit.
	// Must be positive. A value of 1 in this unit equals Scale in the base unit
	// when Offset is zero.
	Scale float64

	// Offset is the additive constant applied after scaling when converting to the
	// base unit. Zero for ratio units; non-zero when the unit shares a scale with
	// the base unit but uses a different zero point (e.g. Celsius 摄氏度 vs kelvin 开尔文).
	Offset float64
}

func (d unitDef) isBase() bool {
	base := d.Dimension.Base()
	return base != "" && d.Unit == base
}

// ToBase converts a value in this unit to the base unit: base = value×Scale + Offset.
func (d unitDef) ToBase(value float64) float64 {
	return value*d.Scale + d.Offset
}

// FromBase converts a value in the base unit to this unit: value = (base−Offset)/Scale.
func (d unitDef) FromBase(base float64) float64 {
	return (base - d.Offset) / d.Scale
}

// unitsProportional reports whether values in left and right convert via the dimension
// base unit with a constant ratio: v_right = k × v_left for all v_left.
func unitsProportional(left, right Unit) bool {
	if left == right {
		return true
	}
	leftDef, ok := left.Def()
	if !ok {
		return false
	}
	rightDef, ok := right.Def()
	if !ok || leftDef.Dimension != rightDef.Dimension {
		return false
	}
	const v1, v2 = 1, 3
	r1 := rightDef.FromBase(leftDef.ToBase(v1))
	r2 := rightDef.FromBase(leftDef.ToBase(v2))
	ratio1 := r1 / v1
	ratio2 := r2 / v2
	return ratio1 == ratio2
}

var unitRegistry = map[Unit]*unitDef{}

func registerUnit(def *unitDef) {
	if def.Scale <= 0 {
		panic(fmt.Sprintf("unit %q: Scale must be positive, got %g", def.Unit, def.Scale))
	}
	if _, exists := unitRegistry[def.Unit]; exists {
		panic(fmt.Sprintf("unit %q: duplicate registration", def.Unit))
	}
	if def.isBase() {
		for _, existing := range unitRegistry {
			if existing.Dimension == def.Dimension && existing.isBase() {
				panic(fmt.Sprintf("dimension %s: duplicate base unit %q and %q", def.Dimension, existing.Unit, def.Unit))
			}
		}
	}
	unitRegistry[def.Unit] = def
	registerUnitSymbol(def.Symbol, def.Unit)
}

// validateRegistry checks that every dimension's base unit is registered.
func validateRegistry() error {
	for u, def := range unitRegistry {
		if def.Scale <= 0 {
			return fmt.Errorf("unit %q: Scale must be positive", u)
		}
	}

	for _, dim := range []Dimension{
		DimLength, DimMass, DimTime, DimCurrent, DimTemperature, DimAmount, DimLuminous,
	} {
		base := dim.Base()
		if base == "" {
			continue
		}
		def, ok := unitRegistry[base]
		if !ok {
			return fmt.Errorf("base unit %q for dimension %s is not registered", base, dim)
		}
		if def.Dimension != dim {
			return fmt.Errorf("base unit %q is registered under dimension %s, want %s", base, def.Dimension, dim)
		}
	}
	return nil
}

// Symbol returns the display symbol for u.
//
// Registered base units return UnitDef.Symbol. Derived units delegate to
// DerivedUnit.Symbol and honor options such as WithNamedSymbol and WithExpSign.
//
// Example:
//
//	Unit(Meter).Symbol()                                 // "m"
//	Unit(ForceUnit.Key()).Symbol()                       // "kg·m·s^-2"
//	Unit(ForceUnit.Key()).Symbol(WithNamedSymbol(true))  // "N"
func (u Unit) Symbol(options ...SymbolOption) string {
	if def, ok := u.Def(); ok && def.Symbol != "" {
		return def.Symbol
	}
	if du, ok := u.DerivedUnit(); ok {
		return du.Symbol(options...)
	}
	return string(u)
}

// Def returns the unit definition registered for u.
func (u Unit) Def() (unitDef, bool) {
	def, ok := unitRegistry[u]
	if !ok {
		return unitDef{}, false
	}
	return *def, true
}

// ToBase converts a value in u to the base unit of its dimension.
func (u Unit) ToBase(value float64) (float64, bool) {
	def, ok := u.Def()
	if !ok {
		return value, false
	}
	return def.ToBase(value), true
}

// FromBase converts a value in the base unit to u.
func (u Unit) FromBase(base float64) (float64, bool) {
	def, ok := u.Def()
	if !ok {
		return base, false
	}
	return def.FromBase(base), true
}
