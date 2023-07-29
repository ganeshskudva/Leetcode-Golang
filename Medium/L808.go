package Medium

func soupServings(n int) float64 {
	if n >= 5000 {
		return 1.0
	}

	var solve func(a, b int) float64
	mp := make(map[string]float64)

	solve = func(a, b int) float64 {
		if a <= 0 && b > 0 {
			return 1.0
		}

		if a <= 0 && b <= 0 {
			return 0.5
		}

		if a > 0 && b <= 0 {
			return 0.0
		}

		key := fmt.Sprintf("%d-%d", a, b)
		if v, ok := mp[key]; ok {
			return v
		}

		mp[key] = 0.25 * (solve(a-100, b) + solve(a-75, b-25) + solve(a-50, b-50) + solve(a-25, b-75))
		return mp[key]
	}

	return solve(n, n)
}
