package Hard

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	result, sign, num := 0, 1, 0

	var st []int
	st = append(st, sign)

	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '+' || s[i] == '-' {
			result += sign * num
			sign = st[len(st)-1]
			if s[i] != '+' {
				sign *= -1
			}
			num = 0
		} else if s[i] == '(' {
			st = append(st, sign)
		} else if s[i] == ')' {
			st = st[:len(st)-1]
		}
	}

	result += sign * num
	return result
}
