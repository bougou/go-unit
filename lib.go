package unit

import "math"

func getExponentScaleSymbol(val float64, prefixMode PrefixMode) (exp int, scale float64, symbol Symbol, oppositeScale float64) {
	val = math.Abs(val)

	switch prefixMode {
	case SI, Auto:
		exp := exponentOfValue(val, scalesSI)
		scaleIndex := exp + len(scalesSI)/2
		symbolIndex := exp + len(symbolsSI)/2
		oppositeScaleIndex := -exp + len(scalesSI)/2

		return exp, scalesSI[scaleIndex], symbolsSI[symbolIndex], scalesSI[oppositeScaleIndex]

	case IEC, SI1024:
		exp := exponentOfValue(val, scalesIEC)
		scaleIndex := exp + len(scalesIEC)/2
		symbolIndex := exp + len(symbolsIEC)/2
		oppositeScaleIndex := -exp + len(scalesIEC)/2

		return exp, scalesIEC[scaleIndex], symbolsIEC[symbolIndex], scalesIEC[oppositeScaleIndex]

	default:
		return 0, One, fakeSymbol, One
	}
}

func getScaleOfSymbol(symbol rune, prefixMode PrefixMode) (scale float64, oppsiteScale float64, err error) {
	switch symbol {
	case altKilo:
		symbol = 'K'
	case altMicro:
		symbol = 'u'
	}

	var symbols []Symbol
	var scales []float64

	switch prefixMode {
	case Auto, SI:
		symbols = symbolsSI
		scales = scalesSI
	case IEC, SI1024:
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
	return One, One, ErrInvalidSymbol
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
