package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Расписал алгоритм поиска 5 релеватных документов из переданных значений
// -- Принцип работы --
// Поиск осуществляется через инвертированный индекс вхождения всех слов в документах
// то есть создается map где ключем является слово, а значением является массив номеров документов.
//
// add - При добавлении нового документам нам необходимо перебрать каждое слово в этом документе
// Проверить вхождение каждого слова в map-у,
// если вхождений нет  тогда создаем массив индексов где добавляем первый элементом id этого документа
// если вхождение есть, мы должны проверить массив на наличие идентификатора нашего документа (O(1)),
// если id документа в массиве не было, тогда добавляем его в массив
// так же запоминаем сколько раз это слово встречалось в этом документе (отдельная map)
//
// search - При поиске мы делаем следующее
// Перебираем каждое слово из документа
// Проверяем их на вхождение в каких либо документах
// Если slice id документов не пустой, тогда мы удостоверившись, что это слово не повторяется (проверяем cache)
// Затем получаем идентификаторы документов и количество этого слова в текущем документе
// Это количетсво мы запоминаем как очки релевантности, по которому в дальнейшем и будем определять топ 5 документов
// Соответственно их необходимо складывать в сумму, чтобы получить какой-то итогововый результат
// Когда мы пробежались по всем нашим документам у нас будет сформирован slice где будут находится документы
// по релевантности
// нам останется их только отсортировать и получить топ 5
// Для сортировки я использовал метод выбора наибольшего элемента из слайса и перемещения его в начало
// Так как мне нужно только 5 элементов, то нет смысла пробегаться по всем элементам
// После сортировки возвращается результат урезанного до длины 5 (если требуется) slice релевантности документов
//
// --Временная сложность --
// n - число документов
//
// Построение индекса (вхождение слов в документы): O(n);
// Сам поиск: O(1)
// Общая сложность O(n)
//
//--Пространственная сложность--
// Алгоритм будет потреблять O(n) памяти
//
// id - https://contest.yandex.ru/contest/24414/run-report/81162270/

type wordInDocumentCount = []map[string]int

type index struct {
	inx            map[string][]int
	counts         wordInDocumentCount
	documentsCount int
}

func newIndex() *index {
	return &index{
		inx: make(map[string][]int),
	}
}

func (idx *index) add(documents []string) {
	idx.documentsCount += len(documents)
	idx.counts = make(wordInDocumentCount, len(documents))

	for id, doc := range documents {
		wordInDocument := make(map[string]int)
		id++
		for _, token := range strings.Split(doc, " ") {
			wordInDocument[token]++

			ids, ok := idx.inx[token]
			if ok && ids[len(ids)-1] == id {
				continue
			} else {
				idx.inx[token] = append(idx.inx[token], id)
			}

		}
		idx.counts[id-1] = wordInDocument
	}
}

func getTop5Relevance(relevance []docRelevance) []docRelevance {
	for i := 0; i < len(relevance); {
		if i == 5 {
			break // дальше не имеет смысла искать, нам нужно только топ 5
		}

		max := relevance[i].relevance
		if max == 0 {
			relevance = append(relevance[:i], relevance[i+1:]...)
			continue
		}

		for j := i + 1; j < len(relevance); j++ {
			curr := relevance[j].relevance
			if curr > max || (curr == max && relevance[j].id < relevance[i].id) {
				relevance[i], relevance[j] = relevance[j], relevance[i]
				max = curr
			}
		}
		i++
	}

	if len(relevance) > 5 {
		relevance = relevance[:5]
	}

	return relevance
}

type docRelevance struct {
	id        int
	relevance int
}

func (idx *index) search(text string) []docRelevance {
	relevance := make([]docRelevance, idx.documentsCount)

	cache := make(map[string]struct{})
	for _, token := range strings.Split(text, " ") {
		ids, ok := idx.inx[token]
		if ok {
			if _, ok := cache[token]; !ok {
				cache[token] = struct{}{}
				for _, id := range ids {
					index := id - 1
					relevance[index].id = id
					rel, _ := idx.counts[index][token]
					relevance[index].relevance += rel
				}
			}
		}
	}

	return getTop5Relevance(relevance)
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	n, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatalln(err)
	}

	documents := make([]string, n)
	for i := 0; i < n; i++ {
		scan.Scan()
		documents[i] = scan.Text()
	}

	inx := newIndex()
	inx.add(documents)

	scan.Scan()
	m, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < m; i++ {
		scan.Scan()
		resp := inx.search(scan.Text())

		var buff strings.Builder
		for j := 0; j < len(resp); j++ {
			curr := resp[j]
			if curr.relevance > 0 {
				buff.WriteString(strconv.Itoa(curr.id))
				buff.WriteString(" ")
			}
		}
		fmt.Println(buff.String())
	}
}
