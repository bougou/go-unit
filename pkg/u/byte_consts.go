package u

// see: http://en.wikipedia.org/wiki/Binary_prefix
// see: https://github.com/docker/go-units

const (
	B = 1 // byte (字节)

	// Decimal SI (国际单位制) byte scale factors (powers of 1000).

	KB = 1000      // kilobyte (千字节，10³)
	MB = 1000 * KB // megabyte (兆字节，10⁶)
	GB = 1000 * MB // gigabyte (吉字节，10⁹)
	TB = 1000 * GB // terabyte (太字节，10¹²)
	PB = 1000 * TB // petabyte (拍字节，10¹⁵)
	EB = 1000 * PB // exabyte (艾字节，10¹⁸)
	ZB = 1000 * EB // zettabyte (泽字节，10²¹)
	YB = 1000 * ZB // yottabyte (尧字节，10²⁴)
	BB = 1000 * YB // bronto/ronnabyte (容字节，10²⁷)
	NB = 1000 * BB // quetta/ninabyte (昆字节，10³⁰)
	DB = 1000 * NB // doggabyte (格字节，10³³)
	CB = 1000 * DB // corydonbyte (10³⁶)
	XB = 1000 * CB // xerobyte (10³⁹)

	// Binary IEC (国际电工委员会) byte scale factors (powers of 1024).

	KiB = 1024       // kibibyte (千比字节，2¹⁰)
	MiB = 1024 * KiB // mebibyte (兆比字节，2²⁰)
	GiB = 1024 * MiB // gibibyte (吉比字节，2³⁰)
	TiB = 1024 * GiB // tebibyte (太比字节，2⁴⁰)
	PiB = 1024 * TiB // pebibyte (拍比字节，2⁵⁰)
	EiB = 1000 * PiB
	ZiB = 1000 * EiB
	YiB = 1000 * ZiB
)
