package Medium

func maxScore(cardPoints []int, k int) int {
	dp := make([]int, k+1)

	for i := len(cardPoints) - 1; i >= len(cardPoints)-k; i-- {
		dp[0] += cardPoints[i]
	}

	maxPoints := dp[0]
	for i := 1; i < k+1; i++ {
		dp[i] = dp[i-1] + cardPoints[i-1] - cardPoints[len(cardPoints)-k+i-1]
		maxPoints = max(maxPoints, dp[i])
	}

	return maxPoints
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
