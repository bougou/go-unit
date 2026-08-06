package u

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/bougou/go-unit/pkg/prefix"
)

// unitTerm is one unit raised to an exponent in a compound unit, e.g. km¹ or h⁻¹.
// Exp > 0 places the term in the numerator; Exp < 0 in the denominator.
// Exp == 0 means the term is not present.
type unitTerm struct {
	unit Unit
	exp  int8
}

// Active reports whether the term participates in the compound unit.
func (t unitTerm) Active() bool {
	return t.exp != 0
}

func derivedDimensionFromTerms(terms []unitTerm) DerivedDimension {
	var dim DerivedDimension
	for _, t := range terms {
		if t.exp == 0 {
			continue
		}
		def, ok := t.unit.Def()
		if !ok {
			continue
		}
		dim = dim.addDimension(def.Dimension, t.exp)
	}
	return dim
}

func factorToBaseFromTerms(terms []unitTerm) float64 {
	factor := 1.0
	for _, t := range terms {
		if t.exp == 0 {
			continue
		}
		def, ok := t.unit.Def()
		if !ok {
			continue
		}
		factor *= math.Pow(def.Scale, float64(t.exp))
	}
	return factor
}

// DerivedUnit (导出单位) is a compound unit: for each SI (国际单位制) dimension, which unit and exponent.
// Symbol, Dim, Key, and FactorToBase are derived on demand.
//
// Build with NewDerivedUnit and chain Length/Mass/Time/…. Call Intern (内化注册) to register
// a canonical instance; otherwise the value is a private temporary object.
//
// Example:
//
//	speed := NewDerivedUnit().Length(Meter.Prefix(prefix.Kilo), 1).Time(Hour, -1)
//	speed.Symbol(WithExpSign(ExpSignSup)) // "km·h⁻¹"
type DerivedUnit struct {
	l unitTerm // length (L)
	m unitTerm // mass (M)
	t unitTerm // time (T)
	i unitTerm // current (I)
	h unitTerm // temperature (H)
	n unitTerm // amount (N)
	j unitTerm // luminous intensity (J)

	// namedSymbol is the SI special name (专用名称) when this compound unit has one,
	// e.g. "N" (牛顿) for kg·m·s⁻². Empty means only the compound form is used.
	// When prefixScale ≠ 1, display is prefix + namedSymbol (e.g. "M"+"Ω" → "MΩ").
	namedSymbol string

	// unitScale is a non-SI multiplier relative to the coherent base composition of
	// the active terms (e.g. 3600 for watt-hour vs joule). Zero means unset (= 1).
	// Distinct from prefixScale, which is only SI decimal prefixes on the special name.
	unitScale float64

	// prefixScale is the SI decimal multiplier relative to the coherent named unit
	// (scale 1). Zero means unset and is treated as 1. Example: Megaohm uses 1e6.
	prefixScale float64
}

// NewDerivedUnit starts an empty derived unit for chained construction.
// The result is not registered until Intern is called.
//
// Example:
//
//	speedUnit := NewDerivedUnit().Length(Meter.Prefix(prefix.Kilo), 1).Time(Hour, -1)
func NewDerivedUnit() *DerivedUnit {
	return &DerivedUnit{}
}

// Length sets the length (L) term. exp > 0 is numerator; exp < 0 is denominator.
// Panics if unit is unknown or belongs to another dimension.
//
// Example: NewDerivedUnit().Length(Meter, 2) // m²
func (u *DerivedUnit) Length(unit LengthUnit, exp int8) *DerivedUnit {
	return u.withUnit(DimLength, Unit(unit), exp)
}

// Mass sets the mass (M) term.
//
// Example: NewDerivedUnit().Mass(Kilogram, 1) // kg
func (u *DerivedUnit) Mass(unit MassUnit, exp int8) *DerivedUnit {
	return u.withUnit(DimMass, Unit(unit), exp)
}

// Time sets the time (T) term.
//
// Example: NewDerivedUnit().Time(Second, -1) // per second (s⁻¹)
func (u *DerivedUnit) Time(unit TimeUnit, exp int8) *DerivedUnit {
	return u.withUnit(DimTime, Unit(unit), exp)
}

// Current sets the electric current (I) term.
func (u *DerivedUnit) Current(unit CurrentUnit, exp int8) *DerivedUnit {
	return u.withUnit(DimCurrent, Unit(unit), exp)
}

