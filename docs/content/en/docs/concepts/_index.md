---
title: "Concepts"
description: "Dimensions, units, quantities, and derived units."
draft: false
weight: 1
toc: true
---
Before using the API, it helps to understand how the library layers concepts:

```
Dimension  →  Unit  →  Quantity
     ↓           ↓
DerivedDimension → DerivedUnit → DerivedQuantity
```

## Layer 1: Dimensions

A **dimension** describes *what kind* of physical quantity you have — length, mass, time, etc. Dimensions do not carry a numeric scale; they only classify quantities.

- [Dimensions](dimensions/) — seven SI base dimensions and derived dimensions

## Layer 2: Units

A **unit** is a specific way to measure a dimension — meter vs kilometer, celsius vs kelvin. Units carry conversion metadata (`Scale`, `Offset`) to the SI base unit of their dimension.

- [Units & conversion](units-and-conversion/) — registry, affine conversion

## Layer 3: Quantities

A **quantity** pairs a `float64` value with a unit. Operations like `Add`, `By`, `Mul`, and `Div` preserve or combine dimensions correctly.

- [Quantities](quantities/) — `Quantity` and typed wrappers

## Layer 4: Derived units

When a quantity spans multiple base dimensions (speed = length/time), you need a **derived unit** — built from base units with exponents, optionally with an SI special name like `N` or `Hz`.

- [Derived units](derived-units/) — `NewDerivedUnit`, builders, SI named units
- [Intern & registry](intern-and-registry/) — canonical instances
- [Symbol formatting](symbol-formatting/) — display options

## Two APIs for the same model

| Style | When to use |
|-------|-------------|
| **Typed** (`LengthQuantity`, …) | Fixed dimension at compile time; safest |
| **Untyped** (`Quantity`) | Dynamic dimension; simpler but less checked |

Both delegate to the same conversion and arithmetic logic.

## What is *not* in scope

- Automatic inference when the unit symbol is ambiguous (e.g. `pc` for parsec vs pica — first registered wins)
- Currency, custom business units without registration

Unregistered derived symbols built from base units (e.g. `km/h`) are parsed via `DerivedUnitParse` when lookup fails. SI special names (`N`, `Hz`) still require registry lookup.

Numeric **prefixes** (K, M, Mi) are documented separately because they operate on plain numbers, not on dimensional quantities. Use `QuantityParse` for physical units; use `PrefixParse` for plain numeric scaling.

Next: [Dimensions →](dimensions/)
