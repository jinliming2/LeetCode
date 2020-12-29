package main

// Runtime: 56 ms, faster than 18.91% of Go online submissions for Coin Change.
// Memory Usage: 6.5 MB, less than 20.63% of Go online submissions for Coin Change.
func coinChange0(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount)

	var sub func(amount int) int
	sub = func(amount int) (result int) {
		if cache := dp[amount-1]; cache != 0 {
			return cache
		}

		defer func() {
			dp[amount-1] = result
		}()

		min := -1
		for _, coin := range coins {
			if amount < coin {
				continue
			}
			if amount == coin {
				return 1
			}
			if c := sub(amount - coin); c >= 0 {
				c++
				if min < 0 || c < min {
					min = c
				}
			}
		}
		return min
	}

	return sub(amount)
}
