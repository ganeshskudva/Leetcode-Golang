package Medium

func wiggleMaxLength(nums []int) int {
            if len(nums) < 2 {
            return len(nums)
        }
        prevdiff := nums[1] - nums[0];
        count := 2
        if prevdiff != 0 {
            count = 2
        } else {
            count = 1
        }
        for i:=2; i < len(nums); i++ {
            diff:= nums[i] - nums[i-1]
            if (diff > 0 && prevdiff <= 0) || (diff < 0 && prevdiff >= 0) {
                count++
                prevdiff = diff
            }
        }
        return count
}
