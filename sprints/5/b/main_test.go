package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestA(t *testing.T) {
	node1 := Node{1, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{8, nil, nil}
	node5 := Node{5, &node3, &node4}
	require.True(t, Solution(&node5))

	node2.value = 5
	require.False(t, Solution(&node5))
}

func TestB(t *testing.T) {
	node7 := NewNode(19, nil, nil)
	node6 := NewNode(9, nil, node7)
	node5 := NewNode(6, nil, nil)
	node4 := NewNode(4, nil, nil)
	node3 := NewNode(1, nil, nil)
	node2 := NewNode(8, node5, node6)
	node1 := NewNode(3, node3, node4)
	node0 := NewNode(5, node1, node2)

	require.True(t, Solution(node0))
}
