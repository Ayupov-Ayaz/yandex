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

func inPlaceQuickSort(arr []int, left, right int) {
	pivot := arr[len(arr)/2]

	for left < right {
		ls := true
		rs := true

		if arr[left] < pivot {
			left++
			ls = false
		}

		if arr[right] > pivot {
			right--
			rs = false
		}

		if ls && rs {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
}

func InPlaceQuickSort(arr []int) {
	inPlaceQuickSort(arr, 0, len(arr)-1)
}
