package Hard

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	mod := int64(1000000007)
	mp := make(map[string]int64)
	var solve func(n, i, minimumProfit int) int64

	solve = func(n, i, minimumProfit int) int64 {
		var (
			ways          int64
			zeroProfitKey = fmt.Sprintf("%d-%d-0", n, i)
			key           = fmt.Sprintf("%d-%d-%d", n, i, minimumProfit)
		)

		if i == len(group) {
			if minimumProfit <= 0 {
				return 1
			}

			return 0
		}

		if minimumProfit <= 0 {
			if v, ok := mp[zeroProfitKey]; ok {
				return v
			}
		} else if minimumProfit > 0 {
			if v, ok := mp[key]; ok {
				return v
			}
		}

		ways += solve(n, i+1, minimumProfit)
		ways %= mod

		if n >= group[i] {
			ways += solve(n-group[i], i+1, minimumProfit-profit[i])
			ways %= mod
		}

		if minimumProfit <= 0 {
			mp[zeroProfitKey] = ways % mod
		} else {
			mp[key] = ways % mod
		}

		return ways % mod
	}

	return int(solve(n, 0, minProfit))
}
