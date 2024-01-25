package Medium

func longestCommonSubsequence(text1 string, text2 string) int {
	var (
		solve func(m, n int) int
	)
	m, n := len(text1), len(text2)
	mp := make(map[string]int)

	solve = func(m, n int) int {
		if m == 0 || n == 0 {
			return 0
		}

		key := fmt.Sprintf("%d#%d", m, n)
		if v, ok := mp[key]; ok {
			return v
		}

		if text1[m-1] == text2[n-1] {
			mp[key] = 1 + solve(m-1, n-1)
		} else {
			mp[key] = max(solve(m-1, n), solve(m, n-1))
		}

		return mp[key]
	}

	return solve(m, n)
}
