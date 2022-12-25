package Hard

func countDigitOne(n int) int {
	str := fmt.Sprintf("%d", n)
	dp := make([][][]int, 10)

	for i := 0; i < 10; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, 10)
			for k := 0; k < 10; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	return solve(str, 0, 1, 0, &dp)
}

func solve(s string, idx, tight, count int, dp *[][][]int) int {
	if idx == len(s) {
		return count
	}

	if (*dp)[idx][tight][count] != -1 {
		return (*dp)[idx][tight][count]
	}

	ans, bound := 0, 9
	if tight == 1 {
		bound = int(s[idx] - '0')
	}

	for i := 0; i <= bound; i++ {
		add := 0
		if i == 1 {
			add = 1
		}

		if i == bound && tight == 1 {
			ans += solve(s, idx+1, tight, count+add, dp)
		} else {
			ans += solve(s, idx+1, 0, count+add, dp)
		}
	}

	(*dp)[idx][tight][count] = ans
	return ans
}
