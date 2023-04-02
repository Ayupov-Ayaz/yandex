package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	WHITE Color = 0
)

type Color int

func (c Color) IsWhite() bool {
	return c == WHITE
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

func makeList(n int) []*Node {
	list := make([]*Node, n)
	for i := 0; i < n; i++ {
		list[i] = NewNode(i + 1)
	}

	return list
}

type Node struct {
	value  int
	points []*Node
}

type Link struct {
	data *Node
	prev *Link
}

func NewLink(data *Node, prev *Link) *Link {
	return &Link{data: data, prev: prev}
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

type Stack struct {
	tail *Link
	size int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	link := NewLink(n, s.tail)
	s.tail = link
	s.size++
}

func (s *Stack) Pop() *Node {
	var resp *Node
	if s.size > 0 {
		s.size--
		resp = s.tail.data
		s.tail = s.tail.prev
	}

	return resp
}

func getKey(index int) int {
	return index - 1
}

func getComponents(list []*Node) [][]int {
	color := Color(0)
	colors := make([]Color, len(list))

	stack := NewStack()
	resp := make([][]int, 0, len(list))

	for i := 0; i < len(list); i++ {
		if !colors[i].IsWhite() {
			continue
		}
		color++
		stack.Push(list[i])

		component := make([]int, 0, 3)
		for stack.size > 0 {
			n := stack.Pop()
			key := getKey(n.value)
			if colors[key].IsWhite() {
				colors[key] = color
				component = append(component, n.value)
			}

			for _, child := range n.points {
				if colors[getKey(child.value)].IsWhite() {
					stack.Push(child)
				}
			}
		}

		resp = append(resp, component)
	}

	return resp
}

func sortComponents(components [][]int) {
	sort.Slice(components, func(i, j int) bool {
		left := components[i]
		right := components[j]

		sort.Slice(left, func(i, j int) bool {
			return left[i] < left[j]
		})

		sort.Slice(right, func(i, j int) bool {
			return right[i] < right[j]
		})

		return left[0] < right[0]
	})
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, m := parseParams(sc.Text())
	list := makeList(n)

	for i := 0; i < m; i++ {
		sc.Scan()
		from, to := parseParams(sc.Text())
		fromN := list[getKey(from)]
		toN := list[getKey(to)]
		fromN.points = append(fromN.points, toN)
		toN.points = append(toN.points, fromN)
	}

	components := getComponents(list)
	sortComponents(components)

	var b strings.Builder
	b.WriteString(strconv.Itoa(len(components)) + "\n")
	for i, component := range components {
		for _, val := range component {
			b.WriteString(strconv.Itoa(val) + " ")
		}

		if i != len(components)-1 {
			b.WriteByte('\n')
		}
	}

	fmt.Println(b.String())
}
