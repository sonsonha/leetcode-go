package validparenthesis

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (c == ')' && top != '(') ||
				(c == ']' && top != '[') ||
				(c == '}' && top != '{') {
				return false
			}
		}
	}
	return len(stack) == 0
}
