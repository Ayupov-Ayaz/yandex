package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// --Принцип работы --
// Из входных данных считываются строка t и n префиксов.
// Для каждого префикса создается префиксное дерево.
// Поиск слова t в префиксных деревьях происходит следующим образом:
// На первом шаге из корня выбирается вершина, которая соответствует первой букве t.
// Затем, для каждой следующей буквы t, выбираются все вершины деревьев,
// находящиеся в потомках выбранной вершины, которые соответствуют этой букве.
// Если на каком-то шаге не найдено ни одной вершины, то поиск прекращается, и выводится ответ "NO".
// Если же на последнем шаге найдена хотя бы одна вершина, помеченная как конечная, то выводится ответ "YES".
//
// Сложность алгоритма составляет O(len(t) * log(n)),
// где len(t) - длина строки t, а n - количество префиксов.
// Каждый поиск требует O(log(n)) времени, так как для каждой буквы происходит обход одного из n деревьев.
// Всего производится len(t) поисков.
//
// Алгоритм:
// 1. Считываем количество префиксов
// 2. Считываем префиксы и создаем префиксное дерево
// 3. Считываем строку t
// 4. Ищем строку t в префиксных деревьях
// 5. Выводим результат
// https://contest.yandex.ru/contest/26133/run-report/85821528/

const (
	SUCCESS = "YES"
	FAIL    = "NO"
)

type Node struct {
	isBlack bool
	tree    PrefixMap
}

type PrefixMap map[string]*Node

type DP []*Node

func main() {
	s := strings.Builder{}

	t, root := initInput(bufio.NewScanner(os.Stdin))

	Solution(t, root, &s)
	fmt.Println(s.String())
}

func Solution(t string, rootNode *Node, s *strings.Builder) {
	if rootNode.tree[string(t[0])] == nil {
		s.WriteString(FAIL)
		return
	}

	dp := make(DP, 1)
	dp[0] = rootNode.tree[string(t[0])]

	for i := 1; i < len(t); i++ {
		c := string(t[i])
		var isSomeBlack bool
		dp, isSomeBlack = findNodes(dp, c)

		if isSomeBlack && rootNode.tree[c] != nil {
			dp = append(dp, rootNode.tree[c])
		}

		if len(dp) == 0 {
			break
		}
	}

	for _, v := range dp {
		if v.isBlack {
			s.WriteString(SUCCESS)
			return
		}
	}

	s.WriteString(FAIL)
}

// findNodes поиск ноды по символу
func findNodes(n DP, char string) (DP, bool) {
	var res DP
	var isBlack bool

	for _, v := range n {
		if v.isBlack {
			isBlack = true
		}
		if v.tree[char] != nil {
			res = append(res, v.tree[char])
		}
	}

	if len(res) == 0 {
		return res, isBlack
	}

	return res, isBlack
}

// addPrefix добавление префикса в ноду
func addPrefix(root *Node, s string) {
	for i := 0; i < len(s); i++ {
		if _, ok := root.tree[string(s[i])]; !ok {
			root.tree[string(s[i])] = &Node{
				isBlack: false,
				tree:    PrefixMap{},
			}
		}

		root = root.tree[string(s[i])]
	}

	root.isBlack = true
}

// initInput парсим input-данные
func initInput(scanner *bufio.Scanner) (t string, root *Node) {
	const bufCapacity = 10000000 // fix long string
	buf := make([]byte, bufCapacity)
	scanner.Buffer(buf, bufCapacity)

	scanner.Scan()
	t = scanner.Text()

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	root = &Node{
		isBlack: false,
		tree:    PrefixMap{},
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		p := scanner.Text()
		if unicode.IsSpace(rune(p[len(p)-1])) {
			addPrefix(root, p[:len(p)-1])
		} else {
			addPrefix(root, p)
		}
	}

	return
}
