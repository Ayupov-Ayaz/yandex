package main

func partition(arr []int, pivot int) (left, center, right []int) {
	for i := 0; i < len(arr); i++ {
		v := arr[i]

		if v < pivot {
			left = append(left, v)
		} else if v == pivot {
			center = append(center, v)
		} else {
			right = append(right, v)
		}
	}

	return
}

func QuickSort(arr []int) []int {
	if len(arr) == 2 {
		return arr
	}

	mid := len(arr) / 2
	pivot := arr[mid]
	left, center, right := partition(arr, pivot)
	resp := make([]int, 0, len(left)+len(center)+len(right))
	resp = append(resp, left...)
	resp = append(resp, center...)
	resp = append(resp, right...)

	return resp
}

func getPivotIndex(arr []int, left, right int) int {
	pivotIndex := left - 1
	pivotElement := arr[right]

	for i := left; i < right; i++ {
		if arr[i] < pivotElement {
			pivotIndex++
			arr[i], arr[pivotIndex] = arr[pivotIndex], arr[i]
		}
	}
	arr[pivotIndex+1], arr[right] = arr[right], arr[pivotIndex+1]
	return pivotIndex + 1
}

func InPlaceQuickSort(arr []int, left, right int) {
	if left < right {
		pivotIndex := getPivotIndex(arr, left, right)
		InPlaceQuickSort(arr, left, pivotIndex-1)
		InPlaceQuickSort(arr, pivotIndex+1, right)
	}
}
