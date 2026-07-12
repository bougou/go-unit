package u

// SI (国际单位制) coherent derived dimensions (一贯导出量纲) and units with special names (专用名称，22 SI derived units).
// Degree Celsius (°C, 摄氏度) is defined as Celsius in unit_si.go (affine temperature 仿射温度，not DerivedUnit).

var (
	// DimAngle is dimensionless (无量纲): plane angle (rad, 弧度) and solid angle (sr, 球面度).
	DimAngle = DerivedDimension{}
	// DimFrequency is T⁻¹: frequency (Hz, 赫兹) and radioactivity (Bq, 贝克勒尔).
	DimFrequency = DerivedDimension{T: -1}
	// DimForce is M¹L¹T⁻² (newton, 牛顿).
	DimForce = DerivedDimension{M: 1, L: 1, T: -2}
	// DimPressure is M¹L⁻¹T⁻² (pascal, 帕斯卡).
	DimPressure = DerivedDimension{M: 1, L: -1, T: -2}
	// DimEnergy is M¹L²T⁻² (joule, 焦耳).
	DimEnergy = DerivedDimension{M: 1, L: 2, T: -2}
	// DimPower is M¹L²T⁻³ (watt, 瓦特).
	DimPower = DerivedDimension{M: 1, L: 2, T: -3}
	// DimElectricCharge is T¹I¹ (coulomb, 库仑).
	DimElectricCharge = DerivedDimension{T: 1, I: 1}
	// DimElectricPotential is M¹L²T⁻³I⁻¹ (volt, 伏特).
	DimElectricPotential = DerivedDimension{M: 1, L: 2, T: -3, I: -1}
	// DimResistance is M¹L²T⁻³I⁻² (ohm, 欧姆).
	DimResistance = DerivedDimension{M: 1, L: 2, T: -3, I: -2}
	// DimConductance is M⁻¹L⁻²T³I² (siemens, 西门子).
	DimConductance = DerivedDimension{M: -1, L: -2, T: 3, I: 2}
	// DimCapacitance is M⁻¹L⁻²T⁴I² (farad, 法拉).
	DimCapacitance = DerivedDimension{M: -1, L: -2, T: 4, I: 2}
	// DimInductance is M¹L²T⁻²I⁻² (henry, 亨利).
	DimInductance = DerivedDimension{M: 1, L: 2, T: -2, I: -2}
	// DimMagneticFluxDensity is M¹T⁻²I⁻¹ (tesla, 特斯拉).
	DimMagneticFluxDensity = DerivedDimension{M: 1, T: -2, I: -1}
	// DimMagneticFlux is M¹L²T⁻²I⁻¹ (weber, 韦伯).
	DimMagneticFlux = DerivedDimension{M: 1, L: 2, T: -2, I: -1}
	// DimLuminousFlux is J¹ (lumen, 流明).
	DimLuminousFlux = DerivedDimension{J: 1}
	// DimIlluminance is J¹L⁻² (lux, 勒克斯).
	DimIlluminance = DerivedDimension{J: 1, L: -2}
	// DimAbsorbedDose is L²T⁻² (gray 戈瑞，sievert 希沃特).
	DimAbsorbedDose = DerivedDimension{L: 2, T: -2}
	// DimCatalyticActivity is N¹T⁻¹ (katal, 开特).
	DimCatalyticActivity = DerivedDimension{N: 1, T: -1}
	// DimSpeed is L¹T⁻¹ (m/s, 米每秒); coherent but without an SI special name.
	DimSpeed = DerivedDimension{L: 1, T: -1}
)

// RadianUnit is the SI unit of plane angle (rad, 弧度).
var RadianUnit *DerivedUnit

// SteradianUnit is the SI unit of solid angle (sr, 球面度).
var SteradianUnit *DerivedUnit

// HertzUnit is the SI unit of frequency (Hz, 赫兹).
var HertzUnit *DerivedUnit

// NewtonUnit is the SI unit of force (N, 牛顿).
var NewtonUnit *DerivedUnit

// PascalUnit is the SI unit of pressure (Pa, 帕斯卡).
var PascalUnit *DerivedUnit

// JouleUnit is the SI unit of energy (J, 焦耳).
var JouleUnit *DerivedUnit

// WattUnit is the SI unit of power (W, 瓦特).
var WattUnit *DerivedUnit

// CoulombUnit is the SI unit of electric charge (C, 库仑).
var CoulombUnit *DerivedUnit

