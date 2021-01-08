package main

import "testing"

var testCase = map[*[]int]bool{
	{1, 5, 11, 5}: true,
	{1, 2, 3, 5}:  false,
	{2, 2, 3, 5}:  false,
	{1, 1}:        true,
}

func TestCanPartitionV1(t *testing.T) {
	for input, output := range testCase {
		if value := canPartitionV1(*input); value != output {
			t.Errorf("Expected canPartitionV1(%v) == %t, But got %t", input, output, value)
		}
	}
}

func TestCanPartitionV2(t *testing.T) {
	for input, output := range testCase {
		if value := canPartitionV2(*input); value != output {
			t.Errorf("Expected canPartitionV2(%v) == %t, But got %t", input, output, value)
		}
	}
}
