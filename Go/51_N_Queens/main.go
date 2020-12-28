package main

var results [][]string

// Runtime: 4 ms, faster than 84.47% of Go online submissions for N-Queens.
// Memory Usage: 3.2 MB, less than 100.00% of Go online submissions for N-Queens.
func solveNQueens(n int) [][]string {
	results = make([][]string, 0)
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	sub(board, 0, n)
	return results
}

func sub(board [][]byte, index, n int) {
	if index == n {
		result := make([]string, 0, n)
		for _, line := range board {
			result = append(result, string(line))
		}
		results = append(results, result)
		return
	}

	for i := 0; i < n; i++ {
		if !check(board, index, i, n) {
			continue
		}
		board[index][i] = 'Q'
		sub(board, index+1, n)
		board[index][i] = '.'
	}
}

func check(board [][]byte, index, i, n int) bool {
	for x, y, z := i-1, index-1, i+1; y >= 0; x, y, z = x-1, y-1, z+1 {
		// Top
		if board[y][i] == 'Q' {
			return false
		}
		// Top Left
		if x >= 0 && board[y][x] == 'Q' {
			return false
		}
		// Top Right
		if z < n && board[y][z] == 'Q' {
			return false
		}
	}
	return true
}
