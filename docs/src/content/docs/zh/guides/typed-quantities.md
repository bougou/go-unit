---
title: "类型化物理量"
description: "类型化物理量"
sidebar:
  order: 1
---
类型化物理量用维度专属的单位类型包装 `Quantity`，在编译期阻止「长度 + 时间」等错误。

## 类型对

| 量纲 | 单位类型 | 物理量类型 | 构造函数 | 解析 |
|------|----------|------------|----------|------|
| 长度 | `LengthUnit` | `LengthQuantity` | `Length(v, u)` / `u.Of(v)` | `LengthQuantityParse(s)` |
| 质量 | `MassUnit` | `MassQuantity` | `Mass(v, u)` / `u.Of(v)` | `MassQuantityParse(s)` |
| 时间 | `TimeUnit` | `TimeQuantity` | `Time(v, u)` / `u.Of(v)` | `TimeQuantityParse(s)` |
| 电流 | `CurrentUnit` | `CurrentQuantity` | `Current(v, u)` / `u.Of(v)` | `CurrentQuantityParse(s)` |
| 温度 | `TemperatureUnit` | `TemperatureQuantity` | `Temperature(v, u)` / `u.Of(v)` | `TemperatureQuantityParse(s)` |
| 物质的量 | `AmountUnit` | `AmountQuantity` | `Amount(v, u)` / `u.Of(v)` | `AmountQuantityParse(s)` |
| 发光强度 | `LuminousUnit` | `LuminousQuantity` | `Luminous(v, u)` / `u.Of(v)` | `LuminousQuantityParse(s)` |

常量（`Meter`、`Kilogram`、`Hour` …）定义在 `unit_si.go`。

## 基本用法

两种等价构造方式——包级函数 `(value, unit)` 或单位方法 `Of(value)`。`prefix.Kilo` 等 SI 十进制因子来自 `github.com/bougou/go-unit/pkg/prefix`：

```go
import "github.com/bougou/go-unit/pkg/prefix"

d := u.Length(42, u.Meter.Prefix(prefix.Kilo))
// 等价于：
d = u.Meter.Prefix(prefix.Kilo).Of(42)

base := d.Base()           // 42000 m
km := base.By(u.Meter.Prefix(prefix.Kilo)) // 42 km
```

SI 导出单位同样有一对：`NewDerivedQuantity(220, u.Volt)` 与 `u.Volt.Of(220)`。详见 [物理量 — 构造](../concepts/quantities/#construction)。

## 解析类型化物理量

从用户输入或配置字符串读取时，使用对应的类型化解析函数以恢复编译期安全：

```go
d, err := u.LengthQuantityParse("10 km")
t := u.TimeQuantityMustParse("30 min")
```

量纲错误时返回 `ErrDimension`——例如 `LengthQuantityParse("5 s")` 会失败，因为 `s` 是时间而非长度。

导出量（`N`、`m/s` 等）请用 `DerivedQuantityParse`。详见 [物理量 — 从文本解析](../concepts/quantities/#parsing-from-text)。

## 同量纲运算

```go
a := u.Length(10, u.Meter)
b := u.Length(5, u.Meter)
sum := a.Add(b)  // 15 m
diff := a.Sub(b) // 5 m
```

量纲不兼容时，`Add` / `Sub` 返回接收者本身。

## 跨量纲运算

`Mul` / `Div` 接受任意 `derivedQuantity`，返回 `DerivedQuantity`：

```go
speed := u.Length(10, u.Meter.Prefix(prefix.Kilo)).Div(u.Time(2, u.Hour))
area := u.Length(3, u.Meter).Mul(u.Length(4, u.Meter))
```

## 标量缩放

```go
u.Length(10, u.Meter).MulV(2) // 20 m
u.Length(10, u.Meter).DivV(4) // 2.5 m
```

`DivV(0)` 返回接收者不变。

## 温度说明

`TemperatureQuantity.Add` / `Sub` 经共享 `Quantity` 逻辑做仿射换算——适用于在兼容仿射单位下表示的温度**间隔**。

## 何时用非类型化 Quantity

量纲仅在运行时确定，或编写泛型工具时使用 `Quantity`。失去编译期检查，换算规则相同。

下一步：[物理量运算 →](arithmetic/)
