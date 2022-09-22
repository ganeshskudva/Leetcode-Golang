package Medium

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	sum := 0

	for _, n := range nums {
		if n%2 == 0 {
			sum += n
		}
	}

	var ret []int
	for _, q := range queries {
		x, y := q[1], q[0]
		if abs(nums[x])%2 == 0 {
			sum -= nums[x]
		}
		if (nums[x]+y)%2 == 0 {
			sum += nums[x] + y
		}

		nums[x] += y
		ret = append(ret, sum)
	}

	return ret
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}
