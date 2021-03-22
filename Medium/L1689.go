package Medium

func minPartitions(n string) int {
	max := 0
	if len(n)== 0 {
		return max
	}
	
	for _, c := range n {
		max = getMax(max, int(c-'0'))
	}
	
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
