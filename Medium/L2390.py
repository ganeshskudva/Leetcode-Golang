package Medium

func removeStars(s string) string {
	st := []rune{}

	for _, c := range s {
		if c != '*' {
			st = append(st, c)
		} else if len(st) > 0 {
			st = st[:len(st)-1]
		}
	}

	return string(st)
}
