package longestsubstring

import "testing"

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "abcabcbb",
			want: 3,
		},
		{
			s:    "bbbbb",
			want: 1,
		},
		{
			s:    "pwwkew",
			want: 3,
		},
	}

	for _, tt := range tests {
		got := LongestSubstring(tt.s)
		if got != tt.want {
			t.Errorf("LongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
