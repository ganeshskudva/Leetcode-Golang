package Medium

func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var (
		found   bool
		q       [][]int
		dfs     func(x, y int)
		isValid func(x, y int) bool
	)

	isValid = func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] && grid[x][y] != 0
	}

	dfs = func(x, y int) {
		if !isValid(x, y) {
			return
		}

		visited[x][y] = true
		q = append(q, []int{x, y})
		for _, d := range dirs {
			dfs(x+d[0], y+d[1])
		}
	}

	for i := 0; i < m; i++ {
		if found {
			break
		}
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				found = true
				break
			}
		}
	}

	step := 0
	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			cur := q[0]
			q = q[1:]
			size--

			for _, d := range dirs {
				x, y := cur[0]+d[0], cur[1]+d[1]
				if x >= 0 && y >= 0 && x < m && y < n && !visited[x][y] {
					if grid[x][y] == 1 {
						return step
					}
					q = append(q, []int{x, y})
					visited[x][y] = true
				}
			}
		}
		step++
	}

	return -1
}
