package main

import "testing"

type inputData struct {
	numbers []int
	target  int
}

func TestTwoSum(t *testing.T) {
	for input, output := range map[*inputData][]int{
		{[]int{2, 7, 11, 15}, 9}: {1, 2},
		{[]int{2, 3, 4}, 6}:      {1, 3},
		{[]int{-1, 0}, -1}:       {1, 2},
	} {
		value := twoSum(input.numbers, input.target)
		if len(value) != len(output) {
			t.Errorf("Expected len(twoSum(%v, %d)) == %d , But got %d", input.target, input.numbers, len(output), len(value))
			continue
		}
		for index, item := range output {
			if value[index] != item {
				t.Errorf("Expected twoSum(%v, %d) == %v, But got %v", input.target, input.numbers, output, value)
				break
			}
		}
	}
}
