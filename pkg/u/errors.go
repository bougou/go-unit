package u

import "fmt"

// Sentinel errors for quantity operations. Prefer errors.Is against these values.
//
// Parse errors also use ErrSyntax and ErrDimension (see quantity/unit parsing).
// Numeric prefix parsing lives in package prefix.
var (
	// ErrIncompatible indicates units share a dimension but are not proportionally
	// related for the requested operation (e.g. overlapping °C/K in derived arithmetic).
	ErrIncompatible = fmt.Errorf("incompatible units")

	// ErrDivByZero indicates division by a zero scalar or zero-magnitude divisor.
	ErrDivByZero = fmt.Errorf("division by zero")

	// ErrInvalidUnit indicates a missing, nil, or unregistered unit.
	ErrInvalidUnit = fmt.Errorf("invalid unit")

	// ErrRoot indicates an nth-root that is undefined for the quantity
	// (n < 2, even root of a negative value, or exponents not divisible by n).
	ErrRoot = fmt.Errorf("root not defined")

	// ErrSyntax indicates the input string could not be parsed.
	ErrSyntax = fmt.Errorf("syntax error")
)
