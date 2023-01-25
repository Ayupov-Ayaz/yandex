package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
