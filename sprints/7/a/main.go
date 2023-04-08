package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Алгоритм решения этой задачи заключается в том,
//что мы проходимся по ценам акций, начиная со второй цены.
//Мы сохраняем цену покупки акции (buyPrice) равной первой цене входных данных,
//а затем проверяем, если текущая цена меньше цены покупки, то мы обновляем цену покупки.
//Если текущая цена больше цены покупки, то мы вычисляем прибыль от продажи
//акции на текущей цене и обновляем максимальную прибыль, если прибыль от продажи больше максимальной.
//В конце мы выводим максимальную прибыль.

// 6
// 7 1 5 3 6 4
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	prices := make([]int, n)
	sc.Scan()
	numbers := strings.Split(sc.Text(), " ")
	for i, n := range numbers {
		m, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		prices[i] = m
	}

	var profit int
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	fmt.Println(profit)
}
