---
title: "数值词头"
description: "数值词头"
sidebar:
  order: 3
---
**数值词头**用 SI（1000）或 IEC（1024）因子缩放纯浮点数，位于包 **`prefix`**（`github.com/bougou/go-unit/pkg/prefix`），与物理单位包 `u` 刻意分离。

```go
import (
	"github.com/bougou/go-unit/pkg/prefix"
	u "github.com/bougou/go-unit/pkg/u"
)
```

> **Important: 不是物理单位**
> `prefix.PrefixParse("1024 G")` 解析的是词头 `G`（10⁹），不是千兆字节。请先去掉 `B`、`bit`、`/s` 等物理单位后缀。
>
## 模式

| 模式 | 基数 | 符号 | 典型用途 |
|------|------|------|----------|
| `prefix.SI` | 1000 | K、M、G（无 `i`） | 十进制数量级 |
| `prefix.IEC` | 1024 | Ki、Mi、Gi（带 `i`） | 二进制存储 |
| `prefix.SI1024` | 1024 | K、M、G（无 `i`） | 1024 进位 + SI 字母 |
| `prefix.Auto` | 自动 | 见下 | 混合输入 |
| `prefix.ForceSI` / `prefix.ForceIEC` | 强制 | 两种符号集 | 严格模式 |

### Auto

- **解析**：末尾有 `i` → IEC，否则 SI
- **格式化**：等同 SI

## 解析 PrefixParse

```go
v, _ := prefix.PrefixParse("1.5G", prefix.SI)    // 1.5e9
v, _ := prefix.PrefixParse("1024Ki", prefix.IEC) // 1048576
```

有效输入：纯数字、数字+词头。

无效示例（需先剥离单位）：

| 输入 | 应改为 |
|------|--------|
| `"1024 MiB"` | `"1024 Mi"` |
| `"1024 Gb/s"` | `"1024 G"` |

数字部分会自动去掉千分位分隔符。包 `u` 另提供 `TrimDelimiter`，可在词头解析之外做同类清理。

## 格式化 PrefixFormat

```go
prefix.PrefixFormat(999_000_000, prefix.SI) // "999 M"
prefix.PrefixFormat(1048576, prefix.IEC)    // "1 Mi"
```

`PrefixFormat2` 分别返回数字与词头。选项：`prefix.WithPrefixSpace`、`prefix.WithPrefixPrecision`、`prefix.WithRoundMethod`、`prefix.WithPrefix` 等（`WithPrefix` 接受 `PrefixSymbol`）。

## 常量

`pkg/prefix/prefix.go` 导出 SI 十进制因子 `prefix.Kilo`、`prefix.Mega` 等（类型 `prefix.SIPrefix`），以及 IEC 二进制因子 `prefix.Kibi`、`prefix.Mebi` 等（`float64`）。`Unit.Prefix` / `ByPrefix` 只接受 `SIPrefix`。

`PrefixSymbol` 是强制指定词头字母时的类型。

字节倍数命名常量见 [字节常量](../reference/byte-constants/)。

## 与物理单位 `Prefix()` 的区别

`PrefixParse` / `PrefixFormat` 缩放**纯数字**。物理单位使用另一套 API — `Anchor.Prefix(factor)` — 返回真正的单位：

```go
u.Meter.Prefix(prefix.Kilo)  // LengthUnit km
u.Gram.Prefix(prefix.Milli)  // MassUnit mg
u.Ohm.Prefix(prefix.Mega)    // *DerivedUnit MΩ
```

带数值的字符串用 `QuantityParse("10 km")` 或 `DerivedQuantityParse("1 MΩ")`。见 [物理量 — Prefix()](../concepts/quantities/)。

## 与物理量解析的区别

`PrefixParse` 缩放**纯数字**（K/M/G 词头），不绑定物理单位。

`"10 km"`、`"5 N"` 等带单位的字符串请用 [物理量解析](../concepts/quantities/#parsing-from-text)（`QuantityParse`、`LengthQuantityParse`、`DerivedQuantityParse`）。

下一步：[数字格式化 →](number-formatting/)
