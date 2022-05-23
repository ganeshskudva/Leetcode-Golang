package Medium

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		one, zero := 0, 0
		for _, c := range s {
			if c == '1' {
				one++
			} else {
				zero++
			}
		}
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				if one <= j && zero <= i {
					dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
				}
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}


// Recursion + Memoization
func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0 {
		return 0
	}

	return solve(strs, make(map[string]int), m, n, 0)
}

func solve(strs []string, mp map[string]int, m, n, index int) int {
	// If we are done with the our remaining zeros and ones return 0 as we cant get any furtherr strings.
	if m == 0 && n == 0 {
		return 0
	}

	if index >= len(strs) {
		return 0
	}

	key := fmt.Sprintf("%d:%d:%d", m, n, index)
	if _, ok := mp[key]; ok {
		return mp[key]
	}

	// For the current index count the required number of zeros and ones .
	totalCnt, curr := 0, strs[index]
	ones, zeroes := 0, 0

	for i := range curr {
		if curr[i] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	takenStrCnt := 0
	// if we have the  required number of zeros and ones  we take the current string and update the
	// remaining zeros and ones and go to the next index
	if ones <= n && zeroes <= m {
		takenStrCnt += 1 + solve(strs, mp, m-zeroes, n-ones, index+1)
	}

	skippedStrNum := solve(strs, mp, m, n, index+1)
	totalCnt = skippedStrNum
	// For every position we also the option to skip the current string
	if takenStrCnt > skippedStrNum {
		totalCnt = takenStrCnt
	}

	mp[key] = totalCnt
	return totalCnt
}
