package Medium

func findWinners(matches [][]int) [][]int {
	var winners, losers []int
	table := make(map[int]int)

	for _, m := range matches {
		if _, ok := table[m[0]]; !ok {
			table[m[0]] = 0
		}
		table[m[1]]++
	}

	for k, v := range table {
		if v == 0 {
			winners = append(winners, k)
		}
		if v == 1 {
			losers = append(losers, k)
		}
	}

	sort.Ints(winners)
	sort.Ints(losers)
	return [][]int{winners, losers}
}
