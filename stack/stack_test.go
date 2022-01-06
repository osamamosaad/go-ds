package stack

import (
	"reflect"
	"testing"
)

func TestConstructorReturnTypeOfStack(t *testing.T) {
	stackConstructor := New()
	var stackType stack

	if reflect.TypeOf(stackConstructor) == reflect.TypeOf(stackType) {
		return
	}

	t.Errorf("stack constructor doesn't return the stack type. Expected: %T, Returned: %T", stackType, stackConstructor)
}

func TestEmptyStateWhileInitializing(t *testing.T) {
	stack := New()

	if stack.IsEmpty() == true {
		return
	}

	t.Errorf("Stack is not an empty while initializing.")
}

func TestEmptyStateAfterPopOutAllStackValues(t *testing.T) {
	data := []int{10, 50, 60, 70}
	stack := New()
	for _, v := range data {
		stack.Push(v)
	}

	for i := 0; i < len(data); i++ {
		stack.Pop()
	}

	if stack.IsEmpty() == true {
		return
	}

	t.Errorf("stack isn't empty after poping out all the data from.")
}

func TestStackValuesOrder(t *testing.T) {
	data := []int{10, 50, 60, 70}
	stack := New()
	for _, v := range data {
		stack.Push(v)
	}

	for i := len(data) - 1; i >= 0; i-- {
		v, _ := stack.Peek()
		if v.(int) != data[i] {
			t.Errorf("this value %d is not the Peek as expected, returned %d.", data[i], v.(int))
			return
		}
		stack.Pop()
	}

}

func TestPeekAfterPopAndPushRandomly(t *testing.T) {
	data := []int{10, 50, 60, 70}
	expected := 80
	stack := New()
	for _, v := range data {
		stack.Push(v)
	}
	stack.Pop()

	stack.Push(100)
	stack.Pop()

	stack.Push(200)
	stack.Push(300)
	stack.Pop()

	stack.Push(expected)
	v, _ := stack.Peek()

	if v.(int) == expected {
		return
	}

	t.Errorf("The Peek is %d not as expected after using pop and push randomly. the expected is %d", v, expected)
}
