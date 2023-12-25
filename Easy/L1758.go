package Easy

func minOperations(s string) int {
	swaps := 0
	for i := 0; i < len(s); i++ {
		if i %2 == 0 {
			if s[i] != '0' {
				swaps++
			}
		} else {
			if s[i] != '1' {
				swaps++
			}
		}
	}
	
	return min(swaps, len(s) - swaps)
}
