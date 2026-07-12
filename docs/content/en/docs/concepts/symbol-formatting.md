---
title: "Symbol formatting"
description: "Symbol formatting"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 6
toc: true
---
Unit and dimension symbols are rendered on demand via `Symbol()` methods and `SymbolOption` functions.

## Default behavior

| Type | Default output |
|------|----------------|
| Registered unit | Registered `Symbol` field (`"km"`, `"°C"`) |
| Derived unit | Compound base-unit expression |
| Derived dimension | Single-letter exponents |

```go
u.Unit(u.Kilometer).Symbol()           // "km"
u.ForceUnit.Symbol()                      // "kg·m·s^-2"
u.ForceUnit.Symbol(u.WithNamedSymbol(true)) // "N"
```

## Symbol options

### WithNamedSymbol

Use the SI special name when set:

```go
u.ForceUnit.Symbol(u.WithNamedSymbol(true)) // "N"
u.HertzUnit.Symbol(u.WithNamedSymbol(true)) // "Hz"
```

Default is `false` — compound form is shown.

### WithExpSign

| Value | Example |
|-------|---------|
| `ExpSignCarat` (default) | `m^2`, `h^-1` |
| `ExpSignSup` | `m²`, `h⁻¹` |

### WithMulSign

| Value | Example |
|-------|---------|
| `MulSignDot` (default) | `km·h^-1` |
| `MulSignSpace` | `km h^-1` |
| `MulSignStar` | `km*h^-1` |

### WithDivSign

| Value | Example |
|-------|---------|
| `DivSignNegative` (default) | `km·h⁻¹` |
| `DivSignSlash` | `km/h`, `kg·m/s^2` |

### WithDimOrder

Order of dimension letters in **dimension** symbols (not unit symbols):

| Value | Order |
|-------|-------|
| `DimOrderMLT` (default) | M, L, T, I, H, N, J |
| `DimOrderTML` | T, L, M, I, H, N, J (ISO 80000 style) |

## Examples

```go
speed := u.NewDerivedUnit().Length(u.Kilometer, 1).Time(u.Hour, -1)

speed.Symbol() // "km·h^-1"

speed.Symbol(
    u.WithExpSign(u.ExpSignSup),
    u.WithDivSign(u.DivSignSlash),
) // would need slash mode: "km/h" via WithDivSign alone

speed.Symbol(u.WithExpSign(u.ExpSignSup)) // "km·h⁻¹"
speed.Symbol(u.WithDivSign(u.DivSignSlash)) // "km/h"
```

Derived quantities call `Symbol(WithExpSign(ExpSignSup))` in their `String()` method for typographic exponents.

## Combining options

Options compose — later functions in the slice override earlier ones when applied through `symbolOptions`:

```go
u.ForceUnit.Symbol(
    u.WithNamedSymbol(false),
    u.WithExpSign(u.ExpSignSup),
    u.WithDivSign(u.DivSignSlash),
) // "kg·m/s²"
```

Next: [Typed quantities guide →](../typed-quantities/)
