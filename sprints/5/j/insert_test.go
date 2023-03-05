package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertA(t *testing.T) {
	node1 := Node{7, nil, nil}
	node2 := Node{8, &node1, nil}
	node3 := Node{7, nil, &node2}
	newHead := insert(&node3, 6)

	require.Equal(t, &node3, newHead)
	require.Equal(t, 6, newHead.left.value)
}
