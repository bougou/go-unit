package u

// Dimension (基本量纲) represents one of the seven independent SI (国际单位制) base dimensions.
type Dimension uint8

const (
	DimLength      Dimension = iota // L, length (长度)
	DimMass                         // M, mass (质量)
	DimTime                         // T, time (时间)
	DimCurrent                      // I, electric current (电流)
	DimTemperature                  // Θ, thermodynamic temperature (热力学温度)
	DimAmount                       // N, amount of substance (物质的量)
	DimLuminous                     // J, luminous intensity (发光强度)
)

// String returns the English name of the dimension, e.g. "length".
func (d Dimension) String() string {
	m := map[Dimension]string{
		DimLength:      "length",
		DimMass:        "mass",
		DimTime:        "time",
		DimCurrent:     "current",
		DimTemperature: "temperature",
		DimAmount:      "amount",
		DimLuminous:    "luminous",
	}
	return m[d]
}

// Symbol returns the single-letter symbol for dimensional analysis.
func (d Dimension) Symbol() string {
	m := map[Dimension]string{
		DimLength:      "L",
		DimMass:        "M",
		DimTime:        "T",
		DimCurrent:     "I",
		DimTemperature: "Θ",
		DimAmount:      "N",
		DimLuminous:    "J",
	}
	return m[d]
}

// Base returns the SI (国际单位制) base unit for this dimension.
//
// Example: DimLength.Base() // Unit(Meter)
func (d Dimension) Base() Unit {
	var dimensionBaseUnit = map[Dimension]Unit{
		DimLength:      Unit(Meter),
		DimMass:        Unit(Kilogram),
		DimTime:        Unit(Second),
		DimCurrent:     Unit(Ampere),
		DimTemperature: Unit(Kelvin),
		DimAmount:      Unit(Mole),
		DimLuminous:    Unit(Candela),
	}
	base, ok := dimensionBaseUnit[d]
	if !ok {
		return ""
	}
	return base
}
