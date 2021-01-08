package main

// Runtime: 32 ms, faster than 54.68% of Go online submissions for Partition Equal Subset Sum.
// Memory Usage: 6.6 MB, less than 41.13% of Go online submissions for Partition Equal Subset Sum.
// https://leetcode.com/submissions/detail/440311031/
func canPartitionV1(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	length := len(nums)
	dp := make([][]bool, length)
	for i, num := range nums {
		dp[i] = make([]bool, sum)
		if i == 0 {
			if num < sum {
				dp[i][num-1] = true
			}
			continue
		}
		for j := 0; j < sum; j++ {
			sub := j - num + 1
			if sub == 0 {
				dp[i][j] = true
			} else if sub < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][sub-1]
			}
		}
	}
	return dp[length-1][sum-1]
}
