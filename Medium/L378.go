package Medium

func kthSmallest(matrix [][]int, k int) int {
	lo, hi := matrix[0][0], matrix[len(matrix)-1][len(matrix)-1]
	for lo <= hi {
		mid := lo + (hi-lo)/2
		cnt := getLessEqual(matrix, mid)
		if cnt < k {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return lo
}

func getLessEqual(matrix [][]int, val int) int {
	res, n, i, j := 0, len(matrix), len(matrix)-1, 0

	for i >= 0 && j < n {
		if matrix[i][j] > val {
			i--
		} else {
			res += i + 1
			j++
		}
	}

	return res
}
