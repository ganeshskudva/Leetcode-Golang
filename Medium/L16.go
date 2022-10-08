package Medium

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum < target {
				l += 1
			} else if sum > target {
				r -= 1
			}
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
