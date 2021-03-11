package Medium

const maxInt = int(^uint(0) >> 1) - 1

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = maxInt
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == maxInt {
		return -1
	}
	return dp[amount]
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
