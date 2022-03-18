package Easy

func sortArrayByParity(nums []int) []int {
	ret := make([]int, len(nums))
	if len(nums) == 0 {
		return ret
	}

	i, j := 0, len(nums)-1
	for _, n := range nums {
		if isOdd(n) {
			ret[j] = n
			j--
		} else {
			ret[i] = n
			i++
		}
	}

	return ret
}

func isOdd(n int) bool {
	return n%2 > 0
}
