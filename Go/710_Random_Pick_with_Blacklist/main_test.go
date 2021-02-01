package main

import "testing"

type inputData struct {
	N         int
	blacklist []int
}

func TestSolution(t *testing.T) {
	for input, output := range map[*inputData][]int{
		{1, []int{}}:               {0},
		{2, []int{}}:               {0, 1},
		{3, []int{1}}:              {0, 2},
		{4, []int{2}}:              {0, 1, 3},
		{10, []int{9, 7, 1, 5, 3}}: {0, 2, 4, 6, 8},
	} {
		obj := Constructor(input.N, input.blacklist)
	testLoop:
		for i := 0; i < 100; i++ {
			value := obj.Pick()
			for _, item := range output {
				if item == value {
					continue testLoop
				}
			}
			t.Errorf("Expected Solution(%v).Pick() returns number in %v, But got %d", *input, output, value)
			break
		}
	}
}
