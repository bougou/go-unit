---
title: "Typed quantities"
description: "Typed quantities"
sidebar:
  order: 1
---
Typed quantities wrap `Quantity` with dimension-specific unit types so mistakes like adding length to time fail at compile time.

## Type pairs

Each SI base dimension has a unit alias and a quantity struct:

| Dimension | Unit type | Quantity type | Constructor | Parse |
|-----------|-----------|---------------|-------------|-------|
| Length | `LengthUnit` | `LengthQuantity` | `Length(v, u)` / `u.Of(v)` | `LengthQuantityParse(s)` |
| Mass | `MassUnit` | `MassQuantity` | `Mass(v, u)` / `u.Of(v)` | `MassQuantityParse(s)` |
| Time | `TimeUnit` | `TimeQuantity` | `Time(v, u)` / `u.Of(v)` | `TimeQuantityParse(s)` |
| Current | `CurrentUnit` | `CurrentQuantity` | `Current(v, u)` / `u.Of(v)` | `CurrentQuantityParse(s)` |
| Temperature | `TemperatureUnit` | `TemperatureQuantity` | `Temperature(v, u)` / `u.Of(v)` | `TemperatureQuantityParse(s)` |
| Amount | `AmountUnit` | `AmountQuantity` | `Amount(v, u)` / `u.Of(v)` | `AmountQuantityParse(s)` |
| Luminous | `LuminousUnit` | `LuminousQuantity` | `Luminous(v, u)` / `u.Of(v)` | `LuminousQuantityParse(s)` |

Constants (`Meter`, `Kilogram`, `Hour`, …) are typed aliases of `Unit` in `unit_si.go`.

## Basic usage

Two equivalent constructors — package function `(value, unit)` or unit method `Of(value)`. SI decimal factors such as `prefix.Kilo` come from `github.com/bougou/go-unit/pkg/prefix`:

```go
import "github.com/bougou/go-unit/pkg/prefix"

d := u.Length(42, u.Meter.Prefix(prefix.Kilo))
// same as:
d = u.Meter.Prefix(prefix.Kilo).Of(42)

fmt.Println(d) // "42 km"

base := d.Base()           // 42000 m
km := base.By(u.Meter.Prefix(prefix.Kilo)) // back to 42 km
```

For SI derived units the same pair exists: `NewDerivedQuantity(220, u.Volt)` and `u.Volt.Of(220)`. See [Quantities — Construction](../concepts/quantities/#construction).

## Parsing typed quantities

When reading user input or config strings, use the matching typed parser to recover compile-time safety:

```go
d, err := u.LengthQuantityParse("10 km")
t := u.TimeQuantityMustParse("30 min")
```

Wrong dimensions return `ErrDimension` — e.g. `LengthQuantityParse("5 s")` fails because `s` is time, not length.

For derived results (`N`, `m/s`, …), use `DerivedQuantityParse` instead. See [Quantities — parsing from text](../concepts/quantities/#parsing-from-text).

## Same-dimension operations

```go
a := u.Length(10, u.Meter)
b := u.Length(5, u.Meter)
sum := a.Add(b)   // 15 m
diff := a.Sub(b)  // 5 m

a.Compatible(b)   // true
```

Incompatible operands: `Add` / `Sub` return the receiver unchanged.

## Cross-dimension operations

`Mul` and `Div` accept any `derivedQuantity` (typed base quantities or `DerivedQuantity`) and return `DerivedQuantity`:

```go
distance := u.Length(10, u.Meter.Prefix(prefix.Kilo))
duration := u.Time(2, u.Hour)
speed := distance.Div(duration) // 5 km/h

area := u.Length(3, u.Meter).Mul(u.Length(4, u.Meter))
```

## Scalar scale

```go
u.Length(10, u.Meter).MulV(2)  // 20 m
u.Length(10, u.Meter).DivV(4)  // 2.5 m
```

`DivV(0)` returns the receiver unchanged.

## Temperature note

`TemperatureQuantity.Add` / `Sub` use affine conversion through the shared `Quantity` logic — suitable for temperature **intervals** expressed in °C when both operands use compatible affine units.

## When to use untyped Quantity

Use `Quantity` when the dimension is only known at runtime, or for generic utilities. You lose compile-time checking but keep the same conversion rules.

Next: [Arithmetic →](arithmetic/)
