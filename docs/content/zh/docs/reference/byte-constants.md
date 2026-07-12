---
title: "字节常量"
description: "字节常量"
date: "2026-01-01T00:00:00+00:00"
draft: false
weight: 3
toc: true
---
`byte_consts.go` 导出字节倍数的命名常量，属于**纯数值**，不是带量纲的物理量。

## 十进制（SI 风格，1000）

`B`、`KB`、`MB` … 至 `XB`，因子为 10 的倍数。扩展名称（BB、NB 等）见源码注释。

## 二进制（IEC 风格，1024）

`KiB`、`MiB` … `PiB` 为 2 的幂；更大常量定义见源码，使用前请核对表达式。

## 示例

```go
size := 512 * u.MiB
u.PrefixFormat(float64(size), u.IEC) // "512 Mi"
```

解析 `"1.5 GiB"` 时去掉 `B`，对 `"1.5Gi"` 调用 `PrefixParse(..., u.IEC)`。

## 参考

- [Binary prefix (Wikipedia)](https://en.wikipedia.org/wiki/Binary_prefix)
- [docker/go-units](https://github.com/docker/go-units)

下一步：[包参考 →](../pkg-reference/)
