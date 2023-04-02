package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value  int
	points []*Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func (n *Node) AddPoint(p *Node) {
	n.points = append(n.points, p)
}

type NodeList []*Node

func (n NodeList) GetOrCreateNode(index int) *Node {
	if len(n) <= index {
		panic(fmt.Errorf("index out of range '%d', len = '%d' ", index, len(n)))
	}

	resp := n[index]
	if resp == nil {
		resp = NewNode(index)
		n[index] = resp
	}

	return resp
}

func parseEdge(str string) int {
	resp, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Errorf("invalid val number '%s' %w", str, err))
	}

	return resp
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	var n, m int
	var err error

	arr := strings.Split(sc.Text(), " ")
	if len(arr) != 2 {
		panic("invalid params")
	}

	n, err = strconv.Atoi(arr[0])
	if err != nil {
		panic(fmt.Errorf("invalid 'm' param: %w", err))
	}

	m, err = strconv.Atoi(arr[1])
	if err != nil {
		panic(fmt.Errorf("invalid 'n' param: %w", err))
	}

	list := make(NodeList, n+1)

	for i := 0; i < m; i++ {
		sc.Scan()

		edge := strings.Split(sc.Text(), " ")

		from := list.GetOrCreateNode(parseEdge(edge[0]))
		to := list.GetOrCreateNode(parseEdge(edge[1]))
		from.AddPoint(to)
	}

	var b strings.Builder
	for i := 1; i < len(list); i++ {
		curr := list[i]

		if curr == nil || len(curr.points) == 0 {
			b.WriteString("0\n")
			continue
		}

		b.WriteString(strconv.Itoa(len(curr.points)) + " ")

		for _, p := range curr.points {
			b.WriteString(strconv.Itoa(p.value) + " ")
		}

		if i != len(list)-1 {
			b.WriteString("\n")
		}
	}

	fmt.Println(b.String())
}
