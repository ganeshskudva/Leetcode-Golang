package Medium

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    sort.Ints(horizontalCuts)
    sort.Ints(verticalCuts)

    maxX := max(h - horizontalCuts[len(horizontalCuts) - 1],
    horizontalCuts[0])
    
    maxY := max(w - verticalCuts[len(verticalCuts) - 1],
    verticalCuts[0])

    for i := 0; i < len(horizontalCuts) - 1; i++ {
        maxX = max(maxX, horizontalCuts[i + 1]-horizontalCuts[i])
    }
    for i := 0; i < len(verticalCuts) - 1; i++ {
        maxY = max(maxY, verticalCuts[i + 1]-verticalCuts[i])
    }
    
    return (maxX * maxY) % 1000000007
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
