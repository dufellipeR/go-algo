package chapter_9

import "algorithms/chapter_14"

type Stack []*chapter_14.Node

func height(s *Stack) int {
	return len(*s)
}

func (s *Stack) Push(item *chapter_14.Node) {
	*s = append(*s, item)
}

func (s *Stack) Read() *chapter_14.Node {
	stackLen := height(s)
	if stackLen == 0 {
		return nil
	}
	index := stackLen - 1
	return (*s)[index]
}

func (s *Stack) Pop() *chapter_14.Node {
	read := s.Read()
	*s = (*s)[:height(s)-1]

	return read
}
