package Medium

func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, len(spells))
	sort.Ints(potions)
	
	for i := 0; i < len(spells); i++ {
		low, high := 0, len(potions) - 1
		for low <= high {
			mid := low + (high - low) / 2
			
			if int64(spells[i] * potions[mid]) >= success {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		res[i] = len(potions) - 1 - high
	}
	
	return res
}
