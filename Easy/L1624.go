package Easy

func maxLengthBetweenEqualCharacters(s string) int {
	indexes, dist := make(map[byte]int), -1

	for i := 0; i < len(s); i++ {
		if v, ok := indexes[s[i]-'a']; ok {
			dist = max(dist, i-v+1)
		} else {
			indexes[s[i]-'a'] = i
		}
	}

	if dist == -1 {
		return dist
	}
	
	return dist - 2
}
