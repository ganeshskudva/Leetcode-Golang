package Easy

func arraySign(nums []int) int {
	sign := 1
	for _, n := range nums {
		if n < 0 {
			sign *= -1
		} else if n == 0 {
			return 0
		}
	}

	return sign
}