// Temperature sets the thermodynamic temperature (Θ) term.
func (u *DerivedUnit) Temperature(unit TemperatureUnit, exp int8) *DerivedUnit {
	return u.withUnit(DimTemperature, Unit(unit), exp)
}

// Amount sets the amount-of-substance (N) term.
func (u *DerivedUnit) Amount(unit AmountUnit, exp int8) *DerivedUnit {
	return u.withUnit(DimAmount, Unit(unit), exp)
}

// Luminous sets the luminous intensity (J) term.
func (u *DerivedUnit) Luminous(unit LuminousUnit, exp int8) *DerivedUnit {
	return u.withUnit(DimLuminous, Unit(unit), exp)
}

// Named sets the SI special name for this derived unit, e.g. "N" for kg·m·s⁻².
func (u *DerivedUnit) Named(namedSymbol string) *DerivedUnit {
	if u == nil {
		return nil
	}
	u.namedSymbol = namedSymbol
	return u
}

// Scale sets a non-SI multiplier relative to the coherent composition of the terms.
// Used for convenience units such as watt-hour (1 W·h = 3600 J). Zero/omitted means 1.
// Prefer SI Prefix for decimal scaling of named units (kW·h = WattHour.Prefix(prefix.Kilo)).
//
// Example: NewDerivedUnit().… .Named("W·h").Scale(3600)
func (u *DerivedUnit) Scale(factor float64) *DerivedUnit {
	if u == nil {
		return nil
	}
	if factor <= 0 {
		panic(fmt.Sprintf("unit scale must be positive, got %g", factor))
	}
	if factor == 1 {
		u.unitScale = 0
		return u
	}
	u.unitScale = factor
	return u
}

// PrefixScale returns the SI prefix factor relative to the coherent named unit.
// Returns 1 when no prefix is applied.
func (u *DerivedUnit) PrefixScale() float64 {
	return u.effectivePrefixScale()
}

func (u *DerivedUnit) effectivePrefixScale() float64 {
	if u == nil || u.prefixScale == 0 {
		return 1
	}
	return u.prefixScale
}

func (u *DerivedUnit) effectiveUnitScale() float64 {
	if u == nil || u.unitScale == 0 {
		return 1
	}
	return u.unitScale
}

// Prefix returns a named derived unit scaled by an SI decimal prefix factor.
// The receiver must already have a special name (e.g. Ohm). Results are interned.
//
// Example:
//
//	Ohm.Prefix(prefix.Mega)   // MΩ, FactorToBase = 1e6
//	Ohm.Prefix(prefix.Micro)  // μΩ
//	Ohm.Prefix(prefix.Mega).Prefix(prefix.Micro) // Ω again (factors cancel)
func (u *DerivedUnit) Prefix(factor prefix.SIPrefix) *DerivedUnit {
	if u == nil {
		panic("Prefix on nil DerivedUnit")
	}
	if u.namedSymbol == "" {
		panic("Prefix requires a named derived unit (call Named first)")
	}
	if factor <= 0 {
		panic(fmt.Sprintf("SI prefix factor must be positive, got %g", factor))
	}
	mustSIPrefixByFactor(factor)

	scale := prefix.SIPrefix(u.effectivePrefixScale() * float64(factor))
	mustSIPrefixByFactor(scale)

	if scale == 1 {
		coherent := u.clone()
		coherent.prefixScale = 0
		coherent.namedSymbol = u.namedSymbol
		return coherent.MustIntern()
	}

	scaled := u.clone()
	scaled.namedSymbol = u.namedSymbol
	scaled.prefixScale = float64(scale)
	return scaled.MustIntern()
}

func (u *DerivedUnit) withUnit(dim Dimension, unit Unit, exp int8) *DerivedUnit {
	if u == nil {
		return nil
	}
	if exp == 0 {
		*u.term(dim) = unitTerm{}
		return u
	}
	def, ok := unit.Def()
	if !ok {
		panic(fmt.Sprintf("unknown unit %q", unit))
	}
	if def.Dimension != dim {
		panic(fmt.Sprintf("unit %q belongs to dimension %s, not %s", unit, def.Dimension, dim))
	}
	term := u.term(dim)
	if term.Active() && term.unit != unit {
		panic(fmt.Sprintf("dimension %s already uses unit %q, cannot switch to %q", dim, term.unit, unit))
	}
	term.unit = unit
	term.exp = exp
	return u
}

