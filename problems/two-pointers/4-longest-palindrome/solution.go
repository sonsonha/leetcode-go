package longestpalindrome

func LongestPalindrome(s string) string {
	left, right := 0, 0

	expandAroundCenter := func(s string, left, right int) (int, int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return left + 1, right - 1
	}

	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)

		if r1-l1 > right-left {
			left, right = l1, r1
		}
		if r2-l2 > right-left {
			left, right = l2, r2
		}
	}
	return s[left : right+1]
}
