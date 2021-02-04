package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	for input, output := range map[*[]int][]int{
		{1, 1, 2}:                      {1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}: {0, 1, 2, 3, 4},
	} {
		i := make([]int, len(*input))
		copy(i, *input)
		value := removeDuplicates(i)
		if value != len(output) {
			t.Errorf("Expected removeDuplicates(%v) == %d , But got %d", input, len(output), value)
			continue
		}
		for index, item := range output {
			if i[index] != item {
				t.Errorf("Expected removeDuplicates(%v)[%d] == %d , But got %d", input, index, item, i[index])
				break
			}
		}
	}
}
