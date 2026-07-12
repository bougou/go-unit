---
title: "Typed quantities"
description: "Typed quantities"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 1
toc: true
---
Typed quantities wrap `Quantity` with dimension-specific unit types so mistakes like adding length to time fail at compile time.

## Type pairs

Each SI base dimension has a unit alias and a quantity struct:

| Dimension | Unit type | Quantity type | Constructor | Parse |
|-----------|-----------|---------------|-------------|-------|
| Length | `LengthUnit` | `LengthQuantity` | `Length(v, u)` | `LengthQuantityParse(s)` |
| Mass | `MassUnit` | `MassQuantity` | `Mass(v, u)` | `MassQuantityParse(s)` |
| Time | `TimeUnit` | `TimeQuantity` | `Time(v, u)` | `TimeQuantityParse(s)` |
| Current | `CurrentUnit` | `CurrentQuantity` | `Current(v, u)` | `CurrentQuantityParse(s)` |
| Temperature | `TemperatureUnit` | `TemperatureQuantity` | `Temperature(v, u)` | `TemperatureQuantityParse(s)` |
| Amount | `AmountUnit` | `AmountQuantity` | `Amount(v, u)` | `AmountQuantityParse(s)` |
| Luminous | `LuminousUnit` | `LuminousQuantity` | `Luminous(v, u)` | `LuminousQuantityParse(s)` |

Constants (`Meter`, `Kilogram`, `Hour`, …) are typed aliases of `Unit` in `unit_si.go`.

## Basic usage

```go
d := u.Length(42, u.Kilometer)
fmt.Println(d) // "42 km"

base := d.Base()              // 42000 m
km := base.By(u.Kilometer) // back to 42 km
```

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
distance := u.Length(10, u.Kilometer)
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

Next: [Arithmetic →](../arithmetic/)
