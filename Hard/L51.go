package Hard

func solveNQueens(n int) [][]string {
	chess := make([][]byte, n)
	for i := range chess {
		chess[i] = make([]byte, n)
		for j := range chess[i] {
			chess[i][j] = '.'
		}
	}
	var res [][]string
	solve(&res, &chess, 0)

	return res
}

func solve(res *[][]string, chess *[][]byte, row int) {
	if row == len(*chess) {
		*res = append(*res, construct(chess))
		return
	}
	
	for col := 0; col < len(*chess); col++ {
		if isValid(chess, row, col) {
			(*chess)[row][col] = 'Q'
			solve(res, chess, row+1)
			(*chess)[row][col] = '.'
		}
	}
}

func isValid(chess *[][]byte, row, col int) bool {
	for i := 0; i < row; i++ {
		if (*chess)[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < len(*chess); i, j = i-1, j+1 {
		if (*chess)[i][j] == 'Q' {
			return false
		}
	}
	
	for i,j := row-1, col-1; i>=0  && j >=0; i,j = i-1, j-1 {
		if (*chess)[i][j] == 'Q' {
			return false
		}
	}
	
	return true
}

func construct (chess *[][]byte) []string {
	var path []string
	for i:= 0; i < len(*chess); i++ {
		path = append(path, string((*chess)[i]))
	}
	
	return path
}
