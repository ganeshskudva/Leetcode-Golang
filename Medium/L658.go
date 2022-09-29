package Medium

func findClosestElements(arr []int, k int, x int) []int {
	lo, hi := 0, len(arr)-1
	for hi-lo >= k {
		if abs(arr[lo]-x) > abs(arr[hi]-x) {
			lo++
		} else {
			hi--
		}
	}
	var res []int
	for i := lo; i <= hi; i++ {
		res = append(res, arr[i])
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
