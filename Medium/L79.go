package Medium

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	found := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			found = false
			if board[i][j] == word[0] {
				found = bt(board, &vis, i, j, 0, word)
				if found {
					return found
				}
			}
		}
	}
	return found
}

func isValid(board [][]byte, x, y int, ch byte) bool {
	m, n := len(board), len(board[0])

	return x >= 0 && x < m && y >= 0 && y < n && board[x][y] == ch
}

func bt(board [][]byte, vis *[][]bool, i, j, index int, word string) bool {
	if index == len(word) {
		return true
	}

	if !isValid(board, i, j, word[index]) {
		return false
	}

	if (*vis)[i][j] {
		return false
	}
	(*vis)[i][j] = true
	found := false
	for _, d := range dirs {
		found = found || bt(board, vis, i+d[0], j+d[1], index+1, word)
	}
	(*vis)[i][j] = false

	return found
}
