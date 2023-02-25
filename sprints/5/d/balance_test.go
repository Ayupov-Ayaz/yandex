package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestA(t *testing.T) {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{10, nil, nil}
	node5 := Node{2, &node3, &node4}
	require.True(t, Solution(&node5))
}

func TestB(t *testing.T) {
	node2 := &Node{2, nil, nil}
	node1 := &Node{1, nil, nil}
	node0 := &Node{0, node1, node2}
	require.True(t, Solution(node0))
}

func TestC(t *testing.T) {
	node3 := &Node{4, nil, nil}
	node2 := &Node{2, nil, nil}
	node1 := &Node{1, node2, node3}
	node0 := &Node{0, nil, node1}
	require.False(t, Solution(node0))
}

func TestD(t *testing.T) {
	node8 := &Node{8, nil, nil}
	node7 := &Node{7, nil, nil}
	node6 := &Node{6, nil, nil}
	node5 := &Node{5, nil, nil}
	node4 := &Node{4, node7, node8}
	node3 := &Node{3, node5, node6}
	node2 := &Node{2, nil, node4}
	node1 := &Node{1, node3, nil}
	node0 := &Node{0, node1, node2}

	require.False(t, Solution(node0))
}
