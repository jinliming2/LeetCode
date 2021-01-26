package main

// Runtime: 36 ms, faster than 25.00% of Go online submissions for Next Greater Element II.
// Memory Usage: 6.4 MB, less than 45.65% of Go online submissions for Next Greater Element II.
// https://leetcode.com/submissions/detail/448099469/

func nextGreaterElements(nums []int) []int {
	stack := []int{}
	for index := len(nums) - 1; index >= 0; index-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[index] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[index])
	}
	for index := len(nums) - 1; index >= 0; index-- {
		v := nums[index]
		for len(stack) > 0 && stack[len(stack)-1] <= v {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			nums[index] = -1
		} else {
			nums[index] = stack[len(stack)-1]
		}
		stack = append(stack, v)
	}
	return nums
}
