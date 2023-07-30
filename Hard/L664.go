package Hard

func strangePrinter(s string) int {
	var (
		solve func(left, right int) int
	)
	mp := make(map[string]int)

	solve = func(left, right int) int {
		if left == right {
			return 1
		}
		key := fmt.Sprintf("%d-%d", left, right)
		if v, ok := mp[key]; ok {
			return v
		}

		res := math.MaxInt32
		for k := left; k < right; k++ {
			res = min(res, solve(left, k)+solve(k+1, right))
		}

		if s[left] == s[right] {
			res--
		}
		mp[key] = res
		return mp[key]
	}

	return solve(0, len(s)-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
