package Easy

type Stack struct {
	arr []byte
}

func NewStack() *Stack {
	return &Stack{arr: make([]byte, 0)}
}

func (s *Stack) Push(b byte) {
	s.arr = append(s.arr, b)
}

func (s *Stack) Pop() byte {
	if s.IsEmpty() {
		return 0
	}

	ret := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return ret
}

func (s *Stack) Top() byte {
	if s.IsEmpty() {
		return 0
	}

	return s.arr[len(s.arr)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	st := NewStack()

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			st.Push(')')
		case '{':
			st.Push('}')
		case '[':
			st.Push(']')
		default:
			if st.IsEmpty() || st.Pop() != s[i] {
				return false
			}
		}
	}

	return st.IsEmpty()
}
