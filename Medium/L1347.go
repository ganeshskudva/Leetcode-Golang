package Medium

func minSteps(s string, t string) int {
	cnt := make([]int, 26)

	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}

	steps := 0
	for _, count := range cnt {
		if count > 0 {
			steps += count
		}
	}

	return steps
}