func (u *DerivedUnit) term(dim Dimension) *unitTerm {
	switch dim {
	case DimLength:
		return &u.l
	case DimMass:
		return &u.m
	case DimTime:
		return &u.t
	case DimCurrent:
		return &u.i
	case DimTemperature:
		return &u.h
	case DimAmount:
		return &u.n
	case DimLuminous:
		return &u.j
	default:
		return nil
	}
}

// terms returns the configured unit terms. Dimensions with Exp == 0 are omitted.
// A nil receiver returns nil.
func (u *DerivedUnit) terms() []unitTerm {
	if u == nil {
		return nil
	}
	var terms []unitTerm
	u.eachActiveTerm(func(term unitTerm) {
		terms = append(terms, term)
	})
	return terms
}

// Dim returns the derived dimension of this unit.
func (u *DerivedUnit) Dim() DerivedDimension {
	if u == nil {
		return DerivedDimension{}
	}
	return derivedDimensionFromTerms(u.terms())
}

// Key returns the registry identifier for this unit.
//
// Unnamed units use the base-unit composition (e.g. "kilogram^1*meter^1*second^-2").
// Named units append "#<symbol>" (e.g. "...#N", "second^-1#Hz"). Dimensionless named
// units use "#<symbol>" only (e.g. "#rad").
func (u *DerivedUnit) Key() Unit {
	return u.registryKey()
}

// compositionKey returns the base-unit composition without any special name suffix.
func (u *DerivedUnit) compositionKey() Unit {
	if u == nil {
		return ""
	}

	terms := u.terms()
	if len(terms) == 0 {
		return ""
	}
	if len(terms) == 2 && terms[0].exp == 1 && terms[1].exp == -1 {
		return Unit(string(terms[0].unit) + "/" + string(terms[1].unit))
	}

	sorted := append([]unitTerm(nil), terms...)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].unit == sorted[j].unit {
			return sorted[i].exp < sorted[j].exp
		}
		return sorted[i].unit < sorted[j].unit
	})

	parts := make([]string, len(sorted))
	for i, t := range sorted {
		parts[i] = fmt.Sprintf("%s^%d", t.unit, t.exp)
	}
	return Unit(strings.Join(parts, "*"))
}

func (u *DerivedUnit) registryKey() Unit {
	if u == nil {
		return ""
	}

	display := u.displayNamedSymbol()
	if len(u.terms()) == 0 {
		if display != "" {
			return Unit("#" + display)
		}
		return ""
	}

	base := u.compositionKey()
	if display != "" {
		return Unit(string(base) + "#" + display)
	}
	return base
}

// FactorToBase returns the multiplier to convert a value in this unit to SI base units.
func (u *DerivedUnit) FactorToBase() float64 {
	if u == nil {
		return 1
	}
	return factorToBaseFromTerms(u.terms()) * u.effectiveUnitScale() * u.effectivePrefixScale()
}

// Base returns an equivalent unit expressed with SI base units for each dimension.
//
// Example: km/h → m·s⁻¹
func (u *DerivedUnit) Base() *DerivedUnit {
	if u == nil {
		return nil
	}
	return siDerivedUnit(u.Dim())
}

// Intern returns the canonical registered copy for this unit (interning, 内化注册).
//
// Unnamed units are keyed by base-unit composition. Named units use composition#symbol
// (or #symbol when dimensionless). When the first registrant for a composition is named,
// the composition key is aliased to that instance so later unnamed Intern calls share it.
//
// Example:
//
//	a, _ := NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 1).Time(Second, -2).Named("N").Intern()
//	b, _ := NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 1).Time(Second, -2).Intern()
//	a == b // true when Newton was registered first in init
func (u *DerivedUnit) Intern() (*DerivedUnit, error) {
	if u == nil {
		return nil, fmt.Errorf("nil derived unit")
	}
	compKey := u.compositionKey()
	key := u.registryKey()

	if u.namedSymbol == "" {
		if existing, ok := derivedRegistry[compKey]; ok {
			return existing, nil
		}
	} else if existing, ok := derivedRegistry[key]; ok {
		return existing, nil
	}

	canonical := u.clone()
	derivedRegistry[key] = canonical
	if u.namedSymbol == "" {
		derivedRegistry[compKey] = canonical
	} else if u.effectivePrefixScale() == 1 {
		// Only coherent (unprefixed) named units alias the composition key.
		if _, ok := derivedRegistry[compKey]; !ok {
			derivedRegistry[compKey] = canonical
		}
	}
	registerDerivedUnitAliases(canonical, key)
	return canonical, nil
}

