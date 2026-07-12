---
title: "Quantities"
description: "Quantities"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 3
toc: true
---
A **quantity** is a numeric value together with a unit. It is the primary user-facing object for measurement data.

## Quantity struct

The untyped `Quantity` holds any registered base-dimension unit:

```go
q := u.Quantity{Value: 3, Unit: u.Unit(u.Meter)}
fmt.Println(q) // "3 m"
```

For compile-time safety, prefer typed constructors — see [Typed quantities](../typed-quantities/).

## Construction

```go
// Typed (recommended)
d := u.Length(10, u.Kilometer)
t := u.Time(30, u.Minute)

// Untyped
q := u.Quantity{Value: 10, Unit: u.Unit(u.Kilometer)}
```

## Conversion

### Base()

Convert to the SI base unit of the same dimension:

```go
u.Length(1, u.Kilometer).Base() // {1000, meter}
```

### By()

Convert to another unit **within the same dimension**:

```go
u.Length(1000, u.Meter).By(u.Kilometer) // 1 km
```

If units are incompatible or unknown, `By` returns the original quantity unchanged.

## Parsing from text {#parsing-from-text}

Parse functions turn `"value unit"` strings into quantities. Rules:

- The **numeric value must not contain spaces** (invalid: `"1 024 m"`).
- **Spaces between value and unit are optional** (`"10 m"` and `"10m"` both work).
- Thousand separators `,` and `_` in the value are accepted (`"1,024 mm"`).
- Unit symbols include registered base units, SI derived special names (`N`, `Hz`), and common compound formats (`kg·m·s^-2`, `m/s`, superscript forms).

### Which parser to use

| Function | Result type | When |
|----------|-------------|------|
| `QuantityParse` | `Quantity` | Generic; accepts base or derived units |
| `LengthQuantityParse`, `TimeQuantityParse`, … | Typed quantity | You know the base dimension at compile time |
| `DerivedQuantityParse` | `DerivedQuantity` | Compound / SI derived units (`N`, `m/s`) |

Typed parsers reject wrong dimensions with `ErrDimension`:

```go
u.LengthQuantityParse("10 km")  // OK
u.LengthQuantityParse("5 N")    // ErrDimension — force is derived, not length
u.DerivedQuantityParse("5 N")   // OK
u.DerivedQuantityParse("10 m") // ErrDimension — use LengthQuantityParse
```

Each parser has a `MustParse` variant that panics on error.

`QuantityParse` is **not** `PrefixParse`: `"1.5G"` is a numeric prefix, not `1.5` gigameters. See [Numeric prefixes](../../guides/numeric-prefixes/).

### Limitations

- SI special names (`N`, `Hz`, …) are resolved via the symbol registry, not by expanding compound factors.
- Duplicate symbols resolve to the first registrant (`pc` → parsec, not pica).
- Compound symbols shared by different named units (e.g. `s^-1` for Hz vs Bq) map to whichever registered first; use special names (`Hz`, `Bq`) to disambiguate.
- `DerivedUnitParse` accepts `/`, `·`, `*`, whitespace, `^`, and Unicode superscript exponents. Adjacent units need an explicit separator; longest-match keeps `ms` as millisecond.

### Derived unit parsing

When a symbol is not pre-registered, `QuantityParse` falls back to `DerivedUnitParse`, which tokenizes base-unit symbols and rebuilds a `DerivedUnit`:

```go
du, err := u.DerivedUnitParse("km/h") // Length(Kilometer,1) · Time(Hour,-1)
q, err := u.QuantityParse("60 km/h")   // same via fallback
```

Only registered **base-unit** symbols participate (`km`, `h`, `kg`, …). Special names such as `"N"` are not expanded algebraically.

## Compatibility

`Compatible` checks same base dimension:

```go
u.Length(1, u.Meter).Compatible(u.Length(2, u.Kilometer)) // true
u.Length(1, u.Meter).Compatible(u.Time(1, u.Second))       // false
```

## String formatting

```go
fmt.Println(u.Mass(2.5, u.Kilogram)) // "2.5 kg"
```

Derived quantities use compound unit symbols — see [Symbol formatting](../symbol-formatting/).

## DerivedQuantity

When dimensions combine (multiply/divide), the result is a `DerivedQuantity`:

```go
speed := u.Length(60, u.Kilometer).Div(u.Time(1, u.Hour))
// DerivedQuantity{60, km/h unit}

area := u.Length(3, u.Meter).Mul(u.Length(4, u.Meter))
// DerivedQuantity in m²
```

Fields:

| Field | Meaning |
|-------|---------|
| `Value` | Magnitude in the attached unit |
| `Unit` | `*DerivedUnit` (nil → dimensionless display) |

### Derived conversion

```go
si := speed.By(u.SpeedUnit)           // full unit conversion
partial := speed.ByLength(u.Meter)    // change only the length term
```

`By` requires matching derived dimensions and proportional units per dimension. Affine temperature terms block partial conversion.

## Dimensionless quantities

Same-dimension division yields a dimensionless result:

```go
ratio := u.Mass(10, u.Kilogram).Div(u.Mass(2, u.Kilogram))
// Uses NoneUnit — generic unnamed dimensionless unit
```

Or construct directly:

```go
u.Dimensionless(3.14)
```

Next: [Derived units →](../derived-units/)
