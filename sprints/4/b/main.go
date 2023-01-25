package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
