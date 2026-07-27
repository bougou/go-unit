package u

import (
	"fmt"
	"math"
	"strconv"
)

type prefixFormatOption struct {
	space           bool        // default false
	precision       int         // default 0
	roundMethod     RoundMethod // applies only when precision is 0
	roundDifference float64     // applies only when roundMethod is RoundMethodDifference
	prefix          Symbol
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

// WithPrefix forces a specific prefix symbol instead of auto-selecting from val.
func WithPrefix(prefix Symbol) prefixFormatOptionFn {
	return func(opt *prefixFormatOption) { opt.prefix = prefix }
}

// WithPrefixSpace inserts a space between the number and prefix in PrefixFormat output.
func WithPrefixSpace(space bool) prefixFormatOptionFn {
	return func(opt *prefixFormatOption) { opt.space = space }
}

// WithPrefixPrecision sets decimal places in PrefixFormat. Zero uses RoundMethod.
func WithPrefixPrecision(precision int) prefixFormatOptionFn {
	return func(opt *prefixFormatOption) { opt.precision = precision }
}

// WithRoundMethod sets how PrefixFormat rounds when precision is zero.
func WithRoundMethod(roundMethod RoundMethod) prefixFormatOptionFn {
	return func(opt *prefixFormatOption) { opt.roundMethod = roundMethod }
}

// WithRoundDifference sets the fractional threshold for RoundMethodDifference (0–1).
func WithRoundDifference(roundDifference float64) prefixFormatOptionFn {
	return func(opt *prefixFormatOption) { opt.roundDifference = roundDifference }
}

// PrefixFormat2 returns the formatted number and numeric prefix separately.
// Example: PrefixFormat2(1048576, IEC) // "1", "Mi"
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
		case IEC, ForceIEC:
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

			// Fractional part used by RoundMethodDifference.
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

// PrefixFormat converts val to a formatted "number prefix" string.
// Example: PrefixFormat(1048576, IEC) // "1 Mi", "1Mi", "1 M", etc. depending on options.
//
// Use PrefixFormat2 to obtain the number and prefix separately.
func PrefixFormat(val float64, prefixMode PrefixMode, prefixFormatOptionFns ...prefixFormatOptionFn) string {
	option := newPrefixFormatOption(prefixFormatOptionFns...)

	number, prefix := PrefixFormat2(val, prefixMode, prefixFormatOptionFns...)

	if prefix == "" {
		return number
	}

	format := "%s%s"
	if option.space {
		format = "%s %s"
	}

	return fmt.Sprintf(format, number, prefix)
}
