---
title: "数值词头"
description: "数值词头"
sidebar:
  order: 3
---
**数值词头**用 SI（1000）或 IEC（1024）因子缩放纯浮点数，与物理单位刻意分离。

> **Important: 不是物理单位**
> `PrefixParse("1024 G")` 解析的是词头 `G`（10⁹），不是千兆字节。请先去掉 `B`、`bit`、`/s` 等物理单位后缀。
>
## 模式

| 模式 | 基数 | 符号 | 典型用途 |
|------|------|------|----------|
| `SI` | 1000 | K、M、G（无 `i`） | 十进制数量级 |
| `IEC` | 1024 | Ki、Mi、Gi（带 `i`） | 二进制存储 |
| `SI1024` | 1024 | K、M、G（无 `i`） | 1024 进位 + SI 字母 |
| `Auto` | 自动 | 见下 | 混合输入 |
| `ForceSI` / `ForceIEC` | 强制 | 两种符号集 | 严格模式 |

### Auto

- **解析**：末尾有 `i` → IEC，否则 SI
- **格式化**：等同 SI

## 解析 PrefixParse

```go
v, _ := u.PrefixParse("1.5G", u.SI)    // 1.5e9
v, _ := u.PrefixParse("1024Ki", u.IEC) // 1048576
```

有效输入：纯数字、数字+词头。

无效示例（需先剥离单位）：

| 输入 | 应改为 |
|------|--------|
| `"1024 MiB"` | `"1024 Mi"` |
| `"1024 Gb/s"` | `"1024 G"` |

数字部分可通过 `TrimDelimiter` 去掉千分位分隔符。

## 格式化 PrefixFormat

```go
u.PrefixFormat(999_000_000, u.SI) // "999 M"
u.PrefixFormat(1048576, u.IEC)    // "1 Mi"
```

`PrefixFormat2` 分别返回数字与词头。选项：`WithPrefixSpace`、`WithPrefixPrecision`、`WithRoundMethod`、`WithPrefix` 等。

## 常量

`prefix.go` 导出 `Kilo`、`Mega`、`Kibi`、`Mebi` 等 `float64` 因子。

字节倍数命名常量见 [字节常量](../reference/byte-constants/)。

## 与物理单位 `Prefix()` 的区别

`PrefixParse` / `PrefixFormat` 缩放**纯数字**。物理单位使用另一套 API — `Anchor.Prefix(prefix)` — 返回真正的单位：

```go
u.Meter.Prefix(u.Kilo)  // LengthUnit km
u.Gram.Prefix(u.Milli)  // MassUnit mg
u.Ohm.Prefix(u.Mega)    // *DerivedUnit MΩ
```

带数值的字符串用 `QuantityParse("10 km")` 或 `DerivedQuantityParse("1 MΩ")`。见 [物理量 — Prefix()](../concepts/quantities/)。

## 与物理量解析的区别

`PrefixParse` 缩放**纯数字**（K/M/G 词头），不绑定物理单位。

`"10 km"`、`"5 N"` 等带单位的字符串请用 [物理量解析](../concepts/quantities/#parsing-from-text)（`QuantityParse`、`LengthQuantityParse`、`DerivedQuantityParse`）。

下一步：[数字格式化 →](number-formatting/)
