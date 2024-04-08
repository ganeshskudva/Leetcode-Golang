package Easy

func countStudents(students []int, sandwiches []int) int {
	ones, zeroes := 0, 0

	for _, student := range students {
		if student == 0 {
			zeroes++
		} else {
			ones++
		}
	}

	for _, sandwich := range sandwiches {
		if sandwich == 0 {
			if zeroes == 0 {
				return ones
			}
			zeroes--
		} else {
			if ones == 0 {
				return zeroes
			}
			ones--
		}
	}

	return 0
}
