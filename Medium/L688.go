package  Medium

func knightProbability(n int, k int, row int, column int) float64 {
	var (
		solve func(K, r, c int) float64
		dir   = [][]int{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}}
		mp    = make(map[string]float64)
	)

	solve = func(K, r, c int) float64 {
		if r < 0 || r > n-1 || c < 0 || c > n-1 {
			return 0
		}
		if K == 0 {
			return 1
		}
		key := fmt.Sprintf("%d-%d-%d", r, c, K)
		if v, ok := mp[key]; ok {
			return v
		}
		var rate float64
		for _, d := range dir {
			rate += 0.125 * solve(K - 1, r + d[0], c + d[1])
		}
		
		mp[key] = rate
		return mp[key]
	}
	
	return solve(k, row, column)
}
