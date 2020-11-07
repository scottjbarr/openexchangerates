package openexchangerates

import (
	"testing"
)

func TestHistoricalParams(t *testing.T) {
	tests := []struct {
		name   string
		params *HistoricalParams
		want   string
	}{
		{
			name:   "nil",
			params: nil,
			want:   "",
		},
		{
			name: "common",
			params: &HistoricalParams{
				Base:    "USD",
				Symbols: []string{"AUD", "GBP"},
			},
			want: "base=USD&prettyprint=false&show_alternative=false&symbols=AUD%2CGBP",
		},
		{
			name: "complete",
			params: &HistoricalParams{
				Base:            "USD",
				Symbols:         []string{"AUD", "GBP"},
				PrettyPrint:     true,
				ShowAlternative: true,
			},
			want: "base=USD&prettyprint=true&show_alternative=true&symbols=AUD%2CGBP",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.Encode()

			if got != tt.want {
				t.Errorf("got\n%s\nwant\n%s", got, tt.want)
			}
		})
	}
}
