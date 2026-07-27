---
title: "SI base units"
description: "SI base units"
sidebar:
  order: 1
---
Constants are defined in `pkg/u/unit_si.go`. Each dimension exports **anchors** and non-SI units only. SI decimal multiples use `Anchor.Prefix(prefix)` (lazy registration).

## Length (`LengthUnit`) — base: meter

| Category | Constants / forms |
|----------|-------------------|
| SI anchor | `Meter`; prefixes via `Meter.Prefix(Kilo)` → km |
| Scientific | `Angstrom`, `AstronomicalUnit`, `LightYear`, `Parsec` |
| Imperial/US | `Inch`, `Foot`, `Yard`, `Mile`, `Mil`, `Hand`, `Fathom`, `Cable`, `NauticalMile`, `Chain`, `Furlong`, `Rod`, `League` |
| Typography | `Point`, `Pica` |
| Chinese traditional | `Chi`, `Cun`, `Zhang`, `Li` |

## Mass (`MassUnit`) — base: kilogram

| Category | Constants / forms |
|----------|-------------------|
| SI | `Kilogram` (base), `Gram` (prefix root) |
| Prefixes | Attach to gram: `Gram.Prefix(Milli)` → mg; `Gram.Prefix(Kilo)` → `Kilogram`; `Kilogram.Prefix(Milli)` → `Gram` |
| Other | `Tonne`, imperial/troy, carat/dalton, jin/liang, … |

## Time (`TimeUnit`) — base: second

`Second` + `Second.Prefix(…)`; also `Minute`, `Hour`, `Day`, `Week`, `Fortnight`, `Month`, `Year`.

## Current (`CurrentUnit`) — base: ampere

`Ampere` + `Ampere.Prefix(…)`, plus CGS `Biot` (1 Bi = 10 A).

## Temperature (`TemperatureUnit`) — base: kelvin

| Type | Units |
|------|-------|
| Ratio | `Kelvin` + `Kelvin.Prefix(…)` |
| Affine | `Celsius`, `Fahrenheit`, `Rankine` (reject `Prefix`) |

## Amount (`AmountUnit`) — base: mole

`Mole` + `Mole.Prefix(…)`, plus `PoundMole`.

## Luminous (`LuminousUnit`) — base: candela

`Candela` + `Candela.Prefix(…)`.

## Symbols

```go
u.Unit(u.Meter.Prefix(u.Kilo)).Symbol() // "km"
u.Unit(u.Gram.Prefix(u.Milli)).Symbol() // "mg"
u.Unit(u.Celsius).Symbol()         // "°C"
```

Full conversion metadata (`Scale`, `Offset`) is in the `siUnitDefs` slice inside `unit_si.go`.

Next: [SI derived units →](si-derived-units/)
