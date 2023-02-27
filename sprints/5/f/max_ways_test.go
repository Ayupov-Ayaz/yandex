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
	require.Equal(t, 3, Solution(&node5))
}

func TestB(t *testing.T) {
	node5 := &Node{2, nil, nil}
	node4 := &Node{2, node5, nil}
	node3 := &Node{2, nil, nil}
	node2 := &Node{2, node3, node4}
	node1 := &Node{1, nil, nil}
	node0 := &Node{1, node1, node2}

	require.Equal(t, 4, Solution(node0))
}
