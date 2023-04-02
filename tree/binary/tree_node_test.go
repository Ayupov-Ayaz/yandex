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

func getBinaryTree() *TreeNode {
	root := NewTreeNode(8)
	root.left = NewTreeNode(4)
	root.left.left = NewTreeNode(2)
	root.left.right = NewTreeNode(6)
	root.left.left.left = NewTreeNode(1)
	root.left.left.right = NewTreeNode(3)

	root.right = NewTreeNode(11)
	root.right.left = NewTreeNode(10)
	root.right.left.left = NewTreeNode(9)
	root.right.right = NewTreeNode(14)
	root.right.right.left = NewTreeNode(13)
	root.right.right.right = NewTreeNode(15)

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

func TestTreeNode_Insert(t *testing.T) {
	values := []int{6, 2, 4, 1, 5, 3, 7}

	root := NewTreeNode(5)
	for _, v := range values {
		root.Insert(v)
	}

	require.Equal(t, 2, root.left.value)
	require.Equal(t, 1, root.left.left.value)
	require.Equal(t, 4, root.left.right.value)
	require.Equal(t, 3, root.left.right.left.value)
	require.Equal(t, 6, root.right.value)
	require.Equal(t, 5, root.right.left.value)
	require.Equal(t, 7, root.right.right.value)
}

func Test_DeleteLeaf(t *testing.T) {
	values := []int{6, 2, 4, 1, 5, 3, 7}

	root := NewTreeNode(5)
	for _, v := range values {
		root.Insert(v)
	}

	require.Equal(t, 3, root.left.right.left.value)
	Delete(root, 3)
	require.Nil(t, root.left.right.left)

	require.Equal(t, 1, root.left.left.value)
	Delete(root, 1)
	require.Nil(t, root.left.left)
}

func Test_DeleteBranchWith2Child(t *testing.T) {
	values := []int{6, 2, 4, 1, 5, 3, 7}

	root := NewTreeNode(4)
	for _, v := range values {
		root.Insert(v)
	}

	require.Equal(t, 2, root.left.value)
	require.Equal(t, 1, root.left.left.value)
	require.Equal(t, 3, root.left.right.value)
	Delete(root, 2)

	require.Equal(t, 1, root.left.value)
	require.Equal(t, 3, root.left.right.value)
}

func Test_DeleteBranchWith1Child_left(t *testing.T) {
	values := []int{6, 2, 4, 1, 5, 3, 7}

	root := NewTreeNode(5)
	for _, v := range values {
		root.Insert(v)
	}

	require.Equal(t, 4, root.left.right.value)
	require.Equal(t, 3, root.left.right.left.value)
	Delete(root, 4)

	require.Equal(t, 3, root.left.right.value)
	require.Nil(t, root.left.right.left)
	require.Nil(t, root.left.right.right)

}

func Test_DeleteBranchWith1Child_right(t *testing.T) {
	values := []int{6, 2, 4, 1, 5, 3, 7, 8}

	root := NewTreeNode(5)
	for _, v := range values {
		root.Insert(v)
	}

	require.Equal(t, 7, root.right.right.value)
	require.Equal(t, 8, root.right.right.right.value)
	Delete(root, 7)

	require.Equal(t, 8, root.right.right.value)
	require.Nil(t, root.right.right.left)
	require.Nil(t, root.right.right.right)

}

func Test_DeleteRoot(t *testing.T) {
	values := []int{6, 2, 4, 1, 5, 3, 7}

	root := NewTreeNode(5)
	for _, v := range values {
		root.Insert(v)
	}

	require.Equal(t, 5, root.value)
	Delete(root, 5)
	require.Equal(t, 4, root.value)
	require.Equal(t, 2, root.left.value)
	require.Equal(t, 1, root.left.left.value)
	require.Equal(t, 3, root.left.right.value)
}

func Test_DeleteRoot1(t *testing.T) {
	root := NewTreeNode(1)
	Delete(root, 1)
}
