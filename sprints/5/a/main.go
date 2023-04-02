package main

//
//Гоша повесил на стену гирлянду в виде бинарного дерева, в узлах которого находятся лампочки.
//У каждой лампочки есть своя яркость. Уровень яркости лампочки соответствует числу,
//расположенному в узле дерева.
//Помогите Гоше найти самую яркую лампочку в гирлянде, то есть такую, у которой яркость наибольшая.

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(data int, left, right *Node) *Node {
	return &Node{
		value: data,
		left:  left,
		right: right,
	}
}

func Solution(root *Node) int {
	max := root.value

	if root.left != nil {
		if leftVal := Solution(root.left); leftVal > max {
			max = leftVal
		}
	}

	if root.right != nil {
		if rightVal := Solution(root.right); rightVal > max {
			max = rightVal
		}
	}

	return max
}

func test() {
	node1 := NewNode(1, nil, nil)
	node2 := NewNode(-5, nil, nil)
	node3 := NewNode(3, node1, node2)
	node4 := NewNode(2, node3, nil)

	if Solution(node4) != 3 {
		panic("WA")
	}
}
