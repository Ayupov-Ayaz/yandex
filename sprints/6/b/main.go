package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseEdge(str string) int {
	resp, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Errorf("invalid val number '%s' %w", str, err))
	}

	return resp
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	var n, m int
	var err error

	arr := strings.Split(sc.Text(), " ")
	if len(arr) != 2 {
		panic("invalid params")
	}

	n, err = strconv.Atoi(arr[0])
	if err != nil {
		panic(fmt.Errorf("invalid 'm' param: %w", err))
	}

	m, err = strconv.Atoi(arr[1])
	if err != nil {
		panic(fmt.Errorf("invalid 'n' param: %w", err))
	}

	list := make([][]int, n)
	for i := 0; i < n; i++ {
		list[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		sc.Scan()

		edge := strings.Split(sc.Text(), " ")

		from := parseEdge(edge[0]) - 1
		to := parseEdge(edge[1]) - 1
		list[from][to] = 1
	}

	var b strings.Builder
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			b.WriteString(strconv.Itoa(list[i][j]) + " ")
		}

		if i != len(list)-1 {
			b.WriteByte('\n')
		}
	}

	fmt.Println(b.String())
}
