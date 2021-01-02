package main

// Runtime: 36 ms, faster than 67.26% of Go online submissions for Binary Search.
// Memory Usage: 6.4 MB, less than 79.37% of Go online submissions for Binary Search.
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2 // n will be in the range [1, 10000].
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
