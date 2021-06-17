package Medium

func numSubarrayBoundedMax(nums []int, left int, right int) int {
   	start, last, res := -1, -1, 0
	for i := range nums {
		if nums[i] > right {
			start, last = i, i
			continue
		}
		
		if nums[i] >= left {
			last = i
		}
		
		res += last - start
	}
	
	return res 
}
