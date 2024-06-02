package unit

import "fmt"

func init() {
	if err := CheckAllValid(ALL); err != nil {
		panic(err)
	}
}

func CheckAllValid(all map[Unit]UnitDef) error {
	dimensionBaseUnit := make(map[Dimension]Unit)
	for u := range all {

		if reason, ok := u.Valid(); !ok {
			return fmt.Errorf("unit not valid, reason: %s", reason)
		}

		_, _, err := u.Base()
		if err != nil {
			return fmt.Errorf("unit (%s) Base method failed, err: %s", u, err)
		}

		if d, ok := u.IsBase(); ok {
			if old, exists := dimensionBaseUnit[d]; !exists {
				dimensionBaseUnit[d] = u
			} else {
				return fmt.Errorf("duplicate base unit for dimension (%s), units (%s) (%s)", d, old, u)
			}
		}
	}

	return nil
}

var (
	ALL = map[Unit]UnitDef{
		LengthMeter: {
			Canonical: "meter",
			Symbol:    "m",
			Dimension: DimensionLength,
		},
		TimeSecond: {
			Canonical: "second",
			Symbol:    "s",
			Alias:     []string{"sec"},
			Dimension: DimensionTime,
		},
		TimeMinute: {
			Canonical: "minute",
			Symbol:    "min",
			Dimension: DimensionTime,
			Relation: &Relation{
				Factor: 60,
				Unit:   TimeSecond,
			},
		},
		TimeHour: {
			Canonical: "hour",
			Symbol:    "h",
			Relation: &Relation{
				Factor: 60,
				Unit:   TimeSecond,
			},
		},
		MassGram: {
			Canonical: "gram",
			Symbol:    "g",
			Dimension: DimensionMass,
		},
		MassKilogram: {
			Canonical: "kilogram",
			Symbol:    "kg",
			Alias:     []string{"kilogramme"},
			Relation: &Relation{
				Unit:   MassGram,
				Factor: Kilo,
			},
		},
	}
)
