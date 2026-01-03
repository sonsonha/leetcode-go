package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 36},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 4, 23, 34, 37},
			target: 37,
			want:   []int{0, 3},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !equal(got, tt.want) {
			t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
