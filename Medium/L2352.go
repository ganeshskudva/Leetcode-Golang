package Medium

func equalPairs(grid [][]int) int {
	pair, tmp, row := 0, 0, 0

	for tmp <= len(grid)-1 {
		mp := make(map[int]int)
		for j := 0; j < len(grid); j++ {
			mp[j] = grid[row][j]
		}
		for i := 0; i < len(grid); i++ {
			curr := 0
			for k := 0; k < len(grid); k++ {
				if v, ok := mp[k]; ok {
					if v != grid[k][i] {
						curr = 0
						break
					} else {
						curr = 1
					}
				}
			}
			pair += curr
		}
		row, tmp = row+1, tmp+1
	}

	return pair
}
