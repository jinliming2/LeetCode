package main

// Runtime: 32 ms, faster than 96.77% of Go online submissions for Capacity To Ship Packages Within D Days.
// Memory Usage: 6.2 MB, less than 95.16% of Go online submissions for Capacity To Ship Packages Within D Days.
// https://leetcode.com/submissions/detail/449378694/

func shipWithinDays(weights []int, D int) int {
	max, sum := 0, 0
	for _, weight := range weights {
		sum += weight
		if weight > max {
			max = weight
		}
	}
	left, right := max, sum+1
	for left < right {
		mid := left + (right-left)/2
		days, stack := 0, 0
		for _, weight := range weights {
			stack += weight
			if stack > mid {
				days++
				stack = weight
			}
		}
		if stack > 0 {
			days++
		}
		if days <= D {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
