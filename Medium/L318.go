package Medium

func maxProduct(words []string) int {
	max, checker := 0, make([]int, len(words))
	for i := range checker {
		num := 0
		for j := 0; j < len(words[i]); j++ {
			num |= 1 << uint64(int(words[i][j])-int('a'))
		}
		checker[i] = num
	}
	
	for i := range words {
		for j:= i+1; j < len(words); j++ {
			if (checker[i] & checker[j]) == 0 {
				max = maxi(max, len(words[i]) * len(words[j]))
			}
		}
	}
	
	return max
}

func maxi (a, b int) int {
	if a > b {
		return a
	}
	
	return b
}
