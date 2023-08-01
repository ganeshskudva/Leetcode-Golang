package Medium

func combine(n int, k int) [][]int {
	var (
		solve func(tmp *[]int, start, k int)
		res   [][]int
		tmp   []int
	)

	solve = func(tmp *[]int, start, k int) {
		if k == 0 {
			tmpCopy := make([]int, len(*tmp))
			copy(tmpCopy, *tmp)
			res = append(res, tmpCopy)
			return
		}

		for i := start; i <= n; i++ {
			*tmp = append(*tmp, i)
			solve(tmp, i+1, k-1)
			*tmp = (*tmp)[:len(*tmp)-1]
		}
	}

	solve(&tmp, 1, k)
	return res
}
