package Easy

func minMoves(nums []int) int {
  min, sum, n := math.MaxInt32, 0, len(nums)
	
	for _, v := range nums {
		if v < min {
			min = v
		}
		sum += v
	}
	
	return sum - (n*min)
}
