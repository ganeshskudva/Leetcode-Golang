package Easy

func checkValidString(s string) bool {
	cMin, cMax := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cMax, cMin = cMax+1, cMin+1
		} else if s[i] == ')' {
			cMax, cMin = cMax-1, cMin-1
		} else if s[i] == '*' {
			cMax, cMin = cMax+1, cMin-1
		}

		if cMax < 0 {
			return false
		}

		cMin = max(cMin, 0)
	}

	return cMin == 0
}
