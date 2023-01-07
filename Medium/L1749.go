package Medium

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxAbsoluteSum(nums []int) int {
	posAns, curPosSum, negAns, curNegSum := 0, 0, 0, 0

	for i := range nums {
		if curPosSum+nums[i] > 0 {
			curPosSum += nums[i]
		} else {
			curPosSum = 0
		}
		posAns = max(posAns, curPosSum)

		if curNegSum+nums[i] < 0 {
			curNegSum += nums[i]
		} else {
			curNegSum = 0
		}
		negAns = min(negAns, curNegSum)
	}

	return max(posAns, -negAns)
}
