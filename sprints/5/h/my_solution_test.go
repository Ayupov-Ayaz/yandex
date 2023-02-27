package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func getTestNode() *Node {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{2, nil, nil}
	node5 := Node{1, &node4, &node3}

	return &node5
}

func TestNode_PreOrder(t *testing.T) {
	const exp = "12 132 131"
	got := getTestNode().PreOrder("")
	require.Equal(t, exp, got)
}

func TestSolutionA(t *testing.T) {
	require.Equal(t, 275, Solution(getTestNode()))
}

func TestSolutionB(t *testing.T) {
	const (
		expSum = 59400
		expStr = "4132 4133 4647 46488"
	)

	node9 := N(8, nil, nil)
	node8 := N(8, node9, nil)
	node7 := N(7, nil, nil)
	node6 := N(3, nil, nil)
	node5 := N(2, nil, nil)
	node4 := N(4, node7, node8)
	node3 := N(3, node5, node6)
	node2 := N(6, nil, node4)
	node1 := N(1, node3, nil)
	node0 := N(4, node1, node2)

	gotStr := node0.PreOrder("")
	require.Equal(t, expStr, gotStr)

	gotSum := Solution(node0)
	require.Equal(t, expSum, gotSum)
}
