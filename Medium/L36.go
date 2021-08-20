package Medium

func isValidSudoku(board [][]byte) bool {
	if len(board) == 0 {
		return false
	}

	index := 0
	for row := 0; row < len(board); row++ {
		rowCheck, colCheck, boxCheck := make([]bool, 9), make([]bool, 9), make([]bool, 9)
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] != '.' {
				index = int((board[row][col] - '0') - 1)

				if rowCheck[index] {
					return false
				}

				rowCheck[index] = true
			}

			if board[col][row] != '.' {
				index = int((board[col][row] - '0') - 1)

				if colCheck[index] {
					return false
				}

				colCheck[index] = true
			}

			boxRow, boxCol := (row/3)*3+col/3, (row%3)*3+col%3
			if board[boxRow][boxCol] != '.' {
				index = int((board[boxRow][boxCol] - '0') - 1)

				if boxCheck[index] {
					return false
				}

				boxCheck[index] = true
			}
		}
	}

	return true
}
