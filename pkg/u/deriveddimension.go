package u

import "sort"

// DerivedDimension (导出量纲) is the dimension of a derived physical quantity, expressed as
// exponents on the seven SI (国际单位制) base dimensions.
//
// In SI, the seven base dimensions (DimLength, DimMass, DimTime, DimCurrent, DimTemperature, DimAmount, DimLuminous) describe
// fundamental quantities. DerivedDimension and DerivedUnit (导出单位) describe derived
// quantities and their concrete unit combinations (e.g. speed: L¹T⁻¹, km·h⁻¹).
//
// Read each field as "this base dimension raised to that power". Omitted dimensions
// are zero. For example:
//
//	{L: 1, T: -1}        → speed     (m/s, km/h)
//	{M: 1, L: 1, T: -2}  → force     (N = kg·m/s²)
//	{M: 1, L: -1, T: -2} → pressure  (Pa)
//
// Dimensional homogeneity (量纲齐次性): two quantities may be added or subtracted
// only when their DerivedDimension values are equal. Multiplication and division
// of quantities always yield a well-defined derived dimension (exponents add or
// subtract); whether that result names a familiar physical quantity is a
// separate concern of physical interpretation, not of dimensional algebra.
//
// Symbol mapping (same as Dimension.Symbol):
//
//	L → length (长度，meter 米)
//	M → mass (质量，kilogram 千克)
//	T → time (时间，second 秒)
//	I → electric current (电流，ampere 安培)
//	Θ → thermodynamic temperature (热力学温度，kelvin 开尔文)
//	N → amount of substance (物质的量，mole 摩尔)
//	J → luminous intensity (发光强度，candela 坎德拉)
type DerivedDimension struct {
	L int8 // length (长度) exponent
	M int8 // mass (质量) exponent
	T int8 // time (时间) exponent
	I int8 // electric current (电流) exponent
	H int8 // thermodynamic temperature (热力学温度) exponent
	N int8 // amount of substance (物质的量) exponent
	J int8 // luminous intensity (发光强度) exponent
}

// Equal reports whether two derived dimensions are identical.
// Quantities can be added or subtracted only when dimensions are equal.
func (d DerivedDimension) Equal(o DerivedDimension) bool {
	return d == o
}

// Add returns the dimension of a sum when d and o are dimensionally homogeneous.
// On success ok is true and the result equals d (and o). If the dimensions differ,
// Add returns d with ok false — dimensional inconsistency has no sum dimension.
//
// Example: DimEnergy.Add(DimEnergy) // DimEnergy, true
// Example: DimEnergy.Add(DimForce)  // DimEnergy, false
func (d DerivedDimension) Add(o DerivedDimension) (DerivedDimension, bool) {
	if !d.Equal(o) {
		return d, false
	}
	return d, true
}

// Sub returns the dimension of a difference when d and o are dimensionally
// homogeneous. Same rules as Add: equal dimensions yield (d, true); otherwise
// (d, false).
func (d DerivedDimension) Sub(o DerivedDimension) (DerivedDimension, bool) {
	if !d.Equal(o) {
		return d, false
	}
	return d, true
}

// Mul returns the product dimension: quantity multiplication adds exponents.
// Always defined in dimensional algebra (e.g. force × length → energy).
//
// Example: DimForce.Mul(DerivedDimension{L: 1}) // DimEnergy
func (d DerivedDimension) Mul(o DerivedDimension) DerivedDimension {
	return DerivedDimension{
		L: d.L + o.L,
		M: d.M + o.M,
		T: d.T + o.T,
		I: d.I + o.I,
		H: d.H + o.H,
		N: d.N + o.N,
		J: d.J + o.J,
	}
}

// Div returns the quotient dimension: quantity division subtracts exponents.
// Always defined; may yield the dimensionless DerivedDimension{}.
//
// Example: DimForce.Div(DerivedDimension{L: 1, T: -2}) // DimMass (≈ M)
func (d DerivedDimension) Div(o DerivedDimension) DerivedDimension {
	return DerivedDimension{
		L: d.L - o.L,
		M: d.M - o.M,
		T: d.T - o.T,
		I: d.I - o.I,
		H: d.H - o.H,
		N: d.N - o.N,
		J: d.J - o.J,
	}
}

