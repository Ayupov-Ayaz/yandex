package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Расписал алгоритм сортировки участников соревнования по принципе in place quick sort (без выделения O(n) памяти)
// Для удобства сортировки была определена структура Member которая имеет следующие поля:
// - login
// - tasks
// - forfeit
// структура Member имеет функцию IsBetter которая поможет нам определить является ли переданный в функцию
// участник лучше текущего участника
// -- Принцип работы --
// 1. запускаем рекурсивную функцию, базовым случаем которого является left >= right;
//    начальные параметры left=0; right=len(members)-1
// 2. ищем опорный элемент
// 	 	для его поиска мы перем элменты по индексу left, right и (left+right)/2
//		и находим между ними среднее число, которое при сортировки была бы между двумя числами
// 3. переносим все элементы меньше опорного влевую сторону
// 4. перемещаем сам опорный элемент на новую позицию (сразу после чисел, которые являются меньше опорного)
// 		по итогу мы получаем массив где слева от опорного элемента будут распологаться числа которые меньше,
//		а справа числа которые больше опорного
// 5. запускаем рекурсию на левую половину массива (от left до опорного, не включая его)
// 6. запускаем рекурсию на правую половину массива (от опорного, не включая его до right)
// 7. сортируем эти половинки, получая опорный элемент и снова запускаем сортировку на левую и правую часть от опорного
// делаем так до тех пор пока не достигнем базового случая, то есть когда left >= right
//
// -- Временная сложность --
// O(n) - так как на каждом стеке вызова функции мы все равно пробегаемся по каждому элементу при сравнениии с опорным
// O(Log n) - сортировка стека
// итоговоая сложность алгоритма составляет:
// O(n) * O(log n) = O(nLog n)
// https://contest.yandex.ru/contest/23815/run-report/80819625/

// Member - участник соревнований
type Member struct {
	login string
	// количество выполненных задач
	tasks int
	// штраф
	forfeit int
}

func NewMember(login string, tasks, forfeit int) Member {
	return Member{
		login:   login,
		tasks:   tasks,
		forfeit: forfeit,
	}
}

// IsBetter - сравнение 2х объектов и возвращает true, если переданный объект лучше, чем текущий
// алгоритм сортировки:
// 1. количество выполненных задач
// 2. если количество задач одинаково, тогда побеждает участник у которого меньше штрафа
// 3. если штрафы тоже совпадают, тогда первым будет тот, у которого логин идет раньше в лексикографическом порядке
func (m Member) IsBetter(that Member) bool {
	tasks := m.tasks - that.tasks
	if tasks > 0 {
		return false
	} else if tasks < 0 {
		return true
	}

	forfeit := m.forfeit - that.forfeit
	if forfeit < 0 {
		return false
	} else if forfeit > 0 {
		return true
	}

	return m.login > that.login
}

// isMedian - определяем является переданная в функция value медианой для
// двух других значений
func isMedian(value, a, b Member) bool {
	if value.IsBetter(a) && !value.IsBetter(b) {
		return true
	}

	return value.IsBetter(b) && !value.IsBetter(a)
}

// getPivotIndex - выдает индекс медианы значения из переданного массива
func getPivotIndex(members []Member, left, right int) int {
	leftV := members[left]
	rightV := members[right]
	center := (left + right) / 2
	centerV := members[center]

	if isMedian(leftV, centerV, rightV) {
		return left
	} else if isMedian(rightV, centerV, leftV) {
		return right
	}

	return center
}

// sort - сортирует элементы относительно опорного элемента
// распологая значения меньше слева,
// а значения больше справа
// возвращает индекс опорного элемента
func sort(members []Member, left, right int) int {
	pivotIndex := getPivotIndex(members, left, right)
	pivot := members[pivotIndex]
	i := left

	for j := left; j <= right; j++ {
		if j == pivotIndex { // не смысла сравнивать
			continue
		}

		if pivot.IsBetter(members[j]) {
			if i == pivotIndex {
				// так как мы сейчас меняем индекс нашего опорного элемента,
				// нам необходимо запомнить его новое расположение
				pivotIndex = j
			}

			members[i], members[j] = members[j], members[i]
			i++
		}
	}

	// ставим наш опорный элемент на свое место (по середине)
	if right-left > 1 && i != pivotIndex {
		members[i], members[pivotIndex] = members[pivotIndex], members[i]
	}

	return i
}

// InPlaceQuickSort - быстрая сортировка без выделения дополнительной памяти
func InPlaceQuickSort(members []Member, left, right int) {
	if left < right {
		pivotIndex := sort(members, left, right)
		InPlaceQuickSort(members, left, pivotIndex-1)
		InPlaceQuickSort(members, pivotIndex+1, right)
	}
}

// parseMember - преобразовывает переданную строку в объект Member
func parseMember(str string) (Member, error) {
	arr := strings.Split(str, " ")
	if len(arr) != 3 {
		return Member{}, errors.New("invalid string")
	}

	tasks, err := strconv.Atoi(arr[1])
	if err != nil {
		return Member{}, err
	}

	forfeits, err := strconv.Atoi(arr[2])
	if err != nil {
		return Member{}, err
	}

	return NewMember(arr[0], tasks, forfeits), nil
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	count, _ := strconv.Atoi(scan.Text())
	members := make([]Member, count)
	for i := 0; i < count; i++ {
		scan.Scan()
		m, err := parseMember(scan.Text())
		if err != nil {
			panic(err)
		}

		members[i] = m
	}

	InPlaceQuickSort(members, 0, len(members)-1)

	for i := 0; i < count; i++ {
		fmt.Println(members[i].login)
	}
}
