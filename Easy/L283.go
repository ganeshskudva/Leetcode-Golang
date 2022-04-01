package Easy

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	insertPos := 0
	for _, n := range nums {
		if n != 0 {
			nums[insertPos] = n
			insertPos++
		}
	}

	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}
