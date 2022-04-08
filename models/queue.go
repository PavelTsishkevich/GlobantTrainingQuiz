package models

import "errors"

type Queue struct {
	elements []int
}

func NewQueue() *Queue {
	return &Queue{make([]int, 0)}
}

func (queue *Queue) Push(element int) {
	queue.elements = append(queue.elements, element)
}

func (queue *Queue) Pull() (int, error) {
	length := len(queue.elements)
	if length == 0 {
		return 0, errors.New("queue is empty")
	}
	result := queue.elements[0]
	queue.elements = queue.elements[1:]
	return result, nil
}

func (queue *Queue) Size() int {
	return len(queue.elements)
}
