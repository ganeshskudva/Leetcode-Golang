package Medium

type Direction int

const (
	DirectionUp      Direction = 0
	DirectionRight   Direction = 1
	DirectionDown    Direction = 2
	DirectionLeft    Direction = 3
	DirectionInvalid Direction = 4
)

func isRobotBounded(instructions string) bool {
	cur, dirs := []int{0, 0}, [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := DirectionUp

	for _, c := range instructions {
		if c == 'G' {
			cur[0] += dirs[dir][0]
			cur[1] += dirs[dir][1]
		} else if c == 'L' {
			dir = (dir + 3) % 4
		} else {
			dir = (dir + 1) % 4
		}
	}

	if cur[0] == 0 && cur[1] == 0 {
		return true
	}

	if dir == 0 && !(cur[0] == 0 && cur[1] == 0) {
		return false
	}

	return true
}
