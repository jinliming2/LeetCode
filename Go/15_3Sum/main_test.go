package main

import "testing"

func TestThreeSum(t *testing.T) {
	for input, output := range map[*[]int][][]int{
		{-1, 0, 1, 2, -1, -4}: {{-1, -1, 2}, {-1, 0, 1}},
		{}:                    {},
		{0}:                   {},
		{-2, 0, 0, 2, 2}:      {{-2, 0, 2}},
		{0, 0, 0}:             {{0, 0, 0}},
	} {
		value := threeSum(*input)
		if len(value) != len(output) {
			t.Errorf("Expected len(threeSum(%v)) == %d , But got %d", input, len(output), len(value))
			continue
		}
	itemLoop:
		for _, item := range output {
		answerLoop:
			for _, answer := range value {
				if len(answer) != len(item) {
					t.Errorf("Expected len(items in threeSum(%v)) == %d , But got %d", input, len(item), len(answer))
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
			t.Errorf("Expected threeSum(%v) == %v, But got %v", input, output, value)
			break
		}
	}
}
