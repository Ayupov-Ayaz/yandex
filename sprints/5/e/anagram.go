package main

// Гоша и Алла играют в игру «Удивительные деревья».
//Помогите ребятам определить, является ли дерево, которое им встретилось, деревом-анаграммой?
// Дерево называется анаграммой, если оно симметрично относительно своего центра.

type Node struct {
	value int
	left  *Node
	right *Node
}

func isOk(a, b bool) bool {
	return a == b
}

func (n *Node) compareNode(curr *Node) bool {
	if n == nil {
		return curr == nil
	} else if curr == nil {
		return false
	}

	if n.value != curr.value {
		return false
	}

	if !isOk(n.left != nil, curr.right != nil) {
		return false
	}

	if !isOk(n.right != nil, curr.left != nil) {
		return false
	}

	return true
}

func isAnagram(a, b *Node) bool {
	aLeft := a.left
	aRight := a.right
	bLeft := b.left
	bRight := b.right

	if !aLeft.compareNode(bRight) {
		return false
	}

	if !aRight.compareNode(bLeft) {
		return false
	}

	return true
}

func solution(left, right *Node) bool {
	if left == nil {
		return right == nil
	} else if right == nil {
		return false
	}

	if !isAnagram(left, right) {
		return false
	}

	if left.left != nil && !solution(left.left, right.right) {
		return false
	}

	if left.right != nil && !solution(left.right, right.left) {
		return false
	}

	return true
}

func Solution(root *Node) bool {
	return solution(root.left, root.right)
}
