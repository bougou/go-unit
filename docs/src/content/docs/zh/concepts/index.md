---
title: "核心概念"
description: "量纲、单位、物理量与导出单位。"
sidebar:
  order: 1
---
使用 API 前，建议先理解库的分层模型：

```
Dimension（量纲）  →  Unit（单位）  →  Quantity（物理量）
         ↓                    ↓
DerivedDimension（导出量纲） → DerivedUnit（导出单位） → DerivedQuantity（导出量）
```

## 第一层：量纲

**量纲**描述物理量的*种类*——长度、质量、时间等。量纲本身不带数值比例，只用于分类。

- [量纲](dimensions/) — 七大 SI 基本量纲与导出量纲

## 第二层：单位

**单位**是量纲的具体度量方式——米与千米、摄氏度与开尔文。单位通过 `Scale`、`Offset` 换算到该量纲的 SI **基本单位**。

- [单位与换算](units-and-conversion/) — 注册表、仿射换算

## 第三层：物理量

**物理量**将 `float64` 数值与单位绑定。`Add`、`By`、`Mul`、`Div` 等运算在量纲规则下组合或保持量纲。

- [物理量](quantities/) — `Quantity` 与类型化包装

## 第四层：导出单位

当物理量跨越多个基本量纲（速度 = 长度/时间），需要 **导出单位**——由基本单位按指数组合，可选 SI 专用名称如 `N`、`Hz`。

- [导出单位](derived-units/) — `NewDerivedUnit`、`NewDerivedQuantity`、`Of`、SI 命名单位
- [内化与注册](intern-and-registry/) — 规范实例
- [符号格式化](symbol-formatting/) — `FormatOption`（`Symbol` / `Format`）

## 两套 API，同一模型

| 风格 | 适用场景 |
|------|----------|
| **类型化**（`LengthQuantity` 等） | 编译期量纲固定，最安全 |
| **非类型化**（`Quantity`） | 运行时量纲，更简单但检查较弱 |

两者共用同一套换算与运算逻辑。

## 不在范围内

- 单位符号存在歧义时的自动推断（如 `pc` 可能是秒差距或派卡 — 先注册者优先）
- 货币、未注册的业务自定义单位

由基本单位符号组成的**未注册导出符号**（如 `km/h`）在 lookup 失败时会通过 `DerivedUnitParse` 解析。SI 专用名（`N`、`Hz`）仍依赖注册表。

**数值词头**（K、M、Mi）单独成章，因其作用于纯数字，而非带量纲的物理量。物理单位用 `QuantityParse`；纯数字缩放用 `PrefixParse`。

下一步：[量纲 →](dimensions/)
