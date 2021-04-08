package Medium

func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) == 0 {
		return ans
	}

	var str string
	mp := make(map[int]string)
	arr := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	i := 2
	for _, s := range arr {
		mp[i] = s
		i++
	}

	bt(&ans, &str, mp, digits, 0)

	return ans
}

func bt(ans *[]string, sb *string, mp map[int]string, digits string, idx int) {
	if idx > len(digits) {
		return
	}

	if len(*sb) == len(digits) {
		*ans = append(*ans, *sb)
		return
	}

	d := int(digits[idx] - '0')
	s := mp[d]
	for i := 0; i < len(s); i++ {
		*sb += string(s[i])
		bt(ans, sb, mp, digits, idx+1)
		*sb = (*sb)[:len(*sb)-1]
	}
}
