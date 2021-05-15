package Hard

func isNumber(s string) bool {
    if len(s) == 0 || len(strings.TrimSpace(s)) == 0 {
		return false
	}

	s = strings.TrimSpace(s)
	seenNum, seenE, seenD := false, false, false

	for i := range s {
		ch := byte(unicode.ToLower(rune(s[i])))
		switch ch {
		case '.':
			if seenD || seenE {
				return false
			}
			seenD = true
		case 'e':
			if seenE || !seenNum {
				return false
			}
			seenE = true
			seenNum = false
		case '+':
			fallthrough
		case '-':
			if i != 0 && s[i-1] != 'e' {
				return false
			}
			seenNum = false
		default:
			if ch-'0' < 0 || ch-'0' > 9 {
				return false
			}
			seenNum = true
		}

	}

	return seenNum
}
