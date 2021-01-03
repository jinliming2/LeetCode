package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber II.
// Memory Usage: 2 MB, less than 50.59% of Go online submissions for House Robber II.
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	length-- // last index
	dp := [4]int{}
	for i, p := range nums {
		if i != length {
			b := dp[0] + p
			if dp[0] < dp[1] {
				dp[0] = dp[1]
			}
			dp[1] = b
		}
		if i != 0 {
			b := dp[2] + p
			if dp[2] < dp[3] {
				dp[2] = dp[3]
			}
			dp[3] = b
		}
	}
	max := 0
	for _, m := range dp {
		if m > max {
			max = m
		}
	}
	return max
}
