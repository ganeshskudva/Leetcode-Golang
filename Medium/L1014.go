package Medium

func maxScoreSightseeingPair(values []int) int {
	mp := make(map[string]int)

	return solve(0, 0, values, mp)
}

func solve(i, taken int, values []int, mp map[string]int) int {
	if taken >= 2 {
		return 0
	}

	if i >= len(values) {
		return math.MinInt32
	}

	key := fmt.Sprintf("%d-%d", i, taken)
	if v, ok := mp[key]; ok {
		return v
	}

	pick, notPick := values[i]+solve(i+1, taken+1, values, mp), solve(i+1, taken, values, mp)

	if taken == 1 {
		pick -= i
	} else {
		pick += i
	}

	mp[key] = max(pick, notPick)
	return mp[key]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
