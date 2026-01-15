package zizagconversion

import (
	"testing"
)

func TestConvet(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		want    string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
	}
	for _, tt := range tests {
		if got := convert(tt.s, tt.numRows); got != tt.want {
			t.Errorf("Convert(%q, %d) = %q, want %q", tt.s, tt.numRows, got, tt.want)
		}
	}
}
