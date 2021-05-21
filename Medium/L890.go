package Medium

func findAndReplacePattern(words []string, pattern string) []string {
	var res []string
	for _, w := range words {
		if match(w, pattern) {
			res = append(res, w)
		}
	}
	
	return res
}

func match(s, p string) bool {
	mp1, mp2 := make(map[byte]byte), make(map[byte]byte)
	
	for i := range  s {
		if _, ok := mp1[s[i]]; !ok {
			mp1[s[i]] = p[i]
		}
		if _, ok := mp2[p[i]]; !ok {
			mp2[p[i]] = s[i]
		}
		if mp1[s[i]] != p[i] || mp2[p[i]] != s[i] {
			return false
		}
	}
	
	return true
}
