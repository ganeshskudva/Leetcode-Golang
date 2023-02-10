package Medium

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func maxDistance(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	var q [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
				vis[i][j] = true
			}
		}
	}

	if len(q) == 0 || len(q) == m*n {
		return -1
	}

	lvl := -1
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			c := q[0]
			q = q[1:]

			for _, d := range dirs {
				newX, newY := c[0]+d[0], c[1]+d[1]

				if isValid(grid, newX, newY) {
					if !vis[newX][newY] {
						q = append(q, []int{newX, newY})
						vis[newX][newY] = true
					}
				}
			}
		}
		lvl++
	}

	if lvl <= 0 {
		return -1
	}

	return lvl
}

func isValid(grid [][]int, x, y int) bool {
	m, n := len(grid), len(grid[0])

	return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 0
}
