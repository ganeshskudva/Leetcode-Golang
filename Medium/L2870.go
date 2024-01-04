package Medium

func minOperations(nums []int) int {
	mp, cnt := make(map[int]int), 0

	for _, n := range nums {
		mp[n]++
	}

	for _, v := range mp {
		if v == 1 {
			return -1
		}

		cnt += v / 3
		if v%3 != 0 {
			cnt++
		}
	}

	return cnt
}
