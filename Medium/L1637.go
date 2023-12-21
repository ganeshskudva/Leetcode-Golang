package Medium

func maxWidthOfVerticalArea(points [][]int) int {
	xCoordMap := make(map[int]bool)
	maxDiff, prevX := 0, math.MinInt32

	for _, p := range points {
		xCoordMap[p[0]] = true
	}

	var keys []int
	for k, _ := range xCoordMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		if prevX != math.MinInt32 {
			maxDiff = max(maxDiff, k-prevX)
		}
		prevX = k
	}

	return maxDiff
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
