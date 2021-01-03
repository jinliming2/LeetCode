package main

import "testing"

func TestRob(t *testing.T) {
	for input, output := range map[*[]int]int{
		{1, 2, 3, 1}:    4,
		{2, 7, 9, 3, 1}: 12,
	} {
		if value := rob(*input); value != output {
			t.Errorf("Expected rob(%v) == %d, But got %d", input, output, value)
		}
	}
}
