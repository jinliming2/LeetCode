package main

import "testing"

type inputData struct {
	nums   []int
	target int
}

func TestSearch(t *testing.T) {
	for input, output := range map[*inputData]int{
		{[]int{-1, 0, 3, 5, 9, 12}, 9}: 4,
		{[]int{-1, 0, 3, 5, 9, 12}, 2}: -1,
	} {
		if value := search(input.nums, input.target); value != output {
			t.Errorf("Expected search(%v, %d) == %d, But got %d", input.nums, input.target, output, value)
		}
	}
}
