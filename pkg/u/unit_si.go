package u

// 长度 (L, length)，基本单位：米 (meter)。

const (
	// SI 词头（国际单位制十进制前缀）

	Picometer  LengthUnit = "picometer"  // 皮米
	Nanometer  LengthUnit = "nanometer"  // 纳米
	Micrometer LengthUnit = "micrometer" // 微米
	Millimeter LengthUnit = "millimeter" // 毫米
	Centimeter LengthUnit = "centimeter" // 厘米
	Decimeter  LengthUnit = "decimeter"  // 分米
	Meter      LengthUnit = "meter"      // 米
	Decameter  LengthUnit = "decameter"  // 十米
	Hectometer LengthUnit = "hectometer" // 百米
	Kilometer  LengthUnit = "kilometer"  // 千米
	Megameter  LengthUnit = "megameter"  // 兆米
	Gigameter  LengthUnit = "gigameter"  // 吉米
	Terameter  LengthUnit = "terameter"  // 太米
	Petameter  LengthUnit = "petameter"  // 拍米

	// 科学

	Angstrom         LengthUnit = "angstrom"          // 埃
	AstronomicalUnit LengthUnit = "astronomical_unit" // 天文单位
	LightYear        LengthUnit = "light_year"        // 光年
	Parsec           LengthUnit = "parsec"            // 秒差距

	// 英制与美制

	Inch         LengthUnit = "inch"          // 英寸
	Foot         LengthUnit = "foot"          // 英尺
	Yard         LengthUnit = "yard"          // 码
	Mile         LengthUnit = "mile"          // 英里
	Mil          LengthUnit = "mil"           // 密耳
	Hand         LengthUnit = "hand"          // 手宽
	Fathom       LengthUnit = "fathom"        // 英寻
	Cable        LengthUnit = "cable"         // 缆（链长单位）
	NauticalMile LengthUnit = "nautical_mile" // 海里
	Chain        LengthUnit = "chain"         // 链
	Furlong      LengthUnit = "furlong"       // 弗隆
	Rod          LengthUnit = "rod"           // 杆
	League       LengthUnit = "league"        // 里格

	// 排版

	Point LengthUnit = "point" // 点
	Pica  LengthUnit = "pica"  // 派卡

	// 市制

	Chi   LengthUnit = "chi"   // 尺
	Cun   LengthUnit = "cun"   // 寸
	Zhang LengthUnit = "zhang" // 丈
	Li    LengthUnit = "li"    // 里
)

// 质量 (M, mass)，基本单位：千克 (kilogram)。

const (
	// SI 词头（国际单位制十进制前缀）

	Nanogram  MassUnit = "nanogram"  // 纳克
	Microgram MassUnit = "microgram" // 微克
	Milligram MassUnit = "milligram" // 毫克
	Gram      MassUnit = "gram"      // 克
	Kilogram  MassUnit = "kilogram"  // 千克
	Megagram  MassUnit = "megagram"  // 兆克
	Gigagram  MassUnit = "gigagram"  // 吉克
	Teragram  MassUnit = "teragram"  // 太克
	Petagram  MassUnit = "petagram"  // 拍克

	// 公制

	Tonne MassUnit = "tonne" // 吨

	// 英制与美制

	Grain     MassUnit = "grain"      // 格令
	Dram      MassUnit = "dram"       // 打兰
	Ounce     MassUnit = "ounce"      // 盎司
	Pound     MassUnit = "pound"      // 磅
	Stone     MassUnit = "stone"      // 英石
	ShortTon  MassUnit = "short_ton"  // 短吨
	LongTon   MassUnit = "long_ton"   // 长吨
	TroyOunce MassUnit = "troy_ounce" // 金衡盎司
	TroyPound MassUnit = "troy_pound" // 金衡磅

	// 科学

	Carat  MassUnit = "carat"  // 克拉
	Dalton MassUnit = "dalton" // 道尔顿

	// 市制

	Jin   MassUnit = "jin"   // 斤
	Liang MassUnit = "liang" // 两
)

// 时间 (T, time)，基本单位：秒 (second)。

