package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Class struct {
	start float64
	end   float64
}

func Filter(classes []Class) []Class {
	var filtered []Class
	if len(classes) == 0 {
		return filtered
	}

	filtered = append(filtered, classes[0])
	prev := classes[0]
	for i := 1; i < len(classes); i++ {
		if prev.end <= classes[i].start {
			filtered = append(filtered, classes[i])
			prev = classes[i]
		}
	}

	return filtered
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	count, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	classes := make([]Class, count)
	for i := 0; i < count; i++ {
		sc.Scan()
		times := strings.Split(sc.Text(), " ")
		start, err := strconv.ParseFloat(times[0], 64)
		if err != nil {
			panic(err)
		}

		end, err := strconv.ParseFloat(times[1], 64)
		if err != nil {
			panic(err)
		}

		classes[i] = Class{
			start: start,
			end:   end,
		}
	}

	sort.Slice(classes, func(i, j int) bool {
		return classes[i].end < classes[j].end
	})

	classes = Filter(classes)
	cnt := len(classes)
	fmt.Println(cnt)
	for i := 0; i < cnt; i++ {
		fmt.Printf("%v %v\n", classes[i].start, classes[i].end)
	}
}
