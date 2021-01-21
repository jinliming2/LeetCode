package main

import "testing"

func TestSolve(t *testing.T) {
	for input, output := range map[*[][]byte][][]byte{
		{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}: {
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		},
		{
			{'O', 'O', 'O'},
			{'O', 'O', 'O'},
			{'O', 'O', 'O'},
		}: {
			{'O', 'O', 'O'},
			{'O', 'O', 'O'},
			{'O', 'O', 'O'},
		},
	} {
		solve(*input)
	loop:
		for i, l := range *input {
			for j, b := range l {
				if output[i][j] != b {
					t.Errorf("Expected solve(%v) == %v, But got %v", input, output, input)
					break loop
				}
			}
		}
	}
}