// VoltUnit is the SI unit of electric potential (V, 伏特).
var VoltUnit *DerivedUnit

// OhmUnit is the SI unit of electrical resistance (Ω, 欧姆).
var OhmUnit *DerivedUnit

// SiemensUnit is the SI unit of electrical conductance (S, 西门子).
var SiemensUnit *DerivedUnit

// FaradUnit is the SI unit of capacitance (F, 法拉).
var FaradUnit *DerivedUnit

// HenryUnit is the SI unit of inductance (H, 亨利).
var HenryUnit *DerivedUnit

// TeslaUnit is the SI unit of magnetic flux density (T, 特斯拉).
var TeslaUnit *DerivedUnit

// WeberUnit is the SI unit of magnetic flux (Wb, 韦伯).
var WeberUnit *DerivedUnit

// LumenUnit is the SI unit of luminous flux (lm, 流明).
var LumenUnit *DerivedUnit

// LuxUnit is the SI unit of illuminance (lx, 勒克斯).
var LuxUnit *DerivedUnit

// BecquerelUnit is the SI unit of radioactivity (Bq, 贝克勒尔).
var BecquerelUnit *DerivedUnit

// GrayUnit is the SI unit of absorbed dose (Gy, 戈瑞).
var GrayUnit *DerivedUnit

// SievertUnit is the SI unit of dose equivalent (Sv, 希沃特).
var SievertUnit *DerivedUnit

// KatalUnit is the SI unit of catalytic activity (kat, 开特).
var KatalUnit *DerivedUnit

// SpeedUnit is the coherent SI unit of speed (m·s⁻¹, 米每秒); no special name.
var SpeedUnit *DerivedUnit

// ForceUnit is an alias for NewtonUnit.
var ForceUnit *DerivedUnit

// SI coherent derived units registered after unit_si.init loads the base unit registry.
func init() {
	RadianUnit = NewDerivedUnit().Named("rad").MustIntern()
	SteradianUnit = NewDerivedUnit().Named("sr").MustIntern()
	HertzUnit = NewDerivedUnit().Time(Second, -1).Named("Hz").MustIntern()
	NewtonUnit = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 1).Time(Second, -2).Named("N").MustIntern()
	PascalUnit = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, -1).Time(Second, -2).Named("Pa").MustIntern()
	JouleUnit = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -2).Named("J").MustIntern()
	WattUnit = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Named("W").MustIntern()
	CoulombUnit = NewDerivedUnit().Time(Second, 1).Current(Ampere, 1).Named("C").MustIntern()
	VoltUnit = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Current(Ampere, -1).Named("V").MustIntern()
	OhmUnit = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -3).Current(Ampere, -2).Named("Ω").MustIntern()
	SiemensUnit = NewDerivedUnit().Mass(Kilogram, -1).Length(Meter, -2).Time(Second, 3).Current(Ampere, 2).Named("S").MustIntern()
	FaradUnit = NewDerivedUnit().Mass(Kilogram, -1).Length(Meter, -2).Time(Second, 4).Current(Ampere, 2).Named("F").MustIntern()
	HenryUnit = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -2).Current(Ampere, -2).Named("H").MustIntern()
	TeslaUnit = NewDerivedUnit().Mass(Kilogram, 1).Time(Second, -2).Current(Ampere, -1).Named("T").MustIntern()
	WeberUnit = NewDerivedUnit().Mass(Kilogram, 1).Length(Meter, 2).Time(Second, -2).Current(Ampere, -1).Named("Wb").MustIntern()
	LumenUnit = NewDerivedUnit().Luminous(Candela, 1).Named("lm").MustIntern()
	LuxUnit = NewDerivedUnit().Luminous(Candela, 1).Length(Meter, -2).Named("lx").MustIntern()
	BecquerelUnit = NewDerivedUnit().Time(Second, -1).Named("Bq").MustIntern()
	GrayUnit = NewDerivedUnit().Length(Meter, 2).Time(Second, -2).Named("Gy").MustIntern()
	SievertUnit = NewDerivedUnit().Length(Meter, 2).Time(Second, -2).Named("Sv").MustIntern()
	KatalUnit = NewDerivedUnit().Amount(Mole, 1).Time(Second, -1).Named("kat").MustIntern()
	SpeedUnit = NewDerivedUnit().Length(Meter, 1).Time(Second, -1).MustIntern()
	ForceUnit = NewtonUnit
}
