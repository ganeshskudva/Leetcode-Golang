package Medium

func decodeString(s string) string {
	index := 0
	return solve(s, &index)
}

func solve(s string, index *int) string {
	var sb strings.Builder
	var num string

	for i := *index; i < len(s); i++ {
		if s[i] != '[' && s[i] != ']' && !isDigit(s[i]) {
			sb.WriteByte(s[i])
		} else if isDigit(s[i]) {
			num += string(rune(s[i]))
		} else if s[i] == '[' {
			*index = i + 1
			next := solve(s, index)
			for n := intVal(num); n > 0; n-- {
				sb.WriteString(next)
			}
			num, i = "", *index
		} else if s[i] == ']' {
			*index = i
			return sb.String()
		}
	}

	return sb.String()
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func intVal(num string) int {
	i, _ := strconv.Atoi(num)

	return i
}
