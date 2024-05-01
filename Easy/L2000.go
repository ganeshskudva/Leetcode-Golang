package Easy

func reversePrefix(word string, ch byte) string {
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			return reverse(word[:i+1]) + word[i+1:]
		}
	}

	return word
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, ch := range s {
		n--
		runes[n] = ch
	}
	return string(runes)
}
