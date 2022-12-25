package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortByBigger(str []string) bool {
	ok := true

	sort.Slice(str, func(i, j int) bool {
		wI := str[i]
		wJ := str[j]

		if len(wI) != len(wJ) {
			ok = false
			return false
		}

		return wI > wJ
	})

	return ok
}

func sortByLen(str []string) {
	sort.Slice(str, func(i, j int) bool {
		wordI := str[i]
		wordJ := str[j]
		equal := true

		if len(wordI) < len(wordJ) {
			for k := 0; k < len(wordI); k++ {
				if wordI[k] > wordJ[k] {
					return true
				} else if wordI[k] != wordJ[k] {
					equal = false
				}
			}
		} else {
			equal = false
			for k := 0; k < len(wordJ); k++ {
				if wordI[k] > wordJ[k] {
					return true
				}
			}
		}

		return equal
	})
}

func build(str []string) string {
	var resp strings.Builder
	for i := 0; i < len(str); i++ {
		resp.WriteString(str[i])
	}
	return resp.String()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	_, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	//
	sc.Scan()
	str := strings.Split(sc.Text(), " ")
	if !sortByBigger(str) {
		sortByLen(str)
	}

	fmt.Println(build(str))
}
