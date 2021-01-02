package main

import "testing"

func TestMinWindow(t *testing.T) {
	for input, output := range map[[2]string]string{
		{"ADOBECODEBANC", "ABC"}: "BANC",
		{"a", "a"}:               "a",
	} {
		if value := minWindow(input[0], input[1]); value != output {
			t.Errorf("Expected minWindow(%s, %s) == %s, But got %s", input[0], input[1], output, value)
		}
	}
}
