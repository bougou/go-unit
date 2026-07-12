---
title: "SI derived units"
description: "SI derived units"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 2
toc: true
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
| `RadianUnit` | rad | dimensionless |
| `SteradianUnit` | sr | dimensionless |
| `HertzUnit` | Hz | s⁻¹ |
| `NewtonUnit` / `ForceUnit` | N | kg·m·s⁻² |
| `PascalUnit` | Pa | kg·m⁻¹·s⁻² |
| `JouleUnit` | J | kg·m²·s⁻² |
| `WattUnit` | W | kg·m²·s⁻³ |
| `CoulombUnit` | C | s·A |
| `VoltUnit` | V | kg·m²·s⁻³·A⁻¹ |
| `OhmUnit` | Ω | kg·m²·s⁻³·A⁻² |
| `SiemensUnit` | S | kg⁻¹·m⁻²·s³·A² |
| `FaradUnit` | F | kg⁻¹·m⁻²·s⁴·A² |
| `HenryUnit` | H | kg·m²·s⁻²·A⁻² |
| `TeslaUnit` | T | kg·s⁻²·A⁻¹ |
| `WeberUnit` | Wb | kg·m²·s⁻²·A⁻¹ |
| `LumenUnit` | lm | cd |
| `LuxUnit` | lx | cd·m⁻² |
| `BecquerelUnit` | Bq | s⁻¹ |
| `GrayUnit` | Gy | m²·s⁻² |
| `SievertUnit` | Sv | m²·s⁻² |
| `KatalUnit` | kat | mol·s⁻¹ |
| `SpeedUnit` | (compound) | m·s⁻¹ |

## Usage

```go
f := u.NewDerivedQuantity(100, u.NewtonUnit)
p := u.NewDerivedQuantity(101325, u.PascalUnit)

fmt.Println(u.ForceUnit.Symbol(u.WithNamedSymbol(true))) // N
fmt.Println(u.ForceUnit.Symbol())                           // kg·m·s^-2
```

## Degree Celsius

`Celsius` is **not** a `DerivedUnit` — it is an affine temperature unit in `unit_si.go` because it uses an offset from kelvin rather than a pure derived composition.

Next: [Byte constants →](../byte-constants/)
