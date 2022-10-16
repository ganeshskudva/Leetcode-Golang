package Hard

func minDifficulty(jobDifficulty []int, d int) int {
	if d > len(jobDifficulty) {
		return -1
	}

	m, n := d-1, len(jobDifficulty)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	return solve(jobDifficulty, d-1, 0, &dp)
}

func solve(jobDifficulty []int, d, pos int, dp *[][]int) int {
	if d == 0 {
		max := jobDifficulty[pos]
		for i := pos; i < len(jobDifficulty); i++ {
			if jobDifficulty[i] > max {
				max = jobDifficulty[i]
			}
		}

		return max
	}

	day := len(*dp) - d
	if (*dp)[day][pos] != -1 {
		return (*dp)[day][pos]
	}

	max, min := math.MinInt32, math.MaxInt32
	for i := pos; i < len(jobDifficulty)-d; i++ {
		if jobDifficulty[i] > max {
			max = jobDifficulty[i]
		}
		compute := solve(jobDifficulty, d-1, i+1, dp)
		if max+compute < min {
			min = max + compute
		}
	}

	(*dp)[day][pos] = min
	return min
}
