package twosum

// Problem: Two Sum
// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
//
// Example:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// TwoSum finds two numbers in nums that add up to target and returns their indices
func TwoSum(nums []int, target int) []int {
	// Map to store value -> index
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, exists := numMap[complement]; exists {
			return []int{idx, i}
		}
		numMap[num] = i
	}

	return nil // Should not reach here per problem constraints
}
