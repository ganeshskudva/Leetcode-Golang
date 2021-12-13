package Easy

func maxPower(s string) int {
	if len(s) == 0 {
		return 0
	}

	cnt, ans := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
			if cnt > ans {
				ans = cnt
			}
		} else {
			cnt = 1
		}
	}

	return ans
}
