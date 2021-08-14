package Hard

func removeBoxes(boxes []int) int {
	n := len(boxes)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
			for k := 0; k < n; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	return solve(boxes, &dp, 0, n-1, 0)
}

func solve(boxes []int, dp *[][][]int, i, j, k int) int {
	if i > j {
		return 0
	}

	if (*dp)[i][j][k] > 0 {
		return (*dp)[i][j][k]
	}

	i0, k0 := i, k
	for ; i+1 <= j && boxes[i+1] == boxes[i]; i, k = i+1, k+1 {

	}

	res := (k+1)*(k+1) + solve(boxes, dp, i+1, j, 0)
	for m := i + 1; m <= j; m++ {
		if boxes[i] == boxes[m] {
			val := solve(boxes, dp, i+1, m-1, 0) + solve(boxes, dp, m, j, k+1)
			if val > res {
				res = val
			}
		}
	}

	(*dp)[i0][j][k0] = res
	return res
}
