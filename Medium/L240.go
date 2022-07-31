package Medium

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	found := false
	for _, m := range matrix {
		if target >= m[0] && target <= m[len(m)-1] {
			found = binSearch(m, 0, len(m)-1, target)
			if found {
				return found
			}
		}
	}

	return found
}

func binSearch(arr []int, low, high, tgt int) bool {
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == tgt {
			return true
		}

		if tgt < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
