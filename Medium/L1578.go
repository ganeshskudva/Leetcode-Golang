package Medium

func minCost(colors string, neededTime []int) int {
	n, res, prev := len(colors), 0, neededTime[0]

	for i := 1; i < n; i++ {
		if colors[i] == colors[i-1] {
			res += min(neededTime[i], prev)
			prev = max(neededTime[i], prev)
		} else {
			prev = neededTime[i]
		}
	}

	return res
}

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
