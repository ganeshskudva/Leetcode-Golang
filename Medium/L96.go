package Medium

func numTrees(n int) int {
	mp := make(map[int]int)

	return solve(n, mp)
}

func solve(n int, mp map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}

	sum := 0
	for i := 1; i <= n; i++ {
		if _, ok := mp[i-1]; !ok {
			mp[i-1] = solve(i-1, mp)
		}
		if _, ok := mp[n-i]; !ok {
			mp[n-i] = solve(n-i, mp)
		}
		sum += mp[i-1] * mp[n-i]
	}
	return sum
}
