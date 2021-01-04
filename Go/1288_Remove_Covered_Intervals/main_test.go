package main

import "testing"

func TestRemoveCoveredIntervals(t *testing.T) {
	for input, output := range map[*[][]int]int{
		{{1, 4}, {3, 6}, {2, 8}}:    2,
		{{1, 4}, {2, 3}}:            1,
		{{0, 10}, {5, 12}}:          2,
		{{3, 10}, {4, 10}, {5, 11}}: 2,
		{{1, 2}, {1, 4}, {3, 4}}:    1,
	} {
		if value := removeCoveredIntervals(*input); value != output {
			t.Errorf("Expected openLockBFS(%v) == %d, But got %d", input, output, value)
		}
	}
}
