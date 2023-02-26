package binary

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func getTestNode() *TreeNode {
	root := NewTreeNode(1)           //head node
	root.left = NewTreeNode(2)       //left subtree
	root.right = NewTreeNode(3)      //right subtree
	root.left.right = NewTreeNode(4) //right subtree of left subtree
	root.right.left = NewTreeNode(5) //left subtree of the left subtree of the right subtree
	root.left.left = NewTreeNode(6)
	root.right.right = NewTreeNode(7)

	return root
}

func TestTreeNode_GetNodeNum(t *testing.T) {
	root := getTestNode()
	require.Equal(t, 7, root.GetNodeNum())
}

func TestTreeNode_GetNodeDegree(t *testing.T) {
	root := getTestNode()
	require.Equal(t, 3, root.GetTreeDegree())
}

func TestTreeNode_PreOrder(t *testing.T) {
	const exp = "1 2 6 4 3 5 7"

	root := getTestNode()
	got := root.PreOrder()

	require.Equal(t, exp, got)
}

func TestTreeNodePostOrder(t *testing.T) {
	const exp = "6 4 2 5 7 3 1"

	root := getTestNode()
	got := root.PostOrder()

	require.Equal(t, exp, got)
}

func TestTreeNode_MidOrder(t *testing.T) {
	const exp = "6 2 4 1 5 3 7"

	root := getTestNode()
	got := root.MidOrder()

	require.Equal(t, exp, got)
}

func TestTreeNode_LayerOrder(t *testing.T) {
	const exp = "1 2 3 6 4 5 7 "

	root := getTestNode()
	got, err := root.LayerOrder()
	require.NoError(t, err)
	require.Equal(t, exp, got)
}
