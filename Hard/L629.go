package Hard

func kInversePairs(n int, k int) int {
	memo := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make([]int, k+1)
		for j := 0; j < k+1; j++ {
			memo[i][j] = -1
		}
	}

	return solve(&memo, 1, n, k)
}

func solve(memo *[][]int, num, n, k int) int {
	mod := 1000000007
	if num == n+1 {
		if k == 0 {
			return 1
		}

		return 0
	}

	if (*memo)[num][k] != -1 {
		return (*memo)[num][k]
	}

	cnt, i := 0, 0
	if k-num+1 > 0 {
		i = k - num + 1
	}
	for ; i <= k; i++ {
		val := solve(memo, num+1, n, i)
		cnt = (cnt + val) % mod
	}

	(*memo)[num][k] = cnt
	return (*memo)[num][k]
}
