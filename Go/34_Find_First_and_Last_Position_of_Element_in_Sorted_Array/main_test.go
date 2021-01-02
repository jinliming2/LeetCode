package main

import "testing"

type inputData struct {
	nums   []int
	target int
}

func TestSearchRange(t *testing.T) {
	for input, output := range map[*inputData][]int{
		{[]int{5, 7, 7, 8, 8, 10}, 8}: {3, 4},
		{[]int{5, 7, 7, 8, 8, 10}, 6}: {-1, -1},
		{[]int{}, 0}:                  {-1, -1},
		{[]int{2, 2}, 3}:              {-1, -1},
	} {
		value := searchRange(input.nums, input.target)
		if len(value) != len(output) {
			t.Errorf("Expected len(searchRange(%v, %d)) == %d , But got %d", input.nums, input.target, len(output), len(value))
			continue
		}
		for i, n := range value {
			if output[i] != n {
				t.Errorf("Expected searchRange(%v, %d) == %v, But got %v", input.nums, input.target, output, value)
				break
			}
		}
	}
}