const (
	// SI 词头（国际单位制十进制前缀）

	Femtosecond TimeUnit = "femtosecond" // 飞秒
	Picosecond  TimeUnit = "picosecond"  // 皮秒
	Nanosecond  TimeUnit = "nanosecond"  // 纳秒
	Microsecond TimeUnit = "microsecond" // 微秒
	Millisecond TimeUnit = "millisecond" // 毫秒
	Centisecond TimeUnit = "centisecond" // 厘秒
	Decisecond  TimeUnit = "decisecond"  // 分秒
	Second      TimeUnit = "second"      // 秒
	Decasecond  TimeUnit = "decasecond"  // 十秒
	Hectosecond TimeUnit = "hectosecond" // 百秒
	Kilosecond  TimeUnit = "kilosecond"  // 千秒
	Megasecond  TimeUnit = "megasecond"  // 兆秒
	Gigasecond  TimeUnit = "gigasecond"  // 吉秒

	// 常用

	Minute    TimeUnit = "minute"    // 分
	Hour      TimeUnit = "hour"      // 时
	Day       TimeUnit = "day"       // 日
	Week      TimeUnit = "week"      // 周
	Fortnight TimeUnit = "fortnight" // 双周
	Month     TimeUnit = "month"     // 月
	Year      TimeUnit = "year"      // 年
)

// 电流 (I, electric current)，基本单位：安培 (ampere)。

const (
	Femtoampere CurrentUnit = "femtoampere" // 飞安
	Picoampere  CurrentUnit = "picoampere"  // 皮安
	Nanoampere  CurrentUnit = "nanoampere"  // 纳安
	Microampere CurrentUnit = "microampere" // 微安
	Milliampere CurrentUnit = "milliampere" // 毫安
	Ampere      CurrentUnit = "ampere"      // 安培
	Kiloampere  CurrentUnit = "kiloampere"  // 千安
	Megaampere  CurrentUnit = "megaampere"  // 兆安
	Gigaampere  CurrentUnit = "gigaampere"  // 吉安
	Biot        CurrentUnit = "biot"        // 毕奥
)

// 温度 (Θ, thermodynamic temperature)，基本单位：开尔文 (kelvin)。

const (
	Microkelvin TemperatureUnit = "microkelvin" // 微开
	Millikelvin TemperatureUnit = "millikelvin" // 毫开
	Kelvin      TemperatureUnit = "kelvin"      // 开尔文
	Celsius     TemperatureUnit = "celsius"     // 摄氏度
	Fahrenheit  TemperatureUnit = "fahrenheit"  // 华氏度
	Rankine     TemperatureUnit = "rankine"     // 兰氏度
)

// 物质的量 (N, amount of substance)，基本单位：摩尔 (mole)。

const (
	Picomole  AmountUnit = "picomole"   // 皮摩
	Nanomole  AmountUnit = "nanomole"   // 纳摩
	Micromole AmountUnit = "micromole"  // 微摩
	Millimole AmountUnit = "millimole"  // 毫摩
	Mole      AmountUnit = "mole"       // 摩尔
	Kilomole  AmountUnit = "kilomole"   // 千摩
	Megamole  AmountUnit = "megamole"   // 兆摩
	PoundMole AmountUnit = "pound_mole" // 磅摩尔
)

// 发光强度 (J, luminous intensity)，基本单位：坎德拉 (candela)。

const (
	Millicandela LuminousUnit = "millicandela" // 毫坎
	Candela      LuminousUnit = "candela"      // 坎德拉
	Kilocandela  LuminousUnit = "kilocandela"  // 千坎
)

