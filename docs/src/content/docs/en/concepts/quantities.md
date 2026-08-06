---
title: "Quantities"
description: "Quantities"
sidebar:
  order: 3
---
A **quantity** is a numeric value together with a unit. It is the primary user-facing object for measurement data.

## Quantity struct

The untyped `Quantity` holds any registered base-dimension unit:

```go
q := u.Quantity{Value: 3, Unit: u.Unit(u.Meter)}
fmt.Println(q) // "3 m"

// Custom display
q.Format(
    u.WithPrecision(2),
    u.WithNumberDelimiter(u.NumberDelimiterComma),
) // "3.00 m"
```

For compile-time safety, prefer typed constructors — see [Typed quantities](../guides/typed-quantities/).

## Construction

There are two equivalent ways to attach a value to a unit. Prefer typed / named units over the untyped `Quantity` struct when the dimension is known at compile time.

### Package constructors `(value, unit)`

```go
import "github.com/bougou/go-unit/pkg/prefix"

// Typed base quantities
d := u.Length(10, u.Meter.Prefix(prefix.Kilo))
t := u.Time(30, u.Minute)
i := u.Current(10, u.Ampere)

// Derived quantities
v := u.NewDerivedQuantity(220, u.Volt)
```

### Unit method `unit.Of(value)`

Every typed base unit and every `*DerivedUnit` has an `Of` method that reads naturally as “this many of that unit”:

```go
d := u.Meter.Prefix(prefix.Kilo).Of(10) // LengthQuantity
t := u.Minute.Of(30)    // TimeQuantity
i := u.Ampere.Of(10)    // CurrentQuantity

v := u.Volt.Of(220)     // DerivedQuantity — same as NewDerivedQuantity(220, u.Volt)
```

`Length(v, u)` and `u.Of(v)` produce the same quantity; likewise `NewDerivedQuantity(v, unit)` and `unit.Of(v)`.

### Untyped

```go
q := u.Quantity{Value: 10, Unit: u.Unit(u.Meter.Prefix(prefix.Kilo))}
```

## Conversion

### Base()

Convert to the SI base unit of the same dimension:

```go
u.Length(1, u.Meter.Prefix(prefix.Kilo)).Base() // {1000, meter}
```

### By()

Convert to another unit **within the same dimension**:

```go
u.Length(1000, u.Meter).By(u.Meter.Prefix(prefix.Kilo)) // 1 km
```

If units are incompatible or unknown, `By` returns the original quantity unchanged.

### Prefix() — SI prefix scaling

`Anchor.Prefix(factor)` scales an anchor unit by an SI decimal prefix from package `prefix` (lazy registration). Prefixed variants are not exported as separate constants:

```go
import "github.com/bougou/go-unit/pkg/prefix"

u.Meter.Prefix(prefix.Kilo)              // km
u.Gram.Prefix(prefix.Milli)              // mg (prefixes attach to gram)
u.Gram.Prefix(prefix.Kilo)               // Kilogram
u.Kilogram.Prefix(prefix.Milli)          // Gram (shortcut)
u.Ohm.Prefix(prefix.Mega)                // MΩ
u.Ohm.Of(1e6).By(u.Ohm.Prefix(prefix.Mega)) // 1 MΩ

u.LengthQuantityParse("10 km")
u.DerivedQuantityParse("1.5 MΩ")
```

Affine units (`Celsius`, …) reject `Prefix`. See also: [Numeric prefixes](../guides/numeric-prefixes/) (plain K/M/G numbers — distinct from unit `P`).

## Parsing from text

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

Each parser has a `MustParse` variant that panics on error (`QuantityMustParse`, `LengthQuantityMustParse`, `DerivedQuantityMustParse`, `DerivedUnitMustParse`, …).

`QuantityParse` is **not** `prefix.PrefixParse`: `"1.5G"` is a numeric prefix, not `1.5` gigameters. See [Numeric prefixes](../guides/numeric-prefixes/).

### Limitations

- SI special names (`N`, `Hz`, …) are resolved via the symbol registry, not by expanding compound factors.
- Duplicate symbols resolve to the first registrant (`pc` → parsec, not pica).
- Compound symbols shared by different named units (e.g. `s^-1` for Hz vs Bq) map to whichever registered first; use special names (`Hz`, `Bq`) to disambiguate.
- `DerivedUnitParse` accepts `/`, `·`, `*`, whitespace, `^`, and Unicode superscript exponents. Adjacent units need an explicit separator; longest-match keeps `ms` as millisecond.

### Derived unit parsing

When a symbol is not pre-registered, `QuantityParse` falls back to `DerivedUnitParse`, which tokenizes base-unit symbols and rebuilds a `DerivedUnit`:

```go
du, err := u.DerivedUnitParse("km/h") // Length(Meter.Prefix(prefix.Kilo),1) · Time(Hour,-1)
q, err := u.QuantityParse("60 km/h")   // same via fallback
```

Only registered **base-unit** symbols participate (`km`, `h`, `kg`, …). Special names such as `"N"` are not expanded algebraically.

## Compatibility

`Compatible` checks same base dimension:

```go
u.Length(1, u.Meter).Compatible(u.Length(2, u.Meter.Prefix(prefix.Kilo))) // true
u.Length(1, u.Meter).Compatible(u.Time(1, u.Second))       // false
```

## String formatting

```go
fmt.Println(u.Mass(2.5, u.Kilogram)) // "2.5 kg"
```

Use `Format` with shared `FormatOption` values for precision, thousands separators, and unit style:

```go
u.Mass(1234.5, u.Kilogram).Format(
    u.WithPrecision(1),
    u.WithNumberDelimiter(u.NumberDelimiterComma),
) // "1,234.5 kg"
```

`DerivedQuantity.String()` uses superscript exponents; `Format()` defaults match `Symbol()` (caret). Details: [Number formatting](../guides/number-formatting/) and [Symbol formatting](symbol-formatting/).

## DerivedQuantity

When dimensions combine (multiply/divide), the result is a `DerivedQuantity`:

```go
speed := u.Length(60, u.Meter.Prefix(prefix.Kilo)).Div(u.Time(1, u.Hour))
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
si := speed.By(u.Speed)           // full unit conversion
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

Next: [Derived units →](derived-units/)
