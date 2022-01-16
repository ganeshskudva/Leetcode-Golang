package Medium

func maxDistToClosest(seats []int) int {
	dist, j := math.MinInt32, -1

	if seats[0] == 1 {
		j = 0
	}

	for i := 1; i < len(seats); i++ {
		if seats[i] == 1 {
			if j == -1 {
				dist = max(dist, i)
			} else {
				dist = max(dist, (i-j)/2)
			}
			j = i
		}
	}

	if seats[len(seats)-1] == 0 {
		dist = max(dist, len(seats)-1-j)
	}

	return dist
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
