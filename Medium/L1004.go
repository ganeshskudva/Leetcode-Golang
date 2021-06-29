package Medium

func longestOnes(nums []int, k int) int {
	max, zeroCnt, i, j := 0, 0, 0, 0
	for ; j < len(nums); j++ {
		if nums[j] == 0 {
			zeroCnt++
		}

		for zeroCnt > k {
			if nums[i] == 0 {
				zeroCnt--
			}
			i++
		}
		if (j - i + 1) > max {
			max = j - i + 1
		}
	}

	return max
}
