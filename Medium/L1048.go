package Medium

func longestStrChain(words []string) int {
	mp, dp := make(map[string]int), make(map[string]int)
	for _, s := range words {
		mp[s] = 1
	}

	cnt := math.MinInt32
	for _, s := range words {
		res := 1 + dfs(s, mp, dp)
		if cnt < res {
			cnt = res
		}
	}

	return cnt
}

func dfs(s string, mp, dp map[string]int) int {
	if len(s) == 0 {
		return 0
	}

	if _, ok := dp[s]; ok {
		return dp[s]
	}
	max := 0
	for i := range s {
		tmp, res := deleteCharAt(s, i), 0
		if _, ok := mp[tmp]; ok {
			res += 1 + dfs(tmp, mp, dp)
		}
		if res > max {
			max = res
		}
	}
	dp[s] = max
	return dp[s]
}

func deleteCharAt(s string, idx int) string {
	if idx == 0 {
		return s[1:]
	}

	if idx == len(s)-1 {
		return s[:len(s)-1]
	}

	return s[:idx] + s[idx+1:]
}
