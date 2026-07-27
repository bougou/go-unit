package one

import (
	"fmt"
	"math"

	"github.com/bougou/go-unit/pkg/u"
)

// Example_capacitorCharge solves Q = C·U.
//
// Given: C = 50 μF, U = 10 V. Find the charge stored on the capacitor.
func Example_capacitorCharge() {
	c := u.Farad.Prefix(u.Micro).Of(50)
	voltage := u.Volt.Of(10)
	q := c.Mul(voltage).Prefix(u.Micro)

	fmt.Println("C =", c.Format())
	fmt.Println("U =", voltage.Format())
	fmt.Println("Q = C·U =", q.Format())
	// Output:
	// C = 50 μF
	// U = 10 V
	// Q = C·U = 500 μC
}

// Example_capacitorEnergy solves W = ½ C U².
//
// Given: C = 100 μF, U = 50 V. Find the electrostatic energy.
func Example_capacitorEnergy() {
	c := u.Farad.Prefix(u.Micro).Of(100)
	voltage := u.Volt.Of(50)
	w := c.Mul(voltage).Mul(voltage).DivV(2)

	fmt.Println("C =", c.Format())
	fmt.Println("U =", voltage.Format())
	fmt.Println("W = ½CU² =", w.Format(u.WithPrecision(3)))
	// Output:
	// C = 100 μF
	// U = 50 V
	// W = ½CU² = 0.125 J
}

// Example_capacitorsSeriesParallel finds equivalent capacitance.
//
// Given: C1 = 3 μF, C2 = 6 μF.
// Series: Ceq = C1 C2 / (C1+C2). Parallel: Ceq = C1+C2.
func Example_capacitorsSeriesParallel() {
	c1 := u.Farad.Prefix(u.Micro).Of(3)
	c2 := u.Farad.Prefix(u.Micro).Of(6)

	series := c1.Mul(c2).Div(c1.Add(c2)).Prefix(u.Micro)
	parallel := c1.Add(c2)

	fmt.Println("series Ceq =", series.Format())
	fmt.Println("parallel Ceq =", parallel.Format())
	// Output:
	// series Ceq = 2 μF
	// parallel Ceq = 9 μF
}

// Example_rcTimeConstant solves τ = R·C.
//
// Given: R = 10 kΩ, C = 100 μF. Find the RC time constant.
func Example_rcTimeConstant() {
	r := u.Ohm.Prefix(u.Kilo).Of(10)
	c := u.Farad.Prefix(u.Micro).Of(100)
	tau := r.Mul(c)

	fmt.Println("R =", r.Format())
	fmt.Println("C =", c.Format())
	fmt.Println("τ = RC =", tau.Format())
	// Output:
	// R = 10 kΩ
	// C = 100 μF
	// τ = RC = 1 s
}

// Example_capacitorReactance solves XC = 1 / (2π f C).
//
// Given: C = 10 μF, f = 50 Hz. Find the capacitive reactance.
// Note: f and ω = 2πf share dimension T⁻¹ — multiply by 2π yourself (MulV).
func Example_capacitorReactance() {
	c := u.Farad.Prefix(u.Micro).Of(10)
	f := u.Hertz.Of(50)
	xc := u.Dimensionless(1).Div(f.Mul(c).MulV(2 * math.Pi)).By(u.Ohm)

	fmt.Println("C =", c.Format())
	fmt.Println("f =", f.Format())
	fmt.Println("XC = 1/(2πfC) =", xc.Format(u.WithPrecision(2)))
	// Output:
	// C = 10 μF
	// f = 50 Hz
	// XC = 1/(2πfC) = 318.31 Ω
}

