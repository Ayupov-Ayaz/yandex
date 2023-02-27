package main

//Алла хочет побывать на разных островах архипелага Алгосы.
//Она составила карту. Карта представлена в виде дерева: корень обозначает центр архипелага,
//узлы –— другие острова. А листья —– это дальние острова, на которые Алла хочет попасть.
//Помогите Алле определить максимальное число островов, через которые ей нужно пройти для
//совершения одной поездки от стартового острова до места назначения, включая начальный и конечный пункты.

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) maxWays() int {
	if n == nil {
		return 0
	}

	var leftMax, rightMax int

	if n.left != nil {
		leftMax = n.left.maxWays()
	}

	if n.right != nil {
		rightMax = n.right.maxWays()
	}

	if leftMax > rightMax {
		return leftMax + 1
	}

	return rightMax + 1
}

func Solution(root *Node) (max int) {
	return root.maxWays()
}
