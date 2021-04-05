package Easy

func maxRepeating(sequence string, word string) int {
	var sb strings.Builder
	cnt := 0
	sb.WriteString(word)
	for strings.Contains(sequence, sb.String()) {
		cnt++
		sb.WriteString(word)
	}
	return cnt
}
