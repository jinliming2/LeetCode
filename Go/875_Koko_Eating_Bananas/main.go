package main

// Runtime: 48 ms, faster than 55.93% of Go online submissions for Koko Eating Bananas.
// Memory Usage: 6.3 MB, less than 38.98% of Go online submissions for Koko Eating Bananas.
// https://leetcode.com/submissions/detail/449374903/

func minEatingSpeed(piles []int, H int) int {
	max := 0
	for _, p := range piles {
		if p > max {
			max = p
		}
	}
	left, right := 1, max+1
	for left < right {
		mid := left + (right-left)/2
		days := 0
		for _, p := range piles {
			days += p / mid
			if p%mid > 0 {
				days++
			}
		}
		if days <= H {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
