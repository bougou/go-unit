---
title: "物理量运算"
description: "物理量运算"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 2
toc: true
---
运算遵循量纲分析：加减要求量纲一致；乘除组合量纲。

## 加减

操作数须具相同（导出）量纲。第二个操作数经基本单位换算到第一个的单位：

```go
a := u.Length(1, u.Kilometer)
b := u.Length(500, u.Meter)
sum := a.Add(b) // 1.5 km
```

量纲不兼容 → 返回接收者不变。

### 导出量

```go
speedUnit := u.NewDerivedUnit().Length(u.Kilometer, 1).Time(u.Hour, -1)
a := u.NewDerivedQuantity(60, speedUnit)
b := u.NewDerivedQuantity(30, speedUnit)
a.Add(b) // 90 km/h
```

内部先归一化到 SI 基本组合，再缩放到接收者单位。

## 乘法

**相加**导出量纲指数，**相乘**数值（含基本单位归一化）：

```go
u.Length(2, u.Kilometer).Mul(u.Length(3, u.Meter))
```

各维单位优先沿用操作数中已有单位，否则回退 SI 基本单位。

## 除法

**相减**指数。 **同量纲**相除得无量纲（`NoneUnit`）：

```go
u.Mass(10, u.Kilogram).Div(u.Mass(2, u.Kilogram)) // 5, 无量纲
```

除数为 0 或 `DivV(0)` 时返回接收者不变。

## 示例：求平均速度

```go
trip := u.Length(120, u.Kilometer)
drive := u.Time(2, u.Hour)
avg := trip.Div(drive)           // 60 km/h
si := avg.By(u.SpeedUnit)     // ≈16.667 m/s
```

## 比例换算检查

内部 `unitsProportional` 验证两单位经基本单位是否为恒定比例，从而排除 °C/K 等仿射对在部分 `By*` 换算中的误用。

## 不支持

- 自动简化为 SI 专用名（除非 `By(ForceUnit)` 或 `WithNamedSymbol`）
- 浮点溢出/下溢处理
- 不确定度传播

下一步：[数值词头 →](../numeric-prefixes/)
