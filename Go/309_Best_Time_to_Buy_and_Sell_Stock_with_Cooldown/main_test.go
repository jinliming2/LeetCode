package main

import "testing"

func TestMaxProfit(t *testing.T) {
	for input, output := range map[*[]int]int{
		{1, 2, 3, 0, 2}: 3,
	} {
		if value := maxProfit(*input); value != output {
			t.Errorf("Expected maxProfit(%v) == %d, But got %d", input, output, value)
		}
	}
}
