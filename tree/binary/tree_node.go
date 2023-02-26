package binary

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

// GetNodeNum - вычисление количества узлов
func (t *TreeNode) GetNodeNum() int {
	if t == nil {
		return 0
	}

	return t.left.GetNodeNum() + t.right.GetNodeNum() + 1
}

func (t *TreeNode) valueToString() string {
	return strconv.Itoa(t.value)
}

// GetTreeDegree - вычисление глубины дерева
func (t *TreeNode) GetTreeDegree() (maxDegree int) {
	if t == nil {
		return 0
	}

	leftDegree := t.left.GetTreeDegree()
	rightDegree := t.right.GetTreeDegree()

	if leftDegree > rightDegree {
		maxDegree = leftDegree
	} else {
		maxDegree = rightDegree
	}

	return maxDegree + 1
}

// PreOrder - Обход дерева в прямом порядке
// head -> left -> right
func (t *TreeNode) PreOrder() string {
	if t == nil {
		return ""
	}

	var buff strings.Builder
	buff.WriteString(t.valueToString())
	left := t.left.PreOrder()
	right := t.right.PreOrder()

	if left != "" {
		buff.WriteString(" ")
		buff.WriteString(left)
	}

	if right != "" {
		buff.WriteString(" ")
		buff.WriteString(right)
	}

	return buff.String()
}

// PostOrder - обход дерева в обратном порядке
// left -> right -> head
func (t *TreeNode) PostOrder() string {
	if t == nil {
		return ""
	}

	left := t.left.PostOrder()
	right := t.right.PostOrder()

	var buff strings.Builder
	if left != "" {
		buff.WriteString(left)
		buff.WriteString(" ")
	}

	if right != "" {
		buff.WriteString(right)
		buff.WriteString(" ")
	}

	buff.WriteString(t.valueToString())

	return buff.String()
}

// MidOrder - не упорядоченный обход
// left -> head -> right
func (t *TreeNode) MidOrder() string {
	if t == nil {
		return ""
	}

	var buff strings.Builder
	left := t.left.MidOrder()
	right := t.right.MidOrder()

	if left != "" {
		buff.WriteString(left)
		buff.WriteString(" ")
	}

	buff.WriteString(t.valueToString())

	if right != "" {
		buff.WriteString(" ")
		buff.WriteString(right)
	}

	return buff.String()
}

// LayerOrder - обход порядка уровней
// root -> left -> right
// Принцип алгоритма:
// Помещаем корневой узел дерева в очередь, так как нам нужно определить очередь таблицы цепочки.
// Удаляем узел из очереди,
// Сначала выведим значения узла,
// Если узел имеет левый узел поддерева, левое поддерево отправляется в стек,
// Если узел имеет правый узел поддерева, правое поддерево отправляется в стек.
func (t *TreeNode) LayerOrder() (string, error) {
	if t == nil {
		return "", nil
	}

	var buff strings.Builder

	queue := NewLinkQueue()
	queue.Add(t)

	for queue.size > 0 {
		curr, err := queue.Remove()
		if err != nil {
			return "", err
		}

		buff.WriteString(curr.valueToString())
		buff.WriteString(" ")
		if curr.left != nil {
			queue.Add(curr.left)
		}

		if curr.right != nil {
			queue.Add(curr.right)
		}
	}

	return buff.String(), nil
}
