package Medium

func uniqueOccurrences(arr []int) bool {
	var exists struct{}
	mp, seen := make(map[int]int), make(map[int]struct{})

	for _, n := range arr {
		mp[n]++
	}
	for _, v := range mp {
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = exists
	}

	return true
}
