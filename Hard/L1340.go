package Hard

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxJumps(arr []int, d int) int {
	n, mp, res := len(arr), make(map[int]int), 1
	
	for i := 0; i < n; i++ {
		res = max(res, solve(arr, mp, n, d, i))
	}
	
	return res
}

func solve(arr []int, mp map[int]int, n, d, idx int) int {
	if v, ok := mp[idx]; ok {
		return v
	}
	
	res := 1
	for j := idx + 1; j <= min(idx + d, n - 1) && arr[j] < arr[idx]; j++ {
		res = max(res, 1 + solve(arr, mp, n, d, j))
	}
	for j := idx - 1; j >= max(idx - d, 0) && arr[j] < arr[idx]; j-- {
		res = max(res, 1 + solve(arr, mp, n, d, j))
	}
	
	mp[idx] = res
	return mp[idx]
}
