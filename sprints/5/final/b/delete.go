package main

// Удали узел
// -- Принцип работы --
// Есть 3 возможных случая:
// 1) Когда удаляемый узел является листочком, то есть не имеет потомков -
// 		В таком случае нам необходимо обнулить его ссылку у parent элемента
// 2) Когда у удаляемого узла есть 2 дочерних элемента -
// 		В таком случае, нам необходимо сначала найти для удаляемого узла приемника.
// 		Приемником в моем алгоритме является самое большое число из левой ветки удаляемого элемента
// 			Чтобы его проставить, сначало необходимо обнулить ссылку у parent на этого приемника
// 		Потом мы можем заменить удаляемый элемент этим самым приемником.
// 	3) Когда у удаляемого элемента 1 дочерний элемент -
// 		В таком случае необходимо проставить у parent ссылку на child удаляемого элемента
//
// -- Временная сложность --
// O(h),  где h - высота дерева
// id = https://contest.yandex.ru/contest/24810/run-report/83415842/

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
