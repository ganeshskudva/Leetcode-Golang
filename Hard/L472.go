package Hard

var exists struct{}

func findAllConcatenatedWordsInADict(words []string) []string {
	var res []string
	preWords := make(map[string]struct{})

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for i := 0; i < len(words); i++ {
		if canConcat(0, words[i], preWords, make(map[int]bool)) {
			res = append(res, words[i])
		}
		preWords[words[i]] = exists
	}

	return res
}

func canConcat(idx int, s string, st map[string]struct{}, mp map[int]bool) bool {
	if idx == len(s) {
		return true
	}
	if len(st) == 0 {
		return false
	}

	if v, ok := mp[idx]; ok {
		return v
	}

	for i := idx + 1; i <= len(s); i++ {
		if _, ok := st[s[idx:i]]; ok {
			if canConcat(i, s, st, mp) {
				mp[idx] = true
				return mp[idx]
			}
		}
	}

	mp[idx] = false
	return mp[idx]
}
