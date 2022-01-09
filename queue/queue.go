package queue

import (
	"errors"
	"fmt"
)

// QueueInterface is queue interface
type QueueInterface interface {
	Enqueue(interface{})
	Dnqueue() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	Display()
}

// Queue type to handle queue data-structure
type Queue struct {
	items []interface{}
}

// Enqueue add item in the queue
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue remove item from queue
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	removedItem := q.items[0]
	q.items = q.items[1:]

	return removedItem, nil
}

// Peek get the queue peek
// returns nil and error when the queue is empty
func (q Queue) Peek() (interface{}, error) {
	if !q.IsEmpty() {
		return q.items[0], nil
	}
	return nil, errors.New("queue is empty")
}

// IsEmpty to check whether the queue is empty
func (q Queue) IsEmpty() bool {
	return len(q.items) < 1
}

// Display queue values
func (q Queue) Display() {
	if q.IsEmpty() {
		return
	}

	for i := 0; i < len(q.items); i++ {
		fmt.Println(q.items[i])
	}
}
