---
title: "Intern & registry"
description: "Intern & registry"
sidebar:
  order: 5
---
**Interning** registers a canonical `*DerivedUnit` in a global registry so identical compositions share one pointer identity and stable lookup keys.

## Why intern?

Without interning, every `NewDerivedUnit()` call creates a new heap object even when the composition is identical:

```go
a := u.NewDerivedUnit().Length(u.Meter, 1).Time(u.Second, -1)
b := u.NewDerivedUnit().Length(u.Meter, 1).Time(u.Second, -1)
a == b // false — different pointers
```

After `Intern()`:

```go
a, _ := u.NewDerivedUnit().Length(u.Meter, 1).Time(u.Second, -1).Intern()
b, _ := u.NewDerivedUnit().Length(u.Meter, 1).Time(u.Second, -1).Intern()
a == b // true
```

Arithmetic on quantities calls interning internally — most application code never needs to call `Intern` explicitly.

## Registry keys

| Unit kind | Key format | Example |
|-----------|------------|---------|
| Unnamed compound | base composition | `kilogram^1*meter^1*second^-2` |
| Named compound | composition + `#symbol` | `...#N` |
| Named dimensionless | `#symbol` only | `#rad` |
| Simple ratio | `a/b` when two terms | `kilometer/hour` |

Lookup after interning:

```go
key := u.Unit(u.Newton.Key())
du, ok := key.DerivedUnit()
```

## Named vs unnamed registration

When the **first** registrant for a composition is **named** (e.g. `Newton` as `"N"`), the bare composition key aliases to that instance. Later unnamed `Intern()` calls for the same composition return the named canonical unit.

```go
// After init registers Newton first:
a, _ := u.NewDerivedUnit().Mass(u.Kilogram, 1).
    Length(u.Meter, 1).Time(u.Second, -2).Named("N").Intern()
b, _ := u.NewDerivedUnit().Mass(u.Kilogram, 1).
    Length(u.Meter, 1).Time(u.Second, -2).Intern()
a == b // true
```

## NoneUnit

`NoneUnit` is the generic **unnamed dimensionless** unit — typically from dividing two equal-dimension quantities:

```go
ratio := u.Mass(10, u.Kilogram).Div(u.Mass(2, u.Kilogram))
ratio.Unit == u.NoneUnit // true (after interning in Div)
```

Compare derived units with `==` only after interning (or when using package globals like `Newton`).

## MustIntern

Package init and internal code use `MustIntern()` — panics on error (e.g. nil receiver):

```go
u.NewDerivedUnit().Named("rad").MustIntern()
```

## Symbol alias map

Interning also registers symbol aliases (`derivedBySign`) so lookups by display symbol can resolve to registry keys — used internally when formatting and resolving signs like `"N"` or `"Hz"`.

Next: [Symbol formatting →](symbol-formatting/)
