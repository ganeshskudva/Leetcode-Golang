package Medium

func findAnagrams(s string, p string) []int {
	var ans []int
	if len(p) > len(s) {
		return ans
	}

	mp := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		mp[p[i]]++
	}

	begin, end, counter := 0, 0, len(mp)

	for end < len(s) {
		if _, ok := mp[s[end]]; ok {
			mp[s[end]]--
			if mp[s[end]] == 0 {
				counter--
			}
		}
		end++

		for counter == 0 {
			if _, ok := mp[s[begin]]; ok {
				mp[s[begin]]++
				if mp[s[begin]] > 0 {
					counter++
				}
			}

			if end-begin == len(p) {
				ans = append(ans, begin)
			}
			begin++
		}
	}

	return ans
}
