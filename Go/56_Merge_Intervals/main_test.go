package main

import "testing"

func TestMerge(t *testing.T) {
	for input, output := range map[*[][]int][][]int{} {
		value := merge(*input)
		if len(value) != len(output) {
			t.Errorf("Expected len(merge(%v)) == %d , But got %d", input, len(output), len(value))
			continue
		}
	itemLoop:
		for _, interval := range output {
			for _, answer := range value {
				if answer[0] == interval[0] && answer[1] == interval[1] {
					continue itemLoop
				}
			}
			t.Errorf("Expected merge(%v) == %v, But got %v", input, output, value)
			break
		}
	}
}
