package longestpalindrome

import "testing"

func TestLongestPlaindroi(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "babad",
			want: "bab",
		},
		{
			s:    "cbbd",
			want: "bb",
		},
		{
			s:    "a",
			want: "a",
		},
	}
	for _, tt := range tests {
		got := LongestPalindrome(tt.s)
		if got != tt.want {
			t.Errorf("LongestPalindrome(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}
