func longestPalindromeSubseq(s string) int {
	mp := make(map[string]int)
	var solve func(left, right int) int

	solve = func(left, right int) int {
		if left == right {
			return 1
		}

		if left > right {
			return 0
		}

		key := fmt.Sprintf("%d-%d", left, right)
		if v, ok := mp[key]; ok {
			return v
		}

		if s[left] == s[right] {
			mp[key] = 2 + solve(left+1, right-1)
		} else {
			mp[key] = max(solve(left+1, right), solve(left, right-1))
		}

		return mp[key]
	}
	
	return solve(0, len(s) - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
