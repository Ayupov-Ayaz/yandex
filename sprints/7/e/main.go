package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	money, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	values := make([]int, count)
	for i, v := range strings.Split(scanner.Text(), " ") {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		values[i] = n
	}

	sort.Ints(values) // сортируем в порядке возрастания номиналов

	dp := make([]int, money+1)
	for i := 1; i <= money; i++ {
		dp[i] = money + 1 // инициализируем максимальным значением
		for j := 0; j < count && values[j] <= i; j++ {
			if dp[i-values[j]]+1 < dp[i] {
				dp[i] = dp[i-values[j]] + 1
			}
		}
	}

	if dp[money] == money+1 {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[money])
	}
}
