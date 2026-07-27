---
title: "Dimensions"
description: "Dimensions"
sidebar:
  order: 1
---
A **dimension** classifies physical quantities by kind. `go-unit` follows the seven SI base dimensions and represents compound (derived) dimensions as exponents on those bases.

## SI base dimensions

The `Dimension` type identifies one of seven independent base dimensions:

| Constant | Symbol | Name | SI base unit |
|----------|--------|------|--------------|
| `DimLength` | L | length | meter (`Meter`) |
| `DimMass` | M | mass | kilogram (`Kilogram`) |
| `DimTime` | T | time | second (`Second`) |
| `DimCurrent` | I | electric current | ampere (`Ampere`) |
| `DimTemperature` | Θ | thermodynamic temperature | kelvin (`Kelvin`) |
| `DimAmount` | N | amount of substance | mole (`Mole`) |
| `DimLuminous` | J | luminous intensity | candela (`Candela`) |

```go
dim := u.DimLength
fmt.Println(dim.String())  // "length"
fmt.Println(dim.Symbol())  // "L"
fmt.Println(dim.Base())    // Unit("meter")
```

> **Note: Temperature symbol**
> In dimensional analysis literature, thermodynamic temperature uses **Θ** (theta). The derived-dimension struct field is named `H` instead: Greek theta is harder to type on a keyboard than the Latin letter H; **H** suggests **Heat**, which is semantically close to temperature; and the shape of **H** resembles **Θ**.
>
## Derived dimensions

A **derived dimension** expresses a compound quantity as exponents on the seven base dimensions. The struct fields `L`, `M`, `T`, `I`, `H`, `N`, `J` hold signed exponents; omitted dimensions are zero.

```go
speed := u.DerivedDimension{L: 1, T: -1}        // L¹T⁻¹
force := u.DerivedDimension{M: 1, L: 1, T: -2}  // M¹L¹T⁻²
```

Common examples:

| Quantity | DerivedDimension | Example units |
|----------|------------------|---------------|
| Speed | `{L:1, T:-1}` | m/s, km/h |
| Force | `{M:1, L:1, T:-2}` | N = kg·m/s² |
| Pressure | `{M:1, L:-1, T:-2}` | Pa |
| Energy | `{M:1, L:2, T:-2}` | J |

Package-level variables such as `DimForce`, `DimPressure`, and `DimSpeed` are predefined in `unit_si_derived.go`.

## Dimension algebra

Derived-dimension arithmetic mirrors quantity arithmetic and the **principle of dimensional homogeneity**:

```go
d1 := u.DimForce                    // M¹L¹T⁻²
d2 := u.DimSpeed                    // L¹T⁻¹

product := d1.Mul(d2)               // multiply → add exponents; always defined
quotient := d1.Div(d2)              // divide → subtract exponents; always defined

sum, ok := d1.Add(u.DimForce)       // add/sub: dimensions must match
diff, ok := d1.Sub(u.DimForce)      // when ok, result dimension is still d1
_, ok = d1.Add(d2)                  // mismatched → ok == false

equal := d1.Equal(u.DimForce)
```

| Operation | Method | Rule |
|-----------|--------|------|
| Add / subtract | `Add` / `Sub` | Meaningful only when dimensions are identical; success returns that dimension with `ok=true` |
| Multiply / divide | `Mul` / `Div` | Exponents add / subtract; always defined algebraically. Whether the result names a familiar quantity is physical interpretation |

## Rendering dimension symbols

`DerivedDimension.Symbol()` renders a human-readable dimension string:

```go
u.DimForce.Symbol()                              // "M·L·T^-2"
u.DimForce.Symbol(u.WithExpSign(u.ExpSignSup)) // "M·L·T⁻²"
```

See [Symbol formatting](symbol-formatting/) for all options.

## Dimensionless quantities

An empty `DerivedDimension{}` represents a dimensionless quantity — for example the result of dividing two lengths, or plane angle (rad) and solid angle (sr), which are treated as dimensionless in this library.

Next: [Units & conversion →](units-and-conversion/)
