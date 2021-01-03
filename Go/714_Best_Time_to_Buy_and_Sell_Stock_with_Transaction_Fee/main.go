package main

import "math"

// Runtime: 92 ms, faster than 74.47% of Go online submissions for Best Time to Buy and Sell Stock with Transaction Fee.
// Memory Usage: 7.4 MB, less than 40.43% of Go online submissions for Best Time to Buy and Sell Stock with Transaction Fee.
func maxProfit(prices []int, fee int) int {
	dp0, dp1 := 0, math.MinInt32
	for _, p := range prices {
		a, b := dp1+p-fee, dp0-p
		if dp0 < a {
			dp0 = a
		}
		if dp1 < b {
			dp1 = b
		}
	}
	return dp0
}
