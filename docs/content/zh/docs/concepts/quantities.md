---
title: "物理量"
description: "物理量"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 3
toc: true
---
**物理量**是数值与单位的组合，是面向用户的主要数据对象。

## Quantity 结构体

非类型化 `Quantity` 可承载任意已注册的基本量纲单位：

```go
q := u.Quantity{Value: 3, Unit: u.Unit(u.Meter)}
fmt.Println(q) // "3 m"
```

需要编译期安全时，请用类型化构造函数 — 见 [类型化物理量](../typed-quantities/)。

## 构造

```go
d := u.Length(10, u.Kilometer)
t := u.Time(30, u.Minute)

q := u.Quantity{Value: 10, Unit: u.Unit(u.Kilometer)}
```

## 换算

### Base()

换算到同一量纲的 SI 基本单位：

```go
u.Length(1, u.Kilometer).Base() // {1000, meter}
```

### By()

在**同一量纲**内换单位：

```go
u.Length(1000, u.Meter).By(u.Kilometer) // 1 km
```

单位不兼容或未知时，`By` 原样返回。

## 从文本解析 {#parsing-from-text}

解析函数将 `"数值 单位"` 字符串转为物理量。规则如下：

- **数值本身不能含空格**（非法：`"1 024 m"`）。
- **数值与单位之间可以有空格**（`"10 m"` 与 `"10m"` 均可）。
- 数值中的千分位 `,`、`_` 可接受（`"1,024 mm"`）。
- 单位符号包括已注册基本单位、SI 导出专用名（`N`、`Hz`）及常见组合形式（`kg·m·s^-2`、`m/s`、上标形式等）。

### 选择哪个解析函数

| 函数 | 返回类型 | 适用场景 |
|------|----------|----------|
| `QuantityParse` | `Quantity` | 通用；接受基本或导出单位 |
| `LengthQuantityParse`、`TimeQuantityParse` 等 | 类型化物理量 | 编译期已知基本量纲 |
| `DerivedQuantityParse` | `DerivedQuantity` | 组合 / SI 导出单位（`N`、`m/s`） |

类型化解析在量纲不匹配时返回 `ErrDimension`：

```go
u.LengthQuantityParse("10 km")  // OK
u.LengthQuantityParse("5 N")    // ErrDimension — 力是导出量，不是长度
u.DerivedQuantityParse("5 N")   // OK
u.DerivedQuantityParse("10 m") // ErrDimension — 应使用 LengthQuantityParse
```

每个 `Parse` 都有对应的 `MustParse`（失败时 panic）。

`QuantityParse` **不同于** `PrefixParse`：`"1.5G"` 是数值词头，不是 1.5 吉米。见 [数值词头](../../guides/numeric-prefixes/)。

### 限制

- SI 专用名（`N`、`Hz` 等）通过符号注册表解析，不会从组合因子代数展开。
- 重复符号以先注册者为准（`pc` → 秒差距，而非派卡）。
- 不同专用名共享同一组合符号时（如 Hz 与 Bq 的 `s^-1`），以先注册者为准；请用专用名（`Hz`、`Bq`）消歧。
- `DerivedUnitParse` 支持 `/`、`·`、`*`、空格、`^` 及 Unicode 上标指数。相邻单位需显式分隔；最长匹配保证 `ms` 表示毫秒。

### 导出单位解析

符号未预注册时，`QuantityParse` 会回退到 `DerivedUnitParse`，对基本单位符号分词并重建 `DerivedUnit`：

```go
du, err := u.DerivedUnitParse("km/h") // Length(Kilometer,1) · Time(Hour,-1)
q, err := u.QuantityParse("60 km/h") // 经 fallback 同样生效
```

仅已注册的**基本单位**符号参与组合（`km`、`h`、`kg` …）。专用名如 `"N"` 不能仅靠语法展开。

## 兼容性

`Compatible` 检查基本量纲是否相同：

```go
u.Length(1, u.Meter).Compatible(u.Length(2, u.Kilometer)) // true
u.Length(1, u.Meter).Compatible(u.Time(1, u.Second))       // false
```

## 字符串

```go
fmt.Println(u.Mass(2.5, u.Kilogram)) // "2.5 kg"
```

## DerivedQuantity（导出量）

量纲组合（乘/除）时得到 `DerivedQuantity`：

```go
speed := u.Length(60, u.Kilometer).Div(u.Time(1, u.Hour))
area := u.Length(3, u.Meter).Mul(u.Length(4, u.Meter))
```

| 字段 | 含义 |
|------|------|
| `Value` | 数值 |
| `Unit` | `*DerivedUnit`（nil 时按无量纲显示） |

### 导出量换算

```go
si := speed.By(u.SpeedUnit)        // 整单位换算
partial := speed.ByLength(u.Meter) // 只换长度项
```

`By` 要求导出量纲一致且各维单位可比例换算。含仿射温度项时，部分维度换算会被拒绝。

## 无量纲

同量纲相除得到无量纲：

```go
ratio := u.Mass(10, u.Kilogram).Div(u.Mass(2, u.Kilogram))
// 使用 NoneUnit
```

或直接构造：

```go
u.Dimensionless(3.14)
```

下一步：[导出单位 →](../derived-units/)
