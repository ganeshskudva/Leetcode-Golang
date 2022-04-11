package Easy

func shiftGrid(grid [][]int, k int) [][]int {
	r, c := len(grid), len(grid[0])
	n := r * c
	sh := k % n

	res := make([][]int, len(grid))
	for i := range grid {
		res[i] = make([]int, len(grid[i]))
		copy(res[i], grid[i])
	}

	reverse(&res, 0, n-sh)
	reverse(&res, n-sh, n)
	reverse(&res, 0, n)

	return res
}

func reverse(grid *[][]int, lo, hi int) {
	for lo < hi {
		hi--
		swap(grid, lo, hi)
		lo++
	}
}

func swap(grid *[][]int, lo, hi int) {
	n := len((*grid)[0])
	(*grid)[lo/n][lo%n], (*grid)[hi/n][hi%n] = (*grid)[hi/n][hi%n], (*grid)[lo/n][lo%n]
}
