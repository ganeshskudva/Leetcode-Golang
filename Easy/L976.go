package Easy

func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		perimeter := nums[i] + nums[i-1] + nums[i-2]

		if nums[i] < nums[i-1]+nums[i-2] {
			return perimeter
		}
	}

	return 0
}
