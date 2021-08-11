package Medium

func minFlipsMonoIncr(s string) int {
	oneCounter, flipCounter := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			oneCounter++
		} else {
			flipCounter++
		}

		if oneCounter < flipCounter {
			flipCounter = oneCounter
		}
	}

	return flipCounter
}
