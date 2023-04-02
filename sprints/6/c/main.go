package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Color int8

const (
	WHITE Color = iota
	GRAY
	BLACK
)

type Colors []Color

type Vertex int

func parseEdge(str string) int {
	resp, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Errorf("invalid val number '%s' %w", str, err))
	}

	return resp
}

func getParams(str string) (int, int) {
	arr := strings.Split(str, " ")
	if len(arr) != 2 {
		panic(fmt.Errorf("invalid params count = %d, should be 2", len(arr)))
	}

	r1 := parseEdge(arr[0])
	r2 := parseEdge(arr[1])

	return r1, r2
}

type node struct {
	data Vertex
	prev *node
}

func newNode(data Vertex, prev *node) *node {
	return &node{data: data, prev: prev}
}

type stack struct {
	size int
	tail *node
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) Pop() Vertex {
	if s.size == 0 {
		return -1
	}

	s.size--
	resp := s.tail.data
	s.tail = s.tail.prev

	return resp
}

func (s *stack) Push(v Vertex) {
	s.size++
	node := newNode(v, s.tail)
	s.tail = node
}

func DFS(arr [][]Vertex, start int) []int {
	colors := make(Colors, len(arr))
	stack := newStack()
	stack.Push(Vertex(start))

	result := make([]int, 0, len(arr)/2)
	for stack.size > 0 {
		v := stack.Pop()
		color := colors[v]
		if color == GRAY {
			colors[v] = BLACK
		}

		if color != WHITE {
			continue
		}

		result = append(result, int(v))
		// красим в серый и добавляем обратно в стек,
		// чтобы проконтролировать момент когда мы вернемся обратно
		colors[v] = GRAY
		stack.Push(v)

		points := arr[v]
		if len(points) > 0 {
			sort.Slice(points, func(i, j int) bool {
				return points[i] > points[j]
			})
			for _, v := range points {
				if colors[v] == WHITE {
					stack.Push(v)
				}
			}
		}
	}

	return result
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	params := strings.Split(sc.Text(), " ")
	n := parseEdge(params[0])
	m := parseEdge(params[1])

	arr := make([][]Vertex, n+1)

	var from, to int
	for i := 0; i < m; i++ {
		sc.Scan()

		from, to = getParams(sc.Text())
		if arr[from] == nil {
			arr[from] = make([]Vertex, 0, 1)
		}
		arr[from] = append(arr[from], Vertex(to))

		if arr[to] == nil {
			arr[to] = make([]Vertex, 0, 1)
		}
		arr[to] = append(arr[to], Vertex(from))
	}

	sc.Scan()
	start := parseEdge(sc.Text())

	var b strings.Builder
	for _, n := range DFS(arr, start) {
		b.WriteString(strconv.Itoa(n) + " ")
	}

	fmt.Println(b.String())
}
