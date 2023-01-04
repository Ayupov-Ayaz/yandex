package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//https://contest.yandex.ru/contest/23991/problems/D/

type classes []string

func (c classes) isExist(class string) bool {
	for _, curr := range c {
		if curr == class {
			return true
		}
	}

	return false
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	count, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}

	classes := make(classes, 0, count)
	for i := 0; i < count; i++ {
		scan.Scan()
		class := scan.Text()

		if !classes.isExist(class) {
			classes = append(classes, class)
		}
	}

	for _, curr := range classes {
		fmt.Println(curr)
	}
}
