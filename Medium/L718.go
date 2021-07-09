package Medium

func findLength(nums1 []int, nums2 []int) int {
	res, m, n := 0, len(nums1), len(nums2)

	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	solve(nums1, nums2, m, n, &memo, &res)
	return res
}

func solve(n1, n2 []int, m, n int, memo *[][]int, res *int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if (*memo)[m][n] != -1 {
		return (*memo)[m][n]
	}

	solve(n1, n2, m, n-1, memo, res)
	solve(n1, n2, m-1, n, memo, res)

	if n1[m-1] == n2[n-1] {
		(*memo)[m][n] = 1 + solve(n1, n2, m-1, n-1, memo, res)
		if (*memo)[m][n] > *res {
			*res = (*memo)[m][n]
		}
		return (*memo)[m][n]
	}
	(*memo)[m][n] = 0

	return (*memo)[m][n]
}
