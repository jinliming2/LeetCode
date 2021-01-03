package main

import "math"

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock IV.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Best Time to Buy and Sell Stock IV.
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp0 := make([]int, k+1)
	dp1 := make([]int, k+1) // dp1 == k is not valid
	minf := math.MinInt32 + prices[0]
	for i := 1; i <= k; i++ {
		dp0[i] = minf
		dp1[i] = minf
	}
	dp1[0] = minf
	for _, p := range prices {
		for i := 0; i <= k; i++ {
			t0, t1 := dp0[i], dp0[i]-p
			if i > 0 {
				t0 = dp1[i-1] + p
			}
			if dp0[i] < t0 {
				dp0[i] = t0
			}
			if dp1[i] < t1 {
				dp1[i] = t1
			}
		}
	}
	max := dp0[0]
	for _, m := range dp0 {
		if m > max {
			max = m
		}
	}
	if max < 0 {
		return 0
	}
	return max
}
