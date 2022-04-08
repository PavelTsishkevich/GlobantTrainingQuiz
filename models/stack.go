package models

import "errors"

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{make([]int, 0)}
}

func (stack *Stack) Push(element int) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack) Pull() (int, error) {
	length := len(stack.elements)
	if length == 0 {
		return 0, errors.New("stack is empty")
	}
	result := stack.elements[(length - 1)]
	stack.elements = stack.elements[:(length - 1)]
	return result, nil
}

func (stack *Stack) Size() int {
	return len(stack.elements)
}
