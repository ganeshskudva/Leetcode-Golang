package Medium

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestContinuousSubstring(s string) int {
	res, cnt := 0, 1
	
	for i := 1; i < len(s); i++ {
		if s[i-1] + 1 == s[i] {
			cnt++
		} else {
			res = max(res, cnt)
			cnt = 1
		}
	}
	
	res = max(res, cnt)
	return res
}
