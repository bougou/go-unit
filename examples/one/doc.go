// Package one contains runnable Go Examples that demonstrate go-unit.
//
// Design goals for examples in this package:
//
//  1. Name by problem, not by order — Example_capacitorCharge, not Example_v3.
//  2. One homework-sized question per Example — state Given/Find in the doc comment.
//  3. Print labeled steps (C = …, Q = C·U = …) so godoc Output reads like a solution.
//  4. Prefer Unit.Of / Prefix for construction; show NewDerivedQuantity only once.
//  5. End with By / Prefix so results use familiar symbols (μF, mA, Ω), not only SI base.
//
// Files:
//
//   - unit_test.go: Ohm's law, prefixes, power and energy
//   - circuit_lc_test.go: capacitor and inductor calculations from circuit coursework
package one
