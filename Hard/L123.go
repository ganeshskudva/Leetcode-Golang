package Hard

func maxProfit(prices []int) int {
	return solve(prices, 0, 2, 0, make(map[string]int))
}

func solve(prices []int, i, k, buyOrSell int, dp map[string]int) int {
	if i >= len(prices) || k == 0 {
		return 0
	}
	key := fmt.Sprintf("%d#%d#%d", i, buyOrSell, k)
	if val, ok := dp[key]; ok {
		return val
	}
	res := 0
	if buyOrSell == 0 {
		buy := solve(prices, i+1, k, 1, dp) - prices[i]
		noBuy := solve(prices, i+1, k, 0, dp)
		res += max(buy, noBuy)
	} else {
		sell := solve(prices, i+1, k-1, 0, dp) + prices[i]
		noSell := solve(prices, i+1, k, 1, dp)
		res += max(sell, noSell)
	}

	dp[key] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
