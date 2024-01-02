package chapter_14

type Stack []*Node

func height(s *Stack) int {
	return len(*s)
}

func (s *Stack) Push(item *Node) {
	*s = append(*s, item)
}

func (s *Stack) Read() *Node {
	stackLen := height(s)
	if stackLen == 0 {
		return nil
	}
	index := stackLen - 1
	return (*s)[index]
}

func (s *Stack) Pop() *Node {
	read := s.Read()
	*s = (*s)[:height(s)-1]

	return read
}
