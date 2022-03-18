package Easy

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}

	st, end := 0, len(s)-1
	for st <= end {
		s[st], s[end] = s[end], s[st]
		st, end = st+1, end-1
	}
}
