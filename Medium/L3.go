package Medium

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	mp, max := make(map[byte]int), 0

	for i, j := 0, 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; ok {
			j = maximum(j, mp[s[i]]+1)
		}
		mp[s[i]] = i
		max = maximum(max, i-j+1)
	}

	return max
}

func maximum(i, j int) int {
	if i > j {
		return i
	}

	return j
}