func registerDerivedUnitAliases(u *DerivedUnit, key Unit) {
	seen := make(map[string]struct{})
	register := func(sym string) {
		if sym == "" {
			return
		}
		if _, ok := seen[sym]; ok {
			return
		}
		seen[sym] = struct{}{}
		registerUnitSymbol(sym, key)
	}

	for _, opts := range [][]FormatOption{
		nil, // special name when set (e.g. "N", "kW·h")
		{WithCompoundSymbol(true)},
		{WithCompoundSymbol(true), WithExpSign(ExpSignSup)},
		{WithCompoundSymbol(true), WithDivSign(DivSignSlash)},
		{WithCompoundSymbol(true), WithMulSign(MulSignStar)},
		{WithCompoundSymbol(true), WithMulSign(MulSignSpace)},
		{WithCompoundSymbol(true), WithDivSign(DivSignSlash), WithExpSign(ExpSignSup)},
		{WithCompoundSymbol(true), WithMulSign(MulSignStar), WithExpSign(ExpSignSup)},
	} {
		register(u.Symbol(opts...))
	}
}

func (u *DerivedUnit) MustIntern() *DerivedUnit {
	interned, err := u.Intern()
	if err != nil {
		panic(err)
	}
	return interned
}

// Of creates a derived quantity with value in u.
//
// Example: Newton.Of(100) // 100 N
func (u *DerivedUnit) Of(value float64) DerivedQuantity {
	return NewDerivedQuantity(value, u)
}

func (u *DerivedUnit) clone() *DerivedUnit {
	return &DerivedUnit{
		l:           u.l,
		m:           u.m,
		t:           u.t,
		i:           u.i,
		h:           u.h,
		n:           u.n,
		j:           u.j,
		namedSymbol: u.namedSymbol,
		unitScale:   u.unitScale,
		prefixScale: u.prefixScale,
	}
}

func (u *DerivedUnit) eachActiveTerm(fn func(term unitTerm)) {
	for _, term := range []unitTerm{u.m, u.l, u.t, u.i, u.h, u.n, u.j} {
		if !term.Active() {
			continue
		}
		fn(term)
	}
}

var derivedRegistry = map[Unit]*DerivedUnit{}

func derivedUnitFromTerms(terms []unitTerm) (*DerivedUnit, error) {
	u := NewDerivedUnit()
	for _, t := range terms {
		if t.exp == 0 {
			continue
		}
		def, ok := t.unit.Def()
		if !ok {
			return nil, fmt.Errorf("unknown unit %q", t.unit)
		}
		switch def.Dimension {
		case DimLength:
			u.Length(LengthUnit(t.unit), t.exp)
		case DimMass:
			u.Mass(MassUnit(t.unit), t.exp)
		case DimTime:
			u.Time(TimeUnit(t.unit), t.exp)
		case DimCurrent:
			u.Current(CurrentUnit(t.unit), t.exp)
		case DimTemperature:
			u.Temperature(TemperatureUnit(t.unit), t.exp)
		case DimAmount:
			u.Amount(AmountUnit(t.unit), t.exp)
		case DimLuminous:
			u.Luminous(LuminousUnit(t.unit), t.exp)
		default:
			return nil, fmt.Errorf("unsupported dimension %s", def.Dimension)
		}
	}
	return u.Intern()
}

func sortedUnitTerms(terms []unitTerm) []unitTerm {
	sorted := append([]unitTerm(nil), terms...)
	sort.Slice(sorted, func(i, j int) bool {
		di, okI := sorted[i].unit.Def()
		dj, okJ := sorted[j].unit.Def()
		if okI && okJ && di.Dimension != dj.Dimension {
			return dimensionOrderRank(di.Dimension, DimOrderMLT) < dimensionOrderRank(dj.Dimension, DimOrderMLT)
		}
		if sorted[i].unit == sorted[j].unit {
			return sorted[i].exp < sorted[j].exp
		}
		return sorted[i].unit < sorted[j].unit
	})
	return sorted
}

