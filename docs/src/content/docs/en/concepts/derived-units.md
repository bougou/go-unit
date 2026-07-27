---
title: "Derived units"
description: "Derived units"
sidebar:
  order: 4
---
A **derived unit** expresses a compound measurement as base units raised to exponents — for example km¹·h⁻¹ for speed, or kg¹·m¹·s⁻² for force.

## Building derived units

Use `NewDerivedUnit()` and chain dimension setters:

```go
speed := u.NewDerivedUnit().
    Length(u.Meter.Prefix(u.Kilo), 1).
    Time(u.Hour, -1)

force := u.NewDerivedUnit().
    Mass(u.Kilogram, 1).
    Length(u.Meter, 1).
    Time(u.Second, -2)
```

| Method | Dimension | exp > 0 | exp < 0 |
|--------|-----------|---------|---------|
| `Length(u, exp)` | L | numerator | denominator |
| `Mass(u, exp)` | M | numerator | denominator |
| `Time(u, exp)` | T | numerator | denominator |
| `Current(u, exp)` | I | … | … |
| `Temperature(u, exp)` | Θ | … | … |
| `Amount(u, exp)` | N | … | … |
| `Luminous(u, exp)` | J | … | … |

Setting `exp == 0` clears that dimension term. Unknown units or wrong dimensions **panic** — derived-unit construction is intended for programmer-defined compositions, not untrusted input.

## SI special names

When a derived unit has an SI special name, set it with `Named`:

```go
newton := u.NewDerivedUnit().
    Mass(u.Kilogram, 1).
    Length(u.Meter, 1).
    Time(u.Second, -2).
    Named("N")
```

Pre-registered globals in `unit_si_derived.go` include `Newton`, `Pascal`, `Hertz`, `Speed`, etc. See [SI derived units](../reference/si-derived-units/).

### Prefix() — SI prefixes on special names

Named derived units support `P` to scale by an SI decimal factor. The result is an interned `*DerivedUnit`:

```go
u.Ohm.Prefix(u.Mega)   // MΩ
u.Newton.Prefix(u.Kilo) // kN
u.Ohm.Prefix(u.Mega).Prefix(u.Micro) // Ohm again (factors cancel)
```

`FactorToBase` includes the prefix factor; `By` converts between `Ohm` and `Ohm.Prefix(Mega)`. Parsing accepts forms such as `"1 MΩ"` and `"2 kN"`.

## Inspecting a derived unit

```go
force := u.Newton

force.Dim()           // DerivedDimension{M:1, L:1, T:-2}
force.FactorToBase()  // multiplier to SI base composition
force.Key()           // registry identifier
force.SpecialSymbol() // "N"
```

### SI()

Rewrite using SI base units per dimension:

```go
u.NewDerivedUnit().Length(u.Meter.Prefix(u.Kilo), 1).Time(u.Hour, -1).SI()
// → m¹·s⁻¹
```

## Derived quantities

Attach a value with either form — they are equivalent — or get one from `Mul` / `Div` on typed quantities:

```go
speedUnit := u.NewDerivedUnit().Length(u.Meter.Prefix(u.Kilo), 1).Time(u.Hour, -1)

v := u.NewDerivedQuantity(60, speedUnit) // (value, unit)
v = speedUnit.Of(60)                     // unit.Of(value) — same result
fmt.Println(v.String())                  // "60 km·h⁻¹"

f := u.Newton.Of(100)                    // 100 N
f = u.NewDerivedQuantity(100, u.Newton)  // same
```

Parse from text with `DerivedQuantityParse` when the input uses a registered derived symbol or a parseable compound such as `km/h`:

```go
force, err := u.DerivedQuantityParse("5 N")
speed, err := u.DerivedQuantityParse("60 km/h") // via DerivedUnitParse fallback
```

Base-dimension strings like `"10 m"` are rejected — use typed parsers instead. Details: [Parsing from text](quantities/#parsing-from-text).

## Predefined dimension variables

Package variables like `DimForce`, `DimEnergy`, and `DimSpeed` document common derived dimensions and match the registered SI derived units.

## Same composition, different names

Dimensionless units can share the same base composition but differ by special name — rad vs sr vs the unnamed `NoneUnit`. See [Intern & registry](intern-and-registry/).

Next: [Intern & registry →](intern-and-registry/)
