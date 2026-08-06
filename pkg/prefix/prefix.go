package prefix

import (
	"fmt"
	"math"
	"strings"
)

// PrefixMode selects SI (1000-based, 国际单位制) decimal, IEC (1024-based, 国际电工委员会) binary, or auto-detection rules.
type PrefixMode int

const (
	// Auto detects the best prefix mode when parsing or formatting.
	//
	// PrefixParse: Auto chooses SI or IEC based on whether the input ends with 'i'.
	// PrefixFormat: Auto behaves as SI mode.
	Auto PrefixMode = iota

	SI       // 1000 base, SI symbols (no 'i' suffix)
	IEC      // 1024 base, IEC symbols (with 'i' suffix)
	SI1024   // 1024 base, SI symbols (no 'i' suffix)
	ForceSI  // force 1000 base for both SI and IEC symbols
	ForceIEC // force 1024 base for both SI and IEC symbols
)

// String returns the name of the prefix mode.
func (m PrefixMode) String() string {
	switch m {
	case Auto:
		return "Auto"
	case SI:
		return "SI"
	case SI1024:
		return "SI1024"
	case IEC:
		return "IEC"
	case ForceSI:
		return "ForceSI"
	case ForceIEC:
		return "ForceIEC"
	default:
		return "<unknown>"
	}
}

// SIPrefix is an SI decimal prefix scale factor used by Unit.Prefix or Quantity.Prefix.
// Only the named constants below are valid inputs to those APIs.
// ref: https://nist.gov/pml/owm/metric-si-prefixes
type SIPrefix float64

const (
	Quecto SIPrefix = 1e-30 // quecto (亏，10⁻³⁰)
	Ronto  SIPrefix = 1e-27 // ronto (柔，10⁻²⁷)
	Yocto  SIPrefix = 1e-24 // yocto (幺，10⁻²⁴)
	Zepto  SIPrefix = 1e-21 // zepto (仄，10⁻²¹)
	Atto   SIPrefix = 1e-18 // atto (阿，10⁻¹⁸)
	Femto  SIPrefix = 1e-15 // femto (飞，10⁻¹⁵)
	Pico   SIPrefix = 1e-12 // pico (皮，10⁻¹²)
	Nano   SIPrefix = 1e-9  // nano (纳，10⁻⁹)
	Micro  SIPrefix = 1e-6  // micro (微，10⁻⁶)
	Milli  SIPrefix = 1e-3  // milli (毫，10⁻³)
	Centi  SIPrefix = 1e-2  // centi (厘，10⁻²); not used in scalesSI
	Deci   SIPrefix = 1e-1  // deci (分，10⁻¹); not used in scalesSI

	One SIPrefix = 1e0 // unity (1); not a prefix, used as the neutral scale

	Deka   SIPrefix = 10    // deka (十，10¹); not used in scalesSI
	Hecto  SIPrefix = 100   // hecto (百，10²); not used in scalesSI
	Kilo   SIPrefix = 1e+3  // kilo (千，10³)
	Mega   SIPrefix = 1e+6  // mega (兆，10⁶)
	Giga   SIPrefix = 1e+9  // giga (吉，10⁹)
	Tera   SIPrefix = 1e+12 // tera (太，10¹²)
	Peta   SIPrefix = 1e+15 // peta (拍，10¹⁵)
	Exa    SIPrefix = 1e+18 // exa (艾，10¹⁸)
	Zetta  SIPrefix = 1e+21 // zetta (泽，10²¹)
	Yotta  SIPrefix = 1e+24 // yotta (尧，10²⁴)
	Ronna  SIPrefix = 1e+27 // ronna (容，10²⁷)
	Quetta SIPrefix = 1e+30 // quetta (昆，10³⁰)
)

// IEC (国际电工委员会) binary (1024-based) prefix scale factors.
const (
	Yocbi float64 = 1.0 / (1 << 80)
	Zepbi float64 = 1.0 / (1 << 70)
	Attbi float64 = 1.0 / (1 << 60)
	Fembi float64 = 1.0 / (1 << 50)
	Picbi float64 = 1.0 / (1 << 40)
	Nanbi float64 = 1.0 / (1 << 30)
	Micbi float64 = 1.0 / (1 << 20)
	Milbi float64 = 1.0 / (1 << 10)

	Kibi float64 = 1 << 10 // kibi (千比，2¹⁰)
	Mebi float64 = 1 << 20 // mebi (兆比，2²⁰)
	Gibi float64 = 1 << 30 // gibi (吉比，2³⁰)
	Tebi float64 = 1 << 40 // tebi (太比，2⁴⁰)
	Pebi float64 = 1 << 50 // pebi (拍比，2⁵⁰)
	Exbi float64 = 1 << 60 // exbi (艾比，2⁶⁰)
	Zebi float64 = 1 << 70 // zebi (泽比，2⁷⁰)
	Yobi float64 = 1 << 80 // yobi (尧比，2⁸⁰)
)

// PrefixSymbol is a single-character SI or IEC numeric prefix (K, M, G, …).
type PrefixSymbol rune

var fakeSymbol PrefixSymbol = '_'

