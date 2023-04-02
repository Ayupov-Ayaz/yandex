package binary

import (
	"errors"
	"sync"
)

var (
	ErrQueueIsEmpty = errors.New("queue is empty")
)

type linkNode struct {
	next *linkNode
	data *TreeNode
}

func newLinkNode(data *TreeNode) *linkNode {
	return &linkNode{
		data: data,
	}
}

type LinkQueue struct {
	size int
	mu   *sync.Mutex
	head *linkNode
	tail *linkNode
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{
		mu: &sync.Mutex{},
	}
}

func (l *LinkQueue) Add(node *TreeNode) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.size++
	lNode := newLinkNode(node)
	if l.head != nil {
		l.tail.next = lNode
		l.tail = lNode
		return
	}

	l.head = lNode
	l.tail = lNode
}

func (l *LinkQueue) Remove() (*TreeNode, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		return nil, ErrQueueIsEmpty
	}

	l.size--

	resp := l.head
	l.head = l.head.next

	return resp.data, nil
}
