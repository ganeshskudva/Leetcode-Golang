package Medium

func updateMatrix(mat [][]int) [][]int {
	var (
		res, q  [][]int
		isValid func(x, y int) bool
	)
	if len(mat) == 0 {
		return res
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m, n := len(mat), len(mat[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	isValid = func(x, y int) bool {
		return x >= 0 && y >= 0 && x < m && y < n && mat[x][y] == math.MaxInt32
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
			} else {
				mat[i][j] = math.MaxInt32
			}
		}
	}

	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		if vis[c[0]][c[1]] {
			continue
		}
		vis[c[0]][c[1]] = true

		for _, d := range dirs {
			newX, newY := c[0]+d[0], c[1]+d[1]
			if isValid(newX, newY) {
				mat[newX][newY] = min(mat[newX][newY], mat[c[0]][c[1]]+1)
				q = append(q, []int{newX, newY})
			}
		}
	}

	return mat
}


func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
