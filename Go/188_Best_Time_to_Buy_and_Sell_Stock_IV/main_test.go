package main

import "testing"

type inputData struct {
	k      int
	prices []int
}

func TestMaxProfit(t *testing.T) {
	for input, output := range map[*inputData]int{
		{2, []int{2, 4, 1}}:          2,
		{2, []int{3, 2, 6, 5, 0, 3}}: 7,
		{2, []int{}}:                 0,
	} {
		if value := maxProfit(input.k, input.prices); value != output {
			t.Errorf("Expected maxProfit(%d, %v) == %d, But got %d", input.k, input.prices, output, value)
		}
	}
}
