package Medium

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][1]-costs[i][0])-(costs[j][1]-costs[j][0]) > 0
	})

	res := 0
	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			res += costs[i][0]
		} else {
			res += costs[i][1]
		}
	}

	return res
}
