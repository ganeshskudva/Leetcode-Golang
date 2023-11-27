package Medium

func getSumAbsoluteDifferences(nums []int) []int {
	result, prefixSum, suffixSum := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))

	prefixSum[0], suffixSum[len(nums)-1] = nums[0], nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
		suffixSum[len(nums)-i-1] = suffixSum[len(nums)-i] + nums[len(nums)-i-1]
	}

	for i := 0; i < len(nums); i++ {
		currAbsDiff := ((nums[i] * i) - prefixSum[i]) + (suffixSum[i] - (nums[i] * (len(nums) - i - 1)))
		result[i] = currAbsDiff
	}
	
	return result
}
