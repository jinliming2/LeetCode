package main

import "testing"

type inputData struct {
	nums []int
	k    int
}

var testCase = map[*inputData][]int{
	{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}: {3, 3, 5, 5, 6, 7},
	{[]int{1}, 1}:                        {1},
	{[]int{1, -1}, 1}:                    {1, -1},
	{[]int{9, 11}, 2}:                    {11},
	{[]int{4, -2}, 2}:                    {4},
	{[]int{1, 3, 1, 2, 0, 5}, 3}:         {3, 3, 2, 5},
}

func TestMaxSlidingWindowV1(t *testing.T) {
	for input, output := range testCase {
		value := maxSlidingWindowV1(input.nums, input.k)
		if len(value) != len(output) {
			t.Errorf("Expected len(maxSlidingWindowV1(%v, %v)) == %d , But got %d", input.nums, input.k, len(output), len(value))
			continue
		}
		for index, item := range output {
			if value[index] != item {
				t.Errorf("Expected maxSlidingWindowV1(%v, %v) == %v, But got %v", input.nums, input.k, output, value)
				break
			}
		}
	}
}

func TestMaxSlidingWindowV2(t *testing.T) {
	for input, output := range testCase {
		value := maxSlidingWindowV2(input.nums, input.k)
		if len(value) != len(output) {
			t.Errorf("Expected len(maxSlidingWindowV2(%v, %v)) == %d , But got %d", input.nums, input.k, len(output), len(value))
			continue
		}
		for index, item := range output {
			if value[index] != item {
				t.Errorf("Expected maxSlidingWindowV2(%v, %v) == %v, But got %v", input.nums, input.k, output, value)
				break
			}
		}
	}
}
