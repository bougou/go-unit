package u

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bougou/go-unit/pkg/prefix"
)

var (
	generatedUnitDefs []unitDef
	symbolIndexMu     sync.Mutex
	symbolIndexDirty  = true
	baseSymbolDirty   = true
)

// Prefix returns this length unit scaled by an SI decimal prefix factor.
//
// Example: Meter.Prefix(prefix.Kilo) // km
func (u LengthUnit) Prefix(factor prefix.SIPrefix) LengthUnit {
	return LengthUnit(prefixedBaseUnit(Unit(u), factor))
}

// Prefix returns this mass unit scaled by an SI decimal prefix factor.
//
// Prefixes attach to Gram (Gram.Prefix(prefix.Milli) → mg, Gram.Prefix(prefix.Kilo) → Kilogram).
// Kilogram.Prefix(prefix.Milli) → Gram is the only Kilogram.Prefix shortcut; other factors panic.
//
// Example: Gram.Prefix(prefix.Milli) // mg
func (u MassUnit) Prefix(factor prefix.SIPrefix) MassUnit {
	if factor != prefix.One && factor != 1 && u == Kilogram && factor != prefix.Milli {
		panic("mass SI prefixes attach to Gram; use Gram.Prefix(...), or Kilogram.Prefix(prefix.Milli) for Gram")
	}
	return MassUnit(prefixedBaseUnit(Unit(u), factor))
}

// Prefix returns this time unit scaled by an SI decimal prefix factor.
//
// Example: Second.Prefix(prefix.Milli) // ms
func (u TimeUnit) Prefix(factor prefix.SIPrefix) TimeUnit {
	return TimeUnit(prefixedBaseUnit(Unit(u), factor))
}

// Prefix returns this current unit scaled by an SI decimal prefix factor.
//
// Example: Ampere.Prefix(prefix.Micro) // μA
func (u CurrentUnit) Prefix(factor prefix.SIPrefix) CurrentUnit {
	return CurrentUnit(prefixedBaseUnit(Unit(u), factor))
}

// Prefix returns this temperature unit scaled by an SI decimal prefix factor.
// Affine units (°C, °F) are rejected.
//
// Example: Kelvin.Prefix(prefix.Milli) // mK
func (u TemperatureUnit) Prefix(factor prefix.SIPrefix) TemperatureUnit {
	return TemperatureUnit(prefixedBaseUnit(Unit(u), factor))
}

// Prefix returns this amount unit scaled by an SI decimal prefix factor.
//
// Example: Mole.Prefix(prefix.Milli) // mmol
func (u AmountUnit) Prefix(factor prefix.SIPrefix) AmountUnit {
	return AmountUnit(prefixedBaseUnit(Unit(u), factor))
}

// Prefix returns this luminous unit scaled by an SI decimal prefix factor.
//
// Example: Candela.Prefix(prefix.Milli) // mcd
func (u LuminousUnit) Prefix(factor prefix.SIPrefix) LuminousUnit {
	return LuminousUnit(prefixedBaseUnit(Unit(u), factor))
}

func prefixedBaseUnit(anchor Unit, factor prefix.SIPrefix) Unit {
	if factor == prefix.One || factor == 1 {
		return anchor
	}
	adef, ok := anchor.Def()
	if !ok {
		panic(fmt.Sprintf("unknown unit %q", anchor))
	}
	if adef.Offset != 0 {
		panic(fmt.Sprintf("Prefix cannot be applied to affine unit %q", anchor))
	}
	mustSIPrefixByFactor(factor)

	targetScale := adef.Scale * float64(factor)
	sym := prefixedBaseSymbol(adef, factor)

	if existing, ok := findUnitBySymbolAndDimension(sym, adef.Dimension); ok {
		return existing
	}
	if existing, ok := findSIFamilyUnitByScale(adef.Dimension, targetScale, sym); ok {
		return existing
	}

	p := mustSIPrefixByFactor(factor)
	name := p.Name + adef.Name
	if adef.Dimension == DimMass {
		name = massPrefixedName(targetScale)
	}
	key := Unit(strings.ReplaceAll(name, " ", "_"))
	if def, ok := key.Def(); ok {
		return def.Unit
	}

	generatedUnitDefs = append(generatedUnitDefs, unitDef{
		Dimension: adef.Dimension,
		Unit:      key,
		Symbol:    sym,
		Name:      name,
		Scale:     targetScale,
		Offset:    0,
	})
	registerUnit(&generatedUnitDefs[len(generatedUnitDefs)-1])
	return key
}

func massPrefixedName(scaleInKg float64) string {
	gramRelative := prefix.SIPrefix(scaleInKg / float64(prefix.Milli))
	if gramRelative == prefix.One || gramRelative == 1 {
		return "gram"
	}
	if gramRelative == prefix.Kilo {
		return "kilogram"
	}
	p := mustSIPrefixByFactor(gramRelative)
	return p.Name + "gram"
}

func prefixedBaseSymbol(adef unitDef, factor prefix.SIPrefix) string {
	targetScale := adef.Scale * float64(factor)
	if adef.Dimension == DimMass {
		gramRelative := prefix.SIPrefix(targetScale / float64(prefix.Milli))
		if gramRelative == prefix.One || gramRelative == 1 {
			return "g"
		}
		p, ok := siPrefixDisplaySymbol(gramRelative)
		if !ok {
			panic(fmt.Sprintf("no SI prefix symbol for mass scale %g kg", targetScale))
		}
		return p + "g"
	}

	base := adef.Dimension.Base()
	baseDef, ok := base.Def()
	if !ok {
		panic(fmt.Sprintf("missing base unit for dimension %s", adef.Dimension))
	}
	relative := prefix.SIPrefix(targetScale / baseDef.Scale)
	if relative == prefix.One || relative == 1 {
		return baseDef.Symbol
	}
	p, ok := siPrefixDisplaySymbol(relative)
	if !ok {
		panic(fmt.Sprintf("no SI prefix symbol for %s scale %g", adef.Dimension, targetScale))
	}
	return p + baseDef.Symbol
}

func findUnitBySymbolAndDimension(symbol string, dim Dimension) (Unit, bool) {
	for u, def := range unitRegistry {
		if def.Symbol == symbol && def.Dimension == dim && def.Offset == 0 {
			return u, true
		}
	}
	return "", false
}

func findSIFamilyUnitByScale(dim Dimension, scale float64, wantSymbol string) (Unit, bool) {
	for u, def := range unitRegistry {
		if def.Dimension != dim || def.Offset != 0 || def.Scale != scale {
			continue
		}
		if def.Symbol == wantSymbol {
			return u, true
		}
	}
	return "", false
}

// isSIPrefixStem reports whether u can take SI decimal prefixes via Prefix / parse.
// Stems are the dimension base unit, plus Gram as the mass naming root.
func isSIPrefixStem(u Unit) bool {
	def, ok := u.Def()
	if !ok || def.Offset != 0 {
		return false
	}
	if u == def.Dimension.Base() {
		return true
	}
	return u == Unit(Gram)
}

func markSymbolIndexesDirty() {
	symbolIndexMu.Lock()
	symbolIndexDirty = true
	baseSymbolDirty = true
	symbolIndexMu.Unlock()
}