func dimensionOrderRank(dim Dimension, order DimOrder) int {
	switch order {
	case DimOrderTML:
		switch dim {
		case DimTime:
			return 0
		case DimLength:
			return 1
		case DimMass:
			return 2
		case DimCurrent:
			return 3
		case DimTemperature:
			return 4
		case DimAmount:
			return 5
		case DimLuminous:
			return 6
		default:
			return 99
		}
	default:
		switch dim {
		case DimMass:
			return 0
		case DimLength:
			return 1
		case DimTime:
			return 2
		case DimCurrent:
			return 3
		case DimTemperature:
			return 4
		case DimAmount:
			return 5
		case DimLuminous:
			return 6
		default:
			return 99
		}
	}
}

func formatSuperscriptExponent(exp int8) string {
	if exp == 0 {
		return ""
	}

	negative := exp < 0
	n := int(exp)
	if negative {
		n = -n
	}

	digits := strconv.Itoa(n)
	var buf strings.Builder
	if negative {
		buf.WriteRune('⁻')
	}
	for _, ch := range digits {
		switch ch {
		case '0':
			buf.WriteRune('⁰')
		case '1':
			buf.WriteRune('¹')
		case '2':
			buf.WriteRune('²')
		case '3':
			buf.WriteRune('³')
		case '4':
			buf.WriteRune('⁴')
		case '5':
			buf.WriteRune('⁵')
		case '6':
			buf.WriteRune('⁶')
		case '7':
			buf.WriteRune('⁷')
		case '8':
			buf.WriteRune('⁸')
		case '9':
			buf.WriteRune('⁹')
		default:
			return fmt.Sprintf("^%d", exp)
		}
	}
	return buf.String()
}

// DerivedUnit returns the registered derived unit for key, if Intern was called.
func (key Unit) DerivedUnit() (*DerivedUnit, bool) {
	u, ok := derivedRegistry[key]
	return u, ok
}

func siDerivedUnit(dim DerivedDimension) *DerivedUnit {
	u := NewDerivedUnit()
	add := func(d Dimension, exp int8) {
		if exp == 0 {
			return
		}
		base := d.Base()
		if base == "" {
			return
		}
		switch d {
		case DimLength:
			u.Length(LengthUnit(base), exp)
		case DimMass:
			u.Mass(MassUnit(base), exp)
		case DimTime:
			u.Time(TimeUnit(base), exp)
		case DimCurrent:
			u.Current(CurrentUnit(base), exp)
		case DimTemperature:
			u.Temperature(TemperatureUnit(base), exp)
		case DimAmount:
			u.Amount(AmountUnit(base), exp)
		case DimLuminous:
			u.Luminous(LuminousUnit(base), exp)
		}
	}
	add(DimLength, dim.L)
	add(DimMass, dim.M)
	add(DimTime, dim.T)
	add(DimCurrent, dim.I)
	add(DimTemperature, dim.H)
	add(DimAmount, dim.N)
	add(DimLuminous, dim.J)
	return u
}

func mustFactorToBase(u *DerivedUnit) float64 {
	return u.FactorToBase()
}

func derivedUnitsConvertible(source, target *DerivedUnit) bool {
	if source == nil || target == nil {
		return false
	}
	for _, dim := range [...]Dimension{
		DimMass, DimLength, DimTime, DimCurrent, DimTemperature, DimAmount, DimLuminous,
	} {
		sourceTerm := source.term(dim)
		targetTerm := target.term(dim)
		if sourceTerm.exp != targetTerm.exp {
			return false
		}
		if !sourceTerm.Active() {
			continue
		}
		if sourceTerm.unit == targetTerm.unit {
			continue
		}
		if !unitsProportional(sourceTerm.unit, targetTerm.unit) {
			return false
		}
	}
	return true
}

// derivedUnitsOverlapProportional reports whether every base dimension that is
// active in both units uses proportionally related units (constant ratio via the
// dimension base). Affine pairs such as °C/K fail. Dimensions present in only
// one operand are ignored — they do not need to agree for Mul/Div.
func derivedUnitsOverlapProportional(a, b *DerivedUnit) bool {
	if a == nil || b == nil {
		return false
	}
	for _, dim := range [...]Dimension{
		DimMass, DimLength, DimTime, DimCurrent, DimTemperature, DimAmount, DimLuminous,
	} {
		aTerm := a.term(dim)
		bTerm := b.term(dim)
		if !aTerm.Active() || !bTerm.Active() {
			continue
		}
		if aTerm.unit == bTerm.unit {
			continue
		}
		if !unitsProportional(aTerm.unit, bTerm.unit) {
			return false
		}
	}
	return true
}

