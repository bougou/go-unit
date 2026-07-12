---
title: "SI base units"
description: "SI base units"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 1
toc: true
---
Constants are defined in `pkg/u/unit_si.go`. Each belongs to one SI base dimension and is registered at init with conversion to the dimension's SI base unit.

## Length (`LengthUnit`) — base: meter

| Category | Constants |
|----------|-----------|
| SI prefixes | `Picometer` … `Petameter` |
| Scientific | `Angstrom`, `AstronomicalUnit`, `LightYear`, `Parsec` |
| Imperial/US | `Inch`, `Foot`, `Yard`, `Mile`, `Mil`, `Hand`, `Fathom`, `Cable`, `NauticalMile`, `Chain`, `Furlong`, `Rod`, `League` |
| Typography | `Point`, `Pica` |
| Chinese traditional | `Chi`, `Cun`, `Zhang`, `Li` |

## Mass (`MassUnit`) — base: kilogram

| Category | Constants |
|----------|-----------|
| SI prefixes | `Nanogram` … `Petagram` |
| Metric | `Tonne` |
| Imperial/US | `Grain`, `Dram`, `Ounce`, `Pound`, `Stone`, `ShortTon`, `LongTon`, `TroyOunce`, `TroyPound` |
| Scientific | `Carat`, `Dalton` |
| Chinese traditional | `Jin`, `Liang` |

## Time (`TimeUnit`) — base: second

| Category | Constants |
|----------|-----------|
| SI prefixes | `Femtosecond` … `Gigasecond` |
| Common | `Minute`, `Hour`, `Day`, `Week`, `Fortnight`, `Month`, `Year` |

## Current (`CurrentUnit`) — base: ampere

`Femtoampere` … `Gigaampere`, plus `Biot` (CGS, 1 Bi = 10 A).

## Temperature (`TemperatureUnit`) — base: kelvin

| Unit | Type |
|------|------|
| `Kelvin`, `Microkelvin`, `Millikelvin` | Ratio (Offset = 0) |
| `Celsius`, `Fahrenheit`, `Rankine` | Affine (Offset ≠ 0) |

## Amount (`AmountUnit`) — base: mole

`Picomole` … `Megamole`, plus `PoundMole`.

## Luminous (`LuminousUnit`) — base: candela

`Millicandela`, `Candela`, `Kilocandela`.

## Symbols

Use `.Symbol()` on `Unit(...)` for display symbols:

```go
u.Unit(u.Kilometer).Symbol() // "km"
u.Unit(u.Celsius).Symbol()   // "°C"
```

Full conversion metadata (`Scale`, `Offset`) is in the `siUnitDefs` slice inside `unit_si.go`.

Next: [SI derived units →](../si-derived-units/)
