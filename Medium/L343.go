package Medium

func integerBreak(n int) int {
	mp := make(map[string]int)
	
	return solve(n, n - 1, mp)
}

func solve(n, idx int, mp map[string]int) int {
	if n == 0 || idx == 0 {
		return 1
	}
	
	key := fmt.Sprintf("%d-%d", n, idx)
	if v, ok := mp[key]; ok {
		return v
	}
	
	if idx > n {
		mp[key] = solve(n, idx - 1, mp)
	} else {
		mp[key] = max(idx * solve(n - idx, idx, mp), solve(n, idx - 1, mp))
	}
	
	return mp[key]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	
	return b
}
