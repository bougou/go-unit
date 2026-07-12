---
title: "Package reference"
description: "Package reference"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 4
toc: true
---
## Online API docs

Full generated documentation:

**[pkg.go.dev/github.com/bougou/go-unit/pkg/u](https://pkg.go.dev/github.com/bougou/go-unit/pkg/u)**

## Source map

All library source lives under `pkg/u/`:

| File | Responsibility |
|------|----------------|
| `pkg/u/dimension.go` | `Dimension` type, seven base dimensions |
| `pkg/u/deriveddimension.go` | `DerivedDimension`, dimension algebra, dimension symbols |
| `pkg/u/unit.go` | `Unit`, `unitDef`, registry, base conversion |
| `pkg/u/unit_si.go` | SI and common base-unit constants |
| `pkg/u/unit_si_derived.go` | SI derived dimensions and unit globals |
| `pkg/u/unit_typed.go` | Typed unit/quantity wrappers |
| `pkg/u/quantity.go` | `Quantity`, `derivedQuantity` interface |
| `pkg/u/quantityparse.go` | `QuantityParse`, symbol index, `ErrDimension` |
| `pkg/u/derivedsymbolparse.go` | `DerivedUnitParse`, derived symbol parsing |
| `pkg/u/quantityparse_typed.go` | Typed `XxxQuantityParse` / `XxxQuantityMustParse` |
| `pkg/u/quantity_arith.go` | Add/Sub/Mul/Div for `Quantity` |
| `pkg/u/derivedunit.go` | `DerivedUnit`, builders, Intern, symbol options |
| `pkg/u/derivedquantity.go` | `DerivedQuantity` and conversions |
| `pkg/u/none_unit.go` | `NoneUnit`, `Dimensionless()` |
| `pkg/u/prefix.go` | Prefix modes and scale constants |
| `pkg/u/prefixparse.go` | `PrefixParse` |
| `pkg/u/prefixformat.go` | `PrefixFormat`, `PrefixFormat2` |
| `pkg/u/roundmethod.go` | Rounding modes for prefix format |
| `pkg/u/comma.go` | Thousands separators |
| `pkg/u/byte_consts.go` | Byte scale constants |

## Key types (quick index)

| Type | Description |
|------|-------------|
| `Dimension` | SI base dimension enum |
| `DerivedDimension` | Compound dimension exponents |
| `Unit` | Unit identifier string |
| `Quantity` | Untyped base-dimension value |
| `LengthQuantity`, … | Typed quantities |
| `DerivedUnit` | Compound unit |
| `DerivedQuantity` | Compound-dimension value |
| `PrefixMode` | SI/IEC parse/format mode |
| `SymbolOption` | Unit/dimension symbol formatting |

## Parse errors

| Error | Meaning |
|-------|---------|
| `ErrSyntax` | Invalid input format, unknown unit symbol, or bad number |
| `ErrDimension` | Parsed unit does not match the expected dimension or quantity kind |

## Quantity parsing

Parse functions accept text of the form `"value unit"`. The numeric value must not contain spaces; optional spaces may appear between the value and the unit.

| API | Returns | Accepts |
|-----|---------|---------|
| `QuantityParse` | `Quantity` | Registered symbols; falls back to derived parse |
| `DerivedUnitParse` | `*DerivedUnit` | Derived base-unit symbols (e.g. `km/h`, `kg·m·s^-2`) |
| `LengthQuantityParse`, … | Typed quantity | Single base dimension only (e.g. `m`, `kg`, `°C`) |
| `DerivedQuantityParse` | `DerivedQuantity` | Derived units only (e.g. `N`, `m/s`, `km/h`) |

Each `Parse` has a matching `MustParse` that panics on error (`QuantityMustParse`, `LengthQuantityMustParse`, `DerivedQuantityMustParse`, …).

```go
d, err := u.LengthQuantityParse("10 km")       // LengthQuantity
force, err := u.DerivedQuantityParse("5 N")    // DerivedQuantity
speed, err := u.DerivedQuantityParse("60 km/h") // derived fallback
q := u.QuantityMustParse("1.5 m")              // Quantity, panic on error
```

See [Quantities — parsing from text](../concepts/quantities/#parsing-from-text) for rules, symbol coverage, and limitations.

## Versioning

Module path: `github.com/bougou/go-unit/pkg/u`

Go version: **1.20+** (see `go.mod`).

## Contributing

Documentation source lives under `docs/content/`. To preview locally:

```bash
make docs-serve
```

Pull requests welcome for unit definitions, docs corrections, and tests.
