package Easy

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	var exists struct{}
	uniqueS := make(map[byte]struct{})
	for i := range s {
		uniqueS[s[i]] = exists
	}
	if len(uniqueS) < len(s) && s == goal {
		return true
	}

	var diff []byte
	for i := 0; i < len(s) && len(diff) <= 4; i++ {
		if s[i] != goal[i] {
			diff = append(diff, s[i])
			diff = append(diff, goal[i])
		}
	}

	return len(diff) == 4 && diff[0] == diff[3] && diff[1] == diff[2]
}
