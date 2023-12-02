package Easy

func countCharacters(words []string, chars string) int {
	var isMatch func(s string) bool
	res := 0

	isMatch = func(s string) bool {
		mp := make(map[byte]int)

		for _, c := range chars {
			mp[byte(c)]++
		}

		for _, c := range s {
			if _, ok := mp[byte(c)]; !ok {
				return false
			}
			mp[byte(c)]--
			if mp[byte(c)] < 0 {
				return false
			}
		}

		return true
	}

	for _, w := range words {
		if isMatch(w) {
			res += len(w)
		}
	}

	return res
}
