package Medium

func getSmallestString(n int, k int) string {
	chars := make([]byte, n)
	for i := 0; i < n; i++ {
		chars[i] = 'a'
	}
	k -= n

	for k > 0 {
		if chars[n-1] < 'z' {
			chars[n-1]++
			k--
		} else {
			n--
		}
	}

	return string(chars)
}
