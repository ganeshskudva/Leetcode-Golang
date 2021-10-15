package Medium

func maxProfit(prices []int) int {
	return solve(prices, 0, make(map[int]int))
}

func solve(prices []int, current int, mp map[int]int) int {
	if current > len(prices) {
		return 0
	}

	if val, ok := mp[current]; ok {
		return val
	}

	max := 0
	for i := current + 1; i < len(prices); i++ {
		sell := prices[i] - prices[current] + solve(prices, i+2, mp)
		if sell > max {
			max = sell
		}
	}

	profit := solve(prices, current+1, mp)
	if profit > max {
		max = profit
	}
	mp[current] = max
	return max
}
