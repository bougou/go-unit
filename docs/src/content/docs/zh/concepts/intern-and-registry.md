---
title: "内化与注册"
description: "内化与注册"
sidebar:
  order: 5
---
**Intern（内化）** 将 `*DerivedUnit` 注册到全局表，使相同组合共享同一指针身份与稳定查找键。

## 为何需要 Intern？

不 Intern 时，每次 `NewDerivedUnit()` 都会新建对象：

```go
a := u.NewDerivedUnit().Length(u.Meter, 1).Time(u.Second, -1)
b := u.NewDerivedUnit().Length(u.Meter, 1).Time(u.Second, -1)
a == b // false
```

`Intern()` 之后：

```go
a, _ := u.NewDerivedUnit().Length(u.Meter, 1).Time(u.Second, -1).Intern()
b, _ := u.NewDerivedUnit().Length(u.Meter, 1).Time(u.Second, -1).Intern()
a == b // true
```

物理量运算内部会间接 Intern——多数业务代码无需显式调用。

## 注册键

| 种类 | 键格式 | 示例 |
|------|--------|------|
| 未命名组合 | 基本单位组合 | `kilogram^1*meter^1*second^-2` |
| 命名组合 | 组合 + `#符号` | `...#N` |
| 命名无量纲 | 仅 `#符号` | `#rad` |
| 简单比 | 两项时 `a/b` | `kilometer/hour` |

Intern 后查找：

```go
key := u.Unit(u.Newton.Key())
du, ok := key.DerivedUnit()
```

## 命名 vs 未命名

某组合的**首个**注册者为**命名**单位（如 init 中的 `Newton`）时，裸组合键会别名到该实例。后续相同组合的未命名 `Intern()` 返回同一规范对象。

## NoneUnit

**未命名无量纲**单位，通常来自同量纲相除：

```go
ratio := u.Mass(10, u.Kilogram).Div(u.Mass(2, u.Kilogram))
ratio.Unit == u.NoneUnit // true
```

比较导出单位请用 `==`（Intern 后或使用 `Newton` 等全局变量）。

## MustIntern

包 init 与内部代码使用 `MustIntern()`，出错时 panic：

```go
u.NewDerivedUnit().Named("rad").MustIntern()
```

## 符号别名

Intern 同时注册 `derivedBySign`，便于按 `"N"`、`"Hz"` 等符号反查注册键。

下一步：[符号格式化 →](symbol-formatting/)
