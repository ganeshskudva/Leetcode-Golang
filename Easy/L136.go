package Easy

func singleNumber(nums []int) int {
	val := 0

	for _, n := range nums {
		val ^= n
	}

	return val
}
