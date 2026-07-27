---
title: "Numeric prefixes"
description: "Numeric prefixes"
sidebar:
  order: 3
---
**Numeric prefixes** scale plain floating-point numbers by SI (1000) or IEC (1024) factors. They are intentionally separate from physical units.

> **Important: Not physical units**
> `PrefixParse("1024 G")` parses the **numeric prefix** `G` (10⁹), not gigabytes. Strip unit suffixes like `B`, `bit`, or `/s` before calling `PrefixParse`.
>
## Prefix modes

| Mode | Base | Symbol style | Typical use |
|------|------|--------------|-------------|
| `SI` | 1000 | K, M, G (no `i`) | metric magnitudes |
| `IEC` | 1024 | Ki, Mi, Gi (`i` suffix) | binary data sizes |
| `SI1024` | 1024 | K, M, G (no `i`) | 1024 with SI letters |
| `Auto` | detect | see below | mixed input |
| `ForceSI` | force 1000 | both symbol sets | strict decimal |
| `ForceIEC` | force 1024 | both symbol sets | strict binary |

### Auto mode

- **Parse**: if input ends with `i`, use IEC; otherwise SI
- **Format**: behaves as SI

## Parsing — PrefixParse

```go
v, err := u.PrefixParse("1.5G", u.SI)    // 1.5e9
v, err := u.PrefixParse("1024Ki", u.IEC) // 1048576
v, err := u.PrefixParse("568m", u.SI)    // 0.568
```

Valid inputs:

- Plain number: `"1024"`, `"0.5"`
- Number + prefix: `"1024 G"`, `"1024Ki"`

Invalid — strip the unit part first:

| Input | Use instead |
|-------|-------------|
| `"1024 MiB"` | `"1024 Mi"` |
| `"1024 Gb/s"` | `"1024 G"` |
| `"1024 Bytes"` | `"1024"` |

Thousands separators in the numeric part are stripped via `TrimDelimiter` (comma, underscore, space, thin space).

### Errors

| Error | Cause |
|-------|-------|
| `ErrSyntax` | Malformed input |
| `ErrInvalidMode` | Unsupported mode |
| `ErrInvalidSymbol` | Unknown prefix letter |

## Formatting — PrefixFormat

```go
u.PrefixFormat(999_000_000, u.SI)  // "999 M"
u.PrefixFormat(1048576, u.IEC)     // "1 Mi"
u.PrefixFormat(0.568, u.SI)        // "568m"
```

### PrefixFormat2

Returns number and prefix separately:

```go
num, prefix := u.PrefixFormat2(1048576, u.IEC)
// num == "1", prefix == "Mi"
```

### Format options

| Option | Effect |
|--------|--------|
| `WithPrefixSpace(true)` | Space between number and prefix (`"1 Mi"`) |
| `WithPrefixPrecision(n)` | Fixed decimal places |
| `WithRoundMethod(m)` | Floor/round/ceil when precision is 0 |
| `WithRoundDifference(d)` | Threshold for `RoundMethodDifference` |
| `WithPrefix(sym)` | Force a specific prefix symbol |

## Scale constants

SI decimal factors (`Kilo`, `Mega`, `Micro`, …) and IEC binary factors (`Kibi`, `Mebi`, …) are exported as `float64` constants in `prefix.go`.

For **byte** multiples with named constants (`KB`, `KiB`, …), see [Byte constants](../reference/byte-constants/).

## vs physical unit `Prefix()`

`PrefixParse` / `PrefixFormat` scale **plain numbers**. Physical units use a different API — `Anchor.Prefix(prefix)` — which returns a real unit:

```go
u.Meter.Prefix(u.Kilo)  // LengthUnit km
u.Gram.Prefix(u.Milli)  // MassUnit mg
u.Ohm.Prefix(u.Mega)    // *DerivedUnit MΩ
```

Use `QuantityParse("10 km")` or `DerivedQuantityParse("1 MΩ")` for value+unit strings. See [Quantities — Prefix()](../concepts/quantities/).

## vs physical unit parsing

`PrefixParse` scales **plain numbers** with K/M/G-style prefixes. It does not attach a physical unit.

For strings like `"10 km"` or `"5 N"`, use [quantity parsing](../concepts/quantities/#parsing-from-text) (`QuantityParse`, `LengthQuantityParse`, `DerivedQuantityParse`) instead.

Next: [Number formatting →](number-formatting/)
