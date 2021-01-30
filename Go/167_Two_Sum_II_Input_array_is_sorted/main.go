package main

// Runtime: 4 ms, faster than 96.02% of Go online submissions for Two Sum II - Input array is sorted.
// Memory Usage: 3 MB, less than 34.83% of Go online submissions for Two Sum II - Input array is sorted.
// https://leetcode.com/submissions/detail/449627725/

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
