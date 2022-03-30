package Medium

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][n-1] == target {
			return true
		}

		if matrix[i][n-1] > target {
			return binSearch(matrix[i], target)
		}
	}

	return false
}

func binSearch(n []int, tgt int) bool {
	st, end := 0, len(n)-1

	for st <= end {
		mid := st + (end-st)/2
		if n[mid] == tgt {
			return true
		}

		if n[mid] < tgt {
			st = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}
