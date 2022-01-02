package Medium

func numPairsDivisibleBy60(time []int) int {
	ans, mp := 0, make(map[int]int)

	for _, t := range time {
		other, reducedTime := 0, t%60
		if reducedTime != 0 {
			other = 60 - reducedTime
		}
		if val, ok := mp[other]; ok {
			ans += val
		}
		mp[reducedTime]++
	}

	return ans
}
