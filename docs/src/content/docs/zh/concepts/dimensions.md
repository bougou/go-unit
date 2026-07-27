---
title: "量纲"
description: "量纲"
sidebar:
  order: 1
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
> 量纲分析中热力学温度用 **Θ**（theta）。导出量纲结构体字段却命名为 `H`：键盘输入希腊字母 Θ 比输入拉丁字母 H 更麻烦；**H** 可理解为 **Heat**（热），与温度语义相近；且 **H** 的字形与 **Θ** 相似。
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

导出量纲的运算与物理量运算对应，并遵循**量纲齐次性**原则：

```go
d1 := u.DimForce                    // M¹L¹T⁻²
d2 := u.DimSpeed                    // L¹T⁻¹

product := d1.Mul(d2)               // 相乘 → 指数相加，始终有定义
quotient := d1.Div(d2)              // 相除 → 指数相减，始终有定义

sum, ok := d1.Add(u.DimForce)       // 加减：量纲必须相同
diff, ok := d1.Sub(u.DimForce)      // ok == true 时结果量纲仍为 d1
_, ok = d1.Add(d2)                  // 量纲不同 → ok == false

equal := d1.Equal(u.DimForce)
```

| 运算 | 方法 | 规则 |
|------|------|------|
| 加减 | `Add` / `Sub` | 仅当量纲完全相同才有意义，成功时返回该量纲与 `ok=true` |
| 乘除 | `Mul` / `Div` | 指数相加/相减，代数上始终有定义；结果是否对应常见物理量取决于物理诠释 |

## 量纲符号渲染

`DerivedDimension.Symbol()` 输出可读的量纲字符串：

```go
u.DimForce.Symbol()                              // "M·L·T^-2"
u.DimForce.Symbol(u.WithExpSign(u.ExpSignSup)) // "M·L·T⁻²"
```

选项见 [符号格式化](symbol-formatting/)。

## 无量纲

空的 `DerivedDimension{}` 表示无量纲——例如两个长度相除，或平面角（rad）、立体角（sr），在本库中均视为无量纲。

下一步：[单位与换算 →](units-and-conversion/)