func derivedUnitConversionFactor(source, target *DerivedUnit) float64 {
	factor := 1.0
	for _, dim := range [...]Dimension{
		DimMass, DimLength, DimTime, DimCurrent, DimTemperature, DimAmount, DimLuminous,
	} {
		sourceTerm := source.term(dim)
		targetTerm := target.term(dim)
		if !sourceTerm.Active() || sourceTerm.unit == targetTerm.unit {
			continue
		}
		sourceDef, ok := sourceTerm.unit.Def()
		if !ok {
			continue
		}
		targetDef, ok := targetTerm.unit.Def()
		if !ok {
			continue
		}
		ratio := sourceDef.Scale / targetDef.Scale
		factor *= math.Pow(ratio, float64(sourceTerm.exp))
	}
	return factor *
		source.effectiveUnitScale() * source.effectivePrefixScale() /
		(target.effectiveUnitScale() * target.effectivePrefixScale())
}

func mustDerivedDim(u *DerivedUnit) DerivedDimension {
	return u.Dim()
}

// MulSign selects the multiplication separator in compound unit symbols.
type MulSign string

const (
	MulSignDot   MulSign = "·" // km·h⁻¹
	MulSignSpace MulSign = " " // km h^-1
	MulSignStar  MulSign = "*" // km*h^-1
	MulSignNone  MulSign = ""  // only valid for DerivedDimension.Symbol
)

// DivSign selects how negative exponents are rendered.
type DivSign string

const (
	DivSignSlash    DivSign = "/" // km/h
	DivSignNegative DivSign = "⁻" // h⁻¹ in km·h⁻¹
)

// ExpSign selects how exponents are rendered.
type ExpSign string

const (
	ExpSignCarat ExpSign = "^"   // m^2
	ExpSignSup   ExpSign = "sup" // m² (Unicode superscript)
)

// DimOrder controls the ordering of dimension symbols in compound units.
type DimOrder string

const (
	DimOrderTML DimOrder = "tml" // T, L, M, I, H, N, J (ISO 80000 convention)
	DimOrderMLT DimOrder = "mlt" // M, L, T, I, H, N, J (textbook convention)
)

// PrecisionAuto selects compact %g-style numeric formatting (Format default).
const PrecisionAuto = -1

// formatOption holds display settings for Symbol and Format.
type formatOption struct {
	// Value formatting (Quantity / DerivedQuantity Format).
	delimiter NumberDelimiter
	precision int // PrecisionAuto means %g

	// Symbol formatting (Unit / DerivedUnit / DerivedDimension Symbol).
	mulSign     MulSign
	divSign     DivSign
	expSign     ExpSign
	dimOrder    DimOrder
	useCompound bool // false (zero value): prefer SI special name when set
}

var defaultFormatOption = formatOption{
	delimiter: NumberDelimiterNone,
	precision: PrecisionAuto,
	mulSign:   MulSignDot,
	divSign:   DivSignNegative,
	expSign:   ExpSignCarat,
	dimOrder:  DimOrderMLT,
	// useCompound stays false: Format/Symbol default to the special name (e.g. "N").
}

// FormatOption configures Symbol and Format output (value + unit symbol).
type FormatOption func(option *formatOption)

// WithNumberDelimiter sets the thousands separator for the numeric value in Format.
// Prefer NumberDelimiterNone, NumberDelimiterComma, or NumberDelimiterUnderscore;
// space separators are not recommended (they break QuantityParse).
func WithNumberDelimiter(delimiter NumberDelimiter) FormatOption {
	return func(option *formatOption) {
		option.delimiter = delimiter
	}
}

// WithPrecision sets fixed decimal places for Format. PrecisionAuto (default) uses %g.
func WithPrecision(precision int) FormatOption {
	return func(option *formatOption) {
		option.precision = precision
	}
}

// WithMulSign sets the multiplication separator.
func WithMulSign(mulSign MulSign) FormatOption {
	return func(option *formatOption) {
		option.mulSign = mulSign
	}
}

