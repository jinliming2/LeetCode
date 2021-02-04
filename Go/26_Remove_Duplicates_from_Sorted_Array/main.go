package main

// Runtime: 4 ms, faster than 99.46% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 4.6 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted Array.
// https://leetcode.com/submissions/detail/451976353/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	p := 0
	for _, num := range nums {
		if num != nums[p] {
			p++
			nums[p] = num
		}
	}
	return p + 1
}
