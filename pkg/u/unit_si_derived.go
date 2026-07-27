package u

// SI (国际单位制) coherent derived dimensions (一贯导出量纲) and units with special names
// (专用名称，22 SI derived units), plus common electrical / magnetic convenience units.
// Degree Celsius (°C, 摄氏度) is defined as Celsius in unit_si.go (affine temperature 仿射温度，not DerivedUnit).
//
// Dim* comments: Chinese name is the physical quantity; text after "SI unit(s):" names the unit(s), not the dimension.

var (
	// DimAngle is plane/solid angle (平面角/立体角), dimensionless.
	// SI units: rad (radian, 弧度), sr (steradian, 球面度).
	DimAngle = DerivedDimension{}
	// DimFrequency is frequency (频率), T⁻¹; also used for radioactivity (放射性活度).
	// SI units: Hz (hertz, 赫兹), Bq (becquerel, 贝克勒尔).
	DimFrequency = DerivedDimension{T: -1}
	// DimForce is force (力), M¹L¹T⁻².
	// SI unit: N (newton, 牛顿).
	DimForce = DerivedDimension{M: 1, L: 1, T: -2}
	// DimPressure is pressure (压强), M¹L⁻¹T⁻².
	// SI unit: Pa (pascal, 帕斯卡).
	DimPressure = DerivedDimension{M: 1, L: -1, T: -2}
	// DimEnergy is energy (能量), M¹L²T⁻².
	// SI unit: J (joule, 焦耳).
	DimEnergy = DerivedDimension{M: 1, L: 2, T: -2}
	// DimPower is power (功率), M¹L²T⁻³; also apparent power (视在功率) and reactive power (无功功率).
	// SI unit: W (watt, 瓦特); also VA (volt-ampere, 伏安), var (var, 乏).
	DimPower = DerivedDimension{M: 1, L: 2, T: -3}
	// DimElectricCharge is electric charge (电荷), T¹I¹.
	// SI unit: C (coulomb, 库仑).
	DimElectricCharge = DerivedDimension{T: 1, I: 1}
	// DimElectricPotential is electric potential / voltage (电压/电势), M¹L²T⁻³I⁻¹.
	// SI unit: V (volt, 伏特).
	DimElectricPotential = DerivedDimension{M: 1, L: 2, T: -3, I: -1}
	// DimResistance is electric resistance (电阻), M¹L²T⁻³I⁻²; also impedance (阻抗) and reactance (电抗).
	// SI unit: Ω (ohm, 欧姆).
	DimResistance = DerivedDimension{M: 1, L: 2, T: -3, I: -2}
	// DimConductance is electric conductance (电导), M⁻¹L⁻²T³I²; also admittance (导纳) and susceptance (电纳).
	// SI unit: S (siemens, 西门子).
	DimConductance = DerivedDimension{M: -1, L: -2, T: 3, I: 2}
	// DimCapacitance is capacitance (电容), M⁻¹L⁻²T⁴I².
	// SI unit: F (farad, 法拉).
	DimCapacitance = DerivedDimension{M: -1, L: -2, T: 4, I: 2}
	// DimInductance is inductance (电感), M¹L²T⁻²I⁻².
	// SI unit: H (henry, 亨利).
	DimInductance = DerivedDimension{M: 1, L: 2, T: -2, I: -2}
	// DimMagneticFluxDensity is magnetic flux density (磁感应强度/磁通密度), M¹T⁻²I⁻¹.
	// SI unit: T (tesla, 特斯拉).
	DimMagneticFluxDensity = DerivedDimension{M: 1, T: -2, I: -1}
	// DimMagneticFlux is magnetic flux (磁通量), M¹L²T⁻²I⁻¹.
	// SI unit: Wb (weber, 韦伯).
	DimMagneticFlux = DerivedDimension{M: 1, L: 2, T: -2, I: -1}
	// DimElectricField is electric field strength (电场强度), M¹L¹T⁻³I⁻¹.
	// SI unit: V/m (volt per metre, 伏特每米).
	DimElectricField = DerivedDimension{M: 1, L: 1, T: -3, I: -1}
	// DimElectricDisplacement is electric displacement (电位移), L⁻²T¹I¹.
	// SI unit: C/m² (coulomb per square metre, 库仑每平方米).
	DimElectricDisplacement = DerivedDimension{L: -2, T: 1, I: 1}
	// DimElectricFlux is electric flux (电通量), M¹L³T⁻³I⁻¹.
	// SI unit: V·m (volt metre, 伏特米).
	DimElectricFlux = DerivedDimension{M: 1, L: 3, T: -3, I: -1}
	// DimCurrentDensity is electric current density (电流密度), L⁻²I¹.
	// SI unit: A/m² (ampere per square metre, 安培每平方米).
	DimCurrentDensity = DerivedDimension{L: -2, I: 1}
	// DimResistivity is electric resistivity (电阻率), M¹L³T⁻³I⁻².
	// SI unit: Ω·m (ohm metre, 欧姆米).
	DimResistivity = DerivedDimension{M: 1, L: 3, T: -3, I: -2}
	// DimConductivity is electric conductivity (电导率), M⁻¹L⁻³T³I².
	// SI unit: S/m (siemens per metre, 西门子每米).
	DimConductivity = DerivedDimension{M: -1, L: -3, T: 3, I: 2}
	// DimPermittivity is permittivity (介电常数/电容率), M⁻¹L⁻³T⁴I².
	// SI unit: F/m (farad per metre, 法拉每米).
	DimPermittivity = DerivedDimension{M: -1, L: -3, T: 4, I: 2}
	// DimPermeability is magnetic permeability (磁导率), M¹L¹T⁻²I⁻².
	// SI unit: H/m (henry per metre, 亨利每米).
	DimPermeability = DerivedDimension{M: 1, L: 1, T: -2, I: -2}
	// DimMagneticFieldStrength is magnetic field strength (磁场强度), L⁻¹I¹.
	// SI unit: A/m (ampere per metre, 安培每米).
	DimMagneticFieldStrength = DerivedDimension{L: -1, I: 1}
	// DimReluctance is magnetic reluctance (磁阻), M⁻¹L⁻²T²I².
	// SI unit: H^-1 (reciprocal henry, 每亨利); also A/Wb.
	DimReluctance = DerivedDimension{M: -1, L: -2, T: 2, I: 2}
	// DimLuminousFlux is luminous flux (光通量), J¹.
	// SI unit: lm (lumen, 流明).
	DimLuminousFlux = DerivedDimension{J: 1}
	// DimIlluminance is illuminance (照度), J¹L⁻².
	// SI unit: lx (lux, 勒克斯).
	DimIlluminance = DerivedDimension{J: 1, L: -2}
	// DimAbsorbedDose is absorbed dose (吸收剂量), L²T⁻²; also dose equivalent (剂量当量).
	// SI units: Gy (gray, 戈瑞), Sv (sievert, 希沃特).
	DimAbsorbedDose = DerivedDimension{L: 2, T: -2}
	// DimCatalyticActivity is catalytic activity (催化活度), N¹T⁻¹.
	// SI unit: kat (katal, 开特).
	DimCatalyticActivity = DerivedDimension{N: 1, T: -1}
	// DimArea is area (面积), L².
	// SI unit: m² (square metre, 平方米).
	DimArea = DerivedDimension{L: 2}
	// DimSpeed is speed (速度), L¹T⁻¹.
	// SI unit: m/s (metre per second, 米每秒); coherent but without an SI special name.
	DimSpeed = DerivedDimension{L: 1, T: -1}
)

