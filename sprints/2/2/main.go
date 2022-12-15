package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
/*
Описал парсер арифметических выражений, обработчик польской натации и
stack для хранения промежуточных данных для вычисления

Структура stack имеет свойства:
* operands - числа которые участвуют в арифметических выражениях, добавление и извлечение по принципу FIFO
* nextIndex - следующий свободный индекс для записи числа

При создании stack необходимо передавать stackLen, это делается для оптимизации работы с памятью
* addOperand - добавление нового числа в slice операндов, при добавление nextIndex инекрементируется
* pop - извлечение последенго элемента массива

--Временная сложность--

Скорость добавления новых операнд в stack:
	* O(n+1) - если nextIndex >= stackLen - будет создан новый slice куда перенесутся
		все имеющиеся на данный момент данные +1 новый
	* О(1) - если nextIndex < stackLen или уже были перенесены данные (1 кейс)
Скорость извлечения данных из stack:
	* O(1) - извлечение идет по nextIndex - 1, по этому скорость будет константная

Скорость подсчета (execute)
	* O(2) - извлечение двух последних элементов
	* O(1) - выполнение арифметической операции над ними
	* O(1) - добавление обратно в stack получившегося результата
	* O(4) - общая сложность execute


Общая сложность O(n)
Потребление памяти O(n)

id - 78966693
https://contest.yandex.ru/contest/22781/run-report/78966693/
*/

const (
	stackLen = 10
)

type stack struct {
	operands  []int
	nextIndex int
}

func newStack(_len int) *stack {
	return &stack{
		operands: make([]int, _len),
	}
}

func (s *stack) addOperand(n int) {
	if s.nextIndex < len(s.operands) {
		s.operands[s.nextIndex] = n
	} else {
		s.operands = append(s.operands, n)
	}

	s.nextIndex++
}

func (s *stack) pop() int {
	if s.nextIndex == 0 {
		return 0
	}

	s.nextIndex--

	return s.operands[s.nextIndex]
}

func negativeNumber(n int) bool {
	return n < 0
}

func execute(operator string, first, second int) int {
	var result int

	switch operator {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "/":
		result = first / second
		if (negativeNumber(first) || negativeNumber(second)) && (first%second != 0) {
			result--
		}
	case "*":
		result = first * second
	}

	return result
}

func calculate(str string) (int, error) {
	values := strings.Split(str, " ")
	stack := newStack(stackLen)

	for i := 0; i < len(values); i++ {
		char := values[i]

		v, err := strconv.Atoi(char)
		if err == nil {
			stack.addOperand(v)
		} else {
			second := stack.pop()
			first := stack.pop()

			stack.addOperand(execute(char, first, second))
		}
	}

	return stack.pop(), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	result, err := calculate(scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
