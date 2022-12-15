package array

import (
	"github.com/ayupov-ayaz/yandex/queue"
	"github.com/stretchr/testify/require"
	"testing"
)

const maxSize = 3

func checkQueue(t *testing.T, q *Queue, size, index, value int) {
	require.Equal(t, size, q.Size())
	require.Equal(t, value, q.values[index])
}

func TestQueue_NewQueue(t *testing.T) {
	q := NewQueue(maxSize)
	require.Equal(t, maxSize, q.maxSize)
	require.Equal(t, maxSize, len(q.values))
	require.Zero(t, q.size)
}

func TestQueue_Push(t *testing.T) {
	q := NewQueue(maxSize)
	for i := 0; i < maxSize; i++ {
		v := i + 134
		require.NoError(t, q.PushBack(v))
		checkQueue(t, q, i+1, i, v)
	}

	err := q.PushBack(1)
	require.ErrorIs(t, err, queue.ErrQueueIsFull)
}

func TestQueue_PopFront(t *testing.T) {
	q := NewQueue(maxSize)
	v, err := q.PopFront()
	require.ErrorIs(t, err, queue.ErrQueueIsEmpty)
	require.Zero(t, v)

	for i := 0; i < maxSize; i++ {
		require.NoError(t, q.PushBack(i+1))
	}

	for i := 0; i < maxSize; i++ {
		v, err = q.PopFront()
		require.NoError(t, err)
		require.Equal(t, i+1, v)
	}
}

func TestQueue_PushAfterPopFront(t *testing.T) {
	q := NewQueue(maxSize)

	for i := 0; i < maxSize; i++ {
		require.NoError(t, q.PushBack(i+1))
	}

	v, err := q.PopFront()
	require.NoError(t, err)
	require.Equal(t, 1, v)
	checkQueue(t, q, 2, 0, 0)
	checkQueue(t, q, 2, 1, 2)
	checkQueue(t, q, 2, 2, 3)

	require.NoError(t, q.PushBack(4))
	checkQueue(t, q, 3, 0, 4)
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue(maxSize)
	v, err := q.Peek()
	require.ErrorIs(t, err, queue.ErrQueueIsEmpty)
	require.Zero(t, v)

	for i := 0; i < maxSize; i++ {
		_ = q.PushBack(i)
	}

	for i := 0; i < maxSize; i++ {
		v, err = q.Peek()
		require.Equal(t, 0, v)
		require.NoError(t, err)
		require.Equal(t, maxSize, q.Size())
	}
}

func TestQueue_PushFront(t *testing.T) {
	q := NewQueue(maxSize)

	for i := 0; i < maxSize; i++ {
		require.NoError(t, q.PushBack(i+1))
	}

	_, err := q.PopFront()
	require.NoError(t, err)
	checkQueue(t, q, 2, 0, 0)

	require.NoError(t, q.PushFront(4))
	checkQueue(t, q, 3, 0, 4)
}

func Test1(t *testing.T) {
	q := NewQueue(4)
	require.NoError(t, q.PushFront(861))
	require.NoError(t, q.PushFront(-819))
	v, err := q.PopBack()
	require.NoError(t, err)
	require.Equal(t, 861, v)
	v, err = q.PopBack()
	require.NoError(t, err)
	require.Equal(t, -819, v)
}
