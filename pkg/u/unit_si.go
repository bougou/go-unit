package u

// 长度 (L, length)，基本单位：米 (meter)。
// SI 词头倍数用 Meter.Prefix(prefix)，例如 Meter.Prefix(Kilo) → km。

const (
	Meter LengthUnit = "meter" // m, 米

	// 科学

	Angstrom         LengthUnit = "angstrom"          // Å, 埃
	AstronomicalUnit LengthUnit = "astronomical_unit" // au, 天文单位
	LightYear        LengthUnit = "light_year"        // ly, 光年
	Parsec           LengthUnit = "parsec"            // pc, 秒差距

	// 英制与美制

	Inch         LengthUnit = "inch"          // in, 英寸
	Foot         LengthUnit = "foot"          // ft, 英尺
	Yard         LengthUnit = "yard"          // yd, 码
	Mile         LengthUnit = "mile"          // mi, 英里
	Mil          LengthUnit = "mil"           // mil, 密耳
	Hand         LengthUnit = "hand"          // hand, 手宽
	Fathom       LengthUnit = "fathom"        // ftm, 英寻
	Cable        LengthUnit = "cable"         // cable, 缆（链长单位）
	NauticalMile LengthUnit = "nautical_mile" // nmi, 海里
	Chain        LengthUnit = "chain"         // ch, 链
	Furlong      LengthUnit = "furlong"       // fur, 弗隆
	Rod          LengthUnit = "rod"           // rd, 杆
	League       LengthUnit = "league"        // lea, 里格

	// 排版

	Point LengthUnit = "point" // pt, 点
	Pica  LengthUnit = "pica"  // pc, 派卡

	// 市制

	Chi   LengthUnit = "chi"   // chi, 尺
	Cun   LengthUnit = "cun"   // cun, 寸
	Zhang LengthUnit = "zhang" // zhang, 丈
	Li    LengthUnit = "li"    // li, 里
)

// 质量 (M, mass)，基本单位：千克 (kilogram)。
// SI 词头挂在 Gram 上：Gram.Prefix(Milli) → mg；Gram.Prefix(Kilo) → Kilogram。
// Kilogram.Prefix(Milli) → Gram 是与克互转的捷径。

const (
	Gram     MassUnit = "gram"     // g, 克（词头附着根）
	Kilogram MassUnit = "kilogram" // kg, 千克（SI 基本单位）

	// 公制

	Tonne MassUnit = "tonne" // t, 吨

	// 英制与美制

	Grain     MassUnit = "grain"      // gr, 格令
	Dram      MassUnit = "dram"       // dr, 打兰
	Ounce     MassUnit = "ounce"      // oz, 盎司
	Pound     MassUnit = "pound"      // lb, 磅
	Stone     MassUnit = "stone"      // st, 英石
	ShortTon  MassUnit = "short_ton"  // ton, 短吨
	LongTon   MassUnit = "long_ton"   // long ton, 长吨
	TroyOunce MassUnit = "troy_ounce" // oz t, 金衡盎司
	TroyPound MassUnit = "troy_pound" // lb t, 金衡磅

	// 科学

	Carat  MassUnit = "carat"  // ct, 克拉
	Dalton MassUnit = "dalton" // Da, 道尔顿

	// 市制

	Jin   MassUnit = "jin"   // jin, 斤
	Liang MassUnit = "liang" // liang, 两
)

// 时间 (T, time)，基本单位：秒 (second)。
// SI 词头倍数用 Second.Prefix(prefix)，例如 Second.Prefix(Milli) → ms。

const (
	Second TimeUnit = "second" // s, 秒

	// 常用

	Minute    TimeUnit = "minute"    // min, 分
	Hour      TimeUnit = "hour"      // h, 时
	Day       TimeUnit = "day"       // d, 日
	Week      TimeUnit = "week"      // wk, 周
	Fortnight TimeUnit = "fortnight" // fn, 双周
	Month     TimeUnit = "month"     // mo, 月
	Year      TimeUnit = "year"      // yr, 年
)

