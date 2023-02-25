package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestA(t *testing.T) {
	node1 := Node{3, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{4, nil, nil}
	node4 := Node{3, nil, nil}
	node5 := Node{2, &node1, &node2}
	node6 := Node{2, &node3, &node4}
	node7 := Node{1, &node5, &node6}

	require.True(t, Solution(&node7))
}

func TestB(t *testing.T) {
	node1 := &Node{3, nil, nil}
	require.True(t, Solution(node1))
}
