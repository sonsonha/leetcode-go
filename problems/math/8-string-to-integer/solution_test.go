package stringtointeger

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"91283472332", 2147483647},
		{"+1", 1},
		{"   +0 123", 0},
		{"-5-", -5},
		{"3.14159", 3},
	}

	for _, tt := range tests {
		got := myAtoi(tt.input)
		if got != tt.want {
			t.Errorf("myAtoi(%q) = %d; want %d", tt.input, got, tt.want)
		}
	}
}
