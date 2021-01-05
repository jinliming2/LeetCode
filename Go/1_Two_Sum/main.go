package main

// Runtime: 4 ms, faster than 94.41% of Go online submissions for Two Sum.
// Memory Usage: 3.2 MB, less than 99.06% of Go online submissions for Two Sum.
// https://leetcode.com/submissions/detail/438940160/
func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := length - 2; i >= 0; i-- {
		numI := nums[i]
		for j := i + 1; j < length; j++ {
			numJ := nums[j]
			if numI+numJ == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
