package main

//Гоше на день рождения подарили два дерева. Тимофей сказал, что они совершенно одинаковые. Но, по мнению Гоши, они отличаются.
//Помогите разрешить этот философский спор!

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) compare(curr *Node) bool {
	if n == nil {
		return curr == nil
	}

	if curr == nil {
		return false
	}

	if n.value != curr.value {
		return false
	}

	if n.left == nil && curr.left != nil || n.left != nil && curr.left == nil {
		return false
	}

	if n.right == nil && curr.right != nil || n.right != nil && curr.right == nil {
		return false
	}

	return true
}

func Solution(root1, root2 *Node) bool {
	if !root1.compare(root2) {
		return false
	}

	if root1 == nil {
		return true
	}

	if !Solution(root1.left, root2.left) {
		return false
	}

	if !Solution(root1.right, root2.right) {
		return false
	}

	return true
}
