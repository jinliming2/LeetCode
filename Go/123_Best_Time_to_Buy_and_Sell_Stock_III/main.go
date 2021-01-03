package main

import "math"

// Runtime: 120 ms, faster than 88.24% of Go online submissions for Best Time to Buy and Sell Stock III.
// Memory Usage: 8.4 MB, less than 63.24% of Go online submissions for Best Time to Buy and Sell Stock III.
func maxProfit(prices []int) int {
	// dp00 always be 0
	dp01, dp10, dp11, dp20 := math.MinInt32, math.MinInt32+prices[0], math.MinInt32, math.MinInt32
	for _, p := range prices {
		b, c, d, e := -p, dp01+p, dp10-p, dp11+p // b == dp00-p
		if dp01 < b {
			dp01 = b
		}
		if dp10 < c {
			dp10 = c
		}
		if dp11 < d {
			dp11 = d
		}
		if dp20 < e {
			dp20 = e
		}
	}
	if dp10 < 0 && dp20 < 0 {
		return 0
	}
	if dp10 > dp20 {
		return dp10
	}
	return dp20
}
