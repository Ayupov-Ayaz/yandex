package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Member struct {
	name  string
	tasks int
	fined int
}

func NewMember(name string, tasks, fined int) Member {
	return Member{
		name:  name,
		tasks: tasks,
		fined: fined,
	}
}

func makeMember(name, s2, s3 string) Member {
	tasks, err := strconv.Atoi(s2)
	if err != nil {
		panic(err)
	}

	fined, err := strconv.Atoi(s3)
	if err != nil {
		panic(err)
	}

	return NewMember(name, tasks, fined)
}

// IsLess - является ли переданный участник хуче, чем текущий
// участник является лучшим, если удовлетворяет одному из условий, по ирерархии:
// 1. Больше решенных задач
// 2. Меньше штрафов
// 3. Лексикографически его фамилия стоит раньше по списку
func (m Member) IsLess(that Member) bool {
	if m.tasks == that.tasks {
		if m.fined == that.fined {
			return that.name > m.name
		}

		return m.fined < that.fined
	}

	return m.tasks > that.tasks
}

// RightOrder - не упорядоченный обход
// right -> head -> left
func (t *Node) RightOrder() string {
	if t == nil {
		return ""
	}

	var buff strings.Builder
	left := t.left.RightOrder()
	right := t.right.RightOrder()

	if right != "" {
		buff.WriteString(right)
		buff.WriteString("\n")
	}

	buff.WriteString(t.data.name)

	if left != "" {
		buff.WriteString("\n")
		buff.WriteString(left)
	}

	return buff.String()
}

type Node struct {
	data  Member
	left  *Node
	right *Node
}

func NewNode(data Member) *Node {
	return &Node{
		data: data,
	}
}

func Insert(root *Node, m Member) {
	node := NewNode(m)

	curr := root
	for curr != nil {
		if curr.data.IsLess(m) {
			if curr.left != nil {
				curr = curr.left
			} else {
				curr.left = node
				break
			}
		} else if curr.right != nil {
			curr = curr.right
		} else {
			curr.right = node
			break
		}
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	var root *Node
	for i := 0; i < n; i++ {
		sc.Scan()
		val := strings.Split(sc.Text(), " ")
		m := makeMember(val[0], val[1], val[2])

		if root != nil {
			Insert(root, m)
		} else {
			root = NewNode(m)
		}
	}

	fmt.Println(root.RightOrder())
}
