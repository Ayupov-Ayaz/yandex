package main

import "fmt"

func BubbleSort(numbers []int) {
	ok := true
AGAIN:
	changed := false
	for i := 0; i < len(numbers)-1; i++ {
		j := i + 1
		if numbers[i] > numbers[j] {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			changed = true
			ok = false
		}
	}

	if changed {
		fmt.Printf("%v\n", numbers)
		goto AGAIN
	}

	if ok {
		fmt.Printf("%v\n", numbers)
	}
}
