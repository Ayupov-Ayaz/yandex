package main

import (
	"bufio"
	"fmt"
	"os"
)

// -- Принцип работы функции по нахождению растояния Левинштейна --
// 1. Создаем матрицу размером (len(s)+1)x(len(t)+1) и заполняем ее начальными значениями.
// 2. Далее проходим по всем элементам матрицы начиная с (1,1) и заполняем их значения.
// 3. Если символы s[i] и t[j] равны, то значение ячейки matrix[i][j] равно значению ячейки matrix[i-1][j-1].
// 4. Если символы s[i] и t[j] не равны, то значение ячейки matrix[i][j]
// равно минимальному из значений ячеек matrix[i-1][j-1], matrix[i-1][j] и matrix[i][j-1] плюс 1.
// 5. Возвращаем значение ячейки matrix[len(s)][len(t)].
//
// Сложность данного алгоритма - O(len(s) * len(t)), так как мы проходим по всем буквам s и t.
// Память, используемая данным алгоритмом, равна len(t)),
// так как используется только один одномерный массив длиной (len(s) + 1),
// который перезаписывается на каждой итерации внешнего цикла.

// https://contest.yandex.ru/contest/25597/run-report/85613584/
func getMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func makeMatrix(lenT int) []int {
	// Создаем одномерный массив размером (lenT + 1) и заполняем его начальными значениями.
	matrix := make([]int, lenT+1)

	// Первый столбец матрицы будет заполнен значениями от 0 до lenT
	for i := 0; i <= lenT; i++ {
		matrix[i] = i
	}

	return matrix
}

func LevenshteinDistance(s, t string) int {
	// Если s и t равны, то расстояние Левенштейна равно 0.
	if s == t {
		return 0
	}

	lenS := len(s)
	lenT := len(t)
	matrix := makeMatrix(lenT)

	// Далее проходим по всем элементам матрицы начиная с (1,1) и заполняем их значения.
	for i := 1; i <= lenS; i++ {
		prevI := i - 1
		prevRowStart := matrix[0] // начальное значение для текущей строки
		matrix[0] = i             // первый элемент текущей строки будет заполнен значением i
		for j := 1; j <= lenT; j++ {
			prevJ := j - 1
			if s[prevI] == t[prevJ] {
				matrix[j], prevRowStart = prevRowStart, matrix[j]
			} else {
				matrix[j], prevRowStart = getMin(matrix[prevJ]+1, matrix[j]+1, prevRowStart+1), matrix[j]
			}
		}
	}

	// ответ будет находиться в последнем элементе массива
	return matrix[lenT]
}

func getString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	s := getString(sc)
	t := getString(sc)
	fmt.Println(LevenshteinDistance(s, t))
}
