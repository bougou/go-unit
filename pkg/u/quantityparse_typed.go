package u

// LengthQuantityParse parses a length quantity from text.
//
// Example: LengthQuantityParse("10 km") // 10 km
func LengthQuantityParse(s string) (LengthQuantity, error) {
	q, err := parseBaseDimensionQuantity(s, DimLength)
	if err != nil {
		return LengthQuantity{}, err
	}
	return LengthQuantity(q), nil
}

// LengthQuantityMustParse parses a length quantity and panics on error.
func LengthQuantityMustParse(s string) LengthQuantity {
	q, err := LengthQuantityParse(s)
	if err != nil {
		panic(err)
	}
	return q
}

// TimeQuantityParse parses a time quantity from text.
//
// Example: TimeQuantityParse("2 h") // 2 h
func TimeQuantityParse(s string) (TimeQuantity, error) {
	q, err := parseBaseDimensionQuantity(s, DimTime)
	if err != nil {
		return TimeQuantity{}, err
	}
	return TimeQuantity(q), nil
}

// TimeQuantityMustParse parses a time quantity and panics on error.
func TimeQuantityMustParse(s string) TimeQuantity {
	q, err := TimeQuantityParse(s)
	if err != nil {
		panic(err)
	}
	return q
}

// MassQuantityParse parses a mass quantity from text.
//
// Example: MassQuantityParse("100 kg") // 100 kg
func MassQuantityParse(s string) (MassQuantity, error) {
	q, err := parseBaseDimensionQuantity(s, DimMass)
	if err != nil {
		return MassQuantity{}, err
	}
	return MassQuantity(q), nil
}

// MassQuantityMustParse parses a mass quantity and panics on error.
func MassQuantityMustParse(s string) MassQuantity {
	q, err := MassQuantityParse(s)
	if err != nil {
		panic(err)
	}
	return q
}

// CurrentQuantityParse parses an electric-current quantity from text.
//
// Example: CurrentQuantityParse("500 mA") // 500 mA
func CurrentQuantityParse(s string) (CurrentQuantity, error) {
	q, err := parseBaseDimensionQuantity(s, DimCurrent)
	if err != nil {
		return CurrentQuantity{}, err
	}
	return CurrentQuantity(q), nil
}

// CurrentQuantityMustParse parses an electric-current quantity and panics on error.
func CurrentQuantityMustParse(s string) CurrentQuantity {
	q, err := CurrentQuantityParse(s)
	if err != nil {
		panic(err)
	}
	return q
}

// TemperatureQuantityParse parses a temperature quantity from text.
//
// Example: TemperatureQuantityParse("-5 °C") // -5 °C
func TemperatureQuantityParse(s string) (TemperatureQuantity, error) {
	q, err := parseBaseDimensionQuantity(s, DimTemperature)
	if err != nil {
		return TemperatureQuantity{}, err
	}
	return TemperatureQuantity(q), nil
}

// TemperatureQuantityMustParse parses a temperature quantity and panics on error.
func TemperatureQuantityMustParse(s string) TemperatureQuantity {
	q, err := TemperatureQuantityParse(s)
	if err != nil {
		panic(err)
	}
	return q
}

// AmountQuantityParse parses an amount-of-substance quantity from text.
//
// Example: AmountQuantityParse("2 mol") // 2 mol
func AmountQuantityParse(s string) (AmountQuantity, error) {
	q, err := parseBaseDimensionQuantity(s, DimAmount)
	if err != nil {
		return AmountQuantity{}, err
	}
	return AmountQuantity(q), nil
}

// AmountQuantityMustParse parses an amount-of-substance quantity and panics on error.
func AmountQuantityMustParse(s string) AmountQuantity {
	q, err := AmountQuantityParse(s)
	if err != nil {
		panic(err)
	}
	return q
}

// LuminousQuantityParse parses a luminous-intensity quantity from text.
//
// Example: LuminousQuantityParse("800 cd") // 800 cd
func LuminousQuantityParse(s string) (LuminousQuantity, error) {
	q, err := parseBaseDimensionQuantity(s, DimLuminous)
	if err != nil {
		return LuminousQuantity{}, err
	}
	return LuminousQuantity(q), nil
}

// LuminousQuantityMustParse parses a luminous-intensity quantity and panics on error.
func LuminousQuantityMustParse(s string) LuminousQuantity {
	q, err := LuminousQuantityParse(s)
	if err != nil {
		panic(err)
	}
	return q
}
