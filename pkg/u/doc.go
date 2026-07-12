// Package u provides physical quantities with SI-aware units, dimensions, and conversions.
//
// SI (International System of Units, 国际单位制) base dimensions and derived units are
// modeled with compile-time typed quantities where possible.
//
// # Core concepts
//
// Unit is the internal identifier for a measurement unit (e.g. "kilometer", 千米).
// UnitDef holds conversion metadata (Scale, Offset) and the display Symbol ("km").
//
// Dimension is one of the seven SI base dimensions (length 长度，mass 质量，time 时间，…).
// DerivedDimension (导出量纲) describes a compound dimension as exponents on those bases,
// e.g. speed is {L:1, T:-1}.
//
// DerivedUnit (导出单位) is a concrete compound unit built from base units, e.g. km·h⁻¹.
// It can also carry an SI special name (专用名称) such as "N" (牛顿) for kg·m·s⁻².
//
// # Typed quantities
//
// Use typed constructors and quantity types for compile-time dimension safety:
//
//	d := Length(10, Kilometer)
//	t := Time(2, Hour)
//	speed := d.Div(t) // DerivedQuantity: 5 km/h
//
// Quantity is the untyped alternative when the dimension is not fixed at compile time.
//
// # Parsing from text
//
// QuantityParse parses "value unit" strings into Quantity. Typed parsers
// (LengthQuantityParse, TimeQuantityParse, …) validate a single base dimension.
// DerivedQuantityParse accepts compound and SI special-name units (N, m/s, …).
// DerivedUnitParse rebuilds a DerivedUnit from base-unit symbols (km/h, …)
// when no pre-registered alias exists. Each Parse has a MustParse variant.
// Use errors.Is(err, ErrDimension) for dimension mismatches.
// PrefixParse is for plain numeric scaling, not physical units.
//
// # Derived units and Intern
//
// NewDerivedUnit builds a compound unit but does not register it globally.
// Intern (内化注册) stores a canonical copy in an internal registry (first registration wins)
// so the same unit always shares one pointer identity:
//
//	u, _ := NewDerivedUnit().Length(Meter, 1).Time(Second, -1).Intern()
//	v, _ := NewDerivedUnit().Length(Meter, 1).Time(Second, -1).Intern()
//	u == v // true
//
// Mul and Div on quantities call Intern indirectly; most callers never need Intern.
//
// # Symbol formatting
//
// Symbol options control how units are rendered. By default DerivedUnit.Symbol
// uses the compound base-unit form. Pass WithNamedSymbol(true) for special names
// such as "N" instead of "kg·m·s^-2".
//
// # Prefix helpers
//
// PrefixParse and PrefixFormat handle SI/IEC (国际电工委员会) numeric prefixes (K, M, Mi, …)
// separately from physical units. Strip unit suffixes before parsing.
package u
