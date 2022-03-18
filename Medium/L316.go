package Medium

func removeDuplicateLetters(s string) string {
	lastIndex := make([]int, 26)
	for i := 0; i < len(s); i++ {
		// track the lastIndex of character presence
		lastIndex[s[i]-'a'] = i
	}

	var st []int
	seen := make([]bool, 26)

	for i := 0; i < len(s); i++ {
		curr := s[i] - 'a'
		if seen[curr] {
			// if seen continue as we need to pick one char only
			continue
		}
		for len(st) > 0 && st[len(st)-1] > int(curr) && i < lastIndex[st[len(st)-1]] {
			// mark unseen
			seen[st[len(st)-1]] = false
			// pop out
			st = st[:len(st)-1]
		}
		st = append(st, int(curr))
		seen[curr] = true
	}

	var sb strings.Builder
	for len(st) > 0 {
		sb.WriteByte(byte(st[len(st)-1] + 'a'))
		st = st[:len(st)-1]
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
