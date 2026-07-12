---
title: "SI 基本单位"
description: "SI 基本单位"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 1
toc: true
---
常量定义于 `unit_si.go`，启动时注册并换算到各量纲 SI 基本单位。

## 长度（`LengthUnit`）— 基本单位：米

| 分类 | 常量 |
|------|------|
| SI 词头 | `Picometer` … `Petameter` |
| 科学 | `Angstrom`、`AstronomicalUnit`、`LightYear`、`Parsec` |
| 英制/美制 | `Inch`、`Foot`、`Yard`、`Mile`、`NauticalMile` 等 |
| 排版 | `Point`、`Pica` |
| 市制 | `Chi`（尺）、`Cun`（寸）、`Zhang`（丈）、`Li`（里） |

## 质量（`MassUnit`）— 基本单位：千克

SI 词头、公制吨、常衡/金衡、克拉/道尔顿、斤/两等。

## 时间（`TimeUnit`）— 基本单位：秒

SI 词头及 `Minute`、`Hour`、`Day`、`Week`、`Month`、`Year` 等。

## 电流（`CurrentUnit`）— 基本单位：安培

含 CGS 单位 `Biot`（1 Bi = 10 A）。

## 温度（`TemperatureUnit`）— 基本单位：开尔文

| 类型 | 单位 |
|------|------|
| 比例 | `Kelvin` 及词头 |
| 仿射 | `Celsius`、`Fahrenheit`、`Rankine` |

## 物质的量（`AmountUnit`）— 摩尔

## 发光强度（`LuminousUnit`）— 坎德拉

## 符号

```go
u.Unit(u.Kilometer).Symbol() // "km"
u.Unit(u.Celsius).Symbol()   // "°C"
```

完整 `Scale`、`Offset` 见 `unit_si.go` 中 `siUnitDefs`。

下一步：[SI 导出单位 →](../si-derived-units/)
