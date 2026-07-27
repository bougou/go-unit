---
title: "SI derived units"
description: "SI derived units"
sidebar:
  order: 2
---
Package globals in `unit_si_derived.go` register coherent SI derived units with special names. Each is built with `NewDerivedUnit()`, `Named()`, and `MustIntern()` during `init`.

## Dimension variables

| Variable | Dimension | SI unit |
|----------|-----------|---------|
| `DimAngle` | dimensionless | rad, sr |
| `DimFrequency` | T⁻¹ | Hz, Bq |
| `DimForce` | M¹L¹T⁻² | N |
| `DimPressure` | M¹L⁻¹T⁻² | Pa |
| `DimEnergy` | M¹L²T⁻² | J |
| `DimPower` | M¹L²T⁻³ | W |
| `DimElectricCharge` | T¹I¹ | C |
| `DimElectricPotential` | M¹L²T⁻³I⁻¹ | V |
| `DimResistance` | M¹L²T⁻³I⁻² | Ω |
| `DimConductance` | M⁻¹L⁻²T³I² | S |
| `DimCapacitance` | M⁻¹L⁻²T⁴I² | F |
| `DimInductance` | M¹L²T⁻²I⁻² | H |
| `DimMagneticFluxDensity` | M¹T⁻²I⁻¹ | T |
| `DimMagneticFlux` | M¹L²T⁻²I⁻¹ | Wb |
| `DimLuminousFlux` | J¹ | lm |
| `DimIlluminance` | J¹L⁻² | lx |
| `DimAbsorbedDose` | L²T⁻² | Gy, Sv |
| `DimCatalyticActivity` | N¹T⁻¹ | kat |
| `DimSpeed` | L¹T⁻¹ | m/s (no special name) |

## Unit globals

| Global | Symbol | Composition |
|--------|--------|-------------|
| `Radian` | rad | dimensionless |
| `Steradian` | sr | dimensionless |
| `Hertz` | Hz | s⁻¹ |
| `Newton` | N | kg·m·s⁻² |
| `Pascal` | Pa | kg·m⁻¹·s⁻² |
| `Joule` | J | kg·m²·s⁻² |
| `Watt` | W | kg·m²·s⁻³ |
| `Coulomb` | C | s·A |
| `Volt` | V | kg·m²·s⁻³·A⁻¹ |
| `Ohm` | Ω | kg·m²·s⁻³·A⁻² |
| `Siemens` | S | kg⁻¹·m⁻²·s³·A² |
| `Farad` | F | kg⁻¹·m⁻²·s⁴·A² |
| `Henry` | H | kg·m²·s⁻²·A⁻² |
| `Tesla` | T | kg·s⁻²·A⁻¹ |
| `Weber` | Wb | kg·m²·s⁻²·A⁻¹ |
| `Lumen` | lm | cd |
| `Lux` | lx | cd·m⁻² |
| `Becquerel` | Bq | s⁻¹ |
| `Gray` | Gy | m²·s⁻² |
| `Sievert` | Sv | m²·s⁻² |
| `Katal` | kat | mol·s⁻¹ |
| `Speed` | (compound) | m·s⁻¹ |

## Usage

```go
f := u.Newton.Of(100)
p := u.NewDerivedQuantity(101325, u.Pascal)

fmt.Println(u.Newton.Symbol(u.WithNamedSymbol(true))) // N
fmt.Println(u.Newton.Symbol())                           // kg·m·s^-2
```

## Degree Celsius

`Celsius` is **not** a `DerivedUnit` — it is an affine temperature unit in `unit_si.go` because it uses an offset from kelvin rather than a pure derived composition.

Next: [Byte constants →](byte-constants/)
