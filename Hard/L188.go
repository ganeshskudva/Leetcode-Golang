package Hard

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit(k int, prices []int) int {
	mp := make(map[string]int)

	return solve(prices, mp, 0, k, true)
}

func solve(prices []int, mp map[string]int, idx, k int, buy bool) int {
	if idx >= len(prices) || k <= 0 {
		return 0
	}

	key := fmt.Sprintf("%d-%d-%d", idx, buy, k)
	if v, ok := mp[key]; ok {
		return v
	}

	if buy {
		mp[key] = max(-prices[idx]+solve(prices, mp, idx+1, k, !buy), solve(prices, mp, idx+1, k, buy))
	} else {
		mp[key] = max(prices[idx]+solve(prices, mp, idx+1, k-1, !buy), solve(prices, mp, idx+1, k, buy))
	}

	return mp[key]
}
