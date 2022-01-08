package stack

import (
	"errors"
)

type stack struct {
	top      int
	elements []interface{}
}

// New to initalize stack and set the top with default value -1 which means empty.
func New() stack {
	return stack{top: -1}
}

// Push to add an item into the stack
func (s *stack) Push(item interface{}) {
	s.top++
	s.elements = append(s.elements, item)
}

// Pop to remove an item from the stack
// returns error if the stack is empty
func (s *stack) Pop() error {
	if !s.IsEmpty() {
		s.elements = s.elements[:s.top]
		s.top--
		return nil
	}
	return errors.New("stack is empty")
}

// Peek get the stack peek
// returns nil and error when the stack is empty
func (s stack) Peek() (interface{}, error) {
	if !s.IsEmpty() {
		return s.elements[s.top], nil
	}
	return nil, errors.New("stack is empty")
}

// IsEmpty to check whether the stack is empty
func (s stack) IsEmpty() bool {
	return s.top < 0
}

// List to retrun the stack data as an array
// returns nil, with error if the stack is empty
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
