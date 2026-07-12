package u

// RoundMethod selects how PrefixFormat rounds fractional values.
type RoundMethod int

const (
	RoundMethodFloor RoundMethod = 0 // floor toward negative infinity (向下取整)
	RoundMethodRound RoundMethod = 1 // round half away from zero (四舍五入)
	RoundMethodCeil  RoundMethod = 2 // ceil toward positive infinity (向上取整)

	// RoundMethodDifference rounds based on the fractional part versus a threshold.
	// For positive values: ceil when fraction >= threshold, else floor.
	// For negative values: floor when fraction >= threshold, else ceil.
	// A threshold of 0.5 behaves like RoundMethodRound.
	RoundMethodDifference RoundMethod = 3
)

// String returns the name of the rounding method.
func (roundMethod RoundMethod) String() string {
	switch roundMethod {
	case RoundMethodFloor:
		return "floor"

	case RoundMethodRound:
		return "round"

	case RoundMethodCeil:
		return "ceil"

	case RoundMethodDifference:
		return "difference"

	default:
		return "<unknown>"
	}
}
