---
title: "Numeric prefixes"
description: "Numeric prefixes"
sidebar:
  order: 3
---
**Numeric prefixes** scale plain floating-point numbers by SI (1000) or IEC (1024) factors. They live in package **`prefix`** (`github.com/bougou/go-unit/pkg/prefix`) and are intentionally separate from physical units in package `u`.

```go
import (
	"github.com/bougou/go-unit/pkg/prefix"
	u "github.com/bougou/go-unit/pkg/u"
)
```

> **Important: Not physical units**
> `prefix.PrefixParse("1024 G")` parses the **numeric prefix** `G` (10⁹), not gigabytes. Strip unit suffixes like `B`, `bit`, or `/s` before calling `PrefixParse`.
>
## Prefix modes

| Mode | Base | Symbol style | Typical use |
|------|------|--------------|-------------|
| `prefix.SI` | 1000 | K, M, G (no `i`) | metric magnitudes |
| `prefix.IEC` | 1024 | Ki, Mi, Gi (`i` suffix) | binary data sizes |
| `prefix.SI1024` | 1024 | K, M, G (no `i`) | 1024 with SI letters |
| `prefix.Auto` | detect | see below | mixed input |
| `prefix.ForceSI` | force 1000 | both symbol sets | strict decimal |
| `prefix.ForceIEC` | force 1024 | both symbol sets | strict binary |

### Auto mode

- **Parse**: if input ends with `i`, use IEC; otherwise SI
- **Format**: behaves as SI

## Parsing — PrefixParse

```go
v, err := prefix.PrefixParse("1.5G", prefix.SI)    // 1.5e9
v, err := prefix.PrefixParse("1024Ki", prefix.IEC) // 1048576
v, err := prefix.PrefixParse("568m", prefix.SI)    // 0.568
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

Thousands separators in the numeric part are stripped automatically (comma, underscore, space, thin space). Package `u` also exposes `TrimDelimiter` for the same kind of cleanup outside prefix parsing.

### Errors

| Error | Cause |
|-------|-------|
| `prefix.ErrSyntax` | Malformed input |
| `prefix.ErrInvalidMode` | Unsupported mode |
| `prefix.ErrInvalidSymbol` | Unknown prefix letter |

## Formatting — PrefixFormat

```go
prefix.PrefixFormat(999_000_000, prefix.SI)  // "999 M"
prefix.PrefixFormat(1048576, prefix.IEC)     // "1 Mi"
prefix.PrefixFormat(0.568, prefix.SI)        // "568m"
```

### PrefixFormat2

Returns number and prefix separately:

```go
num, sym := prefix.PrefixFormat2(1048576, prefix.IEC)
// num == "1", sym == "Mi"
```

### Format options

| Option | Effect |
|--------|--------|
| `prefix.WithPrefixSpace(true)` | Space between number and prefix (`"1 Mi"`) |
| `prefix.WithPrefixPrecision(n)` | Fixed decimal places |
| `prefix.WithRoundMethod(m)` | Floor/round/ceil when precision is 0 |
| `prefix.WithRoundDifference(d)` | Threshold for `RoundMethodDifference` |
| `prefix.WithPrefix(sym)` | Force a specific `PrefixSymbol` |

## Scale constants

SI decimal factors (`prefix.Kilo`, `prefix.Mega`, `prefix.Micro`, …) are `prefix.SIPrefix` constants; IEC binary factors (`prefix.Kibi`, `prefix.Mebi`, …) remain `float64` in `pkg/prefix/prefix.go`. `Unit.Prefix` / `ByPrefix` accept `SIPrefix` only.

`PrefixSymbol` is the single-character type used when forcing a letter via `WithPrefix`.

For **byte** multiples with named constants (`prefix.KB`, `prefix.KiB`, …), see [Byte constants](../reference/byte-constants/).

## vs physical unit `Prefix()`

`PrefixParse` / `PrefixFormat` scale **plain numbers**. Physical units use a different API — `Anchor.Prefix(factor)` — which returns a real unit:

```go
u.Meter.Prefix(prefix.Kilo)  // LengthUnit km
u.Gram.Prefix(prefix.Milli)  // MassUnit mg
u.Ohm.Prefix(prefix.Mega)    // *DerivedUnit MΩ
```

Use `QuantityParse("10 km")` or `DerivedQuantityParse("1 MΩ")` for value+unit strings. See [Quantities — Prefix()](../concepts/quantities/).

## vs physical unit parsing

`PrefixParse` scales **plain numbers** with K/M/G-style prefixes. It does not attach a physical unit.

For strings like `"10 km"` or `"5 N"`, use [quantity parsing](../concepts/quantities/#parsing-from-text) (`QuantityParse`, `LengthQuantityParse`, `DerivedQuantityParse`) instead.

Next: [Number formatting →](number-formatting/)
