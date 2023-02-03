package Medium

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	res, idx, dir := make([]string, numRows), 0, -1

	for i := 0; i < len(s); i++ {
		res[idx] += fmt.Sprintf("%c", s[i])
		if idx == 0 || idx == numRows-1 {
			dir = -dir
		}
		idx += dir
	}

	var sb strings.Builder
	for i := 0; i < len(res); i++ {
		sb.WriteString(res[i])
	}

	return sb.String()
}
