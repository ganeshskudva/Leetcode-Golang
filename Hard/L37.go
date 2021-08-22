package Hard

func solveSudoku(board [][]byte) {
	solve(board, 0, 0)
}

func solve(board [][]byte, row, col int) bool {
	for i := row; i < 9; i, col = i+1, 0 {
		for j := col; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			for num := '1'; num <= '9'; num++ {
				if isValid(board, i, j, byte(num)) {
					board[i][j] = byte(num)
					if solve(board, i, j+1) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}

	return true
}

func isValid(board [][]byte, row, col int, num byte) bool {
	blkrow, blkcol := (row/3)*3, (col/3)*3
	for i := 0; i < 9; i++ {
		if board[i][col] == num || board[row][i] == num ||
			board[blkrow+i/3][blkcol+i%3] == num {
			return false
		}
	}

	return true
}
