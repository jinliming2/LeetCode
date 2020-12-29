package main

// Runtime: 20 ms, faster than 36.68% of Go online submissions for Coin Change.
// Memory Usage: 5.9 MB, less than 71.92% of Go online submissions for Coin Change.
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for dpI := range dp {
		if dpI == 0 { // base
			dp[dpI] = 0
			continue
		}
		dp[dpI] = amount + 1 // Make it impossible
		for _, coin := range coins {
			if dpI < coin {
				continue
			}
			if sub := dp[dpI-coin] + 1; dp[dpI] > sub {
				dp[dpI] = sub
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
