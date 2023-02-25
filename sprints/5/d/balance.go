package main

import "math"

type Node struct {
	value int
	left  *Node
	right *Node
}

func calculate(curr *Node, next func(n *Node) *Node) (count int) {
	for curr != nil {
		count++
		curr = next(curr)
	}

	return count
}

func (n *Node) calculateLeft() (count int) {
	return calculate(n.left, func(n *Node) *Node {
		return n.left
	})
}

func (n *Node) calculateRight() (count int) {
	return calculate(n.right, func(n *Node) *Node {
		return n.right
	})
}

func checkBalance(left, right int) bool {
	return math.Abs(float64(left-right)) <= 1
}

func checkBalanceArr(arr [4]int) bool {
	var (
		i, j int
	)
	for k := 0; k < 3; k++ {
		switch k {
		case 0:
			i, j = 0, 1
		case 1:
			i, j = 0, 3
		case 2:
			i, j = 0, 2
		}

		if !checkBalance(arr[i], arr[j]) {
			return false
		}
	}

	return true
}

func Solution(root *Node) bool {
	data := [4]int{}
	data[0] = root.calculateLeft()
	data[1] = root.calculateRight()
	if root.left != nil {
		data[2] = root.left.calculateRight() + 1
	}
	if root.right != nil {
		data[3] = root.right.calculateLeft() + 1
	}

	return checkBalanceArr(data)
}
