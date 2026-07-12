---
title: "量纲"
description: "量纲"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 1
toc: true
---
**量纲**按种类对物理量分类。`go-unit` 遵循七大 SI 基本量纲，并将组合（导出）量纲表示为基本量纲的指数。

## SI 基本量纲

`Dimension` 类型标识七种独立的基本量纲：

| 常量 | 符号 | 名称 | SI 基本单位 |
|------|------|------|-------------|
| `DimLength` | L | 长度 | 米（`Meter`） |
| `DimMass` | M | 质量 | 千克（`Kilogram`） |
| `DimTime` | T | 时间 | 秒（`Second`） |
| `DimCurrent` | I | 电流 | 安培（`Ampere`） |
| `DimTemperature` | Θ | 热力学温度 | 开尔文（`Kelvin`） |
| `DimAmount` | N | 物质的量 | 摩尔（`Mole`） |
| `DimLuminous` | J | 发光强度 | 坎德拉（`Candela`） |

```go
dim := u.DimLength
fmt.Println(dim.String())  // "length"
fmt.Println(dim.Symbol())  // "L"
fmt.Println(dim.Base())    // Unit("meter")
```

> **Note: 温度符号**
> 量纲分析中热力学温度用 **Θ**（theta）。代码里导出量纲结构体字段名为 `H`，为历史兼容。
>
## 导出量纲

**导出量纲**用七个基本量纲的指数描述组合量。结构体字段 `L`、`M`、`T`、`I`、`H`、`N`、`J` 为有符号指数，未出现的量纲为 0。

```go
speed := u.DerivedDimension{L: 1, T: -1}        // L¹T⁻¹
force := u.DerivedDimension{M: 1, L: 1, T: -2}  // M¹L¹T⁻²
```

常见示例：

| 物理量 | DerivedDimension | 示例单位 |
|--------|------------------|----------|
| 速度 | `{L:1, T:-1}` | m/s、km/h |
| 力 | `{M:1, L:1, T:-2}` | N = kg·m/s² |
| 压强 | `{M:1, L:-1, T:-2}` | Pa |
| 能量 | `{M:1, L:2, T:-2}` | J |

`unit_si_derived.go` 中预定义了 `DimForce`、`DimPressure`、`DimSpeed` 等变量。

## 量纲代数

导出量纲按物理量的乘除组合：

```go
d1 := u.DimForce                    // M¹L¹T⁻²
d2 := u.DimSpeed                    // L¹T⁻¹
product := d1.Add(d2)                  // 相乘 → 指数相加
quotient := d1.Sub(d2)                 // 相除 → 指数相减
equal := d1.Equal(u.DimForce)
```

物理量的**加减**要求导出量纲**相等**；**乘除**通过 `Add` / `Sub` 改变导出量纲。

## 量纲符号渲染

`DerivedDimension.Symbol()` 输出可读的量纲字符串：

```go
u.DimForce.Symbol()                              // "M·L·T^-2"
u.DimForce.Symbol(u.WithExpSign(u.ExpSignSup)) // "M·L·T⁻²"
```

选项见 [符号格式化](../symbol-formatting/)。

## 无量纲

空的 `DerivedDimension{}` 表示无量纲——例如两个长度相除，或平面角（rad）、立体角（sr），在本库中均视为无量纲。

下一步：[单位与换算 →](../units-and-conversion/)
