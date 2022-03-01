package Easy

func countBits(num int) []int {
	arr := make([]int, num+1)

	for i := 1; i < num+1; i++ {
		arr[i] = solve(i&(i-1), &arr) + 1
	}

	return arr
}

func solve(n int, arr *[]int) int {
	if n == 0 {
		return n
	}

	if (*arr)[n&(n-1)] != 0 {
		return (*arr)[n&(n-1)] + 1
	}
	(*arr)[n] = solve(n&(n-1), arr) + 1

	return (*arr)[n]
}
