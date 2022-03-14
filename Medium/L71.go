package Medium

type Stack struct {
	arr []string
}

func NewStack() *Stack {
	return &Stack{arr: make([]string, 0)}
}

func (s *Stack) Push(str string) {
	s.arr = append(s.arr, str)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}

	ret := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return ret
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}

	return s.arr[len(s.arr)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func simplifyPath(path string) string {
	var sb strings.Builder
	st := NewStack()

	for _, s := range strings.Split(path, "/") {
		if s == ".." {
			if !st.IsEmpty() {
				st.Pop()
			}
		} else if s != "" && s != "." {
			st.Push(s)
		}
	}

	if st.Size() == 0 {
		return "/"
	}

	for i := 0; i < st.Size(); i++ {
		sb.WriteString("/")
		sb.WriteString(st.arr[i])
	}

	return sb.String()
}
