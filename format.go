package unit

import (
	"fmt"
	"math"
	"strconv"
)

type formatOption struct {
	space           bool        // default false
	precision       int         // default 0
	roundMethod     RoundMethod // rountMethod takes effect only when precision = 0
	comma           bool
	roundDifference float64 // roundDifference (0 ~ 1.0) takes effect only when roundMethod is RoundMethodPercent
}

type FormatOption = func(opt *formatOption)

func WithSpace(space bool) FormatOption {
	return func(opt *formatOption) { opt.space = space }
}

func WithPrecision(precision int) FormatOption {
	return func(opt *formatOption) { opt.precision = precision }
}

func WithRoundMethod(roundMethod RoundMethod) FormatOption {
	return func(opt *formatOption) { opt.roundMethod = roundMethod }
}

func WithRoundDifference(roundPercentage float64) FormatOption {
	return func(opt *formatOption) { opt.roundDifference = roundPercentage }
}

func Format(val float64, mode Mode, formatOptions ...FormatOption) string {
	var option = formatOption{}
	for _, fmt := range formatOptions {
		fmt(&option)
	}

	if math.IsNaN(val) || math.IsInf(val, 0) {
		return string(strconv.FormatFloat(val, 'f', -1, 64))
	}

	_, scale, symbol, oppositeScale := getExponentScaleSymbol(val, mode)

	if scale >= 1 {
		val = val / scale
	} else {
		val = val * oppositeScale
	}

	var prefix string
	if symbol == fakeSymbol {
		prefix = ""
	} else {
		switch mode {
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

	var format string
	format = fmt.Sprintf("%%.%df", option.precision)
	if option.space {
		format += " "
	}
	format += "%s"

	return fmt.Sprintf(format, val, prefix)
}
