package u

import (
	"testing"
)

func TestDerivedDimensionSymbol(t *testing.T) {
	tests := []struct {
		name    string
		dim     DerivedDimension
		options []SymbolOption
		want    string
	}{
		{
			name: "speed default",
			dim:  DimSpeed,
			want: "L·T^-1",
		},
		{
			name: "speed superscript",
			dim:  DimSpeed,
			options: []SymbolOption{
				WithExpSign(ExpSignSup),
			},
			want: "L·T⁻¹",
		},
		{
			name: "speed slash",
			dim:  DimSpeed,
			options: []SymbolOption{
				WithDivSign(DivSignSlash),
			},
			want: "L/T",
		},
		{
			name: "force default",
			dim:  DerivedDimension{M: 1, L: 1, T: -2},
			want: "M·L·T^-2",
		},
		{
			name: "force slash",
			dim:  DerivedDimension{M: 1, L: 1, T: -2},
			options: []SymbolOption{
				WithDivSign(DivSignSlash),
			},
			want: "M·L/T^2",
		},
		{
			name: "speed tml order",
			dim:  DimSpeed,
			options: []SymbolOption{
				WithDimOrder(DimOrderTML),
			},
			want: "T^-1·L",
		},
		{
			name: "force tml order",
			dim:  DerivedDimension{M: 1, L: 1, T: -2},
			options: []SymbolOption{
				WithDimOrder(DimOrderTML),
			},
			want: "T^-2·L·M",
		},
		{
			name: "force tml slash",
			dim:  DerivedDimension{M: 1, L: 1, T: -2},
			options: []SymbolOption{
				WithDimOrder(DimOrderTML),
				WithDivSign(DivSignSlash),
			},
			want: "L·M/T^2",
		},
		{
			name: "speed no separator",
			dim:  DimSpeed,
			options: []SymbolOption{
				WithMulSign(MulSignNone),
			},
			want: "LT^-1",
		},
		{
			name: "dimensionless",
			dim:  DerivedDimension{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dim.Symbol(tt.options...)
			if got != tt.want {
				t.Fatalf("Symbol() = %q, want %q", got, tt.want)
			}
		})
	}
}
