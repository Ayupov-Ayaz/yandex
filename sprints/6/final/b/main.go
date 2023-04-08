package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// B. Железные дороги
// -- Принцип работы --
// Данный алгоритм решает задачу определения возможности проезда поезда через города, соединенные железными дорогами.
// Предполагается, что между городами могут быть только два типа железных дорог: R  и B.
//
// Алгоритм проверяет, существует ли такой маршрут, что поезд сможет проехать через все города.
// Для этого он применяет поиск в глубину (DFS) для обхода матрицы смежности графа и проверяет, что нет циклов в графе.
// Если все вершины были посещены без образования циклов, то существует маршрут,
// который позволит поезду проехать через все города. Если же есть хотя бы один цикл,
// то поезд не сможет проехать через все города.
//
// Каждая вершина имеет один из трех статусов:
const (
	// WHITE не обработанная вершина
	WHITE Color = iota
	// GRAY мы сюда заходили, но еще не обработали
	GRAY
	// BLACK обработанная вершина
	BLACK
)

//На начальном этапе все вершины белые. Когда вершина посещается, она становится серой,
//и для нее выполняется проверка связности с другими вершинами.
//Если вершина уже была посещена и она все еще серая, то нашелся цикл, и функция возвращает false.
//Когда все соседние вершины были проверены, текущая вершина становится черной.
//Если при проверке всех вершин не было обнаружено циклов, функция возвращает true.
// id = https://contest.yandex.ru/contest/25070/run-report/85345362/

// временная сложность O(n^2),
//	где n - количество городов
//	Внешний цикл перебирает все города, а внутренний цикл для каждого города перебирает все его соседние города.
//	Таким образом, общее количество итераций в циклах равно n^2.

const (
	R   = 'R'
	YES = "YES"
	NO  = "NO"
)

// Matrix матрица смежности
type Matrix [][]int

// Color статусы вершин
type Color int8

// Visited отметки о посещении вершины
type Visited []Color

func (v Visited) IsWhite(i int) bool {
	return v[i] == WHITE
}

func (v Visited) IsGray(i int) bool {
	return v[i] == GRAY
}

// Visit ставит метку, что вершина была посещена
func (v Visited) Visit(i int) {
	v[i] = GRAY
}

// Done ставит метку, что вершина была обработана
func (v Visited) Done(i int) {
	v[i] = BLACK
}

// stack стек для обхода
type stack []int

// push добавляет элемент в стек
func (s stack) push(i int) stack {
	return append(s, i)
}

// shift удаляет элемент из стека
func (s stack) shift() stack {
	return s[:len(s)-1]
}

// pep возвращает последний элемент стека
func (s stack) pep() int {
	return s[len(s)-1]
}

// checkOptimal обход матрицы в поисках путей
func (m Matrix) checkOptimal(visited Visited, n int) bool {
	for i := 1; i <= n; i++ {
		if !visited.IsWhite(i) {
			continue
		}

		stack := stack{i}
		var current int

		for len(stack) > 0 {
			current = stack.pep()

			if visited.IsWhite(current) {
				visited.Visit(current)

				for _, target := range m[current] {
					// если мы еще не обработали вершину, то добавляем ее в стек
					if visited.IsWhite(target) {
						// добавляем в стек
						stack = stack.push(target)
					} else if visited.IsGray(target) {
						return false
					}
				}

			} else {
				visited.Done(current)
				stack = stack.shift()
			}
		}
	}

	return true
}

func main() {
	var builder strings.Builder

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// количество городов
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		log.Fatal(err)
	}

	// матрица смежности
	matrix := make(Matrix, n+1)
	// метки посещенных городов
	visited := make([]Color, n+1)

	// заполняем матрицу
	for i := 1; i <= n; i++ {
		scanner.Scan()
		railData := scanner.Text()

		for j := 0; j < len(railData); j++ {
			target := i + j + 1
			// тип дороги R
			if railData[j] == R {
				matrix[i] = append(matrix[i], target)
			} else { // тип дороги B
				matrix[target] = append(matrix[target], i)
			}
		}
	}

	if matrix.checkOptimal(visited, n) {
		builder.WriteString(YES)
	} else {
		builder.WriteString(NO)
	}

	fmt.Println(builder.String())
}
