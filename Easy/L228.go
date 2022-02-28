package Easy

func summaryRanges(nums []int) []string {
	start, end := 0, 0
	var list []string

	for i := 0; i < len(nums); i++ {
		for i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			end, i = end+1, i+1
		}
		if start == end {
			list = append(list, fmt.Sprintf("%d", nums[start]))
		} else {
			list = append(list, fmt.Sprintf("%d->%d", nums[start], nums[end]))
		}
		end++
		start = end
	}

	return list
}
