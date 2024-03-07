package unit

import (
	"fmt"
	"math"
	"strconv"
)

type prefixFormatOption struct {
	space           bool        // default false
	precision       int         // default 0
	roundMethod     RoundMethod // rountMethod takes effect only when precision = 0
	roundDifference float64     // roundDifference (0 ~ 1.0) takes effect only when roundMethod is RoundMethodPercent
}

type prefixFormatOptionFn = func(opt *prefixFormatOption)

func newPrefixFormatOption(optionFns ...prefixFormatOptionFn) prefixFormatOption {
	var option = prefixFormatOption{
		roundMethod: RoundMethodFloor,
	}

	for _, optionFn := range optionFns {
		optionFn(&option)
	}
	return option
}

func WithSpace(space bool) prefixFormatOptionFn {
	return func(opt *prefixFormatOption) { opt.space = space }
}

func WithPrecision(precision int) prefixFormatOptionFn {
	return func(opt *prefixFormatOption) { opt.precision = precision }
}

func WithRoundMethod(roundMethod RoundMethod) prefixFormatOptionFn {
	return func(opt *prefixFormatOption) { opt.roundMethod = roundMethod }
}

func WithRoundDifference(roundPercentage float64) prefixFormatOptionFn {
	return func(opt *prefixFormatOption) { opt.roundDifference = roundPercentage }
}

// PrefixFormat2 return the formatted 'number' and 'unit prefix'.
// eg: float64(1048576) -> "1", "Mi"
func PrefixFormat2(val float64, prefixMode PrefixMode, prefixFormatOptionFns ...prefixFormatOptionFn) (number string, prefix string) {
	option := newPrefixFormatOption(prefixFormatOptionFns...)

	if math.IsNaN(val) || math.IsInf(val, 0) {
		return string(strconv.FormatFloat(val, 'f', -1, 64)), ""
	}

	_, scale, symbol, oppositeScale := getExponentScaleSymbol(val, prefixMode)

	if scale >= 1 {
		val = val / scale
	} else {
		val = val * oppositeScale
	}

	if symbol != fakeSymbol {
		switch prefixMode {
		case IEC:
			prefix = string(symbol) + "i"
		default:
			prefix = string(symbol)
		}
	}

	if option.precision == 0 {
		switch option.roundMethod {
		case RoundMethodFloor:
			val = math.Floor(val)

		case RoundMethodRound:
			val = math.Round(val)

		case RoundMethodCeil:
			val = math.Ceil(val)

		case RoundMethodDifference:
			var roundDifference float64
			if option.roundDifference < 0 {
				roundDifference = 0
			} else if option.roundDifference > 1.0 {
				roundDifference = 1
			} else {
				roundDifference = option.roundDifference
			}

			// diff := math.Abs(val - float64(int64(val)))
			if val >= 0 {
				difference := val - math.Floor(val)
				if difference >= roundDifference {
					val = math.Ceil(val)
				} else {
					val = math.Floor(val)
				}
			} else {
				diff := math.Ceil(val) - val
				if diff >= roundDifference {
					val = math.Floor(val)
				} else {
					val = math.Ceil(val)
				}
			}

		}
	}

	numberFormat := fmt.Sprintf("%%.%df", option.precision)
	return fmt.Sprintf(numberFormat, val), prefix
}

// PrefixFormat converts float64 val to a formatted string of 'number and unit prefix'.
// eg: float64(1048576) -> "1 Mi" or "1Mi", or "1M", or "1 M" or "1.00 M" or or "1.048 M"  or ...
//
// You can use PrefixFormat2 function to get separate 'number' and 'unit prefix'.
func PrefixFormat(val float64, prefixMode PrefixMode, prefixFormatOptionFns ...prefixFormatOptionFn) string {
	option := newPrefixFormatOption(prefixFormatOptionFns...)

	number, prefix := PrefixFormat2(val, prefixMode, prefixFormatOptionFns...)

	format := "%s"
	if option.space {
		format += " "
	}
	format += "%s"

	return fmt.Sprintf(format, number, prefix)
}
