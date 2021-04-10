package Hard

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func isValid(matrix [][]int, i, j int) bool {
	m, n := len(matrix), len(matrix[0])

	return i >= 0 && i < m && j >= 0 && j < n
}
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	cache := make([][]int, m)
	for i, _ := range cache {
		cache[i] = make([]int, n)
	}
	max := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			max = maxi(max, dfs(matrix, i, j, &cache))
		}
	}
    
    return max
}

func dfs(matrix [][]int, i, j int, cache *[][]int) int {

	if (*cache)[i][j] != 0 {
		return (*cache)[i][j]
	}
	max := 1
	for _, d := range dirs {
		x, y := i+d[0], j+d[1]
		if isValid(matrix, x, y) && matrix[x][y] > matrix[i][j]{
			length := 1 + dfs(matrix, i+d[0], j+d[1], cache)
			max = maxi(max, length)
		}
	}
	(*cache)[i][j] = max

	return (*cache)[i][j]
}

func maxi(a, b int) int {
	if a > b {
		return a
	}

	return b
}
