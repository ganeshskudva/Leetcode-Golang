package Hard

func maximumScore(nums []int, k int) int {
	maxScore, minValue := nums[k], nums[k]
	left, right := k, k

	for left >= 0 && right < len(nums) {
		minValue = min(minValue, min(nums[left], nums[right]))
		maxScore = max(minValue*(right-left+1), maxScore)

		if left == 0 && right < len(nums) {
			right++
		} else if right == len(nums)-1 && left >= 0 {
			left--
		} else if nums[right+1] > nums[left-1] {
			right++
		} else {
			left--
		}
	}

	return maxScore
}
