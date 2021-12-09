package Medium

func canReach(arr []int, start int) bool {
	return solve(arr, start, make(map[int]bool))
}

func solve(arr []int, pos int, seen map[int]bool) bool {
	if pos < 0 || pos >= len(arr) {
		return false
	}

	if _, ok := seen[pos]; ok {
		return false
	}

	if arr[pos] == 0 {
		return true
	}

	seen[pos] = true
	return solve(arr, pos+arr[pos], seen) || solve(arr, pos-arr[pos], seen)
}
