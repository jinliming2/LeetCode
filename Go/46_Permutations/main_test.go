package main

import "testing"

func TestPermute(t *testing.T) {
	for input, output := range map[*[]int][][]int{
		{1, 2, 3}: {{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		{0, 1}:    {{0, 1}, {1, 0}},
		{1}:       {{1}},
	} {
		value := permute(*input)
		if len(output) != len(value) {
			t.Errorf("Expected len(permute(%v)) == %d , But got %d", output, len(output), len(value))
			continue
		}
	itemLoop:
		for _, item := range output {
		answerLoop:
			for _, answer := range value {
				for index, itemV := range item {
					if answer[index] != itemV {
						continue answerLoop
					}
				}
				continue itemLoop
			}
			t.Errorf("Expected permute(%v) == %v, But got %v", input, output, value)
			break itemLoop
		}
	}
}
