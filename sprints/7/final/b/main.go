package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// -- Принцип работы --
// Дан массив чисел, необходимо проверить, можно ли разбить его на две части, сумма элементов которых будет равна.
// Для этого необходимо проверить, что сумма всех элементов делится на 2 и
// существует хотя бы 2 элемента с одинаковой суммой.
// Для этого используется динамическое программирование.
// Для каждого элемента массива суммируются все элементы, которые меньше него.
// Если сумма элементов делится на 2, то можно разбить массив на две части с одинаковой суммой.
// Если сумма элементов не делится на 2, то проверяется, что существует хотя бы 2 элемента с одинаковой суммой.
// Для этого используется массив dp, где dp[i] - количество элементов с суммой i.
// Если dp[i] >= 2, то существует хотя бы 2 элемента с одинаковой суммой.
// Если dp[i] == 1, то существует хотя бы 1 элемент с суммой i.
// Если dp[i] == 0, то не существует элементов с суммой i.
// Для каждого элемента массива суммируются все элементы, которые меньше него.
// Для этого используется массив dpPrev, где dpPrev[i] - количество элементов с суммой i.
// https://contest.yandex.ru/contest/25597/run-report/85471471/
// O(n*m/2), где n - количество элементов, а m - сумма всех элементов в массиве.
const (
	SUCCESS = "True"
	FAIL    = "False"
)

// CheckSum проверка суммы
// n - количество элементов
// sum - сумма всех элементов
// data - массив элементов
// возвращает true, если можно разбить массив на две части с одинаковой суммой
func CheckSum(n, sum int, data []int) bool {
	// базовый случай
	if sum%2 != 0 {
		return false
	}

	dpPrev := make([]int, sum/2+1)
	dp := make([]int, sum/2+1)
	for i := 1; i < n+1; i++ {
		for g := 1; g < sum/2+1; g++ {
			dp[g] = dpPrev[g]
			if g == data[i-1] {
				dp[g]++
			}
			if g > data[i-1] && dpPrev[g-data[i-1]] > 0 {
				dp[g]++
			}
		}
		dpPrev = dp
		dp = make([]int, sum/2+1)
	}

	return dpPrev[sum/2] > 1
}

func getN(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return n
}

func getPoints(str string) (data []int, sum int) {
	points := strings.Split(str, " ")
	data = make([]int, len(points))
	for i, p := range points {
		m, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}

		data[i] = m
		sum += m
	}

	return data, sum
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n := getN(sc.Text())
	sc.Scan()
	data, sum := getPoints(sc.Text())

	if CheckSum(n, sum, data) {
		fmt.Println(SUCCESS)
	} else {
		fmt.Println(FAIL)
	}
}