// Radian (rad, 弧度) is the SI unit of plane angle（平面角）.
var Radian *DerivedUnit

// Steradian (sr, 球面度) is the SI unit of solid angle（立体角）.
var Steradian *DerivedUnit

// Hertz (Hz, 赫兹) is the SI unit of frequency（频率）.
var Hertz *DerivedUnit

// Newton (N, 牛顿) is the SI unit of force（力）.
var Newton *DerivedUnit

// Pascal (Pa, 帕斯卡) is the SI unit of pressure（压强）.
var Pascal *DerivedUnit

// Joule (J, 焦耳) is the SI unit of energy（能量）.
var Joule *DerivedUnit

// WattHour (W·h, 瓦时) is a convenience unit of energy（能量）; 1 W·h = 3600 J.
// Use WattHour.Prefix(Kilo) for kW·h.
var WattHour *DerivedUnit

// Watt (W, 瓦特) is the SI unit of power（功率）.
var Watt *DerivedUnit

// VoltAmpere (VA, 伏安) is the unit of apparent power（视在功率）; same dimension as Watt.
var VoltAmpere *DerivedUnit

// Var (var, 乏) is the unit of reactive power（无功功率）; same dimension as Watt.
var Var *DerivedUnit

// HorsePower (hp, 机械马力) is mechanical (imperial) horsepower（功率）.
// 1 hp = 745.69987158227022 W (≈ 550 ft·lbf/s); commonly rounded to 745.7 W.
var HorsePower *DerivedUnit

