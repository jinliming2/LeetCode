package main

import "testing"

type inputData struct {
	A [][]int
	B [][]int
}

func TestIntervalIntersection(t *testing.T) {
	for input, output := range map[*inputData][][]int{
		{
			[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
		}: {{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		{
			[][]int{{3, 5}, {9, 20}},
			[][]int{{4, 5}, {7, 10}, {11, 12}, {14, 15}, {16, 20}},
		}: {{4, 5}, {9, 10}, {11, 12}, {14, 15}, {16, 20}},
	} {
		value := intervalIntersection(input.A, input.B)
		if len(value) != len(output) {
			t.Errorf("Expected len(intervalIntersection(%v, %v)) == %d , But got %d", input.A, input.B, len(output), len(value))
			continue
		}
	itemLoop:
		for _, item := range output {
			for _, answer := range value {
				if answer[0] == item[0] && answer[1] == item[1] {
					continue itemLoop
				}
			}
			t.Errorf("Expected intervalIntersection(%v, %v) == %v, But got %v", input.A, input.B, output, value)
			break
		}
	}
}
