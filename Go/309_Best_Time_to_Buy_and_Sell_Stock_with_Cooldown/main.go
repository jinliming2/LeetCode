package main

import "math"

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock with Cooldown.
// Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Best Time to Buy and Sell Stock with Cooldown.
func maxProfit(prices []int) int {
	dp00, dp01, dpb0 := 0, math.MinInt32, 0
	for _, p := range prices {
		a, b := dp01+p, dpb0-p
		dpb0 = dp00
		if dp00 < a {
			dp00 = a
		}
		if dp01 < b {
			dp01 = b
		}
	}
	return dp00
}
