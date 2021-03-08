package Easy

func Reverse(s string) string {
	var str strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		str.WriteString(string(s[i]))
	}

	return str.String()
}

func isPalindrome(s string) bool {
	return s == Reverse(s)
}

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	
	if isPalindrome(s) {
		return 1
	} else {
		return 2
	}
}
