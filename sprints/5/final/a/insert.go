package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Пирамидальная сортировка
// -- Принцип работы --
// Сортировка осуществляется через построение бинарного дерева путем добавления в корень новых элементов
// Добавление осущесвляется через функцию - Insert
// Эффективность данного алгоритма целиком и полностью зависит от того в каком порядке будут добавляться данные
// так как бинарное дерево может получиться не сбалансированным, что увеличит время прохода по нему.
// В качестве узлов дерева вуступает структуру *Node, которая имеет следующие поля:
// data Member - данные об участнике соревнования
// left, right - правый и левый узел
//
// Алгоритм вставки (Insert)
// 1) Для первого элемента мы просто создаем *Node и запоминаем его как root элемент
// 2) Все последующие элементы мы будем добавлять в дерево через функцию Insert,
//		которая принимает в качестве аргумента root и значение которое добавляется
// 3) Проходим по дереву пока не найдем свободное место куда можно добавить новый элемент пользуясь принципом:
// 		Если текущий элемент меньше, чем добавляемый, тогда идем в правую ветку (если она есть)
// 		Если текущий элемент больше, чем добавляемый, тогда идем в левую ветку (если она есть)
// 		Если ветки нет, тогда добавляем наш элемент в качестве узла этой ветки
// 	Такой принцип добавляения позволит нам сразу же отсортировать наши данные и держать их в актуальном состоянии
//
// Для сравнения участников используется функция Member.IsLess(m Member) bool
//
// Алгоритм получения отсортированного списка (RightToLeftOrder)
// Мы рекурсивно пробегам по дереву справа на лево
// Сначало мы получаем данные от крайнего правого элемента(участник), добавляем имя участника в буфер с результатом;
// Затем на очереди средний элемент (участник), добавляем имя участника в буфер;
// Потом запоминаем левого участника(элемент дерева)
// Благодаря такому проходу наши участнику будут отсортированы по убыванию мест которые они заняли в соревновании
//
// -- Временная сложность --
// h - высота дерева
//
// Insert:
// 		* В худшем случае скорость будет O(h), зависит порядкя добавления в дерево участников
// 		* Память расходуется O(n)
// RightToLeftOrder
//	* Скорость всегда будет O(n), так как по условию задачи нам нужно вывести абсолютно всех участников в порядке убывания
//
// id = https://contest.yandex.ru/contest/24810/run-report/83425478/

// Member - участник соревнований
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
// участник является хуже, если удовлетворяет одному из условий, по ирерархии:
// 1. Меньше решенных задач
// 2. Больше штрафов
// 3. Лексикографически его фамилия стоит позже по списку
func (m Member) IsLess(that Member) bool {
	if m.tasks == that.tasks {
		if m.fined == that.fined {
			return that.name > m.name
		}

		return m.fined < that.fined
	}

	return m.tasks > that.tasks
}

// RightToLeftOrder - Обратный порядок обхода дерева
// right -> head -> left
func (t *Node) RightToLeftOrder() string {
	if t == nil {
		return ""
	}

	var buff strings.Builder
	left := t.left.RightToLeftOrder()
	right := t.right.RightToLeftOrder()

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

	fmt.Println(root.RightToLeftOrder())
}
