package main

import (
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	for input, output := range map[int][][]string{
		4: {
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		},
		1: {
			{"Q"},
		},
		5: {
			{"Q....", "..Q..", "....Q", ".Q...", "...Q."},
			{"Q....", "...Q.", ".Q...", "....Q", "..Q.."},
			{".Q...", "...Q.", "Q....", "..Q..", "....Q"},
			{".Q...", "....Q", "..Q..", "Q....", "...Q."},
			{"..Q..", "Q....", "...Q.", ".Q...", "....Q"},
			{"..Q..", "....Q", ".Q...", "...Q.", "Q...."},
			{"...Q.", "Q....", "..Q..", "....Q", ".Q..."},
			{"...Q.", ".Q...", "....Q", "..Q..", "Q...."},
			{"....Q", ".Q...", "...Q.", "Q....", "..Q.."},
			{"....Q", "..Q..", "Q....", "...Q.", ".Q..."},
		},
		6: {
			{".Q....", "...Q..", ".....Q", "Q.....", "..Q...", "....Q."},
			{"..Q...", ".....Q", ".Q....", "....Q.", "Q.....", "...Q.."},
			{"...Q..", "Q.....", "....Q.", ".Q....", ".....Q", "..Q..."},
			{"....Q.", "..Q...", "Q.....", ".....Q", "...Q..", ".Q...."},
		},
	} {
		value := solveNQueens(input)
		if len(output) != len(value) {
			t.Errorf("Expected len(solveNQueens(%v)) == %d , But got %d", input, len(output), len(value))
			continue
		}
	itemLoop:
		for _, item := range output {
		answerLoop:
			for _, answer := range value {
				for index, itemV := range item {
					if answer[index] != itemV {
						continue answerLoop
					}
				}
				continue itemLoop
			}
			t.Errorf("Expected solveNQueens(%v) == %v, But got %v", input, output, value)
			break itemLoop
		}
	}
}
