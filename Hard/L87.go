package Hard

func isScramble(s1 string, s2 string) bool {
	var solve func(a, b string) bool
	mp := make(map[string]bool)

	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}

	solve = func(a, b string) bool {
		if a == b {
			return true
		}
		if len(a) <= 1 {
			return false
		}

		key := fmt.Sprintf("%s-%s", a, b)
		if v, ok := mp[key]; ok {
			return v
		}

		n, check := len(a), false
		for i := 1; i < n; i++ {
			swap := solve(a[:i], b[n-i:]) && solve(a[i:], b[0:n-i])
			unSwap := solve(a[0:i], b[0:i]) && solve(a[i:], b[i:])

			if swap || unSwap {
				check = true
				break
			}
		}
		mp[key] = check
		return mp[key]
	}
	
	return solve(s1, s2)
}
