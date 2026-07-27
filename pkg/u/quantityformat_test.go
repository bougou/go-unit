package u

import "testing"

func TestDelimitFloat(t *testing.T) {
	tests := []struct {
		val       float64
		precision int
		delimiter NumberDelimiter
		want      string
	}{
		{1234.5, 1, NumberDelimiterNone, "1234.5"},
		{1234.5, 1, NumberDelimiterComma, "1,234.5"},
		{1234.5, 1, NumberDelimiterUnderscore, "1_234.5"},
		{-1234567.89, 2, NumberDelimiterComma, "-1,234,567.89"},
		{42, 0, NumberDelimiterComma, "42"},
		{123456, 0, NumberDelimiterComma, "123,456"},
	}
	for _, tt := range tests {
		got := DelimitFloat(tt.val, tt.precision, tt.delimiter)
		if got != tt.want {
			t.Errorf("DelimitFloat(%v, %d, %q) = %q, want %q",
				tt.val, tt.precision, tt.delimiter, got, tt.want)
		}
	}
}

func TestQuantityFormat(t *testing.T) {
	q := Quantity{Value: 1234.5, Unit: Unit(Meter)}

	if got := q.Format(); got != "1234.5 m" {
		t.Fatalf("Format() = %q, want %q", got, "1234.5 m")
	}
	if got := q.String(); got != q.Format() {
		t.Fatalf("String() = %q, Format() = %q", got, q.Format())
	}
	if got := q.Format(WithPrecision(1), WithNumberDelimiter(NumberDelimiterComma)); got != "1,234.5 m" {
		t.Fatalf("Format(comma) = %q, want %q", got, "1,234.5 m")
	}
	if got := q.Format(WithPrecision(0), WithNumberDelimiter(NumberDelimiterUnderscore)); got != "1_234 m" {
		t.Fatalf("Format(underscore) = %q, want %q", got, "1_234 m")
	}
}

func TestDerivedQuantityFormat(t *testing.T) {
	q := NewDerivedQuantity(1234.5, Newton)

	if got := q.String(); got != "1234.5 N" {
		t.Fatalf("String() = %q, want %q", got, "1234.5 N")
	}
	if got := q.Format(); got != "1234.5 N" {
		t.Fatalf("Format() = %q, want %q", got, "1234.5 N")
	}
	if got := q.Format(WithCompoundSymbol(true)); got != "1234.5 kg·m·s^-2" {
		t.Fatalf("Format(compound) = %q, want %q", got, "1234.5 kg·m·s^-2")
	}
	if got := q.Format(WithCompoundSymbol(true), WithExpSign(ExpSignSup)); got != "1234.5 kg·m·s⁻²" {
		t.Fatalf("Format(compound+sup) = %q, want %q", got, "1234.5 kg·m·s⁻²")
	}
	if got := q.Format(
		WithPrecision(1),
		WithNumberDelimiter(NumberDelimiterComma),
	); got != "1,234.5 N" {
		t.Fatalf("Format(full) = %q, want %q", got, "1,234.5 N")
	}
	if got := q.Format(WithCompoundSymbol(true), WithDivSign(DivSignSlash)); got != "1234.5 kg·m/s^2" {
		t.Fatalf("Format(slash) = %q, want %q", got, "1234.5 kg·m/s^2")
	}

	nilUnit := NewDerivedQuantity(3.14, nil)
	if got := nilUnit.Format(); got != "3.14" {
		t.Fatalf("Format(nil unit) = %q, want %q", got, "3.14")
	}
}
