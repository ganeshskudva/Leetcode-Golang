package Hard

func longestValidParentheses(s string) int {
    	var st []int
	max, left := 0, -1
	
	for i:= 0; i < len(s); i++ {
		if s[i] == '(' {
			st = append(st, i)
		} else {
			if len(st) == 0 {
				left = i
			} else {
				st = st[:len(st)-1]
				
				if len(st) == 0 {
					max = Max(max, i-left)
				} else {
					max = Max(max, i - st[len(st)-1])
				}
			}
		}
	}
	
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
