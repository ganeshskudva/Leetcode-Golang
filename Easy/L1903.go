package Easy

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if ((num[i] - '0') % 2) > 0 {
			return num[0 : i+1]
		}
	}

	return ""
}
