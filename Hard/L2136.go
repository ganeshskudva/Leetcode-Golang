package Hard

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	pairs := make([][]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = make([]int, 2)
		pairs[i][0], pairs[i][1] = plantTime[i], growTime[i]
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[j][1] > pairs[i][1]
	})

	plantingDays, totalDays := 0, 0
	for _, p := range pairs {
		if plantingDays+p[0]+p[1] > totalDays {
			totalDays = plantingDays + p[0] + p[1]
		}
		plantingDays += p[0]
	}

	return totalDays
}
