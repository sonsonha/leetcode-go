package validparenthesis

import (
	"testing"
)

func TestValidParetnthesis(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()()))({}[]])", false},
		{"(){}[]", true},
		{")()(){]]}", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
	}
	for _, tt := range tests {
		if got := isValid(tt.input); got != tt.want {
			t.Errorf("isValid(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
