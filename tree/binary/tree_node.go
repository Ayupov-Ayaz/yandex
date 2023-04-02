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

// GetMaxNode - функция для поиска минимального значения для узла
func (t *TreeNode) GetMaxNode() *TreeNode {
	curr := t
	for curr.right != nil {
		curr = curr.right
	}

	return curr
}

// Insert - вставка ключа
func (t *TreeNode) Insert(value int) {
	curr := t

	node := NewTreeNode(value)
	for curr != nil {
		if curr.value > value {
			if curr.left == nil {
				curr.left = node
				break
			} else {
				curr = curr.left
			}
		} else {
			if curr.right == nil {
				curr.right = node
				break
			} else {
				curr = curr.right
			}
		}
	}
}

func Delete(root *TreeNode, value int) {
	curr := root

	var parent *TreeNode
	for curr != nil && curr.value != value {
		// обновляем родителя до текущего узла
		parent = curr

		if curr.value > value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	if curr == nil {
		return
	}

	// 1. случай первый: удаляемый узел не имеет дочерних элементов
	if curr.left == nil && curr.right == nil {
		if parent != nil && parent.left == curr {
			parent.left = nil
		} else if parent != nil && parent.right == curr {
			parent.right = nil
		} else if curr == root {
			root = nil
		}

		return
	}

	// 2. Удаляемый узел имеет двух потомков
	if curr.left != nil && curr.right != nil {
		maxLeftValue := curr.left.GetMaxNode().value
		// рекурсивно удаляем приемника
		Delete(curr, maxLeftValue)
		// копируем значение приемника в текущий узел
		curr.value = maxLeftValue
		return
	}

	// 3. Удаляемый узел имеет только одного потомка
	var child *TreeNode
	if curr.left != nil {
		child = curr.left
	} else {
		child = curr.right
	}

	// если удаляемый узел не является корневым узлом, устанавливаем его родителя своему потомку
	if curr != root {
		if curr == parent.left {
			parent.left = child
		} else {
			parent.right = child
		}
	} else {
		root = child
	}
}
