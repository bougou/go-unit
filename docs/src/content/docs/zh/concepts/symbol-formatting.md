---
title: "符号格式化"
description: "符号格式化"
sidebar:
  order: 6
---
单位与量纲符号通过 `Symbol()` 及 `FormatOption` 按需渲染。

## 默认行为

| 类型 | 默认输出 |
|------|----------|
| 注册单位 | `Symbol` 字段（`"km"`、`"°C"`） |
| 导出单位 | 基本单位组合表达式 |
| 导出量纲 | 单字母指数 |

```go
u.Unit(u.Meter.Prefix(u.Kilo)).Symbol()                    // "km"
u.Newton.Symbol()                               // "kg·m·s^-2"
u.Newton.Symbol(u.WithNamedSymbol(true))     // "N"
```

## 选项

### WithNamedSymbol

有 SI 专用名时使用专用名：

```go
u.Newton.Symbol(u.WithNamedSymbol(true)) // "N"
```

默认为 `false`，显示组合形式。

### WithExpSign

| 值 | 示例 |
|----|------|
| `ExpSignCarat`（默认） | `m^2`、`h^-1` |
| `ExpSignSup` | `m²`、`h⁻¹` |

### WithMulSign

| 值 | 示例 |
|----|------|
| `MulSignDot`（默认） | `km·h^-1` |
| `MulSignSpace` | `km h^-1` |
| `MulSignStar` | `km*h^-1` |

### WithDivSign

| 值 | 示例 |
|----|------|
| `DivSignNegative`（默认） | `km·h⁻¹` |
| `DivSignSlash` | `km/h`、`kg·m/s^2` |

### WithDimOrder

**量纲**符号中字母顺序（非单位符号）：

| 值 | 顺序 |
|----|------|
| `DimOrderMLT`（默认） | M, L, T, I, H, N, J |
| `DimOrderTML` | T, L, M, …（ISO 80000 风格） |

## 示例

```go
speed := u.NewDerivedUnit().Length(u.Meter.Prefix(u.Kilo), 1).Time(u.Hour, -1)

speed.Symbol(u.WithExpSign(u.ExpSignSup))   // "km·h⁻¹"
speed.Symbol(u.WithDivSign(u.DivSignSlash)) // "km/h"
```

`DerivedQuantity.String()` 内联上标指数；`Format()` 默认与 `Symbol()` 一致（caret）。同一套 `FormatOption`（`WithExpSign`、`WithNamedSymbol` 等）可同时用于 `Symbol` 与 `Format`。

```go
u.NewDerivedQuantity(10, u.Newton).Format(
    u.WithNamedSymbol(true),
    u.WithPrecision(2),
) // "10.00 N"
```

下一步：[类型化物理量指南 →](../guides/typed-quantities/)
