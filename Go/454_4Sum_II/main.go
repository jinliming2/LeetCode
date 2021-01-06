package main

// Runtime: 64 ms, faster than 81.73% of Go online submissions for 4Sum II.
// Memory Usage: 21.1 MB, less than 70.59% of Go online submissions for 4Sum II.
// https://leetcode.com/submissions/detail/439389933/
func fourSumCount(A []int, B []int, C []int, D []int) int {
	dp := map[int]int{}
	for _, a := range A {
		for _, b := range B {
			dp[a+b]++
		}
	}
	count := 0
	for _, c := range C {
		for _, d := range D {
			count += dp[-(c + d)]
		}
	}
	return count
}
