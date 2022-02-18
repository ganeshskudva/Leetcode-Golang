package Medium

type Stack struct {
	arr []byte
}

func NewStack() *Stack {
	return &Stack{
		arr: make([]byte, 0),
	}
}

func (s *Stack) Push(c byte) {
	s.arr = append(s.arr, c)
}

func (s *Stack) Pop() byte {
	if s.IsEmpty() {
		return 0
	}

	ret := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return ret
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) <= 0
}

func (s *Stack) Peek() byte {
	if s.IsEmpty() {
		return 0
	}

	return s.arr[len(s.arr)-1]
}

func removeKdigits(num string, k int) string {
	if k == 0 {
		return num
	}
	if len(num) == k {
		return "0"
	}

	st := NewStack()
	index := 0

	for index < len(num) {
		for k > 0 && !st.IsEmpty() && st.Peek() > num[index] {
			st.Pop()
			k--
		}
		st.Push(num[index])
		index = index + 1
	}

	for k > 0 {
		k--
		st.Pop()
	}

	var sb strings.Builder
	for !st.IsEmpty() {
		sb.WriteByte(st.Pop())
	}
	smallest := reverse(sb.String())
	for len(smallest) > 1 && smallest[0] == '0' {
		smallest = smallest[1:]
	}

	return smallest
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
