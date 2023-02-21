package Easy

func climbStairs(n int) int {
	mp := make(map[int]int)
	
	mp[1], mp[2], mp[3] = 1,2,3
	return dp(mp, n)
}

func dp(mp map[int]int, n int) int {
	if _, ok := mp[n]; ok {
		return mp[n]
	}
	
	mp[n] = dp(mp, n-1) + dp(mp, n-2)
	
	return mp[n]
}

// With anonymous function
func climbStairs(n int) int {
	mp := map[int]int{
		0: 0,
		1: 1,
		2: 2,
	}

	var solve func(int) int
	solve = func(n int) int {
		if n < 0 {
			return 0
		}

		if v, ok := mp[n]; ok {
			return v
		}

		mp[n] = solve(n-1) + solve(n-2)
		return mp[n]
	}

	return solve(n)
}
