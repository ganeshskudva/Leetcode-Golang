package Easy

func minimumAbsDifference(arr []int) [][]int {
	var res [][]int

	if len(arr) == 0 {
		return res
	}

	sort.Ints(arr)
	min := math.MaxInt64
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff <= min {
			if diff < min {
				min = diff
				res = nil
			}
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}

	return res
}
