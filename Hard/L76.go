package Hard

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	mp := make(map[byte]int)
	for i := range t {
		mp[t[i]]++
	}

	begin, end, counter, length, head := 0, 0, len(mp), math.MaxInt32, 0

	for end < len(s) {
		if v, ok := mp[s[end]]; ok {
			mp[s[end]] = v - 1
			if mp[s[end]] == 0 {
				counter--
			}
		}
		end++

		for counter == 0 {
			if v, ok := mp[s[begin]]; ok {
				mp[s[begin]] = v + 1
				if mp[s[begin]] > 0 {
					counter++
				}
			}

			if end-begin < length {
				length = end - begin
				head = begin
			}
			begin++
		}
	}

	if length == math.MaxInt32 {
		return ""
	}

	return s[head : head+length]
}
