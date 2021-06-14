package Easy

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	sum, trucks := 0, 0
	for _, b := range boxTypes {
		if trucks < truckSize {
			if trucks+b[0] > truckSize {
				sum += (truckSize - trucks) * b[1]
				break
			} else {
				sum += b[0] * b[1]
				trucks += b[0]
			}
		}
	}

	return sum
}