var (
	scalesIEC = []float64{
		Yocbi, Zepbi, Attbi, Fembi, Picbi, Nanbi, Micbi, Milbi,
		float64(One),
		Kibi, Mebi, Gibi, Tebi, Pebi, Exbi, Zebi, Yobi,
	}

	symbolsIEC = []PrefixSymbol{
		'y', 'z', 'a', 'f', 'p', 'n', 'u', 'm',
		fakeSymbol,
		'K', 'M', 'G', 'T', 'P', 'E', 'Z', 'Y',
	}
)

func formatScales(scales []float64) string {
	header := fmt.Sprintf("%-30s   %2s\n", "scale", "exp")
	header += strings.Repeat("=", 36)
	var format string
	for i, scale := range scales {

		exp := i - len(scales)/2
		format += fmt.Sprintf("\n%-30e    %2d", scale, exp)
	}
	return header + format
}

var (
	scalesSI = []float64{
		float64(Quecto), float64(Ronto), float64(Yocto), float64(Zepto), float64(Atto),
		float64(Femto), float64(Pico), float64(Nano), float64(Micro), float64(Milli),
		float64(One),
		float64(Kilo), float64(Mega), float64(Giga), float64(Tera), float64(Peta),
		float64(Exa), float64(Zetta), float64(Yotta), float64(Ronna), float64(Quetta),
	}
	symbolsSI = []PrefixSymbol{
		'q', 'r', 'y', 'z', 'a', 'f', 'p', 'n', 'u', 'm',
		fakeSymbol,
		'K', 'M', 'G', 'T', 'P', 'E', 'Z', 'Y', 'R', 'Q',
	}
)

// AllValidSymbols lists every rune recognized as a prefix symbol by PrefixParse.
var AllValidSymbols = ""

const (
	altKilo  = rune('k') // lowercase 'k' accepted as Kilo (千) prefix
	altMicro = rune('μ') // Greek mu accepted as Micro (微) prefix 'u'
)

func init() {
	if len(scalesIEC)%2 == 0 {
		panic("length of scaleIEC must be odd")
	}

	if len(scalesSI)%2 == 0 {
		panic("length of scaleSI must be odd")
	}

	if len(scalesIEC) != len(symbolsIEC) {
		panic("length of scalesIEC must be equal with length of symbolsIEC")
	}

	if len(scalesSI) != len(symbolsSI) {
		panic("length of scalesSI must be equal with length of symbolsSI")
	}

	allValidSymbols := make([]rune, 0)
	for _, s := range symbolsIEC {
		if s == fakeSymbol {
			continue
		}
		allValidSymbols = append(allValidSymbols, rune(s))
	}
	for _, s := range symbolsSI {
		if s == fakeSymbol {
			continue
		}
		allValidSymbols = append(allValidSymbols, rune(s))
	}
	allValidSymbols = append(allValidSymbols, altKilo, altMicro)

	AllValidSymbols = string(allValidSymbols)
}

func getExponentScaleSymbol(val float64, prefixMode PrefixMode) (exp int, scale float64, symbol PrefixSymbol, oppositeScale float64) {
	val = math.Abs(val)

	switch prefixMode {
	case SI, ForceSI, Auto:
		exp := exponentOfValue(val, scalesSI)
		scaleIndex := exp + len(scalesSI)/2
		symbolIndex := exp + len(symbolsSI)/2
		oppositeScaleIndex := -exp + len(scalesSI)/2

		return exp, scalesSI[scaleIndex], symbolsSI[symbolIndex], scalesSI[oppositeScaleIndex]

	case IEC, ForceIEC, SI1024:
		exp := exponentOfValue(val, scalesIEC)
		scaleIndex := exp + len(scalesIEC)/2
		symbolIndex := exp + len(symbolsIEC)/2
		oppositeScaleIndex := -exp + len(scalesIEC)/2

		return exp, scalesIEC[scaleIndex], symbolsIEC[symbolIndex], scalesIEC[oppositeScaleIndex]

	default:
		return 0, float64(One), fakeSymbol, float64(One)
	}
}

func getScaleOfSymbol(symbol rune, prefixMode PrefixMode) (scale float64, oppsiteScale float64, err error) {
	switch symbol {
	case altKilo:
		symbol = 'K'
	case altMicro:
		symbol = 'u'
	}

	var symbols []PrefixSymbol
	var scales []float64

	switch prefixMode {
	case Auto, SI, ForceSI:
		symbols = symbolsSI
		scales = scalesSI
	case IEC, SI1024, ForceIEC:
		symbols = symbolsIEC
		scales = scalesIEC
	default:
		return 0, 0, ErrInvalidMode
	}

	for i, s := range symbols {
		if symbol == rune(s) {
			oppositeIndex := len(scales) - i - 1
			return scales[i], scales[oppositeIndex], nil
		}
	}
	return float64(One), float64(One), ErrInvalidSymbol
}

func exponentOfValue(val float64, scales []float64) int {
	half := len(scales) / 2
	minExp := -half
	maxExp := half

	var exp = maxExp

	for exp >= minExp {
		scale := scales[exp+half]

		if scale <= val {
			break
		}
		exp--
	}

	if exp < minExp {
		exp = minExp
	}

	return exp
}
