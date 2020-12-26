package main

import "testing"

type inputData struct {
	coins  []int
	amount int
}

func TestCoinChange(t *testing.T) {
	for input, output := range map[*inputData]int{
		{[]int{1, 2, 5}, 11}:             3,
		{[]int{2}, 3}:                    -1,
		{[]int{1}, 0}:                    0,
		{[]int{1}, 1}:                    1,
		{[]int{1}, 2}:                    2,
		{[]int{186, 419, 83, 408}, 6249}: 20,
	} {
		if value := coinChange(input.coins, input.amount); value != output {
			t.Errorf("Expected coinChange(%v, %d) == %d, But got %d", input.coins, input.amount, output, value)
		}
	}
}
