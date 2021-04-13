package Medium

func minSwap(A []int, B []int) int {
	dp := make([][]int, len(A))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
		dp[i][0] = -1
		dp[i][1] = -1
	}
	
	return dfs(A, B, 0, -1, -1, 0, &dp)
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	
	return b
}

func dfs(A, B []int, i, prevA, prevB, swap int, dp *[][]int) int {
	if i == len(A) {
		return 0
	}
	
	if (*dp)[i][swap] != -1 {
		return (*dp)[i][swap]
	}
	
	minSwaps := 100000
	if A[i] > prevA && B[i] > prevB {
		minSwaps = dfs(A, B, i+1, A[i], B[i], 0, dp)
	}
	
	if B[i] > prevA && A[i] > prevB {
		minSwaps = mini(minSwaps, dfs(A, B, i+1, B[i], A[i], 1, dp)+ 1)
	}
	
	(*dp)[i][swap] = minSwaps
	return minSwaps
}
