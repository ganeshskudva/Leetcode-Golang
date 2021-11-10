package Medium

func maxProfit(prices []int) int {
	n := len(prices)

	if n < 2 {
		return 0
	}

	memo := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			memo[i][j] = -1
		}
	}

	return solve(0, 1, prices, &memo)
}

func solve(index, buy int, prices []int, memo *[][]int) int {
	if index >= len(prices) {
		return 0
	}

	if (*memo)[index][buy] != -1 {
		return (*memo)[index][buy]
	}

	if buy == 1 {
		(*memo)[index][buy] = max(-prices[index]+solve(index+1, 0, prices, memo), solve(index+1, 1, prices, memo))
	} else {
		(*memo)[index][buy] = max(prices[index]+solve(index+1, 1, prices, memo), solve(index+1, 0, prices, memo))
	}

	return (*memo)[index][buy]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
