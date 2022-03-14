package Easy

func sumOfUnique(nums []int) int {
	mp, sum := make(map[int]int), 0

	for _, n := range nums {
		mp[n]++
		if mp[n] == 1 {
			sum += n
		} else if mp[n] == 2 {
			sum -= n
		}
	}

	return sum
}
