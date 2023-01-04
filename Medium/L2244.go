package Medium

func minimumRounds(tasks []int) int {
	cnt, mp := 0, make(map[int]int)

	for _, t := range tasks {
		mp[t]++
	}

	for _, v := range mp {
		if v < 2 {
			return -1
		}

		if v%3 == 0 {
			cnt += v / 3
		} else {
			cnt += (v / 3) + 1
		}
	}

	return cnt
}
