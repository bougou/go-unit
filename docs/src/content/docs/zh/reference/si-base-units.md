---
title: "SI 基本单位"
description: "SI 基本单位"
sidebar:
  order: 1
---
常量定义于 `unit_si.go`。各量纲只导出**锚点**与非 SI 单位；SI 十进制倍数一律用 `Anchor.Prefix(factor)`，因子来自包 [`prefix`](../guides/numeric-prefixes/)（惰性注册）。

## 长度（`LengthUnit`）— 基本单位：米

| 分类 | 常量 / 写法 |
|------|------|
| SI 锚点 | `Meter`；词头如 `Meter.Prefix(prefix.Kilo)` → km |
| 科学 | `Angstrom`、`AstronomicalUnit`、`LightYear`、`Parsec` |
| 英制/美制 | `Inch`、`Foot`、`Yard`、`Mile`、`NauticalMile` 等 |
| 排版 | `Point`、`Pica` |
| 市制 | `Chi`（尺）、`Cun`（寸）、`Zhang`（丈）、`Li`（里） |

## 质量（`MassUnit`）— 基本单位：千克

| 分类 | 常量 / 写法 |
|------|------|
| SI | `Kilogram`（基本单位）、`Gram`（词头根） |
| 词头 | 挂在克上：`Gram.Prefix(prefix.Milli)` → mg；`Gram.Prefix(prefix.Kilo)` → `Kilogram`；`Kilogram.Prefix(prefix.Milli)` → `Gram` |
| 其它 | `Tonne`、常衡/金衡、克拉/道尔顿、斤/两等 |

## 时间（`TimeUnit`）— 基本单位：秒

`Second` + `Second.Prefix(…)`；以及 `Minute`、`Hour`、`Day`、`Week`、`Month`、`Year` 等。

## 电流（`CurrentUnit`）— 基本单位：安培

`Ampere` + `Ampere.Prefix(…)`；另有 CGS `Biot`（1 Bi = 10 A）。

## 温度（`TemperatureUnit`）— 基本单位：开尔文

| 类型 | 单位 |
|------|------|
| 比例 | `Kelvin` + `Kelvin.Prefix(…)` |
| 仿射 | `Celsius`、`Fahrenheit`、`Rankine`（不可 `P`） |

## 物质的量（`AmountUnit`）— 摩尔

`Mole` + `Mole.Prefix(…)`；另有 `PoundMole`。

## 发光强度（`LuminousUnit`）— 坎德拉

`Candela` + `Candela.Prefix(…)`。

## 符号

```go
u.Unit(u.Meter.Prefix(prefix.Kilo)).Symbol() // "km"
u.Unit(u.Gram.Prefix(prefix.Milli)).Symbol() // "mg"
u.Unit(u.Celsius).Symbol()         // "°C"
```

完整 `Scale`、`Offset` 见 `unit_si.go` 中 `siUnitDefs`。

下一步：[SI 导出单位 →](si-derived-units/)
