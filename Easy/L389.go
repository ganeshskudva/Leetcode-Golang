package Easy

func findTheDifference(s string, t string) byte {
	sumS, sumT := 0, 0

	for i := 0; i < len(s); i++ {
		sumS += int(s[i])
	}

	for i := 0; i < len(t); i++ {
		sumT += int(t[i])
	}

	return byte(sumT - sumS - 'a')
}
