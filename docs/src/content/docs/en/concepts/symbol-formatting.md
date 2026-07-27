---
title: "Symbol formatting"
description: "Symbol formatting"
sidebar:
  order: 6
---
Unit and dimension symbols are rendered on demand via `Symbol()` methods and `FormatOption` functions.

## Default behavior

| Type | Default output |
|------|----------------|
| Registered unit | Registered `Symbol` field (`"km"`, `"°C"`) |
| Derived unit | SI named symbol when set (e.g. `"N"`); otherwise compound base-unit expression |
| Derived dimension | Single-letter exponents |

```go
u.Unit(u.Meter.Prefix(u.Kilo)).Symbol() // "km"
u.Newton.Symbol()                       // "N"
u.Newton.NamedSymbol()                  // "N" (without SI prefix on the unit)
u.Newton.Symbol(u.WithCompoundSymbol(true)) // "kg·m·s^-2"
```

## Symbol options

### WithCompoundSymbol

Request the compound base-unit expression instead of the SI named symbol:

```go
u.Newton.Symbol() // "N" (default)
u.Newton.Symbol(u.WithCompoundSymbol(true)) // "kg·m·s^-2"
u.Hertz.Symbol(u.WithCompoundSymbol(true))  // "s^-1"
```

Default is `false` — when `Named()` was used, `Symbol()` prefers the named form.

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
speed := u.NewDerivedUnit().Length(u.Meter.Prefix(u.Kilo), 1).Time(u.Hour, -1)

speed.Symbol() // "km·h^-1"

speed.Symbol(
    u.WithExpSign(u.ExpSignSup),
    u.WithDivSign(u.DivSignSlash),
) // would need slash mode: "km/h" via WithDivSign alone

speed.Symbol(u.WithExpSign(u.ExpSignSup)) // "km·h⁻¹"
speed.Symbol(u.WithDivSign(u.DivSignSlash)) // "km/h"
```

`DerivedQuantity.String()` inlines superscript exponents; `Format()` defaults match `Symbol()` (caret).
The same `FormatOption` functions (`WithExpSign`, `WithCompoundSymbol`, …) configure both `Symbol` and `Format`.

## Combining options

Options compose — later functions in the slice override earlier ones:

```go
u.Newton.Symbol(
    u.WithCompoundSymbol(true),
    u.WithExpSign(u.ExpSignSup),
    u.WithDivSign(u.DivSignSlash),
) // "kg·m/s²"
```

Use the same options on quantities:

```go
u.NewDerivedQuantity(10, u.Newton).Format(
    u.WithPrecision(2),
) // "10.00 N"

u.NewDerivedQuantity(10, u.Newton).Format(
    u.WithPrecision(2),
    u.WithCompoundSymbol(true),
) // "10.00 kg·m·s^-2"
```

Next: [Typed quantities guide →](../guides/typed-quantities/)
