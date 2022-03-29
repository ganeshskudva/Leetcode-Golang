package Medium

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	fast = 0
	for fast != slow {
		fast, slow = nums[fast], nums[slow]
	}

	return slow
}
