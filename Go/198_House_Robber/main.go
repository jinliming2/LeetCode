package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
// Memory Usage: 2 MB, less than 22.79% of Go online submissions for House Robber.
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp0, dp1 := 0, nums[0]
	for _, p := range nums[1:] {
		b := dp0 + p
		if dp0 < dp1 {
			dp0 = dp1
		}
		dp1 = b
	}
	if dp0 > dp1 {
		return dp0
	}
	return dp1
}
