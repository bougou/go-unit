package unit

type Unit string

const (
	// 长度
	// 质量
	// 时间
	// 电流
	// 热力学温度
	// 发光强度
	// 物理的量

	// 7个严格定义的基本单位是：长度（米）、质量（千克）、时间（秒）、电流（安培）、热力学温度（开尔文）、物质的量（摩尔）和发光强度（坎德拉）。

	// 时间
	TimeSeconds Unit = "second"

	PPM Unit = "ppm"

	SizeByte Unit = "byte"

	MathRatio      Unit = "ratio"      // 比率，除以 % 后的值（0-1） Better use ratio than percentage as the metric value
	MathPercentage Unit = "percentage" // 百分比，去掉 %之后的数字（0-100）
)

type UnitInfo struct {
	Unit Unit   // 单位的符号
	Sign string // 单位的符号

	Name string // 单位的名称 (秒)

	PhySign string // 物理量符号
}
type SensorUnitType uint8

const (
	SensorUnitType_Unspecified SensorUnitType = 0 // unspecified

	SensorUnitType_DegressC SensorUnitType = 1  // degrees C, Celsius, 摄氏度 ℃
	SensorUnitType_DegreesF SensorUnitType = 2  // degrees F, Fahrenheit, 华氏度
	SensorUnitType_DegreesK SensorUnitType = 3  // degrees K, Kelvins, 开尔文
	SensorUnitType_Volts    SensorUnitType = 4  // Volts, 伏特（电压单位）
	SensorUnitType_Amps     SensorUnitType = 5  // Amps, 安培数
	SensorUnitType_Watts    SensorUnitType = 6  // Watts, 瓦特（功率单位）
	SensorUnitType_Joules   SensorUnitType = 7  // Joules, 焦耳
	SensorUnitType_Coulombs SensorUnitType = 8  // Coulombs, 库伦
	SensorUnitType_VA       SensorUnitType = 9  // VA, 伏安
	SensorUnitType_Nits     SensorUnitType = 10 // Nits, 尼特（光度单位）
	SensorUnitType_Lumen    SensorUnitType = 11 // lumen, 流明（光通量单位）
	SensorUnitType_Lux      SensorUnitType = 12 // lux, 勒克斯（照明单位）
	SensorUnitType_Candela  SensorUnitType = 13 // Candela, 坎, 坎德拉（发光强度单位）
	SensorUnitType_KPa      SensorUnitType = 14 // kPa kilopascal, 千帕, 千帕斯卡
	SensorUnitType_PSI      SensorUnitType = 15 // PSI
	SensorUnitType_Newton   SensorUnitType = 16 // Newton, 牛顿（力的单位）
	SensorUnitType_CFM      SensorUnitType = 17 // CFM, 风量, cubic feet per minute (cu ft/min)
	SensorUnitType_RPM      SensorUnitType = 18 // RPM, 每分钟转数, Revolutions per minute, is the number of turns in one minute
	SensorUnitType_Hz       SensorUnitType = 19 // Hz, 赫兹

	SensorUnitType_MicroSecond SensorUnitType = 20 // microsecond, 微秒
	SensorUnitType_MilliSecond SensorUnitType = 21 // millisecond, 毫秒
	SensorUnitType_Second      SensorUnitType = 22 // second, 秒
	SensorUnitType_Minute      SensorUnitType = 23 // minute, 分
	SensorUnitType_Hour        SensorUnitType = 24 // hour, 时
	SensorUnitType_Day         SensorUnitType = 25 // day, 日
	SensorUnitType_Week        SensorUnitType = 26 // week, 周

	SensorUnitType_Mil SensorUnitType = 27 // mil, 毫升；密耳（千分之一寸）

	SensorUnitType_Inches             SensorUnitType = 28 // inches, 英寸（inch的复数）
	SensorUnitType_Fleet              SensorUnitType = 29 // feet
	SensorUnitType_CuIn               SensorUnitType = 30 // cu in, 立方英寸（cubic inch）
	SensorUnitType_CuFleet            SensorUnitType = 31 // cu feet
	SensorUnitType_MM                 SensorUnitType = 32 // mm, 毫米（millimeter）
	SensorUnitType_CM                 SensorUnitType = 33 // cm, 厘米（centimeter）
	SensorUnitType_M                  SensorUnitType = 34 // m, 米
	SensorUnitType_CuCM               SensorUnitType = 35 // cu cm
	SensorUnitType_Cum                SensorUnitType = 36 // cum
	SensorUnitType_Liters             SensorUnitType = 37 // liters, 公升（容量单位）
	SensorUnitType_FluidOunce         SensorUnitType = 38 // fluid ounce, 液盎司（液体容量单位, 等于 fluidounce）
	SensorUnitType_Radians            SensorUnitType = 39 // radians, 弧度（radian的复数）
	SensorUnitType_vSteradians        SensorUnitType = 40 // steradians, 球面度, 立体弧度（立体角国际单位制, 等于 sterad）
	SensorUnitType_Revolutions        SensorUnitType = 41 // revolutions, 转数（revolution的复数形式）
	SensorUnitType_Cycles             SensorUnitType = 42 // cycles, 周期, 圈
	SensorUnitType_Gravities          SensorUnitType = 43 // gravities, 重力
	SensorUnitType_Ounce              SensorUnitType = 44 // ounce, 盎司
	SensorUnitType_Pound              SensorUnitType = 45 // pound, 英镑
	SensorUnitType_FootPound          SensorUnitType = 46 // ft-lb, 英尺-磅（foot pound）
	SensorUnitType_OzIn               SensorUnitType = 47 // oz-in, 扭力；盎司英寸
	SensorUnitType_Gauss              SensorUnitType = 48 // gauss, 高斯（磁感应或磁场的单位）
	SensorUnitType_Gilberts           SensorUnitType = 49 // gilberts, 吉伯（磁通量的单位）
	SensorUnitType_Henry              SensorUnitType = 50 // henry, 亨利（电感单位）
	SensorUnitType_MilliHenry         SensorUnitType = 51 // millihenry, 毫亨（利）（电感单位）
	SensorUnitType_Farad              SensorUnitType = 52 // farad, 法拉（电容单位）
	SensorUnitType_MicroFarad         SensorUnitType = 53 // microfarad, 微法拉（电容量的实用单位）
	SensorUnitType_Ohms               SensorUnitType = 54 // ohms, 欧姆（Ohm） ：电阻的量度单位, 欧姆值越大, 电阻越大
	SensorUnitType_Siemens            SensorUnitType = 55 // siemens, 西门子, 电导单位
	SensorUnitType_Mole               SensorUnitType = 56 // mole, 摩尔 [化学] 克分子（等于mole）
	SensorUnitType_Becquerel          SensorUnitType = 57 // becquerel, 贝可（放射性活度单位）
	SensorUnitType_PPM                SensorUnitType = 58 // PPM (parts/million), 百万分率, 百万分之…（parts per million）
	SensorUnitType_Reserved           SensorUnitType = 59 // reserved
	SensorUnitType_Decibels           SensorUnitType = 60 // Decibels, 分贝（声音强度单位, decibel的复数）
	SensorUnitType_DbA                SensorUnitType = 61 // DbA, dBA is often used to specify the loudness of the fan used to cool the microprocessor and associated components. Typical dBA ratings are in the neighborhood of 25 dBA, representing 25 A-weighted decibels above the threshold of hearing. This is approximately the loudness of a person whispering in a quiet room.
	SensorUnitType_DbC                SensorUnitType = 62 // DbC
	SensorUnitType_Gray               SensorUnitType = 63 // gray, 核吸收剂量(Gy)
	SensorUnitType_Severt             SensorUnitType = 64 // sievert, 希沃特（辐射效果单位, 简称希）
	SensorUnitType_ColorTempDegK      SensorUnitType = 65 // color temp deg K, 色温
	SensorUnitType_Bit                SensorUnitType = 66 // bit, 比特（二进位制信息单位）
	SensorUnitType_Kilobit            SensorUnitType = 67 // kilobit, 千比特
	SensorUnitType_Megabit            SensorUnitType = 68 // megabit, 兆比特
	SensorUnitType_Gigabit            SensorUnitType = 69 // gigabit, 吉比特
	SensorUnitType_Byte               SensorUnitType = 70 // byte, 字节
	SensorUnitType_Kilobyte           SensorUnitType = 71 // kilobyte, 千字节
	SensorUnitType_Megabyte           SensorUnitType = 72 // megabyte, 兆字节
	SensorUnitType_Gigabyte           SensorUnitType = 73 // gigabyte, 吉字节
	SensorUnitType_Word               SensorUnitType = 74 // word (data), 字
	SensorUnitType_DWord              SensorUnitType = 75 // dword, 双字
	SensorUnitType_QWord              SensorUnitType = 76 // qword, 四字
	SensorUnitType_Line               SensorUnitType = 77 // line (re. mem. line)
	SensorUnitType_Hit                SensorUnitType = 78 // hit, 命中
	SensorUnitType_Miss               SensorUnitType = 79 // miss, 未击中, 未命中
	SensorUnitType_Retry              SensorUnitType = 80 // retry, 重试（次数）
	SensorUnitType_Reset              SensorUnitType = 81 // reset, 重置（次数）
	SensorUnitType_Overrun            SensorUnitType = 82 // overrun) / overflow 满载, 溢出（次数）
	SensorUnitType_Underrun           SensorUnitType = 83 // underrun 欠载
	SensorUnitType_Collision          SensorUnitType = 84 // collision, 冲突
	SensorUnitType_Packet             SensorUnitType = 85 // packets, 包, 数据包
	SensorUnitType_Message            SensorUnitType = 86 // messages, 消息
	SensorUnitType_Characters         SensorUnitType = 87 // characters, 字符
	SensorUnitType_Error              SensorUnitType = 88 // error, 错误
	SensorUnitType_CorrectableError   SensorUnitType = 89 // correctable error 可校正错误
	SensorUnitType_UncorrectableError SensorUnitType = 90 // uncorrectable error 不可校正错误
	SensorUnitType_FatalError         SensorUnitType = 91 // fatal error, 致命错误, 不可恢复的错误
	SensorUnitType_Grams              SensorUnitType = 92 // grams, 克（gram的复数形式）
)
