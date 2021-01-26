package main

import "testing"

func TestNextGreaterElements(t *testing.T) {
	for input, output := range map[*[]int][]int{
		{1, 2, 1}: {2, -1, 2},
	} {
		value := nextGreaterElements(*input)
		if len(value) != len(output) {
			t.Errorf("Expected len(nextGreaterElements(%v)) == %d , But got %d", input, len(output), len(value))
			continue
		}
		for index, item := range output {
			if value[index] != item {
				t.Errorf("Expected nextGreaterElements(%v) == %v, But got %v", input, output, value)
				break
			}
		}
	}
}
