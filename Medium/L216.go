package Medium

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var tmp []int

	if k == 0 || n == 0 {
		return ans
	}

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	backtrack(&ans, &tmp, arr, k, n, 0)

	return ans
}

func backtrack(ans *[][]int, tmp *[]int, arr []int, k, target, start int) {
	if len(*tmp) == k && target == 0 {
		tmpCopy := make([]int, len(*tmp))
		copy(tmpCopy, *tmp)
		*ans = append(*ans, tmpCopy)
	}

	if target < 0 {
		return
	}

	for i := start; i < len(arr); i++ {
		*tmp = append(*tmp, arr[i])
		backtrack(ans, tmp, arr, k, target-arr[i], i+1)
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
