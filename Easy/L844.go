package Easy

func backspaceCompare(s string, t string) bool {
	return eval(s) == eval(t)
}

func eval(s string) string {
	n, cnt := len(s), 0
	var res strings.Builder

	for i := n - 1; i >= 0; i-- {
		if s[i] == '#' {
			cnt++
		} else {
			if cnt > 0 {
				cnt--
			} else {
				res.WriteByte(s[i])
			}
		}
	}

	return res.String()
}
