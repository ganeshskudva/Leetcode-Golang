package Easy

var exists struct{}

func wordPattern(pattern string, s string) bool {
	mp, st := make(map[byte]string), make(map[string]struct{})
	strs := strings.Split(s, " ")

	if len(strs) != len(pattern) {
		return false
	}

	for i := range strs {
		if v, ok := mp[pattern[i]]; ok {
			if v != strs[i] {
				return false
			}
		} else {
			if _, ok := st[strs[i]]; ok {
				return false
			}
			st[strs[i]] = exists
			mp[pattern[i]] = strs[i]
		}
	}

	return true
}
