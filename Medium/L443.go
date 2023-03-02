package Medium

func compress(chars []byte) int {
	index := 0

	updateChars := func(freq int) {
		arr := []byte(fmt.Sprintf("%d", freq))
		for i := range arr {
			chars[index] = arr[i]
			index++
		}
	}

	for i := 0; i < len(chars); {
		j := i
		for j < len(chars) && chars[i] == chars[j] {
			j++
		}
		freq := j - i
		chars[index] = chars[i]
		index++
		if freq > 1 && freq < 10 {
			chars[index] = byte(freq + '0')
			index++
		} else if freq >= 10 {
			updateChars(freq)
		}
		i = j
	}

	return index
}
