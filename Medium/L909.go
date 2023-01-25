package Medium

func snakesAndLadders(board [][]int) int {
	var q []int
	n, move := len(board), 0
	visited := make([]bool, n*(n+1))
	q = append(q, 1)

	for ; len(q) > 0; move++ {
		size := len(q)
		for i := 0; i < size; i++ {
			num := q[0]
			q = q[1:]

			if visited[num] {
				continue
			}
			visited[num] = true
			if num == n*n {
				return move
			}

			for idx := 1; idx <= 6 && num+idx <= (n*n); idx++ {
				next := num + idx
				val := getBoardValue(next, board)
				if val > 0 {
					next = val
				}
				if !visited[next] {
					q = append(q, next)
				}
			}
		}
	}

	return -1
}

func getBoardValue(num int, board [][]int) int {
	n := len(board)
	r := (num - 1) / n

	var x, y int
	x = n - 1 - r
	if r%2 == 0 {
		y = num - 1 - (r * n)
	} else {
		y = n + (r * n) - num
	}

	return board[x][y]
}
