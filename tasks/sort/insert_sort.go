package main

import "fmt"

func InsertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		itemToInsert := arr[i]

		j := i
		for ; j > 0 && itemToInsert < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}

		arr[j] = itemToInsert
		fmt.Printf("step %d sorted %d elements: %v\n", i, i+1, arr)
	}
}

func getArr() []int {
	return []int{10, 9, 7, 5, 4, 2, 1}
}
