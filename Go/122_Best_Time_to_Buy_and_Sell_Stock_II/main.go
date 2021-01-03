package main

import "math"

// Runtime: 4 ms, faster than 96.43% of Go online submissions for Best Time to Buy and Sell Stock II.
// Memory Usage: 3.1 MB, less than 16.67% of Go online submissions for Best Time to Buy and Sell Stock II.
func maxProfit(prices []int) int {
	dp0, dp1 := 0, math.MinInt32
	for _, p := range prices {
		a, b := dp1+p, dp0-p
		if dp0 < a {
			dp0 = a
		}
		if dp1 < b {
			dp1 = b
		}
	}
	return dp0
}
