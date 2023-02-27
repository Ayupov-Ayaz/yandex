package main

import (
	"strconv"
	"strings"
)

// H. Числовые пути
// Вася и его друзья решили поиграть в игру.
// Дано дерево, в узлах которого записаны цифры от 0 до 9. Т
// аким образом, каждый путь от корня до листа содержит число,
// получившееся конкатенацией цифр пути (склеиванием цифр пути в одно число). Нужно найти сумму всех таких чисел в дереве.
// Гарантируется, что ответ не превосходит 20 000.

type Node struct {
	value int
	left  *Node
	right *Node
}

func N(value int, left, right *Node) *Node {
	return &Node{
		value: value,
		left:  left,
		right: right,
	}
}

func (n *Node) PreOrder(prefix string) string {
	if n == nil {
		return ""
	}

	prefix += strconv.Itoa(n.value)
	resp := prefix

	if n.left != nil {
		resp = n.left.PreOrder(prefix)
	}

	if n.right != nil {
		right := n.right.PreOrder(prefix)

		if n.left != nil {
			resp += " " + right
		} else {
			resp = right
		}
	}

	return resp
}

func Solution(root *Node) (sum int) {
	numbers := strings.Split(root.PreOrder(""), " ")
	for _, number := range numbers {
		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		sum += n
	}

	return sum
}
