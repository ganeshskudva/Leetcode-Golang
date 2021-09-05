package Hard

func orderlyQueue(s string, k int) string {
	if len(s) == 0 {
		return s
	}

	n, s := len(s), s+s

	if k >= 2 {
		cnt := make([]int, 26)
		for i := 0; i < n; i++ {
			cnt[s[i]-'a']++
		}
		var sb strings.Builder
		for i := 0; i < 26; i++ {
			for j := 0; j < cnt[i]; j++ {
				sb.WriteByte(byte(i + 'a'))
			}
		}

		return sb.String()
	}

	res, arr, x := make([]byte, n*2), make([]int, n*2), 0
	res[0], arr[0] = s[0], -1

	for i := 1; i < n*2; i++ {
		for x >= 0 && res[arr[x]+1] > s[i] {
			x = arr[x]
		}
		res[x+1] = s[i]

		if x >= 0 {
			idx := arr[x]
			for idx >= 0 && res[idx+1] != s[i] {
				idx = arr[idx]
			}

			if res[idx+1] == s[i] {
				idx++
			}
			arr[x+1] = idx
		}

		x++
	}

	return string(res[:n])
}
