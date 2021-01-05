package main

import "testing"

type inputData struct {
	nums   []int
	target int
}

func TestTwoSum(t *testing.T) {
	for input, output := range map[*inputData][]int{
		{nums: []int{2, 7, 11, 15}, target: 9}: {0, 1},
		{nums: []int{3, 2, 4}, target: 6}:      {1, 2},
		{nums: []int{3, 3}, target: 6}:         {0, 1},
		{nums: []int{2, 5, 5, 11}, target: 10}: {1, 2},
	} {
		value := twoSum(input.nums, input.target)
		if len(value) != len(output) {
			t.Errorf("Expected len(twoSum(%v, %d)) == %d , But got %d", input.nums, input.target, len(output), len(value))
			continue
		}
		if output[0] != value[0] || output[1] != value[1] {
			t.Errorf("Expected twoSum(%v, %d) == %v, But got %v", input.nums, input.target, output, value)
		}
	}
}
