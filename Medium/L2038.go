package Medium

func winnerOfGame(colors string) bool {
	a, b := 0, 0

	for i := 1; i < len(colors)-1; i++ {
		if colors[i] == colors[i-1] && colors[i] == colors[i+1] {
			if colors[i] == 'A' {
				a++
			} else {
				b++
			}
		}
	}

	return a > b
}
