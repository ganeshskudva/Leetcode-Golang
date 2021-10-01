package Medium

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	memo := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			memo[i][j] = -1
		}
	}

	return lcs(text1, text2, m, n, &memo)
}

func lcs(s1, s2 string, m, n int, memo *[][]int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if (*memo)[m][n] != -1 {
		return (*memo)[m][n]
	}

	if s1[m-1] == s2[n-1] {
		(*memo)[m][n] = 1 + lcs(s1, s2, m-1, n-1, memo)
	} else {
		(*memo)[m][n] = max(lcs(s1, s2, m-1, n, memo), lcs(s1, s2, m, n-1, memo))
	}

	return (*memo)[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
