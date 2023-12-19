package Easy

func imageSmoother(img [][]int) [][]int {
	if len(img) == 0 {
		return [][]int{}
	}

	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	rows, cols := len(img), len(img[0])
	var isValid func(x, y int) bool

	isValid = func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	res := make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			cnt, sum := 0, 0
			for _, d := range dirs {
				newRow, newCol := row+d[0], col+d[1]
				if isValid(newRow, newCol) {
					cnt++
					sum += img[newRow][newCol]
				}
			}

			res[row][col] = sum / cnt
		}
	}

	return res
}
