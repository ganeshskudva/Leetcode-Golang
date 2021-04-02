package Easy

func maxDepth(s string) int {
	maxi:= 0
	if len(s) == 0 {
		return maxi
	}
	
	cnt := 0
	for _, c := range s {
		if c == '(' {
			cnt++
		} else if c == ')' {
			cnt--
		}
		maxi = max(maxi, cnt)
	}
	
	return maxi
}

func max(a, b int) int {
	if a > b {
		return a
	}
	
	return b
}
