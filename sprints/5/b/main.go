package main

//Гоша понял, что такое дерево поиска, и захотел написать функцию, которая определяет,
//является ли заданное дерево деревом поиска. Значения в левом поддереве должны быть строго меньше,
//в правом —- строго больше значения в узле.

const (
	initWay int8 = iota
	leftWay
	rightWay
)

func NewNode(value int, left, right *Node) *Node {
	return &Node{
		value: value,
		left:  left,
		right: right,
	}
}

func (n *Node) checkLeftForLeftWay(parent *Node) bool {
	a := parent == nil || parent.value > n.value
	return n.left.value < n.value && a
}

func (n *Node) checkLeftForRightWay(parent *Node) bool {
	a := parent == nil || parent.value < n.value
	return n.left.value < n.value && a
}

func (n *Node) haveChild() bool {
	return n.left != nil || n.right != nil
}

func (n *Node) checkRightForRightWay(parent *Node) bool {
	a := parent == nil || (n.right.value > parent.value)
	return n.right.value > n.value && a
}

func (n *Node) checkRightForLeftWay(parent *Node) bool {
	a := parent == nil || (n.right.value < parent.value)
	return n.right.value > n.value && a
}

func (n *Node) IsValidForRightWay(parent *Node) bool {
	if n.left != nil && !n.checkLeftForRightWay(parent) {
		return false
	}

	if n.right != nil && !n.checkRightForRightWay(parent) {
		return false
	}

	return true
}

func (n *Node) IsValidForLeftWay(parent *Node) bool {
	if n.left != nil && !n.checkLeftForLeftWay(parent) {
		return false
	}

	if n.right != nil && !n.checkRightForLeftWay(parent) {
		return false
	}

	return true
}

func (n *Node) IsValid(parent *Node, way int8) bool {
	switch way {
	case leftWay:
		if !n.IsValidForLeftWay(parent) {
			return false
		}
	case rightWay, initWay:
		if !n.IsValidForRightWay(parent) {
			return false
		}
	}

	return true
}

func solution(root, parent *Node, way int8) bool {
	if !root.IsValid(parent, way) {
		return false
	}

	nextLeftWay := way
	nextRightWay := way

	if way == initWay {
		nextRightWay = rightWay
		nextLeftWay = leftWay
	}

	if root.left != nil && root.left.haveChild() && !solution(root.left, root, nextLeftWay) {
		return false
	}

	if root.right != nil && root.right.haveChild() && !solution(root.right, root, nextRightWay) {
		return false
	}

	return true
}

func Solution(root *Node) bool {
	return solution(root, nil, 0)
}

type Node struct {
	value int
	left  *Node
	right *Node
}
