package main

import "sort"

// Runtime: 12 ms, faster than 60.77% of Go online submissions for 4Sum.
// Memory Usage: 2.7 MB, less than 69.38% of Go online submissions for 4Sum.
// https://leetcode.com/submissions/detail/439381125/
func fourSum(nums []int, target int) [][]int {
	results := [][]int{}
	length := len(nums)
	if length < 4 {
		return results
	}
	sort.Ints(nums)
	var lastN int
	for i, n := range nums[:length-3] {
		if i > 0 && n == lastN {
			continue
		}
		lastN = n
		var lastM int
		for j, m := range nums[i+1:] {
			if j > 0 && m == lastM {
				continue
			}
			lastM = m
			mn := n + m
			left, right := i+j+2, length-1
			for left < right {
				l, r := nums[left], nums[right]
				sum := l + r + mn
				if sum == target {
					results = append(results, []int{n, m, l, r})
					for left < right && nums[left] == l {
						left++
					}
					for left < right && nums[right] == r {
						right--
					}
				} else if sum < target {
					for left < right && nums[left] == l {
						left++
					}
				} else {
					for left < right && nums[right] == r {
						right--
					}
				}
			}
		}
	}
	return results
}
