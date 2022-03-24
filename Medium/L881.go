package Medium

func numRescueBoats(people []int, limit int) int {
	boatCnt, left, right := 0, 0, len(people)-1
	if len(people) == 0 {
		return boatCnt
	}

	sort.Ints(people)

	for left <= right {
		sum := people[left] + people[right]
		if sum <= limit {
			left++
		}
		boatCnt++
		right--
	}

	return boatCnt
}
