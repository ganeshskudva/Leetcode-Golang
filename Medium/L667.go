package Medium

func constructArray(n int, k int) []int {
    	var ans []int
	l, r, i := 1, k+1, 0
	for i < k {
		ans = append(ans, l)
		i++
		l++
		ans = append(ans, r)
		i++
		r--
	}

	if l == r {
		ans = append(ans, r)
		i++
	}

	for i < n {
		ans = append(ans, i+1)
		i++
	}

	return ans
}
