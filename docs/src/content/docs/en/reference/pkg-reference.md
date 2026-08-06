---
title: "Package reference"
description: "Package reference"
sidebar:
  order: 4
---
## Online API docs

Full generated documentation:

- **[pkg.go.dev/github.com/bougou/go-unit/pkg/u](https://pkg.go.dev/github.com/bougou/go-unit/pkg/u)** — units and quantities
- **[pkg.go.dev/github.com/bougou/go-unit/pkg/prefix](https://pkg.go.dev/github.com/bougou/go-unit/pkg/prefix)** — numeric prefixes and byte constants

## Source map

### `pkg/u/` — units and quantities

| File | Responsibility |
|------|----------------|
| `pkg/u/dimension.go` | `Dimension` type, seven base dimensions |
| `pkg/u/deriveddimension.go` | `DerivedDimension`, dimension algebra, dimension symbols |
| `pkg/u/unit.go` | `Unit`, `unitDef`, registry, base conversion |
| `pkg/u/unit_si.go` | SI and common base-unit constants |
| `pkg/u/unit_si_derived.go` | SI derived dimensions and unit globals |
| `pkg/u/siprefix.go` | Maps `prefix.SIPrefix` to unit display symbols for `Unit.Prefix` |
| `pkg/u/unit_typed.go` | Typed unit/quantity wrappers |
| `pkg/u/quantity.go` | `Quantity`, `derivedQuantity` interface |
| `pkg/u/quantityparse.go` | `QuantityParse`, symbol index, `ErrDimension` |
| `pkg/u/derivedsymbolparse.go` | `DerivedUnitParse` / `DerivedUnitMustParse`, derived symbol parsing |
| `pkg/u/quantityparse_typed.go` | Typed `XxxQuantityParse` / `XxxQuantityMustParse` |
| `pkg/u/quantity_arith.go` | Add/Sub/Mul/Div for `Quantity` |
| `pkg/u/derivedunit.go` | `DerivedUnit`, builders, Intern, `FormatOption` |
| `pkg/u/derivedquantity.go` | `DerivedQuantity` and conversions |
| `pkg/u/none_unit.go` | `NoneUnit`, `Dimensionless()` |
| `pkg/u/comma.go` | Thousands separators (`DelimitInt`, `DelimitFloat`) |
| `pkg/u/quantityformat.go` | Quantity `Format` helpers |

### `pkg/prefix/` — numeric prefixes

| File | Responsibility |
|------|----------------|
| `pkg/prefix/prefix.go` | `PrefixMode`, `SIPrefix`, IEC factors, `PrefixSymbol` |
| `pkg/prefix/prefixparse.go` | `PrefixParse` |
| `pkg/prefix/prefixformat.go` | `PrefixFormat`, `PrefixFormat2`, `WithPrefix*` |
| `pkg/prefix/roundmethod.go` | Rounding modes for prefix format |
| `pkg/prefix/byte_consts.go` | Byte scale constants (`KB`, `KiB`, …) |

## Key types (quick index)

| Type | Package | Description |
|------|---------|-------------|
| `Dimension` | `u` | SI base dimension enum |
| `DerivedDimension` | `u` | Compound dimension exponents |
| `Unit` | `u` | Unit identifier string |
| `Quantity` | `u` | Untyped base-dimension value |
| `LengthQuantity`, … | `u` | Typed quantities |
| `DerivedUnit` | `u` | Compound unit |
| `DerivedQuantity` | `u` | Compound-dimension value |
| `FormatOption` | `u` | Shared Format / Symbol display options |
| `PrefixMode` | `prefix` | SI/IEC parse/format mode |
| `SIPrefix` | `prefix` | SI decimal factor for `Unit.Prefix` |
| `PrefixSymbol` | `prefix` | Single-character numeric prefix (K, M, G, …) |

## Constructors

| API | Returns | Example |
|-----|---------|---------|
| `Length`, `Time`, `Mass`, … | Typed quantity | `Length(10, Meter.Prefix(prefix.Kilo))` |
| `NewDerivedQuantity` | `DerivedQuantity` | `NewDerivedQuantity(60, speedUnit)` |
| `DerivedUnit.Of` | `DerivedQuantity` | `Newton.Of(100)` |
| `NewDerivedUnit` | `*DerivedUnit` | `NewDerivedUnit().Length(Meter, 1)` |

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

Each `Parse` has a matching `MustParse` that panics on error (`QuantityMustParse`, `LengthQuantityMustParse`, `DerivedQuantityMustParse`, `DerivedUnitMustParse`, …).

```go
d, err := u.LengthQuantityParse("10 km")       // LengthQuantity
force, err := u.DerivedQuantityParse("5 N")    // DerivedQuantity
speed, err := u.DerivedQuantityParse("60 km/h") // derived fallback
unit := u.DerivedUnitMustParse("km/h")         // *DerivedUnit, panic on error
q := u.QuantityMustParse("1.5 m")              // Quantity, panic on error
```

See [Quantities — parsing from text](../concepts/quantities/#parsing-from-text) for rules, symbol coverage, and limitations.

## Versioning

Module path: `github.com/bougou/go-unit` (packages `pkg/u` and `pkg/prefix`)

Go version: **1.20+** (see `go.mod`).

## Contributing

Documentation source lives under `docs/src/content/docs/`. To preview locally:

```bash
make docs-serve
```

Pull requests welcome for unit definitions, docs corrections, and tests.
