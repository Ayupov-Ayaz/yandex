package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, key int) *Node {
	curr := root

	node := &Node{
		value: key,
	}

	for curr != nil {
		if curr.value > key {
			if curr.left != nil {
				curr = curr.left
			} else {
				curr.left = node
				break
			}
		} else if curr.right != nil {
			curr = curr.right
		} else {
			curr.right = node
			break
		}
	}

	return root
}
