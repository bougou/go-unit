---
title: "导出单位"
description: "导出单位"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 4
toc: true
---
**导出单位**用基本单位的指数组合表示复合度量——如 km¹·h⁻¹（速度）、kg¹·m¹·s⁻²（力）。

## 构建

`NewDerivedUnit()` 链式设置各量纲：

```go
speed := u.NewDerivedUnit().
    Length(u.Kilometer, 1).
    Time(u.Hour, -1)

force := u.NewDerivedUnit().
    Mass(u.Kilogram, 1).
    Length(u.Meter, 1).
    Time(u.Second, -2)
```

| 方法 | 量纲 | exp > 0 | exp < 0 |
|------|------|---------|---------|
| `Length(u, exp)` | L | 分子 | 分母 |
| `Mass(u, exp)` | M | 分子 | 分母 |
| `Time(u, exp)` | T | 分子 | 分母 |
| `Current` / `Temperature` / `Amount` / `Luminous` | I / Θ / N / J | 同上 | 同上 |

`exp == 0` 清除该项。未知单位或量纲错误会 **panic**——构建器面向程序内定义，非不可信输入。

## SI 专用名称

`Named` 设置 SI 专用符号：

```go
newton := u.NewDerivedUnit().
    Mass(u.Kilogram, 1).
    Length(u.Meter, 1).
    Time(u.Second, -2).
    Named("N")
```

`unit_si_derived.go` 预注册 `NewtonUnit`、`PascalUnit`、`HertzUnit`、`SpeedUnit` 等，见 [SI 导出单位](../si-derived-units/)。

## 查询

```go
force := u.ForceUnit
force.Dim()           // DerivedDimension{M:1, L:1, T:-2}
force.FactorToBase()  // 到 SI 基本组合的乘数
force.Key()           // 注册键
force.SpecialSymbol() // "N"
```

### SI()

改写为各维 SI 基本单位：

```go
u.NewDerivedUnit().Length(u.Kilometer, 1).Time(u.Hour, -1).SI()
// → m¹·s⁻¹
```

## 导出量

`NewDerivedQuantity` 或通过 `Mul` / `Div` 得到：

```go
v := u.NewDerivedQuantity(60, speedUnit)
fmt.Println(v.String()) // "60 km·h⁻¹"
```

已注册的导出单位符号，或可解析的组合符号（如 `km/h`），可用 `DerivedQuantityParse` 从文本解析：

```go
force, err := u.DerivedQuantityParse("5 N")
speed, err := u.DerivedQuantityParse("60 km/h") // 经 DerivedUnitParse fallback
```

基本量字符串如 `"10 m"` 会被拒绝——请用类型化解析。详见 [从文本解析](../quantities/#parsing-from-text)。

## 预定义量纲变量

`DimForce`、`DimEnergy`、`DimSpeed` 等文档化常见导出量纲，并与已注册 SI 导出单位对应。

## 同组合、不同名称

无量纲单位可共享同一基本组合但专用名不同——rad、sr 与未命名的 `NoneUnit`。见 [内化与注册](../intern-and-registry/)。

下一步：[内化与注册 →](../intern-and-registry/)
