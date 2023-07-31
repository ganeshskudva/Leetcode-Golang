package Medium

func minimumDeleteSum(s1 string, s2 string) int {
	var (
		solve      func(i, j int) int
		deadEndSum func(str string, i int) int
	)
	mp, m, n := make(map[string]int), len(s1), len(s2)

	deadEndSum = func(str string, i int) int {
		sum := 0
		for ; i < len(str); i++ {
			sum += int(str[i])
		}

		return sum
	}

	solve = func(i, j int) int {
		sum := 0
		if i == m && j == n {
			return 0
		}
		if i == m {
			return deadEndSum(s2, j)
		} else if j == n {
			return deadEndSum(s1, i)
		}

		key := fmt.Sprintf("%d-%d", i, j)
		if v, ok := mp[key]; ok {
			return v
		}

		if s1[i] == s2[j] {
			sum = solve(i+1, j+1)
		} else {
			sum = min(solve(i+1, j)+int(s1[i]), solve(i, j+1)+int(s2[j]))
		}

		mp[key] = sum
		return mp[key]
	}

	return solve(0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
