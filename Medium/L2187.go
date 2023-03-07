package Medium

func minimumTime(time []int, totalTrips int) int64 {
	var (
		lowTime  int64
		highTime int64
	)
	minTripTime := func() int64 {
		mini := math.MaxInt32

		for _, t := range time {
			mini = min(mini, t)
		}

		return int64(mini)
	}

	numTripsForGivenTime := func(givenTime int64) int64 {
		var totTrips int64
		for _, t := range time {
			totTrips += givenTime / int64(t)
		}

		return totTrips
	}
	lowTime, highTime = 1, minTripTime()*int64(totalTrips)

	for lowTime < highTime {
		mid := lowTime + (highTime-lowTime)/2

		if numTripsForGivenTime(mid) >= int64(totalTrips) {
			highTime = mid
		} else {
			lowTime = mid + 1
		}
	}

	return lowTime
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
