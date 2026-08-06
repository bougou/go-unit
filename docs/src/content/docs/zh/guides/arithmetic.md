---
title: "物理量运算"
description: "物理量运算"
sidebar:
  order: 2
---
运算遵循量纲分析：加减要求量纲一致；乘除组合量纲。

## 加减

操作数须具相同（导出）量纲。第二个操作数经基本单位换算到第一个的单位：

```go
import "github.com/bougou/go-unit/pkg/prefix"

a := u.Length(1, u.Meter.Prefix(prefix.Kilo))
b := u.Length(500, u.Meter)
sum := a.Add(b) // 1.5 km
```

量纲不兼容 → 返回接收者不变。

### 导出量

```go
speedUnit := u.NewDerivedUnit().Length(u.Meter.Prefix(prefix.Kilo), 1).Time(u.Hour, -1)
a := u.NewDerivedQuantity(60, speedUnit)
b := speedUnit.Of(30)
a.Add(b) // 90 km/h
```

`DerivedQuantity` 的加减还要求各基本量纲上的单位可**等比例**换算（与 `By` 相同）。量纲相同但含仿射对（如 °C 与 K）时返回接收者不变。比例单位可跨单位相加，例如 `36 km/h + 10 m/s → 72 km/h`。

## 乘法

**相加**导出量纲指数，**相乘**数值（含基本单位归一化）：

```go
u.Length(2, u.Meter.Prefix(prefix.Kilo)).Mul(u.Length(3, u.Meter))
```

各维单位优先沿用操作数中已有单位，否则回退 SI 基本单位。

对 `DerivedQuantity`，两个操作数**共同出现**的每个基本量纲，所用单位也必须等比例；否则返回接收者不变。

## 除法

**相减**指数。 **同量纲**相除得无量纲（`NoneUnit`）：

```go
u.Mass(10, u.Kilogram).Div(u.Mass(2, u.Kilogram)) // 5, 无量纲
```

`DerivedQuantity` 同样校验重叠基本量纲上的等比例关系。除数为 0 或 `DivV(0)` 时返回接收者不变。

## 示例：求平均速度

```go
trip := u.Length(120, u.Meter.Prefix(prefix.Kilo))
drive := u.Time(2, u.Hour)
avg := trip.Div(drive)           // 60 km/h
si := avg.By(u.Speed)     // ≈16.667 m/s
```

## 比例换算检查

内部 `unitsProportional` 验证两单位经基本单位是否为恒定比例。该检查用于 `By*`，也用于 `DerivedQuantity` 的加减乘除：操作数中相同基本量纲上的单位必须等比例，从而排除 °C/K 等仿射对。

## 不支持

- 自动简化为 SI 专用名（除非 `By(Newton)`；有专用名时 `Symbol()` / `Format()` 默认显示专用名）
- 浮点溢出/下溢处理
- 不确定度传播

下一步：[数值词头 →](numeric-prefixes/)