// 电流 (I, electric current)，基本单位：安培 (ampere)。
// SI 词头倍数用 Ampere.Prefix(prefix)，例如 Ampere.Prefix(Micro) → μA。

const (
	Ampere CurrentUnit = "ampere" // A, 安培
	Biot   CurrentUnit = "biot"   // Bi, 毕奥
)

// 温度 (Θ, thermodynamic temperature)，基本单位：开尔文 (kelvin)。
// SI 词头倍数用 Kelvin.Prefix(prefix)，例如 Kelvin.Prefix(Milli) → mK。

const (
	Kelvin     TemperatureUnit = "kelvin"     // K, 开尔文
	Celsius    TemperatureUnit = "celsius"    // °C, 摄氏度
	Fahrenheit TemperatureUnit = "fahrenheit" // °F, 华氏度
	Rankine    TemperatureUnit = "rankine"    // °R, 兰氏度
)

// 物质的量 (N, amount of substance)，基本单位：摩尔 (mole)。
// SI 词头倍数用 Mole.Prefix(prefix)，例如 Mole.Prefix(Milli) → mmol。

const (
	Mole      AmountUnit = "mole"       // mol, 摩尔
	PoundMole AmountUnit = "pound_mole" // lbmol, 磅摩尔
)

// 发光强度 (J, luminous intensity)，基本单位：坎德拉 (candela)。
// SI 词头倍数用 Candela.Prefix(prefix)，例如 Candela.Prefix(Milli) → mcd。

const (
	Candela LuminousUnit = "candela" // cd, 坎德拉
)

var siUnitDefs = []unitDef{
	// 长度 — Meter + non-SI
	{DimLength, Unit(Meter), "m", "meter", 1, 0},
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

	// 质量 — Kilogram (base) + Gram (prefix root) + non-SI
	{DimMass, Unit(Kilogram), "kg", "kilogram", 1, 0},
	{DimMass, Unit(Gram), "g", "gram", 1e-3, 0},
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

	// 时间 — Second + non-SI
	{DimTime, Unit(Second), "s", "second", 1, 0},
	{DimTime, Unit(Minute), "min", "minute", 60, 0},
	{DimTime, Unit(Hour), "h", "hour", 3600, 0},
	{DimTime, Unit(Day), "d", "day", 86400, 0},
	{DimTime, Unit(Week), "wk", "week", 604800, 0},
	{DimTime, Unit(Fortnight), "fn", "fortnight", 1209600, 0},
	{DimTime, Unit(Month), "mo", "month", 2629800, 0},
	{DimTime, Unit(Year), "yr", "year", 31557600, 0},

	// 电流 — Ampere + non-SI
	{DimCurrent, Unit(Ampere), "A", "ampere", 1, 0},
	{DimCurrent, Unit(Biot), "Bi", "biot", 10, 0},

	// 温度 — Kelvin + affine / Rankine
	{DimTemperature, Unit(Kelvin), "K", "kelvin", 1, 0},
	{DimTemperature, Unit(Celsius), "°C", "celsius", 1, 273.15},
	{DimTemperature, Unit(Fahrenheit), "°F", "fahrenheit", 5.0 / 9.0, 255.3722222222222},
	{DimTemperature, Unit(Rankine), "°R", "rankine", 5.0 / 9.0, 0},

	// 物质的量 — Mole + non-SI
	{DimAmount, Unit(Mole), "mol", "mole", 1, 0},
	{DimAmount, Unit(PoundMole), "lbmol", "pound mole", 453.59237, 0},

	// 发光强度 — Candela
	{DimLuminous, Unit(Candela), "cd", "candela", 1, 0},
}

func init() {
	for i := range siUnitDefs {
		registerUnit(&siUnitDefs[i])
	}
	if err := validateRegistry(); err != nil {
		panic(err)
	}
}
