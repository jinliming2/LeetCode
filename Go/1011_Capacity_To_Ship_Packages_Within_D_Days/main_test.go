package main

import "testing"

type inputData struct {
	weights []int
	D       int
}

func TestShipWithinDays(t *testing.T) {
	for input, output := range map[*inputData]int{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5}: 15,
		{[]int{3, 2, 2, 4, 1, 4}, 3}:              6,
		{[]int{1, 2, 3, 1, 1}, 4}:                 3,
	} {
		if value := shipWithinDays(input.weights, input.D); value != output {
			t.Errorf("Expected shipWithinDays(%v, %d) == %d, But got %d", input.weights, input.D, output, value)
		}
	}
}
