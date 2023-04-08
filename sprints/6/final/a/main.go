package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// А. Дорогая сеть
// -- Принцип работы --
// Задача решается через применение алгоритма Дейкстры.
// В качестве очереди используется массив с ребрами и их весом.
// При добавлении ребра в очередь, она просеивается вверх.
// При извлечении ребра из очереди, она просеивается вниз.
// При просеивании вверх, сравнивается вес ребра с весом родителя.
// При просеивании вниз, сравнивается вес ребра с весом наименьшего из детей.
// При извлечении ребра из очереди, оно сравнивается с первым ребром в очереди.
// Если оно меньше, то оно извлекается, иначе извлекается первое ребро.
// При извлечении ребра из очереди, оно помечается как посещенное.
// При добавлении ребра в очередь, оно проверяется на посещенность.
// Если ребро не посещено, то оно добавляется в очередь.
// После извлечения всех ребер из очереди, проверяется, что все вершины посещены.
// Если все вершины посещены, то вес ребер в остовном дереве суммируется.
// сложность: O(|V^2| + |E|)
// id =  https://contest.yandex.ru/contest/25070/run-report/85262868/

var (
	ErrQueueIsEmpty = errors.New("queue is empty")
)

const errStr = "Oops! I did it again"

// Edge ребро
type Edge struct {
	weight int
	node   *Node
}

// Node узел
type Node struct {
	value int
	edges []Edge
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

// NodeList список смежности
type NodeList []*Node

// Add добавление в матрицу
func (n NodeList) Add(v, u, w int) {
	edge := Edge{
		weight: w,
		node:   n[u],
	}

	n[v].edges = append(n[v].edges, edge)
}

// Queue очередь для хранения ребер
type Queue []Edge

// Push добавление в очередь
func (q *Queue) Push(edge Edge) {
	*q = append(*q, edge)
	q.TreeUp(q.Len() - 1)
}

// Pop извлечение из очереди
func (q *Queue) Pop() (*Edge, error) {
	if q.Len() == 1 {
		return nil, ErrQueueIsEmpty
	}

	save := (*q)[1]
	(*q)[1] = (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	TreeDown(*q, 1)

	return &save, nil
}

// Len размер очереди
func (q Queue) Len() int {
	return len(q)
}

// Less сравнение
func (q Queue) Less(a, b int) bool {
	return q[a].weight > q[b].weight
}

// Swap замена
func (q Queue) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
}

// TreeUp просеивание вверх
func (q Queue) TreeUp(index int) {
	if index == 1 {
		return
	}
	parentIndex := index >> 1

	if q.Less(index, parentIndex) {
		q.Swap(index, parentIndex)
		q.TreeUp(parentIndex)
	}
}

// TreeDown просеивание вниз
func TreeDown(q Queue, index int) {
	firstChild := index * 2
	secondChild := firstChild + 1

	if firstChild >= q.Len() {
		return
	}
	best := firstChild

	if secondChild < q.Len() && q.Less(secondChild, firstChild) {
		best = secondChild
	}

	if q.Less(best, index) {
		q.Swap(best, index)
		TreeDown(q, best)
	}
}

// FindMST поиск остового дерева (Дийкстра)
func FindMST(al NodeList) (int, bool) {
	var total = true
	var accum int

	queue := make(Queue, 1, len(al)) // очередь ребер
	visited := make([]bool, len(al)) // посещенные вершины

	visited[1] = true

	// добавляем ребра в очередь
	for _, edge := range al[1].edges {
		queue.Push(edge)
	}

	// извлекаем из очереди и проверям
	for edge, err := queue.Pop(); err == nil; edge, err = queue.Pop() {
		if visited[edge.node.value] {
			continue
		}

		accum += edge.weight
		visited[edge.node.value] = true

		for _, e := range edge.node.edges {
			if !visited[e.node.value] {
				queue.Push(e)
			}
		}
	}

	for i := 1; i < len(visited); i++ {
		if !visited[i] {
			total = false
		}
	}

	return accum, total
}

// парсит строку и получает 2 числа
func parseParams(str string) (n int, m int) {
	var err error

	nmData := strings.Fields(str)
	n, err = strconv.Atoi(nmData[0])
	if err != nil {
		panic(err)
	}

	m, err = strconv.Atoi(nmData[1])
	if err != nil {
		panic(err)
	}

	return n, m
}

func parseEdge(str string) (v int, u int, w int) {
	var err error

	edgeData := strings.Fields(str)
	// первая вершина
	v, err = strconv.Atoi(edgeData[0])
	if err != nil {
		panic(err)
	}
	// вторая вершина
	u, err = strconv.Atoi(edgeData[1])
	if err != nil {
		panic(err)
	}
	// вес
	w, err = strconv.Atoi(edgeData[2])
	if err != nil {
		panic(err)
	}

	return v, u, w
}

func main() {
	var res strings.Builder

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// количество вершин и ребер
	n, m := parseParams(scanner.Text())

	// если нет ребер
	if n > 1 && m == 0 {
		fmt.Println(errStr)
		return
	}

	list := make(NodeList, n+1)

	// создадим вершины
	for i := 1; i <= n; i++ {
		list[i] = NewNode(i)
	}

	// заполняем вершины ребрами
	for i := 1; i <= m; i++ {
		scanner.Scan()
		v, u, w := parseEdge(scanner.Text())

		if v == u {
			continue
		}

		list.Add(v, u, w)
		list.Add(u, v, w)
	}

	// проверим наличие остовного дерева
	if weight, err := FindMST(list); err {
		res.WriteString(strconv.Itoa(weight))
	} else {
		res.WriteString(errStr)
	}

	fmt.Println(res.String())
}
