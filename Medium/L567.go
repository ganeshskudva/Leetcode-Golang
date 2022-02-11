package Medium

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
		return false
	}

	mp := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		mp[s1[i]]++
	}

	counter, begin, end := len(mp), 0, 0
	for end < len(s2) {
		ch := s2[end]

		if _, ok := mp[ch]; ok {
			mp[ch]--
			if mp[ch] == 0 {
				counter--
			}
		}
		end++

		for counter == 0 {
			tmp := s2[begin]
			if _, ok := mp[tmp]; ok {
				mp[tmp]++
				if mp[tmp] > 0 {
					counter++
				}
			}
			if end-begin == len(s1) {
				return true
			}
			begin++
		}
	}

	return false
}
