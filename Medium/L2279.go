package Medium

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	var diff []int
	for i := 0; i < len(rocks); i++ {
		diff = append(diff, capacity[i]-rocks[i])
	}

	sort.Ints(diff)
	ans, cnt := 0, 0
	for _, entry := range diff {
		cnt += entry
		if cnt > additionalRocks {
			break
		}
		ans++
	}

	return ans
}
