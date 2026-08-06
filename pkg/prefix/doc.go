// Package prefix handles SI and IEC numeric prefixes (for example K/M/G and Ki/Mi/Gi)
// and related byte scale constants.
//
// It is separate from physical units and quantities in package u (pkg/u).
//
// # Parsing and formatting
//
// PrefixParse and PrefixFormat (and PrefixFormat2) scale plain float64 values
// with PrefixMode (SI, IEC, Auto, …). They do not attach a physical unit —
// strip suffixes such as "B" or "bit" before parsing.
//
// # Scale factors for units
//
// SIPrefix constants (Kilo, Milli, Mega, …) are the factors accepted by
// u.Unit.Prefix / DerivedUnit.Prefix. IEC binary factors (Kibi, Mebi, …) are
// plain float64 values for numeric scaling.
//
// PrefixSymbol is the single-character (or IEC letter) type used when forcing
// a specific prefix in format options (WithPrefix).
//
// Byte multiples (KB, KiB, MiB, …) live in this package as named int64/float
// scale constants — see byte_consts.go.
package prefix
