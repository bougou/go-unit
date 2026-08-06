package u

import (
	"math"
	"testing"

	"github.com/bougou/go-unit/pkg/prefix"
)

func TestDerivedUnitP(t *testing.T) {
	mega := Ohm.Prefix(prefix.Mega)
	if mega.PrefixScale() != float64(prefix.Mega) {
		t.Fatalf("PrefixScale() = %g, want %g", mega.PrefixScale(), float64(prefix.Mega))
	}
	if mega.Symbol() != "MΩ" {
		t.Fatalf("named symbol = %q, want MΩ", mega.Symbol())
	}
	if mega.FactorToBase() != float64(prefix.Mega) {
		t.Fatalf("FactorToBase() = %g, want %g", mega.FactorToBase(), float64(prefix.Mega))
	}

	micro := Ohm.Prefix(prefix.Micro)
	if micro.Symbol() != "μΩ" {
		t.Fatalf("named symbol = %q, want μΩ", micro.Symbol())
	}

	back := mega.Prefix(prefix.Micro)
	if back != Ohm {
		t.Fatalf("prefix.Mega then prefix.Micro should yield Ohm, got %v (%s)", back, back.Symbol())
	}

	q := Ohm.Of(1e6).By(mega)
	if q.Value != 1 {
		t.Fatalf("1e6 Ω.By(MΩ) value = %g, want 1", q.Value)
	}
	if got := q.Format(); got != "1 MΩ" {
		t.Fatalf("Format = %q, want %q", got, "1 MΩ")
	}
}

func TestParsePrefixedSpecialName(t *testing.T) {
	q, err := DerivedQuantityParse("1.5 MΩ")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if q.Unit.PrefixScale() != float64(prefix.Mega) {
		t.Fatalf("PrefixScale = %g, want %g", q.Unit.PrefixScale(), float64(prefix.Mega))
	}
	if q.Value != 1.5 {
		t.Fatalf("value = %g, want 1.5", q.Value)
	}

	q2, err := QuantityParse("2 kN")
	if err != nil {
		t.Fatalf("parse kN: %v", err)
	}
	du, ok := q2.Unit.DerivedUnit()
	if !ok || du.PrefixScale() != float64(prefix.Kilo) {
		t.Fatalf("kN PrefixScale = %v ok=%v, want prefix.Kilo", du, ok)
	}
}

func TestWattHourAndKilowattHour(t *testing.T) {
	if WattHour.FactorToBase() != 3600 {
		t.Fatalf("WattHour FactorToBase = %g, want 3600", WattHour.FactorToBase())
	}
	kWh := WattHour.Prefix(prefix.Kilo)
	if kWh.FactorToBase() != 3.6e6 {
		t.Fatalf("kW·h FactorToBase = %g, want 3.6e6", kWh.FactorToBase())
	}
	if kWh.Symbol() != "kW·h" {
		t.Fatalf("kW·h symbol = %q, want kW·h", kWh.Symbol())
	}
	if kWh.Symbol(WithCompoundSymbol(true)) != "k(kg·m^2·s^-2)" {
		t.Fatalf("kW·h compound = %q, want k(kg·m^2·s^-2)", kWh.Symbol(WithCompoundSymbol(true)))
	}

	// 20 W × 3600 s → 72000 J → 20 W·h → 0.02 kW·h
	energy := Watt.Of(20).Mul(Second.Of(3600))
	wh := energy.By(WattHour)
	if wh.Value != 20 || wh.Unit != WattHour {
		t.Fatalf("By(WattHour) = %v %q, want 20 W·h", wh.Value, wh.Unit.Symbol())
	}
	got := wh.Prefix(prefix.Kilo)
	if got.Value != 0.02 || got.Unit.Symbol() != "kW·h" {
		t.Fatalf("Prefix(prefix.Kilo) = %v %q, want 0.02 kW·h", got.Value, got.Unit.Symbol())
	}
	if got.Format(WithCompoundSymbol(true)) != "0.02 k(kg·m^2·s^-2)" {
		t.Fatalf("Format(compound) = %q, want 0.02 k(kg·m^2·s^-2)", got.Format(WithCompoundSymbol(true)))
	}

	back := got.By(Joule)
	if back.Value != 72000 || back.Unit != Joule {
		t.Fatalf("By(Joule) = %v %q, want 72000 J", back.Value, back.Unit.Symbol())
	}
}

func TestHorsePowerVariants(t *testing.T) {
	const (
		mechHP = 745.69987158227022
		metric = 735.49875
		elecHP = 746.0
	)
	cases := []struct {
		unit   *DerivedUnit
		symbol string
		watts  float64
	}{
		{HorsePower, "hp", mechHP},
		{MetricHorsepower, "PS", metric},
		{ElectricalHorsePower, "hp(E)", elecHP},
	}
	for _, tc := range cases {
		if tc.unit.FactorToBase() != tc.watts {
			t.Fatalf("%s FactorToBase = %g, want %g", tc.symbol, tc.unit.FactorToBase(), tc.watts)
		}
		if tc.unit.Symbol() != tc.symbol {
			t.Fatalf("symbol = %q, want %q", tc.unit.Symbol(), tc.symbol)
		}
		watts := tc.unit.Of(2).By(Watt)
		if math.Abs(watts.Value-2*tc.watts) > 1e-12 || watts.Unit != Watt {
			t.Fatalf("2 %s → By(Watt) = %v %q, want %g W", tc.symbol, watts.Value, watts.Unit.Symbol(), 2*tc.watts)
		}
		back := watts.By(tc.unit)
		if math.Abs(back.Value-2) > 1e-12 || back.Unit != tc.unit {
			t.Fatalf("By(%s) = %v %q, want 2 %s", tc.symbol, back.Value, back.Unit.Symbol(), tc.symbol)
		}
	}
}