// MetricHorsepower (PS, 公制马力) is metric horsepower（功率）; also cv / ch.
// 1 PS = 735.49875 W (75 kgf·m/s); commonly rounded to 735.5 W.
var MetricHorsepower *DerivedUnit

// ElectricalHorsePower (hp(E), 电气马力) is electrical horsepower（功率）.
// 1 hp(E) = 746 W (exact), used for electric motors.
var ElectricalHorsePower *DerivedUnit

// Coulomb (C, 库仑) is the SI unit of electric charge（电荷）.
var Coulomb *DerivedUnit

// AmpereHour (A·h, 安时) is a convenience unit of electric charge（电荷）; 1 A·h = 3600 C.
// Use AmpereHour.Prefix(Kilo) for kA·h.
var AmpereHour *DerivedUnit

// Volt (V, 伏特) is the SI unit of electric potential / voltage（电压/电势）.
var Volt *DerivedUnit

// Ohm (Ω, 欧姆) is the SI unit of electrical resistance（电阻）.
var Ohm *DerivedUnit

// Impedance is an alias for Ohm; impedance（阻抗）.
var Impedance *DerivedUnit

// Reactance is an alias for Ohm; reactance（电抗）.
var Reactance *DerivedUnit

// Siemens (S, 西门子) is the SI unit of electrical conductance（电导）.
var Siemens *DerivedUnit

// Admittance is an alias for Siemens; admittance（导纳）.
var Admittance *DerivedUnit

// Susceptance is an alias for Siemens; susceptance（电纳）.
var Susceptance *DerivedUnit

// Farad (F, 法拉) is the SI unit of capacitance（电容）.
var Farad *DerivedUnit

// Henry (H, 亨利) is the SI unit of inductance（电感）.
var Henry *DerivedUnit

// Tesla (T, 特斯拉) is the SI unit of magnetic flux density（磁感应强度/磁通密度）.
var Tesla *DerivedUnit

// Weber (Wb, 韦伯) is the SI unit of magnetic flux（磁通量）.
var Weber *DerivedUnit

// ElectricField (V/m, 伏特每米) is the coherent SI unit of electric field strength（电场强度）.
var ElectricField *DerivedUnit

// ElectricDisplacement (C/m², 库仑每平方米) is the coherent SI unit of electric displacement（电位移）.
var ElectricDisplacement *DerivedUnit

// ElectricFlux (V·m, 伏特米) is the coherent SI unit of electric flux（电通量）.
var ElectricFlux *DerivedUnit

// CurrentDensity (A/m², 安培每平方米) is the coherent SI unit of electric current density（电流密度）.
var CurrentDensity *DerivedUnit

// Resistivity (Ω·m, 欧姆米) is the coherent SI unit of electric resistivity（电阻率）.
var Resistivity *DerivedUnit

// Conductivity (S/m, 西门子每米) is the coherent SI unit of electric conductivity（电导率）.
var Conductivity *DerivedUnit

// Permittivity (F/m, 法拉每米) is the coherent SI unit of permittivity（介电常数/电容率）.
var Permittivity *DerivedUnit

// Permeability (H/m, 亨利每米) is the coherent SI unit of magnetic permeability（磁导率）.
var Permeability *DerivedUnit

// MagneticFieldStrength (A/m, 安培每米) is the coherent SI unit of magnetic field strength（磁场强度）.
var MagneticFieldStrength *DerivedUnit

// Reluctance (H^-1, 每亨利) is the coherent SI unit of magnetic reluctance（磁阻）.
var Reluctance *DerivedUnit

// Lumen (lm, 流明) is the SI unit of luminous flux（光通量）.
var Lumen *DerivedUnit

// Lux (lx, 勒克斯) is the SI unit of illuminance（照度）.
var Lux *DerivedUnit

// Becquerel (Bq, 贝克勒尔) is the SI unit of radioactivity（放射性活度）.
var Becquerel *DerivedUnit

