package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// -- Принцип работы --
// 1. Считываем количество строк
// 2. Распаковываем первую строку и записываем ее в префикс
// 3. Распаковываем след строку и записываем ее в текущую
// 4. Находим индекс, с которого начинается различие префикса и текущей строки
// 5. Записываем в префикс подстроку от начала до индекса, с которого начинается различие
// 6. Повторяем шаги 3-5 для всех строк
// 7. Выводим префикс
//
// Распаковка происходит по следующему правилу:
// если в строке встречается число, то следующий за ним символ встречается столько раз, сколько указано в числе.
// Квадратные скобки [ и ] используются как группирующие символы
//
// Распаковка строки происходит по следующему принципу:
// 1. Считываем символ
// 2. Если символ открывающая скобка, то записываем в стек текущую позицию
// 3. Если символ закрывающая скобка, то:
// 3.1. Считываем число из стека
// 3.2. Считываем подстроку из стека
// 3.3. Добавляем подстроку в конец текущей строки
// 3.4. Добавляем подстроку в стек
// 3.5. Добавляем число в стек
// 4. Если символ не открывающая и не закрывающая скобка, то добавляем его в текущую строку
// 5. Повторяем шаги 1-4 для всех символов в строке
// 6. Возвращаем текущую строку
//
// Сложность алгоритма:
// O(n * k), где n - количество строк ввода, k - максимальная длина строки ввода
// Сложность функции unpack зависит от длины входной строки и количества повторений,
// но в худшем случае она также равна O(k).
//
// https://contest.yandex.ru/contest/26133/run-report/85798786/

type Stack [][]byte

const (
	OPEN  = byte('[')
	CLOSE = byte(']')
)

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

// Solution решение задачи
func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)

	const bufCapacity = 10000000 // fix long string
	buf := make([]byte, bufCapacity)
	scanner.Buffer(buf, bufCapacity)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	var prefix, current string

	for i := 0; i < n; i++ {
		scanner.Scan()
		if i == 0 {
			prefix = unpack(scanner.Bytes())
			continue
		}

		current = unpack(scanner.Bytes())
		prefix = current[:findIndex(prefix, current)]
	}

	s.WriteString(prefix)
}

// unpack распаковка строки
func unpack(s []byte) string {
	var stack Stack
	var builder strings.Builder
	var top, curr []byte

	for i := 0; i < len(s); i++ {
		// fix
		if unicode.IsSpace(rune(s[i])) {
			continue
		}

		switch s[i] {
		case CLOSE:
			top = stack.Pop()
			curr = nil

			for !unicode.IsDigit(rune(top[0])) {
				curr = append(top, curr...)
				top = stack.Pop()
			}

			if num, err := strconv.Atoi(string(top)); err != nil {
				log.Fatal(err)
			} else {
				stack.Push(bytes.Repeat(curr, num))
			}
		case OPEN:
			continue
		default:
			stack.Push([]byte(string(s[i])))
		}
	}

	for i := 0; i < len(stack); i++ {
		builder.Write(stack[i])
	}

	return builder.String()
}

// findIndex поиск первого неравного индекса
func findIndex(a, b string) int {
	var max int

	if len(a) > len(b) {
		max = len(a)
	} else {
		max = len(b)
	}

	var i int
	for ; i < max; i++ {
		if i >= len(a) || i >= len(b) || a[i] != b[i] {
			break
		}
	}

	return i
}

// IsEmpty пустой стек
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push добавление в стек
func (s *Stack) Push(str []byte) {
	*s = append(*s, str)
}

// Pop последний элемент из стека
func (s *Stack) Pop() []byte {
	if s.IsEmpty() {
		return nil
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}
