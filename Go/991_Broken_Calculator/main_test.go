package main

import "testing"

func TestBrokenCalc(t *testing.T) {
	for input, output := range map[[2]int]int{
		{2, 3}:    2,
		{5, 8}:    2,
		{3, 10}:   3,
		{1024, 1}: 1023,
	} {
		if value := brokenCalc(input[0], input[1]); value != output {
			t.Errorf("Expected brokenCalc(%d, %d) = %d, But got %d", input[0], input[1], output, value)
		}
	}
}