// WithDivSign sets how division / negative exponents are rendered.
func WithDivSign(divSign DivSign) FormatOption {
	return func(option *formatOption) {
		option.divSign = divSign
	}
}

// WithExpSign sets how exponents are rendered.
//
// Example: unit.Symbol(WithExpSign(ExpSignSup)) // "km·h⁻¹"
func WithExpSign(expSign ExpSign) FormatOption {
	return func(option *formatOption) {
		option.expSign = expSign
	}
}

// WithDimOrder sets the ordering of dimension letters in compound symbols.
func WithDimOrder(dimOrder DimOrder) FormatOption {
	return func(option *formatOption) {
		option.dimOrder = dimOrder
	}
}

// WithCompoundSymbol requests the compound base-unit expression instead of the SI special name.
// Default (false / omitted) prefers the special name when one is set (e.g. "N").
// Pass true for "kg·m·s^-2" instead of "N".
func WithCompoundSymbol(useCompound bool) FormatOption {
	return func(option *formatOption) {
		option.useCompound = useCompound
	}
}

func applyFormatOptions(options ...FormatOption) formatOption {
	opt := defaultFormatOption
	for _, fn := range options {
		fn(&opt)
	}
	return opt
}

// Symbol returns the display symbol for u.
//
// By default returns the SI special name when one is set (e.g. "N").
// Pass WithCompoundSymbol(true) for the compound base-unit expression.
//
// Example:
//
//	Newton.Symbol()                                      // "N"
//	Newton.Symbol(WithCompoundSymbol(true))              // "kg·m·s^-2"
//	Newton.Prefix(prefix.Kilo).Symbol(WithCompoundSymbol(true)) // "k(kg·m·s^-2)"
//	Newton.Symbol(WithCompoundSymbol(true), WithExpSign(ExpSignSup)) // "kg·m·s⁻²"
//	Newton.Symbol(WithCompoundSymbol(true), WithDivSign(DivSignSlash)) // "kg·m/s^2"
func (u *DerivedUnit) Symbol(options ...FormatOption) string {
	return u.symbolWith(applyFormatOptions(options...))
}

func (u *DerivedUnit) symbolWith(opt formatOption) string {
	if u == nil {
		return ""
	}
	if !opt.useCompound && u.namedSymbol != "" {
		return u.displayNamedSymbol()
	}
	compound := symbolFromTerms(u.terms(), &opt)
	scale := u.effectivePrefixScale()
	if scale == 1 || compound == "" {
		return compound
	}
	prefix, ok := siPrefixDisplaySymbol(prefix.SIPrefix(scale))
	if !ok {
		return compound
	}
	// Keep the SI prefix outside the compound so "k" is not glued onto "kg"
	// (e.g. kW·h → "k(kg·m^2·s^-2)", not "kg·m^2·s^-2").
	return prefix + "(" + compound + ")"
}

// NamedSymbol returns the SI special name, if any (without SI prefix).
func (u *DerivedUnit) NamedSymbol() string {
	if u == nil {
		return ""
	}
	return u.namedSymbol
}

// displayNamedSymbol returns the named symbol including any SI prefix (e.g. "MΩ").
func (u *DerivedUnit) displayNamedSymbol() string {
	if u == nil || u.namedSymbol == "" {
		return ""
	}
	prefix, ok := siPrefixDisplaySymbol(prefix.SIPrefix(u.effectivePrefixScale()))
	if !ok {
		return u.namedSymbol
	}
	return prefix + u.namedSymbol
}

func symbolFromTerms(terms []unitTerm, opt *formatOption) string {
	sorted := sortedUnitTerms(terms)
	mul := string(opt.mulSign)

	if opt.divSign == DivSignSlash {
		num := make([]string, 0, len(sorted))
		den := make([]string, 0, len(sorted))
		for _, t := range sorted {
			if t.exp == 0 {
				continue
			}
			def, ok := t.unit.Def()
			if !ok {
				continue
			}
			if t.exp > 0 {
				num = append(num, formatTermSign(def.Symbol, t.exp, opt))
				continue
			}
			den = append(den, formatTermSign(def.Symbol, -t.exp, opt))
		}
		if len(num) == 0 && len(den) == 0 {
			return ""
		}
		if len(den) == 0 {
			return strings.Join(num, mul)
		}
		denStr := strings.Join(den, mul)
		if len(den) > 1 {
			denStr = "(" + denStr + ")"
		}
		if len(num) == 0 {
			return "1/" + denStr
		}
		return strings.Join(num, mul) + "/" + denStr
	}

	parts := make([]string, 0, len(sorted))
	for _, t := range sorted {
		if t.exp == 0 {
			continue
		}
		def, ok := t.unit.Def()
		if !ok {
			continue
		}
		parts = append(parts, formatTermSign(def.Symbol, t.exp, opt))
	}
	return strings.Join(parts, mul)
}

