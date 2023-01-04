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
// 2. ищем опорный элемент попутно сортируя массив
// 2.1 для изначального опорного элемента берем последний элемент из переданной границы,
//		то есть элемент members[right]
// 2.2 запускаем цикл по всем элементам кроме последнего (опорного);
//		заводим переменную lastIndex - который будет указывать на последний переставленный влево индекс элемента
// 2.3 проверяем что текущий элемент имеет значения лучше, чем опорный
// 2.3.3 если это так, тогда мы инкрементируем lastIndex
//		и меняем местами тот элемент который указывает на lastIndex и текущий (то есть i)
// 2.3.4 пройдя циклом по всем элементам мы переставим все элементы которые лучше опорного влевую сторону и будем знать,
// 		что lastIndex указывает на последний такой элемент, который мы переставляли
//		тут мы можем смело переставить наш опорный элемент сразу после lastIndex (инкрементируем lastIndex)
//		а тот элемент который находится в данный момент по индексу lastIndex пока перествим в конец
// 2.3.5 возвращаем наш lastIndex - это наш опорный элемент на текущей итерации
// 3. запускаем нашу фукнкцию сортировки для левой половинки от опорного элемента
// 4. запускаем нашу функцию сортировки для правой половинки от опорного элемента
// 		в пунктах 3 и 4 у нас будет выполняться рекурсивный вызов функции сортировки, на уменьшенных участках массива
// 		для более глубокой сортировки
// -- Временная сложность --
// O(n2) - наихудший случай,
//		если каждый раз будет выбираться в качестве опорного самый максимальный либо минимальный элемент
// O(nLog n) - во всех остальных случаях

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

// getPivotIndex - возвращает индекс опорного элемента
// сортирует элементы относительно этого опорного элемента,
// распологая значения меньше слева,
// а значения больше справа
func getPivotIndex(members []Member, left, right int) int {
	lastIndex := left - 1
	pivotElement := members[right]

	for i := left; i < right; i++ {
		if pivotElement.IsBetter(members[i]) {
			lastIndex++
			members[i], members[lastIndex] = members[lastIndex], members[i]
		}
	}
	members[lastIndex+1], members[right] = members[right], members[lastIndex+1]
	return lastIndex + 1
}

// InPlaceQuickSort - быстрая сортировка без выделения дополнительной памяти
func InPlaceQuickSort(members []Member, left, right int) {
	if left < right {
		pivotIndex := getPivotIndex(members, left, right)
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
