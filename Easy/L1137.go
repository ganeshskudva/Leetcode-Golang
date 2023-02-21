package Easy

func tribonacci(n int) int {
	mp := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}

	var solve func(int) int
	solve = func(n int) int {
		if n < 0 {
			return 0
		}
		
		if v, ok := mp[n]; ok {
			return v
		}

		mp[n] = solve(n - 1) + solve(n - 2) + solve(n - 3)
		return mp[n]
	}
	
	return solve(n)
}