func formatTermSign(sign string, exp int8, opt *formatOption) string {
	if exp == 1 {
		return sign
	}
	if exp == 0 {
		return ""
	}
	if opt.expSign == ExpSignSup {
		return sign + formatSuperscriptExponent(exp)
	}
	return sign + fmt.Sprintf("^%d", exp)
}

func mustInternDerived(u *DerivedUnit) *DerivedUnit {
	interned, err := u.Intern()
	if err != nil {
		return u.clone()
	}
	return interned
}

func derivedUnitFromBaseUnit(u Unit) (*DerivedUnit, bool) {
	def, ok := u.Def()
	if !ok {
		return nil, false
	}
	du := NewDerivedUnit()
	if !setDerivedUnitExp(du, def.Dimension, u, 1) {
		return nil, false
	}
	return mustInternDerived(du), true
}

func mulDerivedUnits(primary, secondary *DerivedUnit) *DerivedUnit {
	return derivedUnitFromDimensionAndPrefs(primary.Dim().Mul(secondary.Dim()), primary, secondary)
}

func divDerivedUnits(primary, secondary *DerivedUnit) *DerivedUnit {
	dim := primary.Dim().Div(secondary.Dim())
	if dim.Equal(DerivedDimension{}) {
		return NoneUnit
	}
	return derivedUnitFromDimensionAndPrefs(dim, primary, secondary)
}

// rootDerivedUnit returns the unit of the nth root of u when every base exponent
// is divisible by n. Dimensionless roots yield NoneUnit.
func rootDerivedUnit(u *DerivedUnit, n int) (*DerivedUnit, bool) {
	if u == nil || n < 2 {
		return nil, false
	}
	dim, ok := u.Dim().Root(n)
	if !ok {
		return nil, false
	}
	if dim.Equal(DerivedDimension{}) {
		return NoneUnit, true
	}
	return derivedUnitFromDimensionAndPrefs(dim, u, nil), true
}

func derivedUnitFromDimensionAndPrefs(dim DerivedDimension, primary, secondary *DerivedUnit) *DerivedUnit {
	u := NewDerivedUnit()
	setIf := func(d Dimension, exp int8) {
		if exp == 0 {
			return
		}
		setDerivedUnitExp(u, d, pickUnitForDimension(primary, secondary, d), exp)
	}
	setIf(DimMass, dim.M)
	setIf(DimLength, dim.L)
	setIf(DimTime, dim.T)
	setIf(DimCurrent, dim.I)
	setIf(DimTemperature, dim.H)
	setIf(DimAmount, dim.N)
	setIf(DimLuminous, dim.J)
	return mustInternDerived(u)
}

func pickUnitForDimension(primary, secondary *DerivedUnit, dim Dimension) Unit {
	if primary != nil {
		if term := primary.term(dim); term != nil && term.Active() {
			return term.unit
		}
	}
	if secondary != nil {
		if term := secondary.term(dim); term != nil && term.Active() {
			return term.unit
		}
	}
	return dim.Base()
}

func setDerivedUnitExp(u *DerivedUnit, dim Dimension, unit Unit, exp int8) bool {
	if u == nil {
		return false
	}
	switch dim {
	case DimLength:
		u.Length(LengthUnit(unit), exp)
	case DimMass:
		u.Mass(MassUnit(unit), exp)
	case DimTime:
		u.Time(TimeUnit(unit), exp)
	case DimCurrent:
		u.Current(CurrentUnit(unit), exp)
	case DimTemperature:
		u.Temperature(TemperatureUnit(unit), exp)
	case DimAmount:
		u.Amount(AmountUnit(unit), exp)
	case DimLuminous:
		u.Luminous(LuminousUnit(unit), exp)
	default:
		return false
	}
	return true
}
