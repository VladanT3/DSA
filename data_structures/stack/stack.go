package stack

type Stack struct {
	elements []any
}

func (s *Stack) Push(item any) {
	s.elements = append(s.elements, item)
}

func (s *Stack) Pop() any {
	item := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return item
}

func (s *Stack) Peek() any {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) IsEmpty() bool {
	if len(s.elements) == 0 || s.elements == nil {
		return true
	}

	return false
}
