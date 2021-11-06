package Medium

func singleNumber(nums []int) []int {
	diff := 0
	for _, n := range nums {
		diff ^= n
	}

	diff = diff & -diff
	rets := []int{0, 0}
	for _, n := range nums {
		if n&diff == 0 {
			rets[0] ^= n
		} else {
			rets[1] ^= n
		}
	}

	return rets
}
