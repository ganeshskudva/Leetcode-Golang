package Medium

func longestPalindrome(s string) string {
	lo, maxLen := 0, 0
	if len(s) < 2 {
		return s
	}

	for i := range s {
		extendPal(s, i, i, &lo, &maxLen)
		extendPal(s, i, i+1, &lo, &maxLen)
	}

	return s[lo:lo+maxLen]
}

func extendPal(s string, left, right int, lo, maxLen *int)  {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	
	if *maxLen < right-left-1 {
		*lo = left+1
		*maxLen = right-left-1
	}
}
