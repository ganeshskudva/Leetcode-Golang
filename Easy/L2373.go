package Easy

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	res := make([][]int, n-2)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n-2)
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			tmp := 0

			for k := i - 1; k <= i + 1; k++ {
				for l := j - 1; l <= j + 1; l++ {
					tmp = max(tmp, grid[k][l])
				}
			}

			res[i-1][j-1] = tmp
		}
	}

	return res
}
