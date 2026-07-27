package u

import "fmt"

// siUnitPrefix is one SI decimal prefix used by Unit.Prefix / DerivedUnit.Prefix.
type siUnitPrefix struct {
	Factor SIPrefix
	Symbol string // unit display form: "k", "M", "μ", …
	Name   string // "kilo", "mega", "micro", …
}

// siUnitPrefixes lists SI prefixes accepted by Prefix, including centi/deci/deka/hecto
// used by base units but omitted from PrefixFormat's scalesSI.
var siUnitPrefixes = []siUnitPrefix{
	{Quecto, "q", "quecto"},
	{Ronto, "r", "ronto"},
	{Yocto, "y", "yocto"},
	{Zepto, "z", "zepto"},
	{Atto, "a", "atto"},
	{Femto, "f", "femto"},
	{Pico, "p", "pico"},
	{Nano, "n", "nano"},
	{Micro, "μ", "micro"},
	{Milli, "m", "milli"},
	{Centi, "c", "centi"},
	{Deci, "d", "deci"},
	{Deka, "da", "deka"},
	{Hecto, "h", "hecto"},
	{Kilo, "k", "kilo"},
	{Mega, "M", "mega"},
	{Giga, "G", "giga"},
	{Tera, "T", "tera"},
	{Peta, "P", "peta"},
	{Exa, "E", "exa"},
	{Zetta, "Z", "zetta"},
	{Yotta, "Y", "yotta"},
	{Ronna, "R", "ronna"},
	{Quetta, "Q", "quetta"},
}

func siPrefixByFactor(factor SIPrefix) (siUnitPrefix, bool) {
	if factor == One || factor == 1 {
		return siUnitPrefix{Factor: 1, Symbol: "", Name: ""}, true
	}
	for _, p := range siUnitPrefixes {
		if p.Factor == factor {
			return p, true
		}
	}
	return siUnitPrefix{}, false
}

func siPrefixDisplaySymbol(factor SIPrefix) (string, bool) {
	p, ok := siPrefixByFactor(factor)
	if !ok {
		return "", false
	}
	return p.Symbol, true
}

func mustSIPrefixByFactor(factor SIPrefix) siUnitPrefix {
	p, ok := siPrefixByFactor(factor)
	if !ok {
		panic(fmt.Sprintf("unsupported SI prefix factor %g", factor))
	}
	return p
}
