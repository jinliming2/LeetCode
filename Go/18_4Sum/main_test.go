package main

import "testing"

type inputData struct {
	nums   []int
	target int
}

func TestFourSum(t *testing.T) {
	for input, output := range map[*inputData][][]int{
		{[]int{1, 0, -1, 0, -2, 2}, 0}: {{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		{[]int{}, 0}:                   {},
		{[]int{-3, -1, 0, 2, 4, 5}, 0}: {{-3, -1, 0, 4}},
	} {
		value := fourSum(input.nums, input.target)
		if len(value) != len(output) {
			t.Errorf("Expected len(fourSum(%v, %d)) == %d , But got %d", input.nums, input.target, len(output), len(value))
			continue
		}
	itemLoop:
		for _, item := range output {
		answerLoop:
			for _, answer := range value {
				if len(answer) != len(item) {
					t.Errorf("Expected len(items in fourSum(%v, %d)) == %d , But got %d", input.nums, input.target, len(item), len(answer))
					continue
				}
			itemVLoop:
				for _, itemV := range item {
					for _, answerV := range answer {
						if itemV == answerV {
							continue itemVLoop
						}
					}
					continue answerLoop
				}
				continue itemLoop
			}
			t.Errorf("Expected fourSum(%v, %d) == %v, But got %v", input.nums, input.target, output, value)
			break
		}
	}
}
