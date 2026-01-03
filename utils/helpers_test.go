package utils

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 2},
		{5, 3, 5},
		{-1, -2, -1},
		{0, 0, 0},
	}

	for _, tt := range tests {
		result := Max(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Max(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 1},
		{5, 3, 3},
		{-1, -2, -2},
		{0, 0, 0},
	}

	for _, tt := range tests {
		result := Min(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Min(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		x, expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
	}

	for _, tt := range tests {
		result := Abs(tt.x)
		if result != tt.expected {
			t.Errorf("Abs(%d) = %d, expected %d", tt.x, result, tt.expected)
		}
	}
}
