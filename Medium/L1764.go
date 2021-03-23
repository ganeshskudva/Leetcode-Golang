package Medium

func canChoose(groups [][]int, nums []int) bool {
	i := 0
	for start := 0; i < len(groups) && len(groups[i])+start <= len(nums); start++ {
		if search(groups[i], nums, start) {
			start += len(groups[i]) - 1
			i++
		}
	}

	return i == len(groups)
}

func search(group, nums []int, start int) bool {
	for i := 0; i < len(group); i++ {
		if group[i] != nums[i+start] {
			return false
		}
	}

	return true
}
