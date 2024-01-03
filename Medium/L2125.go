package Medium

func numberOfBeams(bank []string) int {
	res, prev := 0, 0

	for _, s := range bank {
		cnt := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '1' {
				cnt++
			}
		}
		if cnt > 0 {
			res += prev * cnt
			prev = cnt
		}
	}

	return res
}
