package Medium

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	xDist, yDist := abs(sx-fx), abs(sy-fy)

	if xDist == 0 && yDist == 0 {
		return t != 1
	}

	return xDist <= t && yDist <= t
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
