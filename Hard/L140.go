package Hard

var exists struct{}

func wordBreak(s string, wordDict []string) []string {
	mp, st := make(map[int][]string), make(map[string]struct{})

	for _, w := range wordDict {
		st[w] = exists
	}

	return solve(s, 0, st, mp)
}

func solve(s string, start int, st map[string]struct{}, mp map[int][]string) []string {
	if v, ok := mp[start]; ok {
		return v
	}

	var res []string
	if start == len(s) {
		res = append(res, "")
		return res
	}

	curr := s[start:]
	for w, _ := range st {
		if strings.HasPrefix(curr, w) {
			subLst := solve(s, start+len(w), st, mp)
			for _, sub := range subLst {
				if len(sub) == 0 {
					res = append(res, w+"")
				} else {
					res = append(res, w+" "+sub)
				}
			}
		}
	}

	mp[start] = res
	return mp[start]
}
