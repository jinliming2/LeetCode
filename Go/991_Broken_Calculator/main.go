package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Broken Calculator.
// Memory Usage: 2 MB, less than 21.43% of Go online submissions for Broken Calculator.
// https://leetcode.com/submissions/detail/458866163/

func brokenCalc(X, Y int) int {
	if X >= Y {
		return X - Y
	}
	if Y%2 == 0 {
		return brokenCalc(X, Y/2) + 1
	}
	return brokenCalc(X, Y+1) + 1
}
