package selection_sort

// SelectSort - сортировка выбором
// находим самый минимальный элемент и ставим его в начало
// сдвигаемся на 1 ячейку и повтоярем алгоритм
// скорость O(1/2 * n^2 ), что является быстрее, чем сортировка пузырьком
// 1/2 не принято писать, по этому везде пишут скорость O(n^2)
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}
