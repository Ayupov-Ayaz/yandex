package main

type Node struct {
	value int
	left  *Node
	right *Node
}

// GetMaxNode - функция для поиска минимального значения для узла
func (t *Node) GetMaxNode() *Node {
	curr := t
	for curr.right != nil {
		curr = curr.right
	}

	return curr
}

func remove(root *Node, value int) *Node {
	curr := root

	var parent *Node
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
		return root
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

		return root
	}

	// 2. Удаляемый узел имеет двух потомков
	if curr.left != nil && curr.right != nil {
		maxLeftValue := curr.left.GetMaxNode().value
		// рекурсивно удаляем приемника
		_ = remove(curr, maxLeftValue)
		// копируем значение приемника в текущий узел
		curr.value = maxLeftValue
		return root
	}

	// 3. Удаляемый узел имеет только одного потомка
	var child *Node
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

	return root
}
