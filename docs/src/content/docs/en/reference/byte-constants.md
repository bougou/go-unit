---
title: "Byte constants"
description: "Byte constants"
sidebar:
  order: 3
---

`byte_consts.go` exports named scale factors for byte multiples. These are **numeric constants**, not dimensional quantities.

## Decimal (SI-style, base 1000)

| Constant | Factor | Notes    |
| -------- | ------ | -------- |
| `B`      | 1      | byte     |
| `KB`     | 10³    | kilobyte |
| `MB`     | 10⁶    | megabyte |
| `GB`     | 10⁹    | gigabyte |
| `TB`     | 10¹²   | terabyte |
| `PB`     | 10¹⁵   | petabyte |
| `EB`     | 10¹⁸   | exabyte  |

## Binary (IEC-style, base 1024)

| Constant | Factor | Notes    |
| -------- | ------ | -------- |
| `KiB`    | 2¹⁰    | kibibyte |
| `MiB`    | 2²⁰    | mebibyte |
| `GiB`    | 2³⁰    | gibibyte |
| `TiB`    | 2⁴⁰    | tebibyte |
| `PiB`    | 2⁵⁰    | pebibyte |
| `EiB`    | 2⁶⁰    | exbibyte |

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

Next: [Package reference →](pkg-reference/)
