package Easy

func minCostClimbingStairs(cost []int) int {
	mp := make(map[int]int)
	
	var solve func(int) int 
	solve = func(n int) int {
		if n >= len(cost) {
			return 0
		}
		if v, ok := mp[n]; ok {
			return v
		}
		
		oneStep := cost[n] + solve(n + 1)
		twoStep := cost[n] + solve(n + 2)
		
		mp[n] = min(oneStep, twoStep)
		return mp[n]
	}	
	
	return min(solve(0), solve(1))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
