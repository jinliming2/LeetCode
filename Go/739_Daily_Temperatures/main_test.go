package main

import "testing"

func TestDailyTemperatures(t *testing.T) {
	for input, output := range map[*[]int][]int{
		{73, 74, 75, 71, 69, 72, 76, 73}: {1, 1, 4, 2, 1, 1, 0, 0},
	} {
		value := dailyTemperatures(*input)
		if len(value) != len(output) {
			t.Errorf("Expected len(dailyTemperatures(%v)) == %d , But got %d", input, len(output), len(value))
			continue
		}
		for index, item := range output {
			if value[index] != item {
				t.Errorf("Expected dailyTemperatures(%v) == %v, But got %v", input, output, value)
				break
			}
		}
	}
}
