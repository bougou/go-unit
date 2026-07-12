---
title: "Byte constants"
description: "Byte constants"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 3
toc: true
---
`byte_consts.go` exports named scale factors for byte multiples. These are **numeric constants**, not dimensional quantities.

## Decimal (SI-style, base 1000)

| Constant | Factor | Notes |
|----------|--------|-------|
| `B` | 1 | byte |
| `KB` | 10³ | kilobyte |
| `MB` | 10⁶ | megabyte |
| `GB` | 10⁹ | gigabyte |
| `TB` | 10¹² | terabyte |
| `PB` | 10¹⁵ | petabyte |
| `EB` | 10¹⁸ | exabyte |
| `ZB` | 10²¹ | zettabyte |
| `YB` | 10²⁴ | yottabyte |
| `BB` | 10²⁷ | ronna/bronto byte |
| `NB` | 10³⁰ | quetta byte |
| `DB` | 10³³ | suggested extended name |
| `CB` | 10³⁶ | suggested extended name |
| `XB` | 10³⁹ | suggested extended name |

## Binary (IEC-style, base 1024)

| Constant | Factor |
|----------|--------|
| `KiB` | 2¹⁰ |
| `MiB` | 2²⁰ |
| `GiB` | 2³⁰ |
| `TiB` | 2⁴⁰ |
| `PiB` | 2⁵⁰ |
| `EiB` | 2⁶⁰ × 1000* |
| `ZiB` | … |
| `YiB` | … |

\* Constants above `PiB` use the literal expressions defined in source — verify against your requirements when using extended binary names.

## Example

```go
size := 512 * u.MiB
fmt.Println(u.PrefixFormat(float64(size), u.IEC)) // "512 Mi"
```

For parsing user input like `"1.5 GiB"`, strip the `B` suffix and use `PrefixParse` with IEC mode:

```go
v, _ := u.PrefixParse("1.5Gi", u.IEC)
bytes := v // bytes as float64
```

## References

- [Binary prefix (Wikipedia)](https://en.wikipedia.org/wiki/Binary_prefix)
- Inspired in part by [docker/go-units](https://github.com/docker/go-units)

Next: [Package reference →](../pkg-reference/)
