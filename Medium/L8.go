package Medium

func myAtoi(s string) int {
	index, total, sign := 0, 0, 1

	if len(s) == 0 {
		return 0
	}

	for index < len(s) && s[index] == ' ' {
		index++
	}

	if index == len(s) {
		return 0
	}

	if s[index] == '+' || s[index] == '-' {
		if s[index] == '+' {
			sign = 1
		} else {
			sign = -1
		}
		index++
	}

	for index < len(s) {
		digit := int(s[index] - '0')
		if digit < 0 || digit > 9 {
			break
		}

		if math.MaxInt32/10 < total || (math.MaxInt32/10 == total && math.MaxInt32%10 < digit) {
			if sign == 1 {
				return math.MaxInt32
			}

			return math.MinInt32
		}

		total = total*10 + digit
		index++
	}

	return total * sign
}
