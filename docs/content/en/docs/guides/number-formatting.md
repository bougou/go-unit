---
title: "Number formatting"
description: "Number formatting"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 4
toc: true
---
Separate from physical units, `go-unit` provides helpers for human-readable **integer and float formatting** with thousands separators.

## DelimitInt

Format an `int64` with a configurable delimiter:

```go
u.DelimitInt(1234567, u.NumberDelimiterComma)      // "1,234,567"
u.DelimitInt(1234567, u.NumberDelimiterUnderscore)   // "1_234_567"
u.DelimitInt(1234567, u.NumberDelimiterSpace)      // "1 234 567"
u.DelimitInt(1234567, u.NumberDelimiterThinSpace)    // thin space (U+2009)
u.DelimitInt(1234567, u.NumberDelimiterDot)          // "1.234.567"
u.DelimitInt(1234567, u.NumberDelimiterNone)         // "1234567"
```

Negative numbers preserve a leading minus sign.

## CommaInt / CommaFloat

Convenience wrappers using comma separators:

```go
u.CommaInt(1234567)      // "1,234,567"
u.CommaFloat(1234.5, 1)  // "1,234.5"
```

## TrimDelimiter

Remove common thousands separators before numeric parsing — used internally by `PrefixParse`:

```go
u.TrimDelimiter("1,234,567") // "1234567"
```

Dot separators are **not** removed (they may be decimal points).

## Relation to prefix formatting

| API | Purpose |
|-----|---------|
| `DelimitInt` / `CommaFloat` | Display large counts or fixed-precision decimals |
| `PrefixFormat` | Scale values with SI/IEC prefixes (K, M, Mi, …) |

Use both when presenting dashboard values: prefix for magnitude, delimiter for readability inside the mantissa if needed.

Next: [SI base units reference →](../si-base-units/)