// Gray (Gy, 戈瑞) is the SI unit of absorbed dose（吸收剂量）.
var Gray *DerivedUnit

// Sievert (Sv, 希沃特) is the SI unit of dose equivalent（剂量当量）.
var Sievert *DerivedUnit

// Katal (kat, 开特) is the SI unit of catalytic activity（催化活度）.
var Katal *DerivedUnit

// SquareMeter (m², 平方米) is the coherent SI unit of area（面积）.
// Do not use SquareMeter.Prefix(Kilo) for km² — that scales by 10³, not 10⁶.
// Use SquareKilometer (or Length(Meter.Prefix(Kilo), 2)) instead.
var SquareMeter *DerivedUnit

// SquareKilometer (km², 平方千米) is a common SI area unit; 1 km² = 10⁶ m².
var SquareKilometer *DerivedUnit

// SquareCentimeter (cm², 平方厘米) is a common SI area unit; 1 cm² = 10⁻⁴ m².
var SquareCentimeter *DerivedUnit

// SquareMillimeter (mm², 平方毫米) is a common SI area unit; 1 mm² = 10⁻⁶ m².
var SquareMillimeter *DerivedUnit

// Are (a, 公亩) is a metric area unit; 1 a = 100 m².
var Are *DerivedUnit

// Hectare (ha, 公顷) is a metric area unit; 1 ha = 10⁴ m² = 100 a.
var Hectare *DerivedUnit

// Acre (ac, 英亩) is an imperial/US survey area unit; 1 ac = 4046.8564224 m².
var Acre *DerivedUnit

// SquareInch (in², 平方英寸) is an imperial/US area unit.
var SquareInch *DerivedUnit

// SquareFoot (ft², 平方英尺) is an imperial/US area unit.
var SquareFoot *DerivedUnit

// SquareYard (yd², 平方码) is an imperial/US area unit.
var SquareYard *DerivedUnit

// SquareMile (mi², 平方英里) is an imperial/US area unit.
var SquareMile *DerivedUnit

// Mu (mu, 亩) is a Chinese land-area unit; 1 mu = 2000/3 m².
var Mu *DerivedUnit

// Qing (qing, 顷) is a Chinese land-area unit; 1 qing = 100 mu = 200000/3 m².
var Qing *DerivedUnit

// Barn (b, 靶恩) is a nuclear cross-section area unit; 1 b = 10⁻²⁸ m².
var Barn *DerivedUnit

// Speed (m·s⁻¹, 米每秒) is the coherent SI unit of speed（速度）; no SI special name.
var Speed *DerivedUnit

