package Medium

func findDiagonalOrder(nums [][]int) []int {
	size, mp := 0, make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		numList := make([]int, len(nums[i]))
		copy(numList, nums[i])
		for j := 0; j < len(numList); j++ {
			var lst []int
			index := i + j
			if _, ok := mp[index]; ok {
				lst = mp[index]
			}
			lst = append([]int{numList[j]}, lst...)
			mp[index] = lst
			size++
		}
	}

	maxLen := getMaxLen(mp)
	var resultList []int
	for i := 0; i <= maxLen; i++ {
		diagVal := mp[i]
		resultList = append(resultList, diagVal...)
	}

	return resultList
}

func getMaxLen(mp map[int][]int) int {
	maxLen := math.MinInt32

	for k, _ := range mp {
		maxLen = max(maxLen, k)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
