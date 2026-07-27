---
title: "Number formatting"
description: "Number formatting"
sidebar:
  order: 4
---
`go-unit` formats numbers in two layers: **quantity display** (`Format`) and standalone helpers (`DelimitInt`, `DelimitFloat`).

## Quantity Format

`Quantity`, typed quantities, and `DerivedQuantity` support `Format` with shared `FormatOption` functions:

```go
q := u.Quantity{Value: 1234.5, Unit: u.Unit(u.Meter)}
q.Format() // "1234.5 m"  (same as String)
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

| Option | Effect |
|--------|--------|
| `WithPrecision(n)` | Fixed decimal places (`PrecisionAuto` / default = `%g`) |
| `WithNumberDelimiter(d)` | Thousands separator on the value |
| `WithExpSign` / `WithMulSign` / `WithDivSign` / … | Unit symbol style (same as `Symbol`) |
| `WithCompoundSymbol(true)` | Compound base-unit expression instead of SI named symbol |

For quantity values prefer `NumberDelimiterNone`, `NumberDelimiterComma`, or `NumberDelimiterUnderscore`. Space separators break round-trip parsing.

`String()` uses the same defaults as `Format()` for base quantities. For `DerivedQuantity`, `String()` inlines superscript exponents (`ExpSignSup`); `Format()` defaults to caret exponents like `Symbol()`.

## DelimitInt / DelimitFloat

Standalone helpers (no physical unit):

```go
u.DelimitInt(1234567, u.NumberDelimiterComma)        // "1,234,567"
u.DelimitInt(1234567, u.NumberDelimiterUnderscore)   // "1_234_567"
u.DelimitFloat(1234.5, 1, u.NumberDelimiterComma)    // "1,234.5"
```

Also available: space, thin space (U+2009), dot, and `NumberDelimiterNone`.

## CommaInt / CommaFloat

Convenience wrappers using comma separators:

```go
u.CommaInt(1234567)      // "1,234,567"
u.CommaFloat(1234.5, 1)  // "1,234.5"  (== DelimitFloat with comma)
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
| `Format` / `DelimitFloat` | Display a quantity or fixed-precision number |
| `PrefixFormat` | Scale values with SI/IEC prefixes (K, M, Mi, …) |

Next: [SI base units reference →](../reference/si-base-units/)
