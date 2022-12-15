package array

import "github.com/ayupov-ayaz/yandex/queue"

// Queue - двухсторонняя очередь с ограниченным размеров
type Queue struct {
	size      int
	maxSize   int
	headIndex int
	tailIndex int
	values    []int
}

func NewQueue(maxSize int) *Queue {
	return &Queue{
		maxSize: maxSize,
		values:  make([]int, maxSize),
	}
}

func (q *Queue) availableToPush() error {
	if q.size >= q.maxSize {
		return queue.ErrQueueIsFull
	}

	return nil
}

func (q *Queue) PushBack(v int) error {
	if err := q.availableToPush(); err != nil {
		return err
	}

	q.values[q.tailIndex] = v
	q.tailIndex++
	q.size++

	if q.tailIndex >= q.maxSize {
		q.tailIndex = 0
	}

	return nil
}

func (q *Queue) getNextIndex() int {
	if q.headIndex-1 < 0 {
		return q.maxSize - 1
	}

	return q.headIndex - 1
}

func (q *Queue) PushFront(v int) error {
	if err := q.availableToPush(); err != nil {
		return err
	}

	q.size++
	q.values[q.headIndex] = v

	q.headIndex = q.getNextIndex()

	return nil
}

func (q *Queue) shiftHead() {
	if q.headIndex < q.maxSize-1 {
		q.headIndex++
		return
	}

	q.headIndex = 0
}

func (q *Queue) availableToPop() error {
	if q.size == 0 {
		return queue.ErrQueueIsEmpty
	}

	return nil
}

func (q *Queue) PopFront() (int, error) {
	if err := q.availableToPop(); err != nil {
		return 0, err
	}

	//  this.size += 1;
	//    this.stack[this.head] = value;
	//    this.head = this.head - 1 < 0 ? this.maxN - 1 : this.head - 1;

	index := q.headIndex + 1
	if index < q.maxSize {
		index = q.headIndex + 1
	}

	resp := q.values[index]
	q.values[index] = 0
	q.headIndex = index
	q.size--

	return resp, nil
}

func (q *Queue) PopBack() (int, error) {
	if err := q.availableToPop(); err != nil {
		return 0, err
	}

	index := q.tailIndex - 1
	if q.tailIndex-1 < 0 {
		index = q.maxSize - 1
	}

	resp := q.values[index]
	q.values[index] = 0
	q.tailIndex = index
	q.size--

	return resp, nil
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Peek() (int, error) {
	if q.size == 0 {
		return 0, queue.ErrQueueIsEmpty
	}

	return q.values[q.headIndex], nil
}
