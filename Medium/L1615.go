package Medium

func maximalNetworkRank(n int, roads [][]int) int {
	connected, cnt := make([][]bool, n), make([]int, n)
	for i := 0; i < n; i++ {
		connected[i] = make([]bool, n)
	}

	for _, r := range roads {
		cnt[r[0]]++
		cnt[r[1]]++
		connected[r[0]][r[1]], connected[r[1]][r[0]] = true, true
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if connected[i][j] {
				res = max(res, cnt[i]+cnt[j]-1)
			} else {
				res = max(res, cnt[i]+cnt[j])
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
