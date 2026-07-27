package u

// see: http://en.wikipedia.org/wiki/Binary_prefix
//
// Largest positive integer exponent n such that base^n still fits in the type
// (finite positive range; for float64: finite IEEE-754 value, not Inf):
//
//	uint64  max = 2⁶⁴−1 ≈ 1.84×10¹⁹
//	        10ⁿ → n ≤ 19    (10¹⁹ fits; 10²⁰ overflows)
//	        2ⁿ  → n ≤ 63    (2⁶³ fits; 2⁶⁴ overflows)
//	        eⁿ  → n ≤ 44    (e⁴⁴ ≈ 1.29×10¹⁹; e⁴⁵ overflows)
//
//	int64   max = 2⁶³−1 ≈ 9.22×10¹⁸
//	        10ⁿ → n ≤ 18
//	        2ⁿ  → n ≤ 62
//	        eⁿ  → n ≤ 43    (e⁴³ ≈ 4.73×10¹⁸; e⁴⁴ overflows)
//
//	float64 max ≈ 1.80×10³⁰⁸  (≈ (2−2⁻⁵²)×2¹⁰²³)
//	        10ⁿ → n ≤ 308
//	        2ⁿ  → n ≤ 1023
//	        eⁿ  → n ≤ 709   (e⁷⁰⁹ ≈ 8.22×10³⁰⁷; e⁷¹⁰ → +Inf)

const (
	B = 1 // byte (字节)

	// Decimal SI (国际单位制) byte scale factors (powers of 1000).

	KB = 1000      // kilo (千字节，10³)
	MB = 1000 * KB // mega (兆字节，10⁶)
	GB = 1000 * MB // giga (吉字节，10⁹)
	TB = 1000 * GB // tera (太字节，10¹²)
	PB = 1000 * TB // peta (拍字节，10¹⁵)
	EB = 1000 * PB // exa (艾字节，10¹⁸)
	ZB = 1000 * EB // zetta (泽字节，10²¹)
	YB = 1000 * ZB // yotta (尧字节，10²⁴)
	RB = 1000 * YB // ronna (容字节，10²⁷)
	QB = 1000 * RB // quetta (昆字节，10³⁰)

	// Binary IEC (国际电工委员会) byte scale factors (powers of 1024).

	KiB = 1024       // kibi (千比字节，2¹⁰)
	MiB = 1024 * KiB // mebi (兆比字节，2²⁰)
	GiB = 1024 * MiB // gibi (吉比字节，2³⁰)
	TiB = 1024 * GiB // tebi (太比字节，2⁴⁰)
	PiB = 1024 * TiB // pebi (拍比字节，2⁵⁰)
	EiB = 1024 * PiB // exbi (艾比字节，2⁶⁰)
	ZiB = 1024 * EiB // zebi (泽比字节，2⁷⁰)
	YiB = 1024 * ZiB // yobi (尧比字节，2⁸⁰)
	RiB = 1024 * YiB // robi (容比字节，2⁹⁰)
	QiB = 1024 * RiB // quebi (昆比字节，2¹⁰⁰)
)
