package Hard

func maxDotProduct(nums1 []int, nums2 []int) int {
	dp := make(map[string]int)
	var solve func(n, m int) int

	solve = func(n, m int) int {
		key := fmt.Sprintf("%d#%d", n, m)
		if v, ok := dp[key]; ok {
			return v
		}

		p1 := nums1[n] * nums2[m]
		ans := p1

		if (n < len(nums1)-1) && (m < len(nums2)-1) {
			p1 += solve(n+1, m+1)
		}

		if n < len(nums1)-1 {
			ans = max(ans, solve(n+1, m))
		}

		if m < len(nums2)-1 {
			ans = max(ans, solve(n, m+1))
		}

		dp[key] = max(ans, p1)
		return dp[key]
	}

	return solve(0, 0)
}
