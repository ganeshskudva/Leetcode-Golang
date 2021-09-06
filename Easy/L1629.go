package Easy

func slowestKey(releaseTimes []int, keysPressed string) byte {
	max, key := releaseTimes[0], keysPressed[0]

	for i := 1; i < len(releaseTimes); i++ {
		relTime := releaseTimes[i] - releaseTimes[i-1]
		if relTime >= max {
			if relTime == max {
				if keysPressed[i] > key {
					key = keysPressed[i]
				}
			} else {
				max = relTime
				key = keysPressed[i]
			}
		}
	}

	return key
}
