package Hard

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func isValid(grid [][]byte, x, y int) bool {
	n, m := len(grid), len(grid[0])

	return x >= 0 && x < n && y >= 0 && y < m
}

func containsCycle(grid [][]byte) bool {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i, _ := range visited {
		visited[i] = make([]bool, m)
	}
	
	for i:= 0; i < n; i++ {
		for j:= 0; j < m; j++ {
			if !visited[i][j] && dfs(grid, &visited, i, j, -1, -1) {
				return true
			}
		}
	}

	return false
}

func dfs(grid [][]byte, visited *[][]bool, i, j, x, y int) bool {
	if !isValid(grid, i, j) {
		return false
	}

	if (*visited)[i][j] {
		return true
	}

	(*visited)[i][j] = true

	found := false
	for _, d := range dirs {
		newX, newY := i+d[0], j+d[1]
		if isValid(grid, newX, newY) {
			if grid[newX][newY] == grid[i][j] && !(x == newX && y == newY) {
				found = found || dfs(grid, visited, newX, newY, i, j)
				if found {
					return found
				}
			}
		}
	}

	return found
}