// SI coherent derived units registered after unit_si.init loads the base unit registry.
func init() {
	Radian = NewDerivedUnit().Named("rad").MustIntern()
	Steradian = NewDerivedUnit().Named("sr").MustIntern()
	Hertz = NewDerivedUnit().Time(Second, -1).Named("Hz").MustIntern()
	Newton = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 1).Time(Second, -2).Named("N").MustIntern()
	Pascal = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, -1).Time(Second, -2).Named("Pa").MustIntern()
	Joule = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -2).Named("J").MustIntern()
	WattHour = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -2).Named("W·h").Scale(3600).MustIntern()
	Watt = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Named("W").MustIntern()
	VoltAmpere = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Named("VA").MustIntern()
	Var = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Named("var").MustIntern()
	HorsePower = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Named("hp").Scale(745.69987158227022).MustIntern()
	MetricHorsepower = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Named("PS").Scale(735.49875).MustIntern()
	ElectricalHorsePower = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Named("hp(E)").Scale(746).MustIntern()
	Coulomb = NewDerivedUnit().Time(Second, 1).Current(Ampere, 1).Named("C").MustIntern()
	AmpereHour = NewDerivedUnit().Time(Second, 1).Current(Ampere, 1).Named("A·h").Scale(3600).MustIntern()
	Volt = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Current(Ampere, -1).Named("V").MustIntern()
	Ohm = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Current(Ampere, -2).Named("Ω").MustIntern()
	Impedance = Ohm
	Reactance = Ohm
	Siemens = NewDerivedUnit().Mass(Kilogram, -1).Length(Meter, -2).Time(Second, 3).Current(Ampere, 2).Named("S").MustIntern()
	Admittance = Siemens
	Susceptance = Siemens
	Farad = NewDerivedUnit().Mass(Kilogram, -1).Length(Meter, -2).Time(Second, 4).Current(Ampere, 2).Named("F").MustIntern()
	Henry = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -2).Current(Ampere, -2).Named("H").MustIntern()
	Tesla = NewDerivedUnit().Mass(Kilogram, 1).Time(Second, -2).Current(Ampere, -1).Named("T").MustIntern()
	Weber = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -2).Current(Ampere, -1).Named("Wb").MustIntern()
	ElectricField = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 1).Time(Second, -3).Current(Ampere, -1).Named("V/m").MustIntern()
	ElectricDisplacement = NewDerivedUnit().Length(Meter, -2).Time(Second, 1).Current(Ampere, 1).Named("C/m²").MustIntern()
	ElectricFlux = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 3).Time(Second, -3).Current(Ampere, -1).Named("V·m").MustIntern()
	CurrentDensity = NewDerivedUnit().Length(Meter, -2).Current(Ampere, 1).Named("A/m²").MustIntern()
	Resistivity = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 3).Time(Second, -3).Current(Ampere, -2).Named("Ω·m").MustIntern()
	Conductivity = NewDerivedUnit().Mass(Kilogram, -1).Length(Meter, -3).Time(Second, 3).Current(Ampere, 2).Named("S/m").MustIntern()
	Permittivity = NewDerivedUnit().Mass(Kilogram, -1).Length(Meter, -3).Time(Second, 4).Current(Ampere, 2).Named("F/m").MustIntern()
	Permeability = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 1).Time(Second, -2).Current(Ampere, -2).Named("H/m").MustIntern()
	MagneticFieldStrength = NewDerivedUnit().Length(Meter, -1).Current(Ampere, 1).Named("A/m").MustIntern()
	Reluctance = NewDerivedUnit().Mass(Kilogram, -1).Length(Meter, -2).Time(Second, 2).Current(Ampere, 2).Named("H^-1").MustIntern()
	Lumen = NewDerivedUnit().Luminous(Candela, 1).Named("lm").MustIntern()
	Lux = NewDerivedUnit().Luminous(Candela, 1).Length(Meter, -2).Named("lx").MustIntern()
	Becquerel = NewDerivedUnit().Time(Second, -1).Named("Bq").MustIntern()
	Gray = NewDerivedUnit().Length(Meter, 2).Time(Second, -2).Named("Gy").MustIntern()
	Sievert = NewDerivedUnit().Length(Meter, 2).Time(Second, -2).Named("Sv").MustIntern()
	Katal = NewDerivedUnit().Amount(Mole, 1).Time(Second, -1).Named("kat").MustIntern()
	SquareMeter = NewDerivedUnit().Length(Meter, 2).Named("m²").MustIntern()
	SquareKilometer = NewDerivedUnit().Length(Meter.Prefix(Kilo), 2).Named("km²").MustIntern()
	SquareCentimeter = NewDerivedUnit().Length(Meter.Prefix(Centi), 2).Named("cm²").MustIntern()
	SquareMillimeter = NewDerivedUnit().Length(Meter.Prefix(Milli), 2).Named("mm²").MustIntern()
	Are = NewDerivedUnit().Length(Meter, 2).Named("a").Scale(100).MustIntern()
	Hectare = NewDerivedUnit().Length(Meter, 2).Named("ha").Scale(1e4).MustIntern()
	Acre = NewDerivedUnit().Length(Meter, 2).Named("ac").Scale(4046.8564224).MustIntern()
	SquareInch = NewDerivedUnit().Length(Inch, 2).Named("in²").MustIntern()
	SquareFoot = NewDerivedUnit().Length(Foot, 2).Named("ft²").MustIntern()
	SquareYard = NewDerivedUnit().Length(Yard, 2).Named("yd²").MustIntern()
	SquareMile = NewDerivedUnit().Length(Mile, 2).Named("mi²").MustIntern()
	Mu = NewDerivedUnit().Length(Meter, 2).Named("mu").Scale(2000.0 / 3.0).MustIntern()
	Qing = NewDerivedUnit().Length(Meter, 2).Named("qing").Scale(200000.0 / 3.0).MustIntern()
	Barn = NewDerivedUnit().Length(Meter, 2).Named("b").Scale(1e-28).MustIntern()
	Speed = NewDerivedUnit().Length(Meter, 1).Time(Second, -1).MustIntern()
}
