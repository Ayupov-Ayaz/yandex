package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Color int8

const (
	WHITE Color = iota
	GRAY
	BLACK
)

type Node struct {
	value    int
	children []*Node
}

func NewNode(v int) *Node {
	return &Node{value: v}
}

type Links []*Node

func getKey(value int) int {
	return value - 1
}

func (l Links) GetNode(value int) *Node {
	return l[getKey(value)]
}

type Stack struct {
	head  int
	tail  int
	stack []string
}

func NewStack(cnt int) *Stack {
	return &Stack{
		tail:  cnt - 1,
		stack: make([]string, cnt),
	}
}

func (s *Stack) add(index, v int) {
	s.stack[index] = strconv.Itoa(v)
}

func (s *Stack) PushFront(v int) {
	if s.head <= s.tail {
		s.add(s.head, v)
		s.head++
	}
}

func (s *Stack) PushTail(v int) {
	if s.tail >= s.head {
		s.add(s.tail, v)
		s.tail--
	}
}

func (s *Stack) ToString() string {
	return strings.Join(s.stack, " ")
}

func parseParams(str string) (int, int) {
	arr := strings.Split(str, " ")

	a, err := strconv.Atoi(arr[0])
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(arr[1])
	if err != nil {
		panic(err)
	}

	return a, b
}

func Sort(n *Node, colors []Color, stack *Stack) {
	key := getKey(n.value)

	colors[key] = GRAY

	for _, child := range n.children {
		if colors[getKey(child.value)] == WHITE {
			Sort(child, colors, stack)
		}
	}

	colors[key] = BLACK
	stack.PushTail(n.value)
}

func TopSort(links Links) string {
	colors := make([]Color, len(links))
	stack := NewStack(len(links))

	for i := len(colors) - 1; i > -1; i-- {
		if colors[i] == WHITE {
			Sort(links[i], colors, stack)
		}
	}

	return stack.ToString()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	arr := strings.Split(sc.Text(), " ")
	if len(arr) != 2 {
		panic("invalid params number")
	}

	n, m := parseParams(sc.Text())
	links := make(Links, n)

	for i := 1; i <= n; i++ {
		links[i-1] = NewNode(i)
	}

	var from, to int
	var fromN, toN *Node
	for i := 0; i < m; i++ {
		sc.Scan()
		from, to = parseParams(sc.Text())
		fromN = links.GetNode(from)
		toN = links.GetNode(to)

		fromN.children = append(fromN.children, toN)
	}

	fmt.Println(TopSort(links))
}
