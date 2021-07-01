package Medium

func grayCode(n int) []int {
	var res []int
	res = append(res, 0)
	for i := 0; i < n; i++ {
		size := len(res)
		for k := size - 1; k >= 0; k-- {
			res = append(res, res[k]|1<<i)
		}
	}

	return res
}
