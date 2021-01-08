package main

// Runtime: 24 ms, faster than 64.29% of Go online submissions for Partition Equal Subset Sum.
// Memory Usage: 2.5 MB, less than 83.50% of Go online submissions for Partition Equal Subset Sum.
// https://leetcode.com/submissions/detail/440319319/
func canPartitionV2(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	if num := nums[0]; num <= sum {
		dp[num] = true
	}
	for _, num := range nums[1:] {
		for i := sum; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[sum]
}
