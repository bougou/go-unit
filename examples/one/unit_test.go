package one

import (
	"fmt"

	"github.com/bougou/go-unit/pkg/prefix"
	"github.com/bougou/go-unit/pkg/u"
)

// Example_ohmLaw computes R = U/I and shows named, compound, and dimension symbols.
func Example_ohmLaw() {
	v := u.Volt.Of(220)
	i := u.Ampere.Of(10)
	r := v.Div(i)

	fmt.Println("U =", v.Format())
	fmt.Println("I =", i.Format())
	fmt.Println("R = U/I =", r.Format())
	fmt.Println("R (compound) =", r.Format(u.WithCompoundSymbol(true), u.WithExpSign(u.ExpSignSup)))
	fmt.Println("dim =", r.Unit.Dim().Symbol(u.WithExpSign(u.ExpSignSup)))
	// Output:
	// U = 220 V
	// I = 10 A
	// R = U/I = 22 Ω
	// R (compound) = 22 kg·m²·s⁻³·A⁻²
	// dim = M·L²·T⁻³·I⁻²
}

// Example_constructors shows the two equivalent ways to attach a value to a unit.
func Example_constructors() {
	v := u.NewDerivedQuantity(220, u.Volt) // (value, unit)
	i := u.Current(10, u.Ampere)           // typed base quantity
	r := v.Div(i)

	fmt.Println(r.Format())
	// Output: 22 Ω
}

// Example_resistanceWithPrefixes builds a megaohm from μA-scale current.
func Example_resistanceWithPrefixes() {
	v := u.Volt.Of(1)
	i := u.Ampere.Prefix(prefix.Micro).Of(1)
	r := v.Div(i).By(u.Ohm).Prefix(prefix.Mega)
	r2 := v.Div(i).Prefix(prefix.Mega)

	fmt.Println("R =", r.Format())
	fmt.Println("R2 =", r2.Format())
	// Output:
	// R = 1 MΩ
	// R2 = 1 MΩ
}

// Example_siPrefix converts a bare value into a prefixed unit (km, MΩ).
func Example_siPrefix() {
	d := u.Length(5, u.Meter.Prefix(prefix.Kilo))
	r := u.Ohm.Of(2e6).Prefix(prefix.Mega)

	fmt.Println(d)
	fmt.Println(r.Format())
	// Output:
	// 5 km
	// 2 MΩ
}

// Example_currentFromPowerAndResistance solves I = √(P/R).
func Example_currentFromPowerAndResistance() {
	p := u.Watt.Prefix(prefix.Milli).Of(20)
	r := u.Ohm.Of(5000).Prefix(prefix.Kilo)
	i := p.Div(r).Sqrt().Prefix(prefix.Milli)

	fmt.Println("P =", p.Format())
	fmt.Println("R =", r.Format())
	fmt.Println("I = √(P/R) =", i.Format())
	// Output:
	// P = 20 mW
	// R = 5 kΩ
	// I = √(P/R) = 2 mA
}

// Example_energyFromPowerAndTime computes W = P·t and converts J → kW·h.
//
// Prefer By(WattHour) for energy display — ByTime(Hour) would rewrite the time
// exponent of the energy dimension, not apply a “per hour” convenience unit.
func Example_energyFromPowerAndTime() {
	p := u.Watt.Of(20)
	t := u.Second.Of(3600)
	w := p.Mul(t)

	fmt.Println("P·t =", w.Format())

	w = w.By(u.WattHour).Prefix(prefix.Kilo)
	fmt.Println("→", w.Format())

	w = w.By(u.Joule)
	fmt.Println("→", w.Format())
	// Output:
	// P·t = 72000 J
	// → 0.02 kW·h
	// → 72000 J
}

// Example_yearlyEnergyKilowattHour estimates annual energy use of a 60 W load.
func Example_yearlyEnergyKilowattHour() {
	p := u.Watt.Of(60)
	t := u.Day.Of(365)
	e := p.Mul(t).By(u.WattHour).Prefix(prefix.Kilo)

	fmt.Println("E =", e.Format())
	// Output:
	// E = 525.6 kW·h
}

// Example_runtimeFromEnergyAndPower solves t = E/P for a battery-like energy store.
func Example_runtimeFromEnergyAndPower() {
	p := u.Watt.Of(340)
	e := u.WattHour.Prefix(prefix.Kilo).Of(4)
	t := e.Div(p).ByTime(u.Hour)

	fmt.Println("E =", e.Format())
	fmt.Println("P =", p.Format())
	fmt.Println("t = E/P =", t.Format(u.WithPrecision(2)))
	// Output:
	// E = 4 kW·h
	// P = 340 W
	// t = E/P = 11.76 h
}

// Example_electricalHorsepowerEnergy converts hp(E)·h into kW·h.
func Example_electricalHorsepowerEnergy() {
	p := u.ElectricalHorsePower.Of(5)
	t := u.Hour.Of(2)
	e := p.Mul(t).By(u.WattHour).Prefix(prefix.Kilo)

	fmt.Println("E =", e.Format())
	// Output:
	// E = 7.46 kW·h
}

// Example_motorCurrentFromHorsepower finds line current from electrical hp and efficiency.
//
// P_in = P_out / η, then I = P_in / U (DC / resistive model).
func Example_motorCurrentFromHorsepower() {
	pOut := u.ElectricalHorsePower.Of(2)
	pIn := pOut.DivV(0.75).By(u.Watt)
	v := u.Volt.Of(220)
	i := pIn.Div(v)

	fmt.Println("P_in =", pIn.Format(u.WithPrecision(2)))
	fmt.Println("I = P_in/U =", i.Format(u.WithPrecision(2)))
	// Output:
	// P_in = 1989.33 W
	// I = P_in/U = 9.04 A
}
