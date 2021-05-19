package Medium

func minMoves2(nums []int) int {
  sort.Ints(nums)
	i, j, cnt := 0, len(nums)-1, 0
	
	for i < j {
		cnt += nums[j] - nums[i]
		i++
		j--
	}
	
	return cnt
}
