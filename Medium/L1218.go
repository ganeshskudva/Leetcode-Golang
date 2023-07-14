package Medium

func longestSubsequence(arr []int, difference int) int {
	mp, ans := make(map[int]int), 0

	for i := range arr {
		if v, ok := mp[arr[i]-difference]; ok {
			ans = max(ans, v+1)
			mp[arr[i]] = v + 1
		} else {
			ans = max(ans, 1)
			mp[arr[i]] = 1
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
