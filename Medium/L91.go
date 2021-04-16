package Medium

func numDecodings(s string) int {
	mp := make(map[int]int)

	return memo(s, 0, mp)
}

func memo(s string, idx int, mp map[int]int) int {
	if idx == len(s) {
		return 1
	}

	if s[idx] == '0' {
		return 0
	}

	if _, ok := mp[idx]; ok {
		return mp[idx]
	}

	way := memo(s, idx+1, mp)

	if idx < len(s)-1 && (s[idx] == '1' || s[idx] == '2' && s[idx+1] < '7') {
		way += memo(s, idx+2, mp)
	}
	mp[idx] = way

	return mp[idx]
}
