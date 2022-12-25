package binary_search

// BinarySearch - рекурсивный бинарный поиск
// в худшем случае скорость Log(n)
// в лучшем случае скорость O(1) - случай когда объект по середине
// left - левая граница, которая входит в диапазон
// right - правая граница, которая не входит в диапазон
// если объект найден, то возвращается индекс объекта
// если объект не найден, то вернется чисто -1
// Обязательное условие для бинарного поиска:
// * массив должен быть отсортирован
func BinarySearch(arr []int, need int, left, right int) int {
	// базовый случай
	// это означает, что нужного элемента в массиве нету
	if left == right {
		return -1
	}

	// смотрим середина массива
	// если это нужный нам объект, тогда возвращем его
	mid := (left + right) / 2
	if arr[mid] == need {
		return mid
	}

	// если объект по середине будет больше запрошенного
	// тогда берем левую сторону
	// иначе берем правую сторону
	if arr[mid] > need {
		return BinarySearch(arr, need, left, mid)
	} else {
		return BinarySearch(arr, need, mid+1, right)
	}
}