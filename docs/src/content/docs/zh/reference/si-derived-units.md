---
title: "SI 导出单位"
description: "SI 导出单位"
sidebar:
  order: 2
---
`unit_si_derived.go` 在 `init` 中注册带 SI 专用名称的一贯导出单位。

## 量纲变量

| 变量 | 量纲 | SI 单位 |
|------|------|---------|
| `DimAngle` | 无量纲 | rad、sr |
| `DimFrequency` | T⁻¹ | Hz、Bq |
| `DimForce` | M¹L¹T⁻² | N（牛顿） |
| `DimPressure` | M¹L⁻¹T⁻² | Pa（帕斯卡） |
| `DimEnergy` | M¹L²T⁻² | J（焦耳） |
| `DimPower` | M¹L²T⁻³ | W（瓦特） |
| `DimElectricCharge` | T¹I¹ | C（库仑） |
| `DimElectricPotential` | M¹L²T⁻³I⁻¹ | V（伏特） |
| `DimResistance` | M¹L²T⁻³I⁻² | Ω（欧姆） |
| `DimConductance` | M⁻¹L⁻²T³I² | S（西门子） |
| `DimCapacitance` | M⁻¹L⁻²T⁴I² | F（法拉） |
| `DimInductance` | M¹L²T⁻²I⁻² | H（亨利） |
| `DimMagneticFluxDensity` | M¹T⁻²I⁻¹ | T（特斯拉） |
| `DimMagneticFlux` | M¹L²T⁻²I⁻¹ | Wb（韦伯） |
| `DimLuminousFlux` | J¹ | lm（流明） |
| `DimIlluminance` | J¹L⁻² | lx（勒克斯） |
| `DimAbsorbedDose` | L²T⁻² | Gy、Sv |
| `DimCatalyticActivity` | N¹T⁻¹ | kat（开特） |
| `DimSpeed` | L¹T⁻¹ | m/s（无专用名） |

## 单位全局变量

`Radian`、`Newton`、`Pascal`、`Hertz`、`Speed` 等 — 完整表见英文参考或与源码对照。

## 用法

```go
f := u.Newton.Of(100)
u.Newton.Symbol(u.WithNamedSymbol(true)) // "N"
u.Newton.Symbol()                           // "kg·m·s^-2"
```

## 摄氏度

`Celsius` **不是** `DerivedUnit`，而是 `unit_si.go` 中带 Offset 的仿射温度单位。

下一步：[字节常量 →](byte-constants/)
