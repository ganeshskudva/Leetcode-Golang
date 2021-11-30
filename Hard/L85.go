package Hard

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	ans, m, n, height := 0, len(matrix), len(matrix[0]), make([]int, len(matrix[0]))

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
				continue
			}
			height[j]++
			for cur, pre := j-1, height[j]; cur >= 0; cur-- {
				if height[cur] == 0 {
					break
				}

				if height[cur] < pre {
					pre = height[cur]
				}

				if pre*(j-cur+1) > ans {
					ans = pre * (j - cur + 1)
				}
			}

			if height[j] > ans {
				ans = height[j]
			}
		}
	}

	return ans
}
