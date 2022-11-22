package Medium

var (
	exists struct{}
	dirs   = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func nearestExit(maze [][]byte, entrance []int) int {
	vis := make(map[string]struct{})

	var q [][]int
	q = append(q, entrance)
	vis[key(entrance[0], entrance[1])] = exists

	cnt := 0
	for len(q) > 0 {
		cnt++
		size := len(q)
		for i := 0; i < size; i++ {
			curr := q[0]
			q = q[1:]

			for _, d := range dirs {
				x, y := curr[0]+d[0], curr[1]+d[1]

				if !isValid(maze, x, y) {
					continue
				}
				if _, ok := vis[key(x, y)]; ok {
					continue
				}

				if isExit(maze, x, y) {
					return cnt
				}

				vis[key(x, y)] = exists
				q = append(q, []int{x, y})
			}
		}
	}

	return -1
}

func isValid(maze [][]byte, x, y int) bool {
	m, n := len(maze), len(maze[0])

	return x >= 0 && x < m && y >= 0 && y < n && maze[x][y] == '.'
}

func isExit(maze [][]byte, x, y int) bool {
	m, n := len(maze), len(maze[0])

	return x == 0 || y == 0 || x == m-1 || y == n-1
}

func key(x, y int) string {
	return fmt.Sprintf("%dx%d", x, y)
}
