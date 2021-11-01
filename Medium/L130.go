package Medium

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func isValid(board *[][]byte, x, y int) bool {
	m, n := len(*board), len((*board)[0])

	return x >= 0 && x < m && y >= 0 && y < n && (*board)[x][y] == 'O'
}

func bfs(board *[][]byte, i, j int) {
	var q [][]int
	q = append(q, []int{i, j})

	for len(q) > 0 {
		tmp := q[0]
		q = q[1:]

		x, y := tmp[0], tmp[1]
		(*board)[x][y] = '1'

		for _, d := range dirs {
			newX, newY := x+d[0], y+d[1]
			if isValid(board, newX, newY) {
				(*board)[newX][newY] = '1'
				q = append(q, []int{newX, newY})
			}
		}
	}
}

func solve(board [][]byte) {
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			bfs(&board, 0, i)
		}
	}

	for i := 0; i < len(board[0]); i++ {
		if board[len(board)-1][i] == 'O' {
			bfs(&board, len(board)-1, i)
		}
	}

	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			bfs(&board, i, 0)
		}
	}

	for i := 0; i < len(board); i++ {
		if board[i][len(board[0])-1] == 'O' {
			bfs(&board, i, len(board[0])-1)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}
