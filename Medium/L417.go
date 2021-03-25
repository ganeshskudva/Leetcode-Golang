package Medium

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func isValid(x, y int, visited *[][]bool, matrix [][]int) bool {
	n, m := len(matrix), len(matrix[0])

	if x >= 0 && x < n && y >= 0 && y < m && !(*visited)[x][y] {
		return true
	}

	return false
}
func pacificAtlantic(matrix [][]int) [][]int {
	var res [][]int

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	n, m := len(matrix), len(matrix[0])
	pacific, atlantic := makeArr(n, m), makeArr(n, m)

	for i := 0; i < n; i++ {
		dfs(matrix, &pacific, 0, i, 0)
		dfs(matrix, &atlantic, 0, i, m-1)
	}

	for i := 0; i < m; i++ {
		dfs(matrix, &pacific, 0, 0, i)
		dfs(matrix, &atlantic, 0, n-1, i)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func dfs(matrix [][]int, visited *[][]bool, height, x, y int) {
	if !isValid(x, y, visited, matrix) {
		return
	}

	if matrix[x][y] < height {
		return
	}

	(*visited)[x][y] = true

	for _, d := range dirs {
		dfs(matrix, visited, matrix[x][y], x+d[0], y+d[1])
	}
}

func makeArr(n, m int) [][]bool {
	a := make([][]bool, n)

	for i := range a {
		a[i] = make([]bool, m)
	}

	return a
}
