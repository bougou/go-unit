---
title: "单位与换算"
description: "单位与换算"
sidebar:
  order: 2
---
**单位**是量纲的具体度量。`go-unit` 在初始化时注册 SI 及常用单位，并通过各量纲的 SI **基本单位**进行换算。

## 单位标识

`Unit` 为字符串类型，作为稳定的内部标识：

```go
km := u.Unit(u.Meter.Prefix(u.Kilo))
fmt.Println(km.Symbol()) // "km"
```

显示符号来自注册的 `unitDef`，或来自组合单位的 `DerivedUnit`。

## 注册表与 unitDef

每个已注册单位包含：

| 字段 | 含义 |
|------|------|
| `Dimension` | 所属基本量纲 |
| `Unit` | 内部注册键 |
| `Symbol` | 短符号（`"km"`、`"°C"`） |
| `Name` | 英文全称 |
| `Scale` | 到基本单位的乘性因子（须 > 0） |
| `Offset` | 缩放后的加性常数 |

### 比例换算（Offset = 0）

纯比例变换——大多数单位：

```
基本单位值 = 数值 × Scale
数值 = 基本单位值 / Scale
```

示例：1 km → 1000 m（`Scale = 1000`，`Offset = 0`）。

### 仿射换算（Offset ≠ 0）

与开尔文同比例但零点不同——温度：

```
基本单位值 = 数值 × Scale + Offset
数值 = (基本单位值 − Offset) / Scale
```

示例：摄氏度 → 开尔文，`Scale = 1`，`Offset = 273.15`。

> **Warning: 比例 vs 仿射**
> `°C` 与 `K` 量纲相同但**不成比例**（Offset ≠ 0）。库会拒绝二者之间的 naive 比例换算。请用 `Base()` / `By()` 走完整仿射映射。
>
## 换算 API

```go
q := u.Length(1, u.Hour)
base := q.Base()           // 3600 s
back := base.By(u.Hour) // 1 h

val, ok := u.Unit(u.Meter.Prefix(u.Kilo)).ToBase(2.5)    // 2500
back, ok := u.Unit(u.Meter.Prefix(u.Kilo)).FromBase(2500) // 2.5
```

`Def()` 返回完整定义：

```go
def, ok := u.Unit(u.Celsius).Def()
// def.Scale == 1, def.Offset == 273.15
```

## 已注册单位

`unit_si.go` 按量纲分组，含：

- 长度 — SI 词头、英制/美制、航海、中国传统（尺、寸、丈、里）
- 质量 — SI、吨、常衡/金衡、斤/两
- 时间 — SI 词头、分至年
- 电流、温度、物质的量、发光强度 — SI 及常用变体

完整列表见 [SI 基本单位](../reference/si-base-units/)。

## 注册单位 vs 导出单位

| 种类 | 示例 | 查询 |
|------|------|------|
| 注册基本量纲单位 | `Meter.Prefix(Kilo)` | `Unit.Def()` |
| 导出组合单位 | km/h、N | `Intern` 后 `Unit.DerivedUnit()` |

组合单位见 [导出单位](derived-units/)。

## 校验

`unit_si.go` 启动时注册全部定义并执行 `validateRegistry()`——每个量纲须恰好有一个 `Scale = 1`、`Offset = 0` 的基本单位。

下一步：[物理量 →](quantities/)
