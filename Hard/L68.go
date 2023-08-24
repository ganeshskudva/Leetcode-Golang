package Hard

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	left := 0

	for left < len(words) {
		right := findRight(words, left, maxWidth)
		res = append(res, justify(words, left, right, maxWidth))
		left = right + 1
	}
	
	return res
}

func findRight(words []string, left, maxWidth int) int {
	right := left
	sum := len(words[right])
	right++

	for right < len(words) && (sum + 1 + len(words[right])) <= maxWidth {
		sum += 1 + len(words[right])
		right++
	}

	return right - 1
}

func justify(words []string, left, right, maxWidth int) string {
	if right - left == 0 {
		return padResult(words[left], maxWidth)
	}
	
	lastLine := isLastLine(words, right)
	numSpaces := right - left
	totalSpace := maxWidth - wordsLength(words, left, right)
	
	space, remainder := " ", 0
	if !lastLine {
		space = blank(totalSpace / numSpaces)
		remainder = totalSpace % numSpaces
	}
	
	var sb strings.Builder
	for i := left; i <= right; i++ {
		sb.WriteString(words[i])
		sb.WriteString(space)
		
		if remainder > 0 {
			sb.WriteString(" ")
		}
		remainder--
	}
	
	return padResult(strings.Trim(sb.String(), " "), maxWidth)
}

func isLastLine(words []string, right int) bool {
	return right == len(words) - 1
}

func wordsLength(words []string, left, right int) int {
	wordsLen := 0

	for i := left; i <= right; i++ {
		wordsLen += len(words[i])
	}

	return wordsLen
}

func padResult(result string, maxWidth int) string {
	var sb strings.Builder
	
	sb.WriteString(result)
	sb.WriteString(blank(maxWidth - len(result)))
	
	return sb.String()
}

func blank(length int) string {
	var sb strings.Builder
	
	for i := 0; i < length; i++ {
		sb.WriteString(" ")
	}
	
	return sb.String()
}
