package Medium

func findMinArrowShots(points [][]int) int {
    if len(points) == 0 {
		return 0
	}
	
	sort.Slice(points, func(i, j int) bool {
		if points[i][1] < points[j][1]  {
			return true
		}
		
		return false
	})
	
	end, res := points[0][1], 1
	for _, p := range points {
		if p[0] <= end {
			continue
		}
		res++
		end = p[1]
	}
	return res
}
