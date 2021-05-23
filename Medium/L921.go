package Medium

func minAddToMakeValid(s string) int {
    	stSize, missMatch := 0,0
	
	for i := range s {
		if s[i] == '(' {
			stSize++
		} else if s[i] == ')' && stSize > 0 {
			stSize--
		} else {
			missMatch++
		}
	}
	
	return stSize + missMatch
}
