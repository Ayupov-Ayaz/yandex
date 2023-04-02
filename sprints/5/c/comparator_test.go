package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestA(t *testing.T) {
	node1 := Node{1, nil, nil}
	node2 := Node{2, nil, nil}
	node3 := Node{3, &node1, &node2}

	node4 := Node{1, nil, nil}
	node5 := Node{2, nil, nil}
	node6 := Node{3, &node4, &node5}

	require.True(t, Solution(&node3, &node6))
}

func TestB(t *testing.T) {
	a1 := &Node{1, nil, nil}
	b1 := &Node{0, nil, nil}

	require.False(t, Solution(a1, b1))
}
