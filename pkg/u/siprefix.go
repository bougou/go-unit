package u

import (
	"fmt"

	"github.com/bougou/go-unit/pkg/prefix"
)

// siUnitPrefix is one SI decimal prefix used by Unit.Prefix / DerivedUnit.Prefix.
type siUnitPrefix struct {
	Factor prefix.SIPrefix
	Symbol string // unit display form: "k", "M", "μ", …
	Name   string // "kilo", "mega", "micro", …
}

// siUnitPrefixes lists SI prefixes accepted by Prefix, including centi/deci/deka/hecto
// used by base units but omitted from PrefixFormat's scalesSI.
var siUnitPrefixes = []siUnitPrefix{
	{prefix.Quecto, "q", "quecto"},
	{prefix.Ronto, "r", "ronto"},
	{prefix.Yocto, "y", "yocto"},
	{prefix.Zepto, "z", "zepto"},
	{prefix.Atto, "a", "atto"},
	{prefix.Femto, "f", "femto"},
	{prefix.Pico, "p", "pico"},
	{prefix.Nano, "n", "nano"},
	{prefix.Micro, "μ", "micro"},
	{prefix.Milli, "m", "milli"},
	{prefix.Centi, "c", "centi"},
	{prefix.Deci, "d", "deci"},
	{prefix.Deka, "da", "deka"},
	{prefix.Hecto, "h", "hecto"},
	{prefix.Kilo, "k", "kilo"},
	{prefix.Mega, "M", "mega"},
	{prefix.Giga, "G", "giga"},
	{prefix.Tera, "T", "tera"},
	{prefix.Peta, "P", "peta"},
	{prefix.Exa, "E", "exa"},
	{prefix.Zetta, "Z", "zetta"},
	{prefix.Yotta, "Y", "yotta"},
	{prefix.Ronna, "R", "ronna"},
	{prefix.Quetta, "Q", "quetta"},
}

func siPrefixByFactor(factor prefix.SIPrefix) (siUnitPrefix, bool) {
	if factor == prefix.One || factor == 1 {
		return siUnitPrefix{Factor: 1, Symbol: "", Name: ""}, true
	}
	for _, p := range siUnitPrefixes {
		if p.Factor == factor {
			return p, true
		}
	}
	return siUnitPrefix{}, false
}

func siPrefixDisplaySymbol(factor prefix.SIPrefix) (string, bool) {
	p, ok := siPrefixByFactor(factor)
	if !ok {
		return "", false
	}
	return p.Symbol, true
}

func mustSIPrefixByFactor(factor prefix.SIPrefix) siUnitPrefix {
	p, ok := siPrefixByFactor(factor)
	if !ok {
		panic(fmt.Sprintf("unsupported SI prefix factor %g", factor))
	}
	return p
}
