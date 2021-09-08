package Medium

func shiftingLetters(s string, shifts []int) string {
	var sb strings.Builder
	shift := 0
	for i := len(s) - 1; i >= 0; i-- {
		shift = (shift + shifts[i]) % 26
		sb.WriteByte(byte((int(s[i]-'a')+shift)%26 + 'a'))
	}

	return reverse(sb.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