// Example_capacitorCurrentFromVoltageRamp solves i = C · ΔU/Δt (constant ramp).
//
// Given: C = 47 μF, voltage rises 20 V in 10 ms. Find the capacitor current.
func Example_capacitorCurrentFromVoltageRamp() {
	c := u.Farad.Prefix(u.Micro).Of(47)
	dU := u.Volt.Of(20)
	dt := u.Second.Prefix(u.Milli).Of(10)
	i := c.Mul(dU.Div(dt)).Prefix(u.Milli)

	fmt.Println("C =", c.Format())
	fmt.Println("ΔU =", dU.Format())
	fmt.Println("Δt =", dt.Format())
	fmt.Println("i = C·ΔU/Δt =", i.Format())
	// Output:
	// C = 47 μF
	// ΔU = 20 V
	// Δt = 10 ms
	// i = C·ΔU/Δt = 94 mA
}

// Example_inductorReactance solves XL = 2π f L.
//
// Given: L = 0.1 H, f = 50 Hz. Find the inductive reactance.
func Example_inductorReactance() {
	l := u.Henry.Of(0.1)
	f := u.Hertz.Of(50)
	xl := f.Mul(l).MulV(2 * math.Pi).By(u.Ohm)

	fmt.Println("L =", l.Format())
	fmt.Println("f =", f.Format())
	fmt.Println("XL = 2πfL =", xl.Format(u.WithPrecision(2)))
	// Output:
	// L = 0.1 H
	// f = 50 Hz
	// XL = 2πfL = 31.42 Ω
}

// Example_inductorEnergy solves W = ½ L I².
//
// Given: L = 200 mH, I = 2 A. Find the magnetic energy stored in the inductor.
func Example_inductorEnergy() {
	l := u.Henry.Prefix(u.Milli).Of(200)
	i := u.Ampere.Of(2)
	w := l.Mul(i).Mul(i).DivV(2)

	fmt.Println("L =", l.Format())
	fmt.Println("I =", i.Format())
	fmt.Println("W = ½LI² =", w.Format())
	// Output:
	// L = 200 mH
	// I = 2 A
	// W = ½LI² = 0.4 J
}

// Example_lrTimeConstant solves τ = L/R.
//
// Given: L = 0.5 H, R = 100 Ω. Find the LR time constant.
func Example_lrTimeConstant() {
	l := u.Henry.Of(0.5)
	r := u.Ohm.Of(100)
	tau := l.Div(r).Prefix(u.Milli)

	fmt.Println("L =", l.Format())
	fmt.Println("R =", r.Format())
	fmt.Println("τ = L/R =", tau.Format())
	// Output:
	// L = 0.5 H
	// R = 100 Ω
	// τ = L/R = 5 ms
}

// Example_inductorVoltageFromCurrentRamp solves u = L · ΔI/Δt (constant ramp).
//
// Given: L = 0.2 H, current rises 5 A in 10 ms. Find the inductor voltage.
func Example_inductorVoltageFromCurrentRamp() {
	l := u.Henry.Of(0.2)
	dI := u.Ampere.Of(5)
	dt := u.Second.Prefix(u.Milli).Of(10)
	voltage := l.Mul(dI.Div(dt)).By(u.Volt)

	fmt.Println("L =", l.Format())
	fmt.Println("ΔI =", dI.Format())
	fmt.Println("Δt =", dt.Format())
	fmt.Println("u = L·ΔI/Δt =", voltage.Format())
	// Output:
	// L = 0.2 H
	// ΔI = 5 A
	// Δt = 10 ms
	// u = L·ΔI/Δt = 100 V
}

// Example_fluxLinkage solves Φ = L·I.
//
// Given: L = 50 mH, I = 2 A. Find the flux linkage (weber).
func Example_fluxLinkage() {
	l := u.Henry.Prefix(u.Milli).Of(50)
	i := u.Ampere.Of(2)
	phi := l.Mul(i).By(u.Weber)

	fmt.Println("L =", l.Format())
	fmt.Println("I =", i.Format())
	fmt.Println("Φ = L·I =", phi.Format())
	// Output:
	// L = 50 mH
	// I = 2 A
	// Φ = L·I = 0.1 Wb
}
