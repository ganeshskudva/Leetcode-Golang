package Medium

func countNicePairs(nums []int) int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		rev := revInt(nums[i])
		diff := nums[i] - rev
		mp[diff]++
	}

	res, mod := 0, 1000000007
	for _, v := range mp {
		if v == 1 {
			continue
		}
		res = (res + v*(v-1)/2) % mod
	}

	return res
}

func revInt(n int) int {
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}
