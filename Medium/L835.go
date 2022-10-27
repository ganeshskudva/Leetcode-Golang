package Medium

func largestOverlap(img1 [][]int, img2 [][]int) int {
	ans := 0

	for row := -len(img1); row < len(img1); row++ {
		for col := -len(img1[0]); col < len(img1[0]); col++ {
			max := overlap(img1, img2, row, col)
			if max > ans {
				ans = max
			}
		}
	}

	return ans
}

func overlap(img1, img2 [][]int, rowOffset, colOffset int) int {
	ans := 0
	for row := 0; row < len(img1); row++ {
		for col := 0; col < len(img1[0]); col++ {
			if isValid(row, col, rowOffset, colOffset, img1) {
				if img1[row][col] == 1 && img2[row+rowOffset][col+colOffset] == 1 {
					ans++
				}
			}
		}
	}

	return ans
}

func isValid(row, col, rowOffset, colOffset int, img [][]int) bool {
	return row+rowOffset >= 0 &&
		row+rowOffset < len(img) &&
		col+colOffset >= 0 &&
		col+colOffset < len(img)
}
Console
