package main

import "sort"

// Runtime: 32 ms, faster than 72.50% of Go online submissions for 3Sum.
// Memory Usage: 7 MB, less than 71.58% of Go online submissions for 3Sum.
// https://leetcode.com/submissions/detail/439373925/
func threeSum(nums []int) [][]int {
	results := [][]int{}
	length := len(nums)
	if length < 3 {
		return results
	}
	sort.Ints(nums)
	var last int
	for i, n := range nums[:length-2] {
		if i > 0 && n == last {
			continue
		}
		last = n
		findTwoSum(nums[i+1:], n, &results)
	}
	return results
}

func findTwoSum(nums []int, n int, results *[][]int) {
	left, right := 0, len(nums)-1
	for left < right {
		l, r := nums[left], nums[right]
		sum := l + r
		if sum == -n {
			*results = append(*results, []int{n, l, r})
			for left < right && nums[left] == l {
				left++
			}
			for left < right && nums[right] == r {
				right--
			}
		} else if sum < -n {
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
