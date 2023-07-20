package Medium

func asteroidCollision(asteroids []int) []int {
	var res []int
	if len(asteroids) == 0 {
		return res
	}
	
	for _, a := range asteroids {
		if a > 0 {
			res = append(res, a)
		} else {
			for len(res) > 0 && 
				res[len(res) - 1] > 0 &&
				-a > res[len(res) - 1] {
				res = res[:len(res) - 1]
			}
			if len(res) == 0 || res[len(res) - 1] < 0 {
				res = append(res, a)
			} else if res[len(res) - 1] == -a {
				res = res[:len(res) - 1]
			}
		}
	}
	
	return res
}
