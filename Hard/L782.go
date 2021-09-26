package Hard

func movesToChessboard(board [][]int) int {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (board[0][0] ^ board[i][0] ^ board[0][j] ^ board[i][j]) == 1 {
				return -1
			}
		}
	}

	rowSum, colSum, rowMisplaced, colMisplaced := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		rowSum += board[0][i]
		colSum += board[i][0]
		if board[i][0] == i%2 {
			rowMisplaced++
		}
		if board[0][i] == i%2 {
			colMisplaced++
		}
	}

	if rowSum != n/2 && rowSum != (n+1)/2 {
		return -1
	}
	if colSum != n/2 && colSum != (n+1)/2 {
		return -1
	}

	if n%2 == 1 {
		if colMisplaced%2 == 1 {
			colMisplaced = n - colMisplaced
		}
		if rowMisplaced%2 == 1 {
			rowMisplaced = n - rowMisplaced
		}
	} else {
		colMisplaced = min(n-colMisplaced, colMisplaced)
		rowMisplaced = min(n-rowMisplaced, rowMisplaced)
	}

	return (colMisplaced + rowMisplaced) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
