package Easy

func majorityElement(nums []int) int {
	cnt, candidate := 0, -1

	for _, n := range nums {
		if cnt == 0 {
			candidate = n
		}

		if candidate == n {
			cnt++
		} else {
			cnt -= 1
		}
	}

	return candidate
}
