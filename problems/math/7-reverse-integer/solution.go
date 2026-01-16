package reverseinteger

func reverse(x int) int {
	const (
		maxInt32 = 1<<31 - 1
		minInt32 = -1 << 31
	)
	res := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		if res > maxInt32/10 || (res == maxInt32/10 && digit > 7) {
			return 0
		}
		if res < minInt32/10 || (res == minInt32/10 && digit < -8) {
			return 0
		}
		res = res*10 + digit
	}
	return res
}
