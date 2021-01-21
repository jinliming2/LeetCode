package main

// Runtime: 20 ms, faster than 96.21% of Go online submissions for Surrounded Regions.
// Memory Usage: 6 MB, less than 23.48% of Go online submissions for Surrounded Regions.
// https://leetcode.com/submissions/detail/445933119/
func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	y, x := len(board), len(board[0])
	for i := range board {
		convert(board, i, 0, y-1, x-1)
		convert(board, i, x-1, y-1, x-1)
	}
	for j := range board[0] {
		convert(board, 0, j, y-1, x-1)
		convert(board, y-1, j, y-1, x-1)
	}
	for i, line := range board {
		for j, b := range line {
			switch b {
			case 'O':
				board[i][j] = 'X'
			case '#':
				board[i][j] = 'O'
			}
		}
	}
}

func convert(board [][]byte, i, j, y, x int) {
	if i < 0 || i > y || j < 0 || j > x || board[i][j] != 'O' {
		return
	}
	board[i][j] = '#'
	convert(board, i-1, j, y, x)
	convert(board, i+1, j, y, x)
	convert(board, i, j-1, y, x)
	convert(board, i, j+1, y, x)
}
