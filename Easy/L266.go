package Easy

var exists struct{}

func canPermutePalindrome(s string) bool {
	mp := make(map[byte]struct{})

	for i := range s {
		if _, ok := mp[s[i]]; !ok {
			mp[s[i]] = exists
		} else {
			delete(mp, s[i])
		}
	}

	return len(mp) <= 1
}
