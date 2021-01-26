package main

// Runtime: 92 ms, faster than 37.91% of Go online submissions for Daily Temperatures.
// Memory Usage: 7 MB, less than 65.36% of Go online submissions for Daily Temperatures.
// https://leetcode.com/submissions/detail/448108031/

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := []int{}
	for index := len(T) - 1; index >= 0; index-- {
		v := T[index]
		for len(stack) > 0 && T[stack[len(stack)-1]] <= v {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[index] = stack[len(stack)-1] - index
		}
		stack = append(stack, index)
	}
	return res
}
