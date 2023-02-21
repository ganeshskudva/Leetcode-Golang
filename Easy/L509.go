package Easy

func fib(n int) int {
	mp := make(map[int]int)
	
	mp[0], mp[1], mp[2] = 0, 1, 1
	
	return fibo(mp, n)
}

func fibo(mp map[int]int, n int) int {
	if _, ok := mp[n]; ok {
		return mp[n]
	}
	
	mp[n] = fibo(mp, n-1) + fibo(mp, n-2)
	
	return mp[n]
}

// With anonymous function
func fib(n int) int {
	mp := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}

	var solve func(n int) int
	solve = func(n int) int {
		if v, ok := mp[n]; ok {
			return v
		}

		mp[n] = solve(n-1) + solve(n-2)

		return mp[n]
	}

	return solve(n)
}
