package Easy

func tictactoe(moves [][]int) string {
	n := 3
	rows, cols := make([]int, n), make([]int, n)
	d1, d2, curr := 0, 0, 1

	for _, currMove := range moves {
		rows[currMove[0]] += curr
		cols[currMove[1]] += curr

		if currMove[0] == currMove[1] {
			d1 += curr
		}

		if currMove[0]+currMove[1] == n-1 {
			d2 += curr
		}

		if abs(rows[currMove[0]]) == n ||
			abs(cols[currMove[1]]) == n ||
			abs(d1) == n ||
			abs(d2) == n {
			if curr == 1 {
				return "A"
			}
			return "B"
		}

		// 1 is 'A', -1 is 'B'
		curr *= -1
	}

	if len(moves) < 9 {
		return "Pending"
	}

	return "Draw"
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
