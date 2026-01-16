package reverseinteger

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{1534236469, 0},
	}
	for _, tt := range tests {
		if got := reverse(tt.x); got != tt.want {
			t.Errorf("reverse(%d) = %d, want %d", tt.x, got, tt.want)
		}
	}
}
