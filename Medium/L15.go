package Medium

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i+2 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k, tgt := i+1, len(nums)-1, -nums[i]
		for j < k {
			if nums[j]+nums[k] == tgt {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[j]+nums[k] > tgt {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
