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
	WHITE Color = iota
	GRAY
	BLACK
)

type Color int8

func (c Color) IsWhite() bool {
	return c == WHITE
}

type Node struct {
	data   int
	points []*Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

type Link struct {
	data *Node
	next *Link
}

func NewLink(data *Node) *Link {
	return &Link{data: data}
}

type Queue struct {
	head *Link
	tail *Link
	size int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(n *Node) {
	q.size++
	link := NewLink(n)

	if q.size == 1 {
		q.head = link
	} else {
		q.tail.next = link
	}

	q.tail = link
}

func (q *Queue) Pop() (resp *Node) {
	if q.size > 0 {
		q.size--
		resp = q.head.data
		q.head = q.head.next
	}

	return resp
}

func getKey(index int) int {
	return index - 1
}

func parseParams(str string, isIndex bool) (int, int) {
	arr := strings.Split(str, " ")

	a, err := strconv.Atoi(arr[0])
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(arr[1])
	if err != nil {
		panic(err)
	}

	if isIndex {
		a = getKey(a)
		b = getKey(b)
	}

	return a, b
}

func BFS(links []*Node, start int) []int {
	planed := NewQueue()
	colors := make([]Color, len(links))

	planed.Push(links[start])
	colors[getKey(links[start].data)] = GRAY

	result := make([]int, 0, len(links))
	for planed.size > 0 {
		curr := planed.Pop()
		colors[getKey(curr.data)] = BLACK
		result = append(result, curr.data)

		for _, point := range curr.points {
			key := getKey(point.data)
			if colors[key].IsWhite() {
				colors[key] = GRAY
				planed.Push(point)
			}
		}
	}

	return result
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, m := parseParams(sc.Text(), false)

	links := make([]*Node, n)
	for i := 0; i < n; i++ {
		links[i] = NewNode(i + 1)
	}

	for i := 0; i < m; i++ {
		sc.Scan()
		from, to := parseParams(sc.Text(), true)

		links[from].points = append(links[from].points, links[to])
		links[to].points = append(links[to].points, links[from])
	}

	for i := 0; i < n; i++ {
		curr := links[i].points

		sort.Slice(curr, func(i, j int) bool {
			return curr[i].data < curr[j].data
		})
	}

	sc.Scan()
	start, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	result := BFS(links, getKey(start))
	var b strings.Builder
	for _, r := range result {
		b.WriteString(strconv.Itoa(r) + " ")
	}

	fmt.Println(b.String())
}
