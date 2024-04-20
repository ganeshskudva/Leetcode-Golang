package Medium

func findFarmland(land [][]int) [][]int {
	var (
		lst   [][]int
		solve func(x, y int)
	)
	row, col := len(land), len(land[0])

	solve = func(x, y int) {
		endX, endY := x, y

		for endY+1 < len(land[0]) && land[x][endY+1] == 1 {
			land[x][endY+1] = 0
			endY++
		}

		for endX+1 < len(land) && land[endX+1][y] == 1 {
			land[endX+1][y] = 0
			endX++
		}

		for i := x; i <= endX; i++ {
			for j := y; j <= endY; j++ {
				land[i][j] = 0
			}
		}

		lst = append(lst, []int{x, y, endX, endY})
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if land[i][j] == 1 {
				solve(i, j)
			}
		}
	}

	return lst
}
