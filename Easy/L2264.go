package Easy

func largestGoodInteger(num string) string {
	res := -1
	for i := 0; i+2 < len(num); i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			res = max(res, int(num[i]-'0'))
		}
	}

	if res == -1 {
		return ""
	}

	return fmt.Sprintf("%d%d%d", res, res, res)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
