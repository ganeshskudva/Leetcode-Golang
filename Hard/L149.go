package Hard

func maxPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	
	res := 0
	for i := 0; i < len(points); i++ { 
		mp, dup, maxi := make(map[string]int), 0, 0
		for j := i + 1; j < len(points); j++ {
			deltaX := points[j][0] - points[i][0]
			deltaY := points[j][1] - points[i][1]
			
			if deltaX == 0 && deltaY == 0 {
				dup++
				continue
			}
			
			gcd := greatestCommonDiv(deltaX, deltaY)
			dX, dY := deltaX / gcd, deltaY / gcd
			
			key := fmt.Sprintf("%d-%d", dX, dY)
			mp[key]++
			maxi = max(maxi, mp[key])
		}
		res = max(res, maxi + dup + 1)
	}
	
	return res
}

func greatestCommonDiv(a, b int) int {
	if b == 0 {
		return a
	}
	
	return greatestCommonDiv(b, a % b)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
