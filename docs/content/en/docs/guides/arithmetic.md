---
title: "Arithmetic"
description: "Arithmetic"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 2
toc: true
---
Quantity arithmetic follows dimensional analysis: add/subtract require matching dimensions; multiply/divide combine dimensions.

## Addition and subtraction

Operands must share the same (derived) dimension. The second operand converts to the first's unit via the base unit:

```go
a := u.Length(1, u.Kilometer)
b := u.Length(500, u.Meter)
sum := a.Add(b) // 1.5 km — 500 m converted to 0.5 km
```

Incompatible dimensions → receiver returned unchanged:

```go
u.Length(1, u.Meter).Add(u.Time(1, u.Second)) // unchanged
```

### Derived quantities

```go
speedUnit := u.NewDerivedUnit().Length(u.Kilometer, 1).Time(u.Hour, -1)
a := u.NewDerivedQuantity(60, speedUnit)
b := u.NewDerivedQuantity(30, speedUnit)
a.Add(b) // 90 km/h (same unit)
```

Values normalize through SI base composition internally before re-scaling to the receiver's unit.

## Multiplication

Multiplying quantities **adds** derived-dimension exponents and **multiplies** numeric values (with base-unit normalization):

```go
u.Length(2, u.Kilometer).Mul(u.Length(3, u.Meter))
// 6000 m·m → derived unit km·m or equivalent
```

Unit choice per dimension prefers units already present in operands, then falls back to SI base units.

## Division

Dividing quantities **subtracts** exponents. **Same-dimension** division yields a **dimensionless** quantity with `NoneUnit`:

```go
ratio := u.Mass(10, u.Kilogram).Div(u.Mass(2, u.Kilogram))
// Dimensionless, Value: 5
```

Division by zero (zero divisor quantity or `DivV(0)`) returns the receiver unchanged.

## Speed example (end-to-end)

```go
trip := u.Length(120, u.Kilometer)
drive := u.Time(2, u.Hour)
avg := trip.Div(drive)

fmt.Println(avg) // 60 km·h⁻¹

si := avg.By(u.SpeedUnit)
// ~16.667 m/s
```

## Proportional conversion check

Internal helper `unitsProportional` verifies two units convert with a constant ratio through the base unit. This excludes affine pairs like °C/K for partial `ByLength`-style conversions on derived units.

## What arithmetic does not do

- Automatic simplification to SI special names (force stays as kg·m·s⁻² unless you `By(ForceUnit)` or format with `WithNamedSymbol`)
- Unchecked float overflow/underflow handling
- Uncertainty propagation

Next: [Numeric prefixes →](../numeric-prefixes/)
