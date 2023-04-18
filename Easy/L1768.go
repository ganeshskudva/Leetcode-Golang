package Easy

func mergeAlternately(word1 string, word2 string) string {
	i, j, n, m := 0, 0, len(word1), len(word2)
	res := make([]byte, n + m)
	
	for k := 0; k < n + m; k++ {
		if i < n && j < m {
			res[k] = word1[i]
			k, i = k + 1, i + 1
		} else if i < n {
			res[k] = word1[i]
			i++
		}
		if j < m {
			res[k] = word2[j]
			j++
		}
	}
	
	return string(res)
}
