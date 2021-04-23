package Easy

func countBinarySubstrings(s string) int {
	zeroes, ones, last, cnt := 0, 0, -1, 0

	for _, c := range s {
		if c == '0' {
			if last == 1 {
				zeroes = 0
			}
			zeroes++
			last = 0
		} else {
			if last == 0 {
				ones = 0
			}
			ones++
			last = 1
		}
		if (last == 1 && zeroes >= ones) || (last == 0 && ones >= zeroes) {
			cnt++
		}
	}

	return cnt
}
