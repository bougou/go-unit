---
title: "字节常量"
description: "字节常量"
sidebar:
  order: 3
---

`byte_consts.go` 导出字节倍数的命名常量，属于**纯数值**，不是带量纲的物理量。

## 十进制（SI 风格，1000）

| 常量 | 因子 | 说明               |
| ---- | ---- | ------------------ |
| `B`  | 1    | 字节               |
| `KB` | 10³  | kilobyte（千字节） |
| `MB` | 10⁶  | megabyte（兆字节） |
| `GB` | 10⁹  | gigabyte（吉字节） |
| `TB` | 10¹² | terabyte（太字节） |
| `PB` | 10¹⁵ | petabyte（拍字节） |
| `EB` | 10¹⁸ | exabyte（艾字节）  |

## 二进制（IEC 风格，1024）

| 常量  | 因子 | 说明                 |
| ----- | ---- | -------------------- |
| `KiB` | 2¹⁰  | kibibyte（千比字节） |
| `MiB` | 2²⁰  | mebibyte（兆比字节） |
| `GiB` | 2³⁰  | gibibyte（吉比字节） |
| `TiB` | 2⁴⁰  | tebibyte（太比字节） |
| `PiB` | 2⁵⁰  | pebibyte（拍比字节） |
| `EiB` | 2⁶⁰  | exbibyte（艾比字节） |

## 示例

```go
size := 512 * u.MiB
u.PrefixFormat(float64(size), u.IEC) // "512 Mi"
```

解析 `"1.5 GiB"` 时去掉 `B`，对 `"1.5Gi"` 调用 `PrefixParse(..., u.IEC)`。

## 参考

- [Binary prefix (Wikipedia)](https://en.wikipedia.org/wiki/Binary_prefix)
- [docker/go-units](https://github.com/docker/go-units)

下一步：[包参考 →](pkg-reference/)
