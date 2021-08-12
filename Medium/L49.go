package Medium

func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	if len(strs) == 0 {
		return ans
	}

	mp := make(map[string][]string)

	for _, s := range strs {
		k := getKey(s)
		mp[k] = append(mp[k], s)
	}

	for _, v := range mp {
		ans = append(ans, v)
	}

	return ans
}

func getKey(s string) string {
	arr := make([]byte, 26)

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}

	var sb strings.Builder
	for _, v := range arr {
		sb.WriteByte(v)
	}

	return sb.String()
}
