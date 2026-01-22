package containerwithmostwater

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		h := min(height[left], height[right])
		res = max(res, h*(right-left))
		if height[right] < height[left] {
			right--
		} else {
			left++
		}
	}
	return res
}
