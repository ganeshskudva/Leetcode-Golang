package Medium

/*Recursion*/
func mini(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minSteps(n int) int {
	steps := 10000
	if n == 1 {
		return 0
	}
	dfs(n, 1, 1, 1, &steps)

	return steps
}

func dfs(n, stp, curr, prev int, steps *int) {
	if curr > n {
		return
	}

	if curr == n {
		*steps = mini(*steps, stp)
	}

	dfs(n, stp+1, curr+prev, prev, steps)
	dfs(n, stp+2, 2*curr, curr, steps)
}


/*DP*/
func mini(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minSteps(n int) int {
	dp := make([][]int, 1001)
	for i, _ := range dp {
		dp[i] = make([]int, 1001)
		for j := 0; j < 1001; j++ {
			dp[i][j] = -1
		}
	}
	if n == 1 {
		return 0
	}
	return dfs(n, 1, 1, 1, &dp)
}

func dfs(n, stp, curr, prev int, dp *[][]int) int {
	if stp > n || curr > n {
		return 10001
	}
	
	if curr == n {
		return stp
	}
	
	if (*dp)[stp][curr] != -1 {
		return (*dp)[stp][curr]
	}
	
	(*dp)[stp][curr] = mini(dfs(n, stp+1, curr+prev, prev, dp), dfs(n, stp+2, 2*curr, curr, dp))
	return (*dp)[stp][curr]
}
