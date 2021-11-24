package Medium

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	if len(firstList) == 0 || len(secondList) == 0 {
		return res
	}

	i, j, m, n := 0, 0, len(firstList), len(secondList)
	max, min := math.MinInt64, math.MaxInt64

	for i < m && j < n {
		max = maximum(firstList[i][0], secondList[j][0])
		min = minimum(firstList[i][1], secondList[j][1])
		if max <= min {
			res = append(res, []int{max, min})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return res
}

func maximum(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minimum(a, b int) int {
	if a < b {
		return a
	}

	return b
}
