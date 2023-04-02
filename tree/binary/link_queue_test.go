package binary

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLinkQueue_Add(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	queue := NewLinkQueue()

	for i, n := range numbers {
		queue.Add(NewTreeNode(n))
		require.Equal(t, i+1, queue.size)
	}

	var count int
	curr := queue.head
	for curr != nil {
		require.Equal(t, numbers[count], curr.data.value)
		count++
		curr = curr.next
	}

	require.Equal(t, len(numbers), count)
}

func TestLinkQueue_Remove(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	queue := NewLinkQueue()

	node, err := queue.Remove()
	require.ErrorIs(t, err, ErrQueueIsEmpty)
	require.Nil(t, node)

	for _, n := range numbers {
		queue.Add(NewTreeNode(n))
	}

	for i := 0; i < len(numbers); i++ {
		node, err = queue.Remove()
		require.NoError(t, err)
		require.Equal(t, numbers[i], node.value)
		require.Equal(t, len(numbers)-1-i, queue.size)
	}

	node, err = queue.Remove()
	require.ErrorIs(t, err, ErrQueueIsEmpty)
	require.Nil(t, node)
}
