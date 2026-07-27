---
title: "包参考"
description: "包参考"
sidebar:
  order: 4
---
## 在线 API 文档

**[pkg.go.dev/github.com/bougou/go-unit/pkg/u](https://pkg.go.dev/github.com/bougou/go-unit/pkg/u)**

## 源码文件

库源码位于 `pkg/u/`：

| 文件 | 职责 |
|------|------|
| `pkg/u/dimension.go` | 基本量纲 |
| `pkg/u/deriveddimension.go` | 导出量纲与量纲符号 |
| `pkg/u/unit.go` | 单位注册与换算 |
| `pkg/u/unit_si.go` | SI 及常用基本单位 |
| `pkg/u/unit_si_derived.go` | SI 导出单位 |
| `pkg/u/unit_typed.go` | 类型化 API |
| `pkg/u/quantity.go` / `pkg/u/quantity_arith.go` | 物理量与运算 |
| `pkg/u/quantityparse.go` | `QuantityParse`、符号索引、`ErrDimension` |
| `pkg/u/derivedsymbolparse.go` | `DerivedUnitParse`、导出符号反向解析 |
| `pkg/u/quantityparse_typed.go` | 类型化 `XxxQuantityParse` / `XxxQuantityMustParse` |
| `pkg/u/derivedunit.go` / `pkg/u/derivedquantity.go` | 导出单位、导出量、`FormatOption` |
| `pkg/u/none_unit.go` | 无量纲 |
| `pkg/u/prefix*.go` / `pkg/u/roundmethod.go` | 数值词头 |
| `pkg/u/comma.go` | 千分位格式化（`DelimitInt`、`DelimitFloat`） |
| `pkg/u/quantityformat.go` | 量的 `Format` 辅助函数 |
| `pkg/u/byte_consts.go` | 字节常量 |

## 主要类型

`Dimension`、`DerivedDimension`、`Unit`、`Quantity`、`DerivedUnit`、`DerivedQuantity`、`PrefixMode`、`FormatOption` 等。

## 构造函数

| API | 返回类型 | 示例 |
|-----|----------|------|
| `Length`、`Time`、`Mass` 等 | 类型化物理量 | `Length(10, Meter.Prefix(Kilo))` |
| `NewDerivedQuantity` | `DerivedQuantity` | `NewDerivedQuantity(60, speedUnit)` |
| `DerivedUnit.Of` | `DerivedQuantity` | `Newton.Of(100)` |
| `NewDerivedUnit` | `*DerivedUnit` | `NewDerivedUnit().Length(Meter, 1)` |

## 解析错误

| 错误 | 含义 |
|------|------|
| `ErrSyntax` | 格式非法、未知单位符号或数值无效 |
| `ErrDimension` | 单位量纲与期望的类型不匹配 |

## 物理量字符串解析

输入格式为 `"数值 单位"`。数值本身不能含空格；数值与单位之间可以有空格。

| API | 返回类型 | 接受 |
|-----|----------|------|
| `QuantityParse` | `Quantity` | 已注册符号；失败时回退到组合解析 |
| `DerivedUnitParse` | `*DerivedUnit` | 基本单位组成的导出符号（如 `km/h`、`kg·m·s^-2`） |
| `LengthQuantityParse` 等 | 类型化物理量 | 仅单一基本量纲（如 `m`、`kg`、`°C`） |
| `DerivedQuantityParse` | `DerivedQuantity` | 仅导出单位（如 `N`、`m/s`、`km/h`） |

每个 `Parse` 都有对应的 `MustParse`（失败时 panic），如 `QuantityMustParse`、`LengthQuantityMustParse`、`DerivedQuantityMustParse`。

```go
d, err := u.LengthQuantityParse("10 km")       // LengthQuantity
force, err := u.DerivedQuantityParse("5 N")    // DerivedQuantity
speed, err := u.DerivedQuantityParse("60 km/h") // 组合符号 fallback
q := u.QuantityMustParse("1.5 m")              // Quantity
```

详见 [物理量 — 从文本解析](../concepts/quantities/#parsing-from-text)。

## 版本

模块：`github.com/bougou/go-unit/pkg/u`，Go **1.20+**。

## 本地文档

```bash
make docs-serve
```

文档源码在 `docs/src/content/docs/` 目录，欢迎 PR。
