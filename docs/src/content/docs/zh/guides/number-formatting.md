---
title: "数字格式化"
description: "数字格式化"
sidebar:
  order: 4
---
库在两层处理数字显示：**量的 Format**，以及独立的 `DelimitInt` / `DelimitFloat`。

## Quantity Format

`Quantity`、各类型量以及 `DerivedQuantity` 支持 `Format`，选项与 `Symbol` 共用 `FormatOption`：

```go
q := u.Quantity{Value: 1234.5, Unit: u.Unit(u.Meter)}
q.Format() // "1234.5 m"（与 String 相同）
q.Format(
    u.WithPrecision(1),
    u.WithNumberDelimiter(u.NumberDelimiterComma),
) // "1,234.5 m"

force := u.NewDerivedQuantity(1234.5, u.Newton)
force.Format(
    u.WithPrecision(1),
    u.WithNumberDelimiter(u.NumberDelimiterUnderscore),
) // "1_234.5 N"
```

| 选项 | 作用 |
|------|------|
| `WithPrecision(n)` | 固定小数位（默认 `PrecisionAuto` = `%g`） |
| `WithNumberDelimiter(d)` | 数值千分位 |
| `WithExpSign` 等 | 单位符号样式（与 `Symbol` 相同） |
| `WithCompoundSymbol(true)` | 使用基本单位组合，而非 SI 专用名 |

量显示推荐只用 `None` / `Comma` / `Underscore`；空格分隔会破坏解析往返。

基本量的 `String()` 与 `Format()` 默认一致。`DerivedQuantity.String()` 内联上标指数（`ExpSignSup`）；`Format()` 默认与 `Symbol()` 一样使用 caret 指数。

## DelimitInt / DelimitFloat

```go
u.DelimitInt(1234567, u.NumberDelimiterComma)      // "1,234,567"
u.DelimitFloat(1234.5, 1, u.NumberDelimiterComma)  // "1,234.5"
```

另支持下划线、空格、细空格、点号等。

## CommaInt / CommaFloat

逗号分隔的便捷封装；`CommaFloat` 等价于 `DelimitFloat(..., Comma)`。

## TrimDelimiter

去掉常见千分位符（`prefix.PrefixParse` 也会自行做同类剥离）。不删除 `.`（可能是小数点）。

## 与词头格式化的关系

| API | 用途 |
|-----|------|
| `Format` / `DelimitFloat` | 量或固定小数位展示 |
| `prefix.PrefixFormat` | SI/IEC 词头缩放 — 见包 [`prefix`](numeric-prefixes/) |

下一步：[SI 基本单位参考 →](../reference/si-base-units/)
