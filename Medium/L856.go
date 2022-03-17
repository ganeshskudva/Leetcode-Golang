package Medium

func scoreOfParentheses(s string) int {
	var st []int
	cur := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st = append(st, cur)
			cur = 0
		} else {
			cur = st[len(st)-1] + max(cur*2, 1)
			st = st[:len(st)-1]
		}
	}

	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
