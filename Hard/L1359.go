package Hard

func countOrders(n int) int {
	dp := make([][]int64, 501)
	for i := 0; i < 501; i++ {
		dp[i] = make([]int64, 501)
		for j := 0; j < 501; j++ {
			dp[i][j] = -1
		}
	}

	return int(solve(n, n, &dp))
}

func solve(pick, del int, dp *[][]int64) int64 {
	if del < 0 || pick < 0 {
		return 0
	}

	if del < pick {
		return 0
	}

	if pick == 0 && del == 0 {
		return 1
	}

	if (*dp)[pick][del] != -1 {
		return (*dp)[pick][del]
	}

	var ans int64
	mod := 1000000007

	ans += int64(pick) * solve(pick-1, del, dp)
	ans %= int64(mod)
	
	ans += int64(del-pick) * solve(pick, del-1, dp)
	ans %= int64(mod)

	(*dp)[pick][del] = ans
	
	return (*dp)[pick][del]
}
