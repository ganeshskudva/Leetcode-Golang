package Medium

func wonderfulSubstrings(word string) int64 {
	var (
		res int64
		mask int
	)
	cnt := make([]int64, 1024)
	cnt[0] = 1
	
	for i := 0; i < len(word); i++ {
		mask ^= 1 << (word[i] - 'a')
		res += cnt[mask]
		for i := 0; i < 10; i++ {
			res += cnt[mask ^ (1 << i)]
		}
		cnt[mask]++
	}
	
	return res
}
