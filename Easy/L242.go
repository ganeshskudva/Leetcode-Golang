package Easy

func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	for i := range s {
		alphabet[s[i]-'a']++
	}
	for i := range t {
		alphabet[t[i]-'a']--
		if alphabet[t[i]-'a'] < 0 {
			return false
		}
	}

	for _, a := range alphabet {
		if a != 0 {
			return false
		}
	}

	return true
}
