package t51

// N皇后问题
// 同行同列对角线上不能有皇后

// 回溯
// 注意逐层判断，同时判断只要向上判断
func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i] = append(board[i], '.')
		}
	}
	var helper func(board [][]rune, row int)
	helper = func(board [][]rune, row int) {
		if row == n {
			tmp := make([]string, 0)
			for i := 0; i < n; i++ {
				tmp = append(tmp, string(board[i]))
			}
			result = append(result, tmp)
		}
		for i := 0; i < n; i++ {
			if isValid(board, row, i) {
				board[row][i] = 'Q'
				helper(board, row+1)
				board[row][i] = '.'
			}
		}
	}
	helper(board, 0)
	return result
}

func isValid(board [][]rune, row, col int) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
