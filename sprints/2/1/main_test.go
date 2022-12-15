package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_process(t *testing.T) {
	const (
		maxSize = 6
	)

	q := NewQueue(maxSize)
	require.NoError(t, q.PushFront(-201))
	require.Equal(t, maxSize-1, q.headIndex)
	require.Equal(t, 1, q.size)
	require.Equal(t, -201, q.values[0])
	//
	require.NoError(t, q.PushBack(959))
	require.Equal(t, maxSize-1, q.headIndex)
	require.Equal(t, 2, q.size)
	require.Equal(t, 2, q.tailIndex)
	require.Equal(t, 959, q.values[1])
	//
	require.NoError(t, q.PushBack(102))
	require.Equal(t, maxSize-1, q.headIndex)
	require.Equal(t, 3, q.size)
	require.Equal(t, 3, q.tailIndex)
	require.Equal(t, 102, q.values[2])
	//
	v, err := q.PopFront()
	require.NoError(t, err)
	require.Equal(t, -201, v)
	//
	v, err = q.PopBack()
	require.NoError(t, err)
	require.Equal(t, 102, v)
}
