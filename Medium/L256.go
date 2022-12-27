package Medium

func minCost(costs [][]int) int {
	mp := make(map[string]int)

	mini := math.MaxInt32
	for i := 0; i < len(costs[0]); i++ {
		mini = min(mini, solve(costs, mp, 0, i))
	}

	return mini
}

func solve(costs [][]int, mp map[string]int, house, color int) int {
	if house >= len(costs) {
		return 0
	}

	if color >= len(costs[0]) {
		return 0
	}

	key := fmt.Sprintf("%d-%d", house, color)
	if v, ok := mp[key]; ok {
		return v
	}

	mini := math.MaxInt32
	for i := 0; i < len(costs[0]); i++ {
		if i == color {
			continue
		}
		mini = min(mini, costs[house][color]+solve(costs, mp, house+1, i))
	}

	mp[key] = mini
	return mp[key]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
