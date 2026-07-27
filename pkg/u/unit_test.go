package u

import "testing"

func TestValidateRegistry(t *testing.T) {
	if err := validateRegistry(); err != nil {
		t.Fatal(err)
	}
}

func TestTimeHourToBase(t *testing.T) {
	got := Time(1, Hour).Base()
	if got.Value != 3600 {
		t.Fatalf("1 hour in seconds = %v, want 3600", got.Value)
	}
	if got.Unit != Unit(Second) {
		t.Fatalf("base unit = %v, want second", got.Unit)
	}
}

func TestUnitSymbol(t *testing.T) {
	tests := []struct {
		unit Unit
		want string
	}{
		{Unit(Meter), "m"},
		{Unit(Meter.Prefix(Kilo)), "km"},
		{Unit(Hour), "h"},
		{Unit(Kelvin), "K"},
		{Unit(Celsius), "°C"},
	}

	for _, tt := range tests {
		if got := tt.unit.Symbol(); got != tt.want {
			t.Fatalf("Unit(%q).Symbol() = %q, want %q", tt.unit, got, tt.want)
		}
	}

	newtonKey := Unit(Newton.Key())
	if newtonKey.Symbol() != "N" {
		t.Fatalf("Unit(force).Symbol() = %q, want N", newtonKey.Symbol())
	}
	if newtonKey.Symbol(WithCompoundSymbol(true)) != "kg·m·s^-2" {
		t.Fatalf("Unit(force).Symbol(compound) = %q, want kg·m·s^-2", newtonKey.Symbol(WithCompoundSymbol(true)))
	}
}
