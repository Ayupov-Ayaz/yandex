package queue

import "errors"

var (
	ErrQueueIsFull  = errors.New("queue is full")
	ErrQueueIsEmpty = errors.New("queue is empty")
)

type Queue interface {
	Push(v int) error   // add element to tail queue
	Pop() (int, error)  // get element from head queue with remove
	Size() int          // get queue size
	Peek() (int, error) // get element from head queue without remove
}
