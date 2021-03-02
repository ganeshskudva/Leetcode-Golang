package Easy

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var idx int
	cnt := 0

	if ruleKey == "type" {
		idx = 0
	} else if ruleKey == "color" {
		idx = 1
	} else {
		idx = 2
	}

	for _, itm:= range items {
		if itm[idx] == ruleValue {
			cnt++
		}
	}

	return cnt
}
