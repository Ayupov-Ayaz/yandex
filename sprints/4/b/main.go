package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Расписал алгоритм  hash таблицы
// Hash таблица имеет размер n/2, где n - количество вводимых операций
// если учесть, что операций на put не божет быть ровно n,
// то можно с уверенностью предположить, что таблица не будет заполнена
// -- Принцип работы --
// Создается hash таблица - slice с заранее заданной длиной n в которой хранятся узлы
//
// Put - при добавлении нового объекта алгоритм следующий:
// 1) высчитывается hash от ключа (алгоритм ниже)
// 2) смотрим есть ли какой-то объект в slice по данному hash-y
// 2.1) если объекта нет, тогда записываем наше значение
// 2.2) если объект есть, тогда разрешаем коллизии методом цепочек - двигаемся по цепочке от головы узла до его конца,
// попутно проверяя нет ли совпадения по ключам у какой-нибудь узла
// если найдено совпадение, тогда заменяем у объекта значение
// иначе доходим до конца списка и создаем узел с переданным ключем и значением
//
// Get - при получении объекта алгоритм следующий:
//  1. вычисляем hash от ключа
//  2. проверяем существование узла по данному ключу
//  3. если объекта нет, тогда сообщаем об этом
//  4. если объект есть проверяем ключ текущего узла,
//     если ключ текущего узла совпадает с запрошенным, тогда возвращаем значение
//     если ключ не совпадает, проверяем есть ли у даного узла следующий элемент и проверяем его ключ
//     делаем так до тех пор пока не найдем нужный ключ,
//     либо не дойдем до конца списка (в таком случае сообщаем, что элемента нету)
//
// Delete - при удалении объекта алгоритм следующий
// 1. вычисляем hash от ключа
// 2. проверяем существование узла в slice по данному ключу
// 3. если узла нет, тогда сообщаем об этом
// 4. если узел есть, тогда двигаемся по нему до конца списка
// пока не встретим нужный нам ключ и не удалим его из списка (поменяв порядок)
// 5. возвращаем значение удаленного узла
//
// hash функция очень простая
// мы следим, за тем чтобы нам не попадали отрицательные значения (умножает на -1, если число меньше 0)
// и возвращаем остаток от деления на длину slice
//
// -- Временная сложность --
// O(1) - вставка
// O(1) - удаление
// O(1) - получение
// (при условии, что коллизий нет)
// -- Расход памяти --
// O(n) - так как нужно создавать узлы для каждого узла
//
// id - https://contest.yandex.ru/contest/24414/run-report/81335072/
type node struct {
	key  int
	val  int
	next *node
}

func newNode(key, val int) *node {
	return &node{
		key: key,
		val: val,
	}
}

type HashTable struct {
	table    []*node
	size     int
	hashFunc HashFunc
}

type HashFunc func(key int, size int) int

func NewHashTable(size int, hashFunc HashFunc) *HashTable {
	return &HashTable{
		table:    make([]*node, size),
		size:     size,
		hashFunc: hashFunc,
	}
}

func SimpleHashFunc(key, size int) int {
	if key < 0 {
		key *= -1
	}

	return key % size
}

func (h *HashTable) Add(key, val int) {
	hashKey := h.hashFunc(key, h.size)

	curr := h.table[hashKey]
	if curr == nil {
		h.table[hashKey] = newNode(key, val)
		return
	}

	var prev *node
	for curr != nil {
		if curr.key == key {
			curr.val = val
			return
		}

		prev = curr
		curr = curr.next
	}

	prev.next = newNode(key, val)

	return
}

func (h *HashTable) Get(key int) (int, bool) {
	hashKey := h.hashFunc(key, h.size)

	curr := h.table[hashKey]
	for curr != nil {
		if curr.key == key {
			return curr.val, true
		}

		curr = curr.next
	}

	return 0, false
}

func (h *HashTable) Remove(key int) (int, bool) {
	hashKey := h.hashFunc(key, h.size)

	curr := h.table[hashKey]
	if curr != nil && curr.key == key && curr.next == nil {
		h.table[hashKey] = nil
		return curr.val, true
	}

	var prev *node
	for curr != nil {
		if curr.key == key {
			if prev == nil {
				h.table[hashKey] = curr.next
			} else {
				prev.next = curr.next
			}
			return curr.val, true
		}
		// >>
		prev = curr
		curr = curr.next
	}

	return 0, false
}

func iterator(ht *HashTable, next func() string) {
	for {
		text := next()
		if text == "" {
			return
		}

		cmd := strings.Split(text, " ")

		key, err := strconv.Atoi(cmd[1])
		if err != nil {
			panic(err)
		}

		switch cmd[0] {
		case "put":
			val, err := strconv.Atoi(cmd[2])
			if err != nil {
				panic(err)
			}

			ht.Add(key, val)
		case "get":
			val, ok := ht.Get(key)
			if ok {
				fmt.Println(val)
			} else {
				fmt.Println("None")
			}
		case "delete":
			val, ok := ht.Remove(key)
			if !ok {
				fmt.Println("None")
			} else {
				fmt.Println(val)
			}
		default:
			panic(fmt.Sprintf("invalid command %s", cmd[0]))
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

	size := n / 2

	ht := NewHashTable(size, SimpleHashFunc)

	k := 0
	iterator(ht, func() string {
		k++
		if k <= n {
			sc.Scan()
			return sc.Text()
		}

		return ""
	})
}
