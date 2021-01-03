package main

import "testing"

type inputData struct {
	prices []int
	fee    int
}

func TestMaxProfit(t *testing.T) {
	for input, output := range map[*inputData]int{
		{[]int{1, 3, 2, 8, 4, 9}, 2}: 8,
	} {
		if value := maxProfit(input.prices, input.fee); value != output {
			t.Errorf("Expected maxProfit(%v, %d) == %d, But got %d", input.prices, input.fee, output, value)
		}
	}
}
