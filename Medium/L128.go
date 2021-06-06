package Medium

func longestConsecutive(nums []int) int {
    numMap := map[int]bool{}
    consecutiveNumCnt := 0

    for _, num := range nums{
        numMap[num] = true
    }
    
    for num := range numMap{
        if numMap[num - 1]{
            continue
        }
        
        cur := num
        for numMap[cur+1]{
            cur++
        }
        
        consecutiveNumCnt = max(consecutiveNumCnt, cur - num+1)
    }
    
    return consecutiveNumCnt
}

func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}
