package Medium

func checkPossibility(nums []int) bool {
    if len(nums) <= 2 {
        return true
    }
    
    modified := false
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] <= nums[i+1] {
            continue
        }
        if modified {
            return false
        }
        if i > 0 && nums[i-1] > nums[i+1] {
            nums[i+1] = nums[i]
        }
        modified = true
    }
    return true
}
