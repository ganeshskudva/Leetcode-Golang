package Medium

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	var ans [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start, end := intervals[0][0], intervals[0][1]
	for _, i := range intervals {
		if i[0] <= end {
			end = max(end, i[1])
		} else {
			ans = append(ans, []int{start, end})
			start, end = i[0], i[1]
		}
	}
	ans = append(ans, []int{start, end})

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
