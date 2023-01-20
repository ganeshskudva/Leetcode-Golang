package Medium

func subarraysDivByK(nums []int, k int) int {
	mp := make(map[int]int)
	mp[0] = 1
	cnt, sum := 0, 0

	for _, n := range nums {
		sum = (sum + n) % k
		if sum < 0 {
			sum += k
		}
		cnt += mp[sum]
		mp[sum]++
	}

	return cnt
}
