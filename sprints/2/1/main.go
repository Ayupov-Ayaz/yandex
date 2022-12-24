package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
/*
-- Принцип работы --
Описал структуру Queue которая работает по принципу очереди на кольцевом буфере
При создании структуры необходимо проставлять максимальную длину
Структура имеет свойства:
* maxSize - максимальный размер
* size - текущий размер
* headIndex - индекс на первый элемент
* tailIndex - индекс на следующий элемент который будет добавляться в конец

PushFront - добавляет элемент в начало списка
PushBack - добавляет элемент в конец списка
перед добавлением проверяет наличие свободного места, если его нет сразу выдает ошибку;
иначе увеличивает счетчик size на 1 и добавляет элемент в head

PopFront - извлекает элемент с начала списка, если элемента нет, тогда возвращет ошибку, size уменьшается на 1
PopBack - извлекает элемент с конца списка, если элемента нет, тогда возвращает ошибку, size уменьшается на 1

--Временная сложность--

Скорость добавления О(1)
Скорость удаления О(1)
Общая сложность O(n)
Потребление памяти O(n)


id - 79553249
https://contest.yandex.ru/contest/22781/run-report/79553249/
*/

var (
	errQueue = errors.New("error")
)

// Queue - двухсторонняя очередь с ограниченным размеров
type Queue struct {
	size      int
	maxSize   int
	headIndex int
	tailIndex int
	values    []int
}

func NewQueue(maxSize int) *Queue {
	return &Queue{
		maxSize:   maxSize,
		tailIndex: 1,
		values:    make([]int, maxSize),
	}
}

func (q *Queue) availableToPush() error {
	if q.size >= q.maxSize {
		return errQueue
	}

	return nil
}

func (q *Queue) availableToPop() error {
	if q.size == 0 {
		return errQueue
	}

	return nil
}

func (q *Queue) getNextIndex(i, delta int) int {
	x := i + delta

	if x < 0 {
		return q.maxSize + delta
	} else if x < q.maxSize {
		return x
	} else {
		return 0
	}
}

func (q *Queue) PopFront() (int, error) {
	if err := q.availableToPop(); err != nil {
		return 0, err
	}

	index := q.getNextIndex(q.headIndex, +1)
	resp := q.values[index]
	q.headIndex = index
	q.size--

	return resp, nil
}

func (q *Queue) PushBack(v int) error {
	if err := q.availableToPush(); err != nil {
		return err
	}

	q.values[q.tailIndex] = v
	q.size++

	q.tailIndex = q.getNextIndex(q.tailIndex+1, 0)

	return nil
}

func (q *Queue) PopBack() (int, error) {
	if err := q.availableToPop(); err != nil {
		return 0, err
	}

	index := q.getNextIndex(q.tailIndex, -1)
	resp := q.values[index]
	q.tailIndex = index
	q.size--

	return resp, nil
}

func (q *Queue) PushFront(v int) error {
	if err := q.availableToPush(); err != nil {
		return err
	}

	q.size++
	q.values[q.headIndex] = v
	q.headIndex = q.getNextIndex(q.headIndex, -1)

	return nil
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 16 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readN(scan *bufio.Scanner) int {
	scan.Scan()
	v, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}

	return v
}

var pop map[string]func() (int, error)

var push map[string]func(v int) error

func init() {
	pop = make(map[string]func() (int, error), 2)
	push = make(map[string]func(v int) error)
}

func newQueue(maxSize int) func(commands string) error {
	q := NewQueue(maxSize)
	pop["pop_back"] = q.PopBack
	pop["pop_front"] = q.PopFront
	push["push_front"] = q.PushFront
	push["push_back"] = q.PushBack

	var (
		v   int
		err error
	)

	return func(commands string) error {
		text := strings.Split(commands, " ")
		command := text[0]

		if popFunc, ok := pop[command]; ok {
			v, err = popFunc()
			if err == nil {
				fmt.Println(v)
				return nil
			}
		}

		if pushFunc, ok := push[command]; ok {
			v, err = strconv.Atoi(text[1])
			if err == nil {
				err = pushFunc(v)
			}
		}

		return err
	}
}

func main() {
	scan := makeScanner()
	n := readN(scan)
	maxSize := readN(scan)

	q := newQueue(maxSize)
	for i := 0; i < n; i++ {
		scan.Scan()
		if err := q(scan.Text()); err != nil {
			fmt.Println(err)
		}
	}
}
