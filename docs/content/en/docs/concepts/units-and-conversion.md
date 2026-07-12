---
title: "Units & conversion"
description: "Units & conversion"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 2
toc: true
---
A **unit** is the concrete measure attached to a dimension. `go-unit` registers SI and commonly used units at init time and converts between them via each dimension's SI **base unit**.

## Unit identity

`Unit` is a string type — the stable internal identifier:

```go
km := u.Unit(u.Kilometer)
fmt.Println(km.Symbol()) // "km"
```

Display symbols come from the registered `unitDef` or from a `DerivedUnit` when the unit is compound.

## Registry and unitDef

Each registered unit stores:

| Field | Role |
|-------|------|
| `Dimension` | Which base dimension the unit belongs to |
| `Unit` | Internal registry key |
| `Symbol` | Short display symbol (`"km"`, `"°C"`) |
| `Name` | Full English name |
| `Scale` | Multiplicative factor to base unit (must be > 0) |
| `Offset` | Additive constant after scaling |

### Ratio conversion (Offset = 0)

Pure scale change — most units:

```
base = value × Scale
value = base / Scale
```

Example: 1 km → 1000 m (`Scale = 1000`, `Offset = 0`).

### Affine conversion (Offset ≠ 0)

Same scale as kelvin but different zero point — temperature:

```
base = value × Scale + Offset
value = (base − Offset) / Scale
```

Example: Celsius → Kelvin uses `Scale = 1`, `Offset = 273.15`.

> **Warning: Proportional vs affine**
> `°C` and `K` share a dimension but are **not proportional** (offset ≠ 0). The library rejects naive ratio conversion between them. Use `Base()` / `By()` which apply the full affine map.
>
## Conversion API

```go
// From a registered base-dimension unit
q := u.Length(1, u.Hour)
base := q.Base()                    // 3600 s
back := base.By(u.Hour)          // 1 h

// Low-level on Unit
val, ok := u.Unit(u.Kilometer).ToBase(2.5)   // 2500
back, ok := u.Unit(u.Kilometer).FromBase(2500) // 2.5
```

`Def()` returns the full definition when the unit is registered:

```go
def, ok := u.Unit(u.Celsius).Def()
// def.Scale == 1, def.Offset == 273.15
```

## Registered unit catalog

All SI base-dimension constants live in `unit_si.go`, grouped by dimension:

- Length — SI prefixes, imperial, nautical, Chinese traditional (chi, cun, …)
- Mass — SI, tonnes, avoirdupois, troy, jin/liang
- Time — SI prefixes, minute through year
- Current, temperature, amount, luminous — SI and common variants

See [SI base units](../si-base-units/) for the full list.

## Derived vs registered units

| Kind | Example | Lookup |
|------|---------|--------|
| Registered base unit | `Kilometer` | `Unit.Def()` |
| Derived compound unit | km/h, N | `Unit.DerivedUnit()` after `Intern` |

Compound units are covered in [Derived units](../derived-units/).

## Validation

At startup, `unit_si.go` registers all definitions and runs `validateRegistry()` — every dimension must have exactly one base unit with `Scale = 1` and `Offset = 0`.

Next: [Quantities →](../quantities/)
