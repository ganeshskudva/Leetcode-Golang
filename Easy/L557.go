package Easy

func reverseWords(s string) string {
	var sb strings.Builder

	strs := strings.Split(s, " ")
	for i := 0; i < len(strs); i++ {
		sb.WriteString(reverse(strs[i]))
		if i != len(strs)-1 {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
