package Easy

func makeGood(s string) string {
	var st []byte

	for i := range s {
		if len(st) > 0 && areEqual(st[len(st)-1], s[i]) {
			st = st[:len(st)-1]
			continue
		}
		st = append(st, s[i])
	}

	return string(st)
}

func areEqual(a, b byte) bool {
	if (isUpperCase(a) && isLowerCase(b)) ||
		(isLowerCase(a) && isUpperCase(b)) {
		return byte(unicode.ToLower(rune(a))) == byte(unicode.ToLower(rune(b)))
	}

	return false
}

func isUpperCase(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

func isLowerCase(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}
