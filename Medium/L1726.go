package Medium

func tupleSameProduct(nums []int) int {
    res := 0
	mulFreq := make(map[int]int)
	
	for i:= 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			prod := nums[i] * nums[j]
			c := mulFreq[prod]
			res += c
			mulFreq[prod] = c+1
		}
	}
	
	return res*8
}
