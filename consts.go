package unit

// SI factors
const (
	Quecto float64 = 1e-30
	Ronto  float64 = 1e-27
	Yocto  float64 = 1e-24
	Zepto  float64 = 1e-21
	Atto   float64 = 1e-18
	Femto  float64 = 1e-15
	Pico   float64 = 1e-12
	Nano   float64 = 1e-9
	Micro  float64 = 1e-6 // 1000^-2
	Milli  float64 = 1e-3 // 1000^-1
	// Centi  float64 = 1e-2
	// Deci   float64 = 1e-1

	Unit float64 = 1e0 // Not a standard SI prefix.

	Kilo   float64 = 1e+3  // 1000^1
	Mega   float64 = 1e+6  // 1000^2
	Giga   float64 = 1e+9  // 1000^3
	Tera   float64 = 1e+12 // 1000^4
	Peta   float64 = 1e+15 // 1000^5
	Exa    float64 = 1e+18 // 1000^6
	Zetta  float64 = 1e+21
	Yotta  float64 = 1e+24
	Ronna  float64 = 1e+27
	Quetta float64 = 1e+30
)

// IEC factors
const (
	yocbi float64 = 1.0 / (1 << 80)
	zepbi float64 = 1.0 / (1 << 70)
	attbi float64 = 1.0 / (1 << 60)
	fembi float64 = 1.0 / (1 << 50)
	picbi float64 = 1.0 / (1 << 40)
	nanbi float64 = 1.0 / (1 << 30)
	micbi float64 = 1.0 / (1 << 20)
	milbi float64 = 1.0 / (1 << 10)

	Kibi float64 = 1 << 10
	Mebi float64 = 1 << 20
	Gibi float64 = 1 << 30
	Tebi float64 = 1 << 40
	Pebi float64 = 1 << 50
	Exbi float64 = 1 << 60
	Zebi float64 = 1 << 70
	Yobi float64 = 1 << 80
)

// Symbol is the single char which represents a specific unit prefix
type Symbol rune

var fakeSymbol Symbol = '_'

var (
	scalesIEC = []float64{
		yocbi, zepbi, attbi, fembi, picbi, nanbi, micbi, milbi,
		Unit,
		Kibi, Mebi, Gibi, Tebi, Pebi, Exbi, Zebi, Yobi,
	}

	symbolsIEC = []Symbol{
		'y', 'z', 'a', 'f', 'p', 'n', 'u', 'm',
		fakeSymbol,
		'K', 'M', 'G', 'T', 'P', 'E', 'Z', 'Y',
	}
)

var (
	scalesSI = []float64{
		Quecto, Ronto, Yocto, Zepto, Atto, Femto, Pico, Nano, Micro, Milli,
		Unit,
		Kilo, Mega, Giga, Tera, Peta, Exa, Zetta, Yotta, Ronna, Quetta,
	}
	symbolsSI = []Symbol{
		'q', 'r', 'y', 'z', 'a', 'f', 'p', 'n', 'u', 'm',
		fakeSymbol,
		'K', 'M', 'G', 'T', 'P', 'E', 'Z', 'Y', 'R', 'Q',
	}
)

// initialized by init function
var AllValidSymbols = ""

const (
	altKilo  = rune('k') // 'K'
	altMicro = rune('μ') // 'u'
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
