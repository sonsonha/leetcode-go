package longestsubstring

func LongestSubstring(s string) int {
	m := make(map[byte]int, 255)
	res := 0
	left := 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		if prev, ok := m[c]; ok && prev >= left {
			left = prev + 1
		}
		m[c] = right
		res = max(res, right-left+1)
	}
	return res
}
