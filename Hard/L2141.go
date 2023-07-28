package Hard

func maxRunTime(n int, batteries []int) int64 {
	var (
		canFit                    func(k int64) bool
		lower, upper, res, batSum int64
	)

	for _, b := range batteries {
		batSum += int64(b)
	}
	lower, upper, res = 0, batSum/int64(n), -1
	
	canFit = func(k int64) bool {
		var (
			currBatSum, target int64
		)
		
		target = int64(n) * k
		for _, b := range batteries {
			if int64(b) < k {
				currBatSum += int64(b)
			} else {
				currBatSum += k
			}
			
			if currBatSum >= target {
				return true
			}
		}
		
		return currBatSum >= target
	}
	
	for lower <= upper {
		var mid int64
		mid = lower + (upper - lower) / 2
		
		if canFit(mid) {
			res, lower = mid, mid + 1
		} else {
			upper = mid - 1
		}
	}

	return res
}
