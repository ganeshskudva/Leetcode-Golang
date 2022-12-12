package Easy

func maxProfit(prices []int) int {
	mp := make(map[string]int)

	return solve(prices, 0, 1, true, mp)
}

func solve(prices []int, i, k int, buy bool, mp map[string]int) int {
	if i >= len(prices) || k <= 0 {
		return 0
	}

	key := fmt.Sprintf("%d-%t", i, buy)
	if v, ok := mp[key]; ok {
		return v
	}

	if buy {
		mp[key] = max(-prices[i]+solve(prices, i+1, k, !buy, mp), solve(prices, i+1, k, buy, mp))
		return mp[key]
	}

	mp[key] = max(prices[i]+solve(prices, i+1, k-1, !buy, mp), solve(prices, i+1, k, buy, mp))
	return mp[key]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
