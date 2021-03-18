package Easy

func mostVisited(n int, rounds []int) []int {
    	var res []int
	
	for i:= rounds[0]; i <= rounds[len(rounds) - 1]; i++ {
		res = append(res, i)
	}
	
	if len(res) > 0 {
		return res
	}
	
	for i:= 1; i <= rounds[len(rounds) -1]; i++ {
		res = append(res, i)
	}
	
	for i := rounds[0]; i <= n; i++ {
		res = append(res, i)
	}
	
	return res
}