// Root returns the nth-root dimension when every base exponent is divisible by n.
// n must be >= 2. On failure ok is false and the original dimension is returned.
//
// Example: DerivedDimension{I: 2}.Root(2) // {I: 1}, true (ampere-squared → ampere)
func (d DerivedDimension) Root(n int) (DerivedDimension, bool) {
	if n < 2 {
		return d, false
	}
	divisible := func(e int8) bool { return int(e)%n == 0 }
	if !divisible(d.L) || !divisible(d.M) || !divisible(d.T) ||
		!divisible(d.I) || !divisible(d.H) || !divisible(d.N) || !divisible(d.J) {
		return d, false
	}
	return DerivedDimension{
		L: int8(int(d.L) / n),
		M: int8(int(d.M) / n),
		T: int8(int(d.T) / n),
		I: int8(int(d.I) / n),
		H: int8(int(d.H) / n),
		N: int8(int(d.N) / n),
		J: int8(int(d.J) / n),
	}, true
}

func (d DerivedDimension) addDimension(dim Dimension, exp int8) DerivedDimension {
	switch dim {
	case DimLength:
		d.L += exp
	case DimMass:
		d.M += exp
	case DimTime:
		d.T += exp
	case DimCurrent:
		d.I += exp
	case DimTemperature:
		d.H += exp
	case DimAmount:
		d.N += exp
	case DimLuminous:
		d.J += exp
	}
	return d
}

// Symbol renders the dimension as single-letter exponents, e.g. "M·L·T^-2".
//
// Example: d.Symbol(WithExpSign(ExpSignSup)) // "M·L·T⁻²"
func (d DerivedDimension) Symbol(options ...FormatOption) string {
	opt := applyFormatOptions(options...)
	return symbolFromDimensionExponents(dimensionExponents(d, &opt), &opt)
}

type dimensionExponent struct {
	dimension Dimension
	exp       int8
}

func dimensionExponents(d DerivedDimension, opt *formatOption) []dimensionExponent {
	candidates := []dimensionExponent{
		{DimMass, d.M},
		{DimLength, d.L},
		{DimTime, d.T},
		{DimCurrent, d.I},
		{DimTemperature, d.H},
		{DimAmount, d.N},
		{DimLuminous, d.J},
	}
	terms := make([]dimensionExponent, 0, len(candidates))
	for _, t := range candidates {
		if t.exp == 0 {
			continue
		}
		terms = append(terms, t)
	}
	sort.Slice(terms, func(i, j int) bool {
		return dimensionOrderRank(terms[i].dimension, opt.dimOrder) <
			dimensionOrderRank(terms[j].dimension, opt.dimOrder)
	})
	return terms
}

func symbolFromDimensionExponents(terms []dimensionExponent, opt *formatOption) string {
	mul := string(opt.mulSign)

	if opt.divSign == DivSignSlash {
		num := make([]string, 0, len(terms))
		den := make([]string, 0, len(terms))
		for _, t := range terms {
			sign := t.dimension.Symbol()
			if t.exp > 0 {
				num = append(num, formatTermSign(sign, t.exp, opt))
				continue
			}
			den = append(den, formatTermSign(sign, -t.exp, opt))
		}
		if len(num) == 0 && len(den) == 0 {
			return ""
		}
		if len(den) == 0 {
			return joinSymbolParts(num, mul)
		}
		denStr := joinSymbolParts(den, mul)
		if len(den) > 1 {
			denStr = "(" + denStr + ")"
		}
		if len(num) == 0 {
			return "1/" + denStr
		}
		return joinSymbolParts(num, mul) + "/" + denStr
	}

	parts := make([]string, 0, len(terms))
	for _, t := range terms {
		parts = append(parts, formatTermSign(t.dimension.Symbol(), t.exp, opt))
	}
	return joinSymbolParts(parts, mul)
}

func joinSymbolParts(parts []string, mul string) string {
	if len(parts) == 0 {
		return ""
	}
	out := parts[0]
	for _, part := range parts[1:] {
		out += mul + part
	}
	return out
}
