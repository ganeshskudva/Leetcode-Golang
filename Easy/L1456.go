package Easy

var exists struct{}

func max(a, b int) int {
	if a > b {
		return a
	}
	
	return b
}
func maxVowels(s string, k int) int {
	vowels, ans, cnt := "aeiou", 0, 0
	st := make(map[byte]struct{})
	for i := range vowels {
		st[vowels[i]] = exists
	}
	
	for i := range s {
		if _, ok := st[s[i]]; ok {
			cnt++
		}
		if i >= k {
			if _, ok := st[s[i - k]]; ok {
				cnt--
			}
		}
		ans = max(ans, cnt)
	}
	
	return ans
}
