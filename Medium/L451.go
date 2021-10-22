package Medium

func frequencySort(s string) string {
	mp := make(map[byte]int)
	for i := range s {
		mp[s[i]]++
	}

	bucket := make([][]byte, len(s)+1)
	for i := 0; i < len(mp)+1; i++ {
		bucket[i] = make([]byte, 0)
	}
	for k, freq := range mp {
		bucket[freq] = append(bucket[freq], k)
	}

	var sb strings.Builder
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) != 0 {
			for idx := 0; idx < len(bucket[i]); idx++ {
				for pos := 0; pos < i; pos++ {
					sb.WriteByte(bucket[i][idx])
				}
			}
		}
	}

	return sb.String()
}
