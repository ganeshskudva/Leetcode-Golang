package Easy

func countNegatives(grid [][]int) int {
	res := 0

	for _, r := range grid {
		res += binarySearch(r)
	}

	return res
}

func binarySearch(row []int) int {
	l, r := 0, len(row)

	for l < r {
		m := l + (r-l)/2
		if row[m] < 0 {
			r = m
		} else {
			l = m + 1
		}
	}

	return len(row) - l
}
