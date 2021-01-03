package main

import "testing"

func TestMaxProfit(t *testing.T) {
	for input, output := range map[*[]int]int{
		{3, 3, 5, 0, 0, 3, 1, 4}: 6,
		{1, 2, 3, 4, 5}:          4,
		{7, 6, 4, 3, 1}:          0,
		{1}:                      0,
	} {
		if value := maxProfit(*input); value != output {
			t.Errorf("Expected maxProfit(%v) == %d, But got %d", input, output, value)
		}
	}
}
