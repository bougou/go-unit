# go-unit

[![Go Reference](https://pkg.go.dev/badge/github.com/bougou/go-unit/pkg/u.svg)](https://pkg.go.dev/github.com/bougou/go-unit/pkg/u)
[![Documentation](https://img.shields.io/badge/docs-GitHub%20Pages-blue)](https://bougou.github.io/go-unit/)

**go-unit** 是一个 Go 物理量与单位库，支持 SI（国际单位制）量纲、单位换算、导出单位，以及编译期量纲安全。

[English README](README.md) · [在线文档](https://bougou.github.io/go-unit/)

## 特性

- **七大 SI 基本量纲** — 长度、质量、时间、电流、温度、物质的量、发光强度
- **导出量纲与导出单位** — 组合单位（如 km/h、N、Pa），可选 SI 专用名称
- **类型化物理量** — `LengthQuantity`、`TimeQuantity` 等，编译期量纲检查
- **仿射换算** — 比例单位（km → m）与零点偏移单位（°C → K）
- **符号格式化** — 通过 `FormatOption` 配置组合符号（`km·h⁻¹`、`km/h`、`N`）
- **数值词头** — SI（1000 进位）与 IEC（1024 进位）解析/格式化，与物理单位分离
- **字符串解析** — `QuantityParse`、类型化 `XxxQuantityParse`、`DerivedQuantityParse`
- **数字格式化** — 量的 `Format`，以及 `DelimitInt` / `DelimitFloat` / `CommaFloat`

## 安装

```bash
go get github.com/bougou/go-unit/pkg/u
```

需要 Go 1.20 及以上。

## 快速上手

```go
package main

import (
	"fmt"

	u "github.com/bougou/go-unit/pkg/u"
)

func main() {
	// 类型化物理量，编译期量纲安全
	d := u.Length(10, u.Meter.Prefix(u.Kilo))
	t := u.Time(2, u.Hour)
	speed := d.Div(t) // DerivedQuantity: 5 km/h
	fmt.Println(speed.String())

	// 同量纲单位换算
	fmt.Println(u.Length(1000, u.Meter).By(u.Meter.Prefix(u.Kilo))) // 1 km
	fmt.Println(u.Length(1000, u.Meter).Prefix(u.Kilo))           // 1 km（等价写法）
	fmt.Println(u.Ohm.Of(2e6).Prefix(u.Mega).Format()) // 2 MΩ

	// SI 专用名称 vs 组合符号
	fmt.Println(u.Newton.Symbol()) // N
	fmt.Println(u.Newton.Symbol(u.WithCompoundSymbol(true))) // kg·m·s^-2

	// 从文本解析物理量
	d, _ := u.LengthQuantityParse("10 km")
	force, _ := u.DerivedQuantityParse("5 N")
	fmt.Println(d, force)

	// 数值词头（非物理单位）
	v, _ := u.PrefixParse("1.5G", u.SI)
	fmt.Println(v) // 1.5e9
	fmt.Println(u.PrefixFormat(1048576, u.IEC)) // 1 Mi
}
```

## 文档

完整文档通过 GitHub Pages 发布：

| 语言 | 链接 |
|------|------|
| English | [bougou.github.io/go-unit/](https://bougou.github.io/go-unit/) |
| 中文 | [bougou.github.io/go-unit/zh/](https://bougou.github.io/go-unit/zh/) |

建议阅读顺序：

1. [量纲](https://bougou.github.io/go-unit/zh/concepts/dimensions/) — 基本量纲与导出量纲
2. [单位与换算](https://bougou.github.io/go-unit/zh/concepts/units-and-conversion/) — 注册表、Scale、Offset
3. [物理量](https://bougou.github.io/go-unit/zh/concepts/quantities/) — 带单位的数值
4. [导出单位](https://bougou.github.io/go-unit/zh/concepts/derived-units/) — 组合单位与 Intern
5. [指南](https://bougou.github.io/go-unit/zh/guides/typed-quantities/) — 类型化物理量、运算、词头

API 参考：[pkg.go.dev/github.com/bougou/go-unit/pkg/u](https://pkg.go.dev/github.com/bougou/go-unit/pkg/u)

## 本地预览文档

安装 [Node.js](https://nodejs.org/) 20+ 与 [pnpm](https://pnpm.io/installation)，然后：

```bash
make docs-serve
```

打开：

- 英文：[http://127.0.0.1:4321/go-unit/](http://127.0.0.1:4321/go-unit/)
- 中文：[http://127.0.0.1:4321/go-unit/zh/](http://127.0.0.1:4321/go-unit/zh/)

## 许可证

Apache License 2.0 — 见 [LICENSE](LICENSE)。
