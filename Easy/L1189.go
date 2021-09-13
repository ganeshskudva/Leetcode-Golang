package Easy

func maxNumberOfBalloons(text string) int {
	mp := make(map[byte]int)
	arr := []byte{'b', 'a', 'l', 'o', 'n'}

	for i := range text {
		if isValid(text[i]) {
			mp[text[i]]++
		}
	}

	min := math.MaxInt64
	for i := range arr {
		if val, ok := mp[arr[i]]; ok {
			if arr[i] == 'l' || arr[i] == 'o' {
				if val < 2 {
					return 0
				}
				val = val / 2
			}
			if val < min {
				min = val
			}
		} else {
			return 0
		}
	}

	if min == math.MaxInt64 {
		return 0
	}

	return min
}

var validMap = map[byte]bool{
	'b': true,
	'a': true,
	'l': true,
	'o': true,
	'n': true,
}

func isValid(ch byte) bool {
	if _, ok := validMap[ch]; !ok {
		return false
	}

	return true
}
