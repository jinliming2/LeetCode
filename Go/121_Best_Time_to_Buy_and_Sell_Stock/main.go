package main

import "math"

// Runtime: 4 ms, faster than 97.03% of Go online submissions for Best Time to Buy and Sell Stock.
// Memory Usage: 3.1 MB, less than 17.66% of Go online submissions for Best Time to Buy and Sell Stock.
func maxProfit(prices []int) int {
	dp0, dp1 := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		a, b := dp1+prices[i], -prices[i]
		if dp0 < a {
			dp0 = a
		}
		if dp1 < b {
			dp1 = b
		}
	}
	return dp0
}
