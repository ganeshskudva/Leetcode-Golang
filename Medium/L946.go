package Medium

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	popCounter := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) != 0 && popped[popCounter] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			popCounter++
			if popCounter == len(popped) {
				break
			}
		}
	}

	return len(stack) == 0
}
