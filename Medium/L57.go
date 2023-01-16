package Medium

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	start, end := newInterval[0], newInterval[1]

	for _, itvl := range intervals {
		currStart, currEnd := itvl[0], itvl[1]
		if currEnd < start {
			res = append(res, []int{currStart, currEnd})
		} else if currStart > end {
			res = append(res, []int{start, end})
			start, end = currStart, currEnd
		} else {
			start = min(start, currStart)
			end = max(end, currEnd)
		}
	}
	res = append(res, []int{start, end})

	return res
}
