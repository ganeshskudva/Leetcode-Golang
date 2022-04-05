package Medium

func maxArea(height []int) int {
	l, r, mini, maxi := 0, len(height)-1, math.MaxInt32, math.MinInt32

	for l <= r {
		mini = min(height[l], height[r])
		maxi = max(maxi, mini*(r-l))
		if l <= r && height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxi
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
