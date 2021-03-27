package Medium

func countSubstrings(s string) int {
	cnt := 0
	if len(s) == 0 {
		return cnt
	}
	
	for i:= range s {
		extendPal(s, i, i , &cnt)
		extendPal(s, i, i+1, &cnt)
	}
	
	return cnt
}

func extendPal(s string, left, right int, cnt *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		*cnt++
		left--
		right++
	}
}
