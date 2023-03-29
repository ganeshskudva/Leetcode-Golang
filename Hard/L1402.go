package Hard

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	mp := make(map[string]int)
	var solve func(idx, time int) int

	solve = func(idx, time int) int {
		if idx == len(satisfaction) {
			return 0
		}

		key := fmt.Sprintf("%d-%d", idx, time)
		if v, ok := mp[key]; ok {
			return v
		}

		includeDish := (satisfaction[idx] * (time + 1)) + solve(idx+1, time+1)
		excludeDish := solve(idx+1, time)

		mp[key] = max(includeDish, excludeDish)
		return mp[key]
	}

	return solve(0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	
	return b
}