var siUnitDefs = []unitDef{
	// 长度
	{DimLength, Unit(Picometer), "pm", "picometer", 1e-12, 0},
	{DimLength, Unit(Nanometer), "nm", "nanometer", 1e-9, 0},
	{DimLength, Unit(Micrometer), "μm", "micrometer", 1e-6, 0},
	{DimLength, Unit(Millimeter), "mm", "millimeter", 1e-3, 0},
	{DimLength, Unit(Centimeter), "cm", "centimeter", 1e-2, 0},
	{DimLength, Unit(Decimeter), "dm", "decimeter", 1e-1, 0},
	{DimLength, Unit(Meter), "m", "meter", 1, 0},
	{DimLength, Unit(Decameter), "dam", "decameter", 10, 0},
	{DimLength, Unit(Hectometer), "hm", "hectometer", 100, 0},
	{DimLength, Unit(Kilometer), "km", "kilometer", 1e3, 0},
	{DimLength, Unit(Megameter), "Mm", "megameter", 1e6, 0},
	{DimLength, Unit(Gigameter), "Gm", "gigameter", 1e9, 0},
	{DimLength, Unit(Terameter), "Tm", "terameter", 1e12, 0},
	{DimLength, Unit(Petameter), "Pm", "petameter", 1e15, 0},
	{DimLength, Unit(Angstrom), "Å", "angstrom", 1e-10, 0},
	{DimLength, Unit(AstronomicalUnit), "au", "astronomical unit", 149597870700, 0},
	{DimLength, Unit(LightYear), "ly", "light year", 9.4607304725808e15, 0},
	{DimLength, Unit(Parsec), "pc", "parsec", 3.08567758149137e16, 0},
	{DimLength, Unit(Inch), "in", "inch", 0.0254, 0},
	{DimLength, Unit(Foot), "ft", "foot", 0.3048, 0},
	{DimLength, Unit(Yard), "yd", "yard", 0.9144, 0},
	{DimLength, Unit(Mile), "mi", "mile", 1609.344, 0},
	{DimLength, Unit(Mil), "mil", "mil", 2.54e-5, 0},
	{DimLength, Unit(Hand), "hand", "hand", 0.1016, 0},
	{DimLength, Unit(Fathom), "ftm", "fathom", 1.8288, 0},
	{DimLength, Unit(Cable), "cable", "cable", 185.2, 0},
	{DimLength, Unit(NauticalMile), "nmi", "nautical mile", 1852, 0},
	{DimLength, Unit(Chain), "ch", "chain", 20.1168, 0},
	{DimLength, Unit(Furlong), "fur", "furlong", 201.168, 0},
	{DimLength, Unit(Rod), "rd", "rod", 5.0292, 0},
	{DimLength, Unit(League), "lea", "league", 4828.032, 0},
	{DimLength, Unit(Point), "pt", "point", 0.0254 / 72, 0},
	{DimLength, Unit(Pica), "pc", "pica", 0.0254 / 6, 0},
	{DimLength, Unit(Chi), "chi", "chi", 1.0 / 3, 0},
	{DimLength, Unit(Cun), "cun", "cun", 1.0 / 30, 0},
	{DimLength, Unit(Zhang), "zhang", "zhang", 10.0 / 3, 0},
	{DimLength, Unit(Li), "li", "li", 500, 0},

	// 质量
	{DimMass, Unit(Nanogram), "ng", "nanogram", 1e-12, 0},
	{DimMass, Unit(Microgram), "μg", "microgram", 1e-9, 0},
	{DimMass, Unit(Milligram), "mg", "milligram", 1e-6, 0},
	{DimMass, Unit(Gram), "g", "gram", 1e-3, 0},
	{DimMass, Unit(Kilogram), "kg", "kilogram", 1, 0},
	{DimMass, Unit(Megagram), "Mg", "megagram", 1e3, 0},
	{DimMass, Unit(Gigagram), "Gg", "gigagram", 1e6, 0},
	{DimMass, Unit(Teragram), "Tg", "teragram", 1e9, 0},
	{DimMass, Unit(Petagram), "Pg", "petagram", 1e12, 0},
	{DimMass, Unit(Tonne), "t", "tonne", 1e3, 0},
	{DimMass, Unit(Grain), "gr", "grain", 6.479891e-5, 0},
	{DimMass, Unit(Dram), "dr", "dram", 1.7718451953125e-3, 0},
	{DimMass, Unit(Ounce), "oz", "ounce", 0.028349523125, 0},
	{DimMass, Unit(Pound), "lb", "pound", 0.45359237, 0},
	{DimMass, Unit(Stone), "st", "stone", 6.35029318, 0},
	{DimMass, Unit(ShortTon), "ton", "short ton", 907.18474, 0},
	{DimMass, Unit(LongTon), "long ton", "long ton", 1016.0469088, 0},
	{DimMass, Unit(TroyOunce), "oz t", "troy ounce", 0.0311034768, 0},
	{DimMass, Unit(TroyPound), "lb t", "troy pound", 0.3732417216, 0},
	{DimMass, Unit(Carat), "ct", "carat", 2e-4, 0},
	{DimMass, Unit(Dalton), "Da", "dalton", 1.66053906660e-27, 0},
	{DimMass, Unit(Jin), "jin", "jin", 0.5, 0},
	{DimMass, Unit(Liang), "liang", "liang", 0.05, 0},

	// 时间
	{DimTime, Unit(Femtosecond), "fs", "femtosecond", 1e-15, 0},
	{DimTime, Unit(Picosecond), "ps", "picosecond", 1e-12, 0},
	{DimTime, Unit(Nanosecond), "ns", "nanosecond", 1e-9, 0},
	{DimTime, Unit(Microsecond), "μs", "microsecond", 1e-6, 0},
	{DimTime, Unit(Millisecond), "ms", "millisecond", 1e-3, 0},
	{DimTime, Unit(Centisecond), "cs", "centisecond", 1e-2, 0},
	{DimTime, Unit(Decisecond), "ds", "decisecond", 1e-1, 0},
	{DimTime, Unit(Second), "s", "second", 1, 0},
	{DimTime, Unit(Decasecond), "das", "decasecond", 10, 0},
	{DimTime, Unit(Hectosecond), "hs", "hectosecond", 100, 0},
	{DimTime, Unit(Kilosecond), "ks", "kilosecond", 1e3, 0},
	{DimTime, Unit(Megasecond), "Ms", "megasecond", 1e6, 0},
	{DimTime, Unit(Gigasecond), "Gs", "gigasecond", 1e9, 0},
	{DimTime, Unit(Minute), "min", "minute", 60, 0},
	{DimTime, Unit(Hour), "h", "hour", 3600, 0},
	{DimTime, Unit(Day), "d", "day", 86400, 0},
	{DimTime, Unit(Week), "wk", "week", 604800, 0},
	{DimTime, Unit(Fortnight), "fn", "fortnight", 1209600, 0},
	{DimTime, Unit(Month), "mo", "month", 2629800, 0},
	{DimTime, Unit(Year), "yr", "year", 31557600, 0},

	// 电流
	{DimCurrent, Unit(Femtoampere), "fA", "femtoampere", 1e-15, 0},
	{DimCurrent, Unit(Picoampere), "pA", "picoampere", 1e-12, 0},
	{DimCurrent, Unit(Nanoampere), "nA", "nanoampere", 1e-9, 0},
	{DimCurrent, Unit(Microampere), "μA", "microampere", 1e-6, 0},
	{DimCurrent, Unit(Milliampere), "mA", "milliampere", 1e-3, 0},
	{DimCurrent, Unit(Ampere), "A", "ampere", 1, 0},
	{DimCurrent, Unit(Kiloampere), "kA", "kiloampere", 1e3, 0},
	{DimCurrent, Unit(Megaampere), "MA", "megaampere", 1e6, 0},
	{DimCurrent, Unit(Gigaampere), "GA", "gigaampere", 1e9, 0},
	{DimCurrent, Unit(Biot), "Bi", "biot", 10, 0},

	// 温度
	{DimTemperature, Unit(Microkelvin), "μK", "microkelvin", 1e-6, 0},
	{DimTemperature, Unit(Millikelvin), "mK", "millikelvin", 1e-3, 0},
	{DimTemperature, Unit(Kelvin), "K", "kelvin", 1, 0},
	{DimTemperature, Unit(Celsius), "°C", "celsius", 1, 273.15},
	{DimTemperature, Unit(Fahrenheit), "°F", "fahrenheit", 5.0 / 9.0, 255.3722222222222},
	{DimTemperature, Unit(Rankine), "°R", "rankine", 5.0 / 9.0, 0},

	// 物质的量
	{DimAmount, Unit(Picomole), "pmol", "picomole", 1e-12, 0},
	{DimAmount, Unit(Nanomole), "nmol", "nanomole", 1e-9, 0},
	{DimAmount, Unit(Micromole), "μmol", "micromole", 1e-6, 0},
	{DimAmount, Unit(Millimole), "mmol", "millimole", 1e-3, 0},
	{DimAmount, Unit(Mole), "mol", "mole", 1, 0},
	{DimAmount, Unit(Kilomole), "kmol", "kilomole", 1e3, 0},
	{DimAmount, Unit(Megamole), "Mmol", "megamole", 1e6, 0},
	{DimAmount, Unit(PoundMole), "lbmol", "pound mole", 453.59237, 0},

	// 发光强度
	{DimLuminous, Unit(Millicandela), "mcd", "millicandela", 1e-3, 0},
	{DimLuminous, Unit(Candela), "cd", "candela", 1, 0},
	{DimLuminous, Unit(Kilocandela), "kcd", "kilocandela", 1e3, 0},
}

func init() {
	for i := range siUnitDefs {
		registerUnit(&siUnitDefs[i])
	}
	if err := validateRegistry(); err != nil {
		panic(err)
	}
}
