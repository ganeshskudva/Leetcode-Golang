package Medium

func dailyTemperatures(T []int) []int {
    	ret, n := make([]int, len(T)), len(T)
	var st []int

	for i := n - 1; i >= 0; i-- {
		cur := T[i]
		for len(st) > 0 && cur >= T[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ret[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}

	return ret
}
