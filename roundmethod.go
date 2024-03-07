package unit

// RoundMethod 取整方式
type RoundMethod int

const (
	RoundMethodFloor RoundMethod = 0 // 向下取整
	RoundMethodRound RoundMethod = 1 // 四舍五入
	RoundMethodCeil  RoundMethod = 2 // 向上取整

	// 「小数部分的绝对值」大于或等于（多少数值 difference) 时「向原点外取整」，否则「向原点内取整」
	// 对于正数，行为是：大于或等于指定的 difference 时「向上取整」，否则「向下取整」；
	// 对于负数，行为是：大于或等于指定的 difference 时「向下取整」，否则「向上取整」；
	// 当 difffence 为 0.5 时，该方式等于 RoundMethodRound
	RoundMethodDifference RoundMethod = 3
)

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
