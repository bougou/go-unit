# go-unit

[![Go Reference](https://pkg.go.dev/badge/github.com/bougou/go-unit/pkg/u.svg)](https://pkg.go.dev/github.com/bougou/go-unit/pkg/u)
[![Documentation](https://img.shields.io/badge/docs-GitHub%20Pages-blue)](https://bougou.github.io/go-unit/)

**go-unit** is a Go library for physical quantities with SI-aware units, dimensions, conversions, and compile-time dimension safety.

[中文 README](README.zh-CN.md) · [Documentation](https://bougou.github.io/go-unit/)

## Features

- **Seven SI base dimensions** — length, mass, time, current, temperature, amount, and luminous intensity
- **Derived dimensions and units** — build compound units (e.g. km/h, N, Pa) with optional SI special names
- **Typed quantities** — `LengthQuantity`, `TimeQuantity`, … for compile-time dimension checks
- **Affine conversion** — ratio units (km → m) and offset units (°C → K)
- **Symbol formatting** — configurable compound symbols (`km·h⁻¹`, `km/h`, `N`) via `FormatOption`
- **Numeric prefixes** — SI (1000) and IEC (1024) parse/format helpers in `pkg/prefix`, separate from physical units
- **String parsing** — `QuantityParse`, typed `XxxQuantityParse`, and `DerivedQuantityParse`
- **Number formatting** — quantity `Format`, plus `DelimitInt` / `DelimitFloat` / `CommaFloat`

## Installation

```bash
go get github.com/bougou/go-unit/pkg/u
```

Requires Go 1.20 or later. Numeric prefix APIs live in the same module under `pkg/prefix` (imported as needed; one `go get` of the module is enough).

## Quick start

```go
package main

import (
	"fmt"

	"github.com/bougou/go-unit/pkg/prefix"
	u "github.com/bougou/go-unit/pkg/u"
)

func main() {
	// Typed quantities with compile-time dimension safety
	d := u.Length(10, u.Meter.Prefix(prefix.Kilo))
	t := u.Time(2, u.Hour)
	speed := d.Div(t) // DerivedQuantity: 5 km/h
	fmt.Println(speed.String())

	// Unit conversion within the same dimension
	fmt.Println(u.Length(1000, u.Meter).By(u.Meter.Prefix(prefix.Kilo))) // 1 km
	fmt.Println(u.Length(1000, u.Meter).Prefix(prefix.Kilo))             // 1 km (equivalent)
	fmt.Println(u.Ohm.Of(2e6).Prefix(prefix.Mega).Format())              // 2 MΩ

	// SI special name vs compound symbol
	fmt.Println(u.Newton.Symbol()) // N
	fmt.Println(u.Newton.Symbol(u.WithCompoundSymbol(true))) // kg·m·s^-2

	// Parse quantities from text
	d, _ := u.LengthQuantityParse("10 km")
	force, _ := u.DerivedQuantityParse("5 N")
	fmt.Println(d, force)

	// Numeric prefix helpers (not physical units) — package prefix
	v, _ := prefix.PrefixParse("1.5G", prefix.SI)
	fmt.Println(v) // 1.5e9
	fmt.Println(prefix.PrefixFormat(1048576, prefix.IEC)) // 1 Mi
}
```

## Documentation

Full documentation is published on GitHub Pages:

| Language | Link |
|----------|------|
| English | [bougou.github.io/go-unit/](https://bougou.github.io/go-unit/) |
| 中文 | [bougou.github.io/go-unit/zh/](https://bougou.github.io/go-unit/zh/) |

Conceptual reading order:

1. [Dimensions](https://bougou.github.io/go-unit/en/concepts/dimensions/) — base and derived dimensions
2. [Units & conversion](https://bougou.github.io/go-unit/en/concepts/units-and-conversion/) — registry, scale, offset
3. [Quantities](https://bougou.github.io/go-unit/en/concepts/quantities/) — values with units
4. [Derived units](https://bougou.github.io/go-unit/en/concepts/derived-units/) — compound units and Intern
5. [Guides](https://bougou.github.io/go-unit/en/guides/typed-quantities/) — typed quantities, arithmetic, prefixes

API reference: [pkg.go.dev/github.com/bougou/go-unit/pkg/u](https://pkg.go.dev/github.com/bougou/go-unit/pkg/u) · [pkg/prefix](https://pkg.go.dev/github.com/bougou/go-unit/pkg/prefix)

## Local docs preview

Install [Node.js](https://nodejs.org/) 20+ and [pnpm](https://pnpm.io/installation), then:

```bash
make docs-serve
```

Open:

- English: [http://127.0.0.1:4321/go-unit/](http://127.0.0.1:4321/go-unit/)
- 中文：[http://127.0.0.1:4321/go-unit/zh/](http://127.0.0.1:4321/go-unit/zh/)

## License

Apache License 2.0 — see [LICENSE](LICENSE).
