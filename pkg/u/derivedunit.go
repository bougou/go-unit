package u

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
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
//	speed := NewDerivedUnit().Length(Kilometer, 1).Time(Hour, -1)
//	speed.Symbol(WithExpSign(ExpSignSup)) // "km·h⁻¹"
type DerivedUnit struct {
	l unitTerm // length (L)
	m unitTerm // mass (M)
	t unitTerm // time (T)
	i unitTerm // current (I)
	h unitTerm // temperature (H)
	n unitTerm // amount (N)
	j unitTerm // luminous intensity (J)

	// specialSymbol is the SI special name (专用名称) when this compound unit has one,
	// e.g. "N" (牛顿) for kg·m·s⁻². Empty means only the compound form is used.
	specialSymbol string
}

// NewDerivedUnit starts an empty derived unit for chained construction.
// The result is not registered until Intern is called.
//
// Example:
//
//	speedUnit := NewDerivedUnit().Length(Kilometer, 1).Time(Hour, -1)
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
func (u *DerivedUnit) Named(specialSymbol string) *DerivedUnit {
	if u == nil {
		return nil
	}
	u.specialSymbol = specialSymbol
	return u
}

// SpecialSymbol returns the SI special name, if any.
func (u *DerivedUnit) SpecialSymbol() string {
	if u == nil {
		return ""
	}
	return u.specialSymbol
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

	if len(u.terms()) == 0 {
		if u.specialSymbol != "" {
			return Unit("#" + u.specialSymbol)
		}
		return ""
	}

	base := u.compositionKey()
	if u.specialSymbol != "" {
		return Unit(string(base) + "#" + u.specialSymbol)
	}
	return base
}

// FactorToBase returns the multiplier to convert a value in this unit to SI base units.
func (u *DerivedUnit) FactorToBase() float64 {
	if u == nil {
		return 1
	}
	return factorToBaseFromTerms(u.terms())
}

// SI returns an equivalent unit expressed with SI base units for each dimension.
//
// Example: km/h → m·s⁻¹
func (u *DerivedUnit) SI() *DerivedUnit {
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
//	a == b // true when ForceUnit was registered first in init
func (u *DerivedUnit) Intern() (*DerivedUnit, error) {
	if u == nil {
		return nil, fmt.Errorf("nil derived unit")
	}
	compKey := u.compositionKey()
	key := u.registryKey()

	if u.specialSymbol == "" {
		if existing, ok := derivedRegistry[compKey]; ok {
			return existing, nil
		}
	} else if existing, ok := derivedRegistry[key]; ok {
		return existing, nil
	}

	canonical := u.clone()
	derivedRegistry[key] = canonical
	if u.specialSymbol == "" {
		derivedRegistry[compKey] = canonical
	} else if _, ok := derivedRegistry[compKey]; !ok {
		derivedRegistry[compKey] = canonical
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

	for _, opts := range [][]SymbolOption{
		nil,
		{WithExpSign(ExpSignSup)},
		{WithDivSign(DivSignSlash)},
		{WithMulSign(MulSignStar)},
		{WithMulSign(MulSignSpace)},
		{WithNamedSymbol(true)},
		{WithDivSign(DivSignSlash), WithExpSign(ExpSignSup)},
		{WithMulSign(MulSignStar), WithExpSign(ExpSignSup)},
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

func (u *DerivedUnit) clone() *DerivedUnit {
	return &DerivedUnit{
		l:             u.l,
		m:             u.m,
		t:             u.t,
		i:             u.i,
		h:             u.h,
		n:             u.n,
		j:             u.j,
		specialSymbol: u.specialSymbol,
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
	return factor
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

type symbolOption struct {
	mulSign  MulSign
	divSign  DivSign
	expSign  ExpSign
	dimOrder DimOrder
	useNamed bool
}

var defaultSymbolOption = symbolOption{
	mulSign:  MulSignDot,
	divSign:  DivSignNegative,
	expSign:  ExpSignCarat,
	dimOrder: DimOrderMLT,
	useNamed: false,
}

// SymbolOption configures Symbol output for DerivedUnit and DerivedDimension.
type SymbolOption func(option *symbolOption)

// WithMulSign sets the multiplication separator.
func WithMulSign(mulSign MulSign) SymbolOption {
	return func(option *symbolOption) {
		option.mulSign = mulSign
	}
}

// WithDivSign sets how division / negative exponents are rendered.
func WithDivSign(divSign DivSign) SymbolOption {
	return func(option *symbolOption) {
		option.divSign = divSign
	}
}

// WithExpSign sets how exponents are rendered.
//
// Example: unit.Symbol(WithExpSign(ExpSignSup)) // "km·h⁻¹"
func WithExpSign(expSign ExpSign) SymbolOption {
	return func(option *symbolOption) {
		option.expSign = expSign
	}
}

// WithDimOrder sets the ordering of dimension letters in compound symbols.
func WithDimOrder(dimOrder DimOrder) SymbolOption {
	return func(option *symbolOption) {
		option.dimOrder = dimOrder
	}
}

// WithNamedSymbol controls whether Symbol returns the SI special name when one is set.
// Default (no options) uses the compound base-unit expression. Pass true for "N" instead of "kg·m·s^-2".
func WithNamedSymbol(useNamed bool) SymbolOption {
	return func(option *symbolOption) {
		option.useNamed = useNamed
	}
}

func symbolOptions(options []SymbolOption) symbolOption {
	opt := defaultSymbolOption
	for _, fn := range options {
		fn(&opt)
	}
	return opt
}

// Symbol returns the display symbol for u.
//
// By default returns the compound base-unit expression (useNamed is false).
// Pass WithNamedSymbol(true) for the SI special name when set (e.g. "N").
//
// Example:
//
//	ForceUnit.Symbol()                           // "kg·m·s^-2"
//	ForceUnit.Symbol(WithNamedSymbol(true))      // "N"
//	ForceUnit.Symbol(WithExpSign(ExpSignSup))    // "kg·m·s⁻²"
//	ForceUnit.Symbol(WithDivSign(DivSignSlash))  // "kg·m/s^2"
func (u *DerivedUnit) Symbol(options ...SymbolOption) string {
	if u == nil {
		return ""
	}
	opt := symbolOptions(options)
	if opt.useNamed && u.specialSymbol != "" {
		return u.specialSymbol
	}
	return symbolFromTerms(u.terms(), &opt)
}

func symbolFromTerms(terms []unitTerm, opt *symbolOption) string {
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

func formatTermSign(sign string, exp int8, opt *symbolOption) string {
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
	return derivedUnitFromDimensionAndPrefs(primary.Dim().Add(secondary.Dim()), primary, secondary)
}

func divDerivedUnits(primary, secondary *DerivedUnit) *DerivedUnit {
	dim := primary.Dim().Sub(secondary.Dim())
	if dim.Equal(DerivedDimension{}) {
		return NoneUnit
	}
	return derivedUnitFromDimensionAndPrefs(dim, primary, secondary)
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
