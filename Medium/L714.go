package Medium

func maxProfit(prices []int, fee int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	length := len(prices)
	buy := -prices[0]
	sell := 0

	for i := 1; i < length; i++ {
		tmp := buy
		buy = max(buy, sell-prices[i])
		sell = max(sell, tmp+prices[i]-fee)
	}

	return max(buy, sell)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
