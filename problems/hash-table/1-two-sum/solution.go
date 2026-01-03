package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		remain := target - num
		if j, ok := m[remain]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}
