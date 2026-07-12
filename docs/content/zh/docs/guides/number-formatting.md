---
title: "数字格式化"
description: "数字格式化"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 4
toc: true
---
除物理单位外，库提供带**千分位分隔符**的整数/浮点格式化。

## DelimitInt

```go
u.DelimitInt(1234567, u.NumberDelimiterComma)    // "1,234,567"
u.DelimitInt(1234567, u.NumberDelimiterUnderscore) // "1_234_567"
```

支持逗号、下划线、空格、细空格（U+2009）、点号等；`NumberDelimiterNone` 不分隔。

## CommaInt / CommaFloat

```go
u.CommaInt(1234567)
u.CommaFloat(1234.5, 1) // "1,234.5"
```

## TrimDelimiter

去掉常见千分位符，`PrefixParse` 内部使用。不删除 `.`（可能是小数点）。

## 与词头格式化的关系

| API | 用途 |
|-----|------|
| `DelimitInt` / `CommaFloat` | 大整数或固定小数位展示 |
| `PrefixFormat` | SI/IEC 词头缩放 |

下一步：[SI 基本单位参考 →](../si-base-units/)
