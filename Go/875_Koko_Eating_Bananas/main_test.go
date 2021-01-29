package main

import "testing"

type inputData struct {
	piles []int
	H     int
}

func TestMinEatingSpeed(t *testing.T) {
	for input, output := range map[*inputData]int{
		{[]int{3, 6, 7, 11}, 8}:       4,
		{[]int{30, 11, 23, 4, 20}, 5}: 30,
		{[]int{30, 11, 23, 4, 20}, 6}: 23,
		{[]int{312884470}, 312884469}: 2,
	} {
		if value := minEatingSpeed(input.piles, input.H); value != output {
			t.Errorf("Expected minEatingSpeed(%v, %d) == %d, But got %d", input.piles, input.H, output, value)
		}
	}
}
