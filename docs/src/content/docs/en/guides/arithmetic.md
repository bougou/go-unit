---
title: "Arithmetic"
description: "Arithmetic"
sidebar:
  order: 2
---
Quantity arithmetic follows dimensional analysis: add/subtract require matching dimensions; multiply/divide combine dimensions.

## Addition and subtraction

Operands must share the same (derived) dimension. The second operand converts to the first's unit via the base unit:

```go
a := u.Length(1, u.Meter.Prefix(u.Kilo))
b := u.Length(500, u.Meter)
sum := a.Add(b) // 1.5 km — 500 m converted to 0.5 km
```

Incompatible dimensions → receiver returned unchanged:

```go
u.Length(1, u.Meter).Add(u.Time(1, u.Second)) // unchanged
```

### Derived quantities

```go
speedUnit := u.NewDerivedUnit().Length(u.Meter.Prefix(u.Kilo), 1).Time(u.Hour, -1)
a := u.NewDerivedQuantity(60, speedUnit)
b := speedUnit.Of(30)
a.Add(b) // 90 km/h (same unit)
```

`DerivedQuantity` add/sub also requires **proportionally related** units on every base dimension (same rule as `By`). Same derived dimension with an affine pair (°C vs K) leaves the receiver unchanged. Proportional units may differ, e.g. `36 km/h + 10 m/s → 72 km/h`.

## Multiplication

Multiplying quantities **adds** derived-dimension exponents and **multiplies** numeric values (with base-unit normalization):

```go
u.Length(2, u.Meter.Prefix(u.Kilo)).Mul(u.Length(3, u.Meter))
// 6000 m·m → derived unit km·m or equivalent
```

Unit choice per dimension prefers units already present in operands, then falls back to SI base units.

For `DerivedQuantity`, every base dimension active in **both** operands must use proportionally related units; otherwise the receiver is returned unchanged.

## Division

Dividing quantities **subtracts** exponents. **Same-dimension** division yields a **dimensionless** quantity with `NoneUnit`:

```go
ratio := u.Mass(10, u.Kilogram).Div(u.Mass(2, u.Kilogram))
// Dimensionless, Value: 5
```

`DerivedQuantity` likewise requires proportional units on overlapping base dimensions. Division by zero (zero divisor quantity or `DivV(0)`) returns the receiver unchanged.

## Speed example (end-to-end)

```go
trip := u.Length(120, u.Meter.Prefix(u.Kilo))
drive := u.Time(2, u.Hour)
avg := trip.Div(drive)

fmt.Println(avg) // 60 km·h⁻¹

si := avg.By(u.Speed)
// ~16.667 m/s
```

## Proportional conversion check

Internal helper `unitsProportional` verifies two units convert with a constant ratio through the base unit. The same check applies to `By*` and to all `DerivedQuantity` arithmetic: shared base dimensions must use proportional units, excluding affine pairs such as °C/K.

## What arithmetic does not do

- Automatic simplification to SI special names (force stays as kg·m·s⁻² unless you `By(Newton)`; `Symbol()` / `Format()` default to the named form when set)
- Unchecked float overflow/underflow handling
- Uncertainty propagation

Next: [Numeric prefixes →](numeric-prefixes/)
