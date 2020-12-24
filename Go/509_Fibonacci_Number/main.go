package main

func fib(n int) int {
	dp := [2]int{0, 1}

	if n < 2 {
		return dp[n]
	}

	x := 1
	y := -1

	for i := 2; i <= n; i, x, y = i+1, x+y, y*-1 {
		dp[x+y] = dp[0] + dp[1]
	}

	return dp[x]
}
