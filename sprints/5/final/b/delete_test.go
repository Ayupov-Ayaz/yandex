package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_remove(t *testing.T) {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}
	newHead := remove(&node7, 10)
	require.Equal(t, 5, newHead.value)
	require.NotEqual(t, &node5, newHead.right)
	require.NotEqual(t, 8, newHead.right.value)
}

func Test_remove1(t *testing.T) {
	root := Node{6, nil, nil}
	newHead := remove(&root, 6)
	require.Nil(t, newHead)
}
