package stack

import (
	"errors"
)

type stack struct {
	top      int
	elements []interface{}
}

func New() stack {
	return stack{top: -1}
}

func (s *stack) Push(item interface{}) {
	s.top++
	s.elements = append(s.elements, item)
}

func (s *stack) Pop() error {
	if !s.IsEmpty() {
		s.elements = s.elements[:s.top]
		s.top--
		return nil
	}
	return errors.New("stack is empty")
}

func (s stack) Peek() (interface{}, error) {
	if !s.IsEmpty() {
		return s.elements[s.top], nil
	}
	return nil, errors.New("stack is empty")
}

func (s stack) IsEmpty() bool {
	return s.top < 0
}

func (s stack) List() ([]interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	var list []interface{}
	for i := s.top; i >= 0; i-- {
		list = append(list, s.elements[i])
	}

	return list, nil
}
