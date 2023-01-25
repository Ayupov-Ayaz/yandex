package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIndex_add(t *testing.T) {
	index := newIndex()

	documets := []string{
		"a donut on a glass plate only the donuts",
		"donut is a donut",
	}

	index.add(documets)
	require.Equal(t, 9, len(index.inx))
	require.Equal(t, 2, index.documentsCount)
	require.Equal(t, []int{1, 2}, index.inx["a"])
	require.Equal(t, []int{1, 2}, index.inx["donut"])
	require.Equal(t, []int{1}, index.inx["on"])
	require.Equal(t, []int{1}, index.inx["glass"])
	require.Equal(t, []int{1}, index.inx["only"])
	require.Equal(t, []int{1}, index.inx["plate"])
	require.Equal(t, []int{2}, index.inx["is"])
}

func TestIndex_search(t *testing.T) {
	index := newIndex()
	documents := []string{
		"Studies that estimate and rank the most common words in English examine texts written in English",
		"The wildcat is a species complex comprising two small wild cat species, the European wildcat (Felis silvestris) and the African wildcat (F. lybica).",
		"Catopuma is a genus containing two Asian small wild cat species, the Asian golden cat (C. temminckii) and the bay cat.",
		"\nA list of 100 words that occur most frequently in written English is given below,",
	}

	index.add(documents)

	result := index.search("Small wild Cat")
	require.Equal(t, 2, len(result))
	require.Equal(t, 2, result[0].id)
	require.Equal(t, 3, result[1].id)
}

func TestA(t *testing.T) {
	tests := []struct {
		name      string
		documents []string
		search    []string
		exp       [][]int
	}{
		{
			name: "1",
			documents: []string{
				"Studies that estimate and rank the most common words in English examine texts written in English",
				"The wildcat is a species complex comprising two small wild cat species, the European wildcat (Felis silvestris) and the African wildcat (F. lybica).",
				"Catopuma is a genus containing two Asian small wild cat species, the Asian golden cat (C. temminckii) and the bay cat.",
				"\nA list of 100 words that occur most frequently in written English is given below,",
			},
			search: []string{
				"Small wild Cat",
			},
			exp: [][]int{
				{2, 3},
			},
		},
		{
			name: "2",
			documents: []string{
				"i love coffee",
				"coffee with milk and sugar",
				"free tea for everyone",
			},
			search: []string{
				"i like black coffee without milk",
				"everyone loves new year",
				"mary likes black coffee without milk",
			},
			exp: [][]int{
				{1, 2},
				{3},
				{2, 1},
			},
		},
		{
			name: "3",
			documents: []string{
				"tjegerxbyk pdvmj wulmqfrx",
				"pndygsm dvjihmxr tcdtqsmfe",
				"txamzxqzeq dxkxwq aua",
				"hsciljsrdo fipazun kngi",
				"xtkomk aua wulmqfrx ydkbncmzee",
				"pndygsm cqvffye pyrhcxbcef",
				"szyc uffqhayg ccktodig",
				"ntr wpvlifrgjg htywpe",
				"kngi tjegerxbyk zsnfd",
				"tqilkkd gq qc fipazun",
			},
			search: []string{
				"dxkxwq htywpe",
				"aua tjegerxbyk",
				"xtkomk tjegerxbyk",
				"szyc fipazun",
				"xtkomk tjegerxbyk",
			},
			exp: [][]int{
				{3, 8},
				{1, 3, 5, 9},
				{1, 5, 9},
				{4, 7, 10},
				{1, 5, 9},
			},
		},
		{
			name: "4",
			documents: []string{
				"i like dfs and bfs",
				"i like dfs dfs",
				"i like bfs with bfs and bfs",
			},
			search: []string{
				"dfs dfs dfs dfs bfs",
			},
			exp: [][]int{{3, 1, 2}},
		},
		{
			name: "5",
			documents: []string{
				"tjegerxbyk pdvmj wulmqfrx",
				"pndygsm dvjihmxr tcdtqsmfe",
				"txamzxqzeq dxkxwq aua",
				"hsciljsrdo fipazun kngi",
				"xtkomk aua wulmqfrx ydkbncmzee",
				"pndygsm cqvffye pyrhcxbcef",
				"szyc uffqhayg ccktodig",
				"ntr wpvlifrgjg htywpe",
				"kngi tjegerxbyk zsnfd",
				"tqilkkd gq qc fipazun",
			},
			search: []string{
				"dxkxwq htywpe",
				"aua tjegerxbyk",
				"xtkomk tjegerxbyk",
				"szyc fipazun",
				"xtkomk tjegerxbyk",
			},
			exp: [][]int{
				{3, 8},
				{1, 3, 5, 9},
				{1, 5, 9},
				{4, 7, 10},
				{1, 5, 9},
			},
		},
		{

			name: "6",
			documents: []string{
				"buy flat in moscow",
				"rent flat in moscow",
				"sell flat in moscow",
				"want flat in moscow like crazy",
				"clean flat in moscow on weekends",
				"renovate flat in moscow",
			},
			search: []string{
				"flat in moscow for crazy weekends",
			},
			exp: [][]int{
				{4, 5, 1, 2, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inx := newIndex()
			inx.add(tt.documents)

			for i, text := range tt.search {
				resp := inx.search(text)
				_resp := make([]int, len(resp))
				for i := 0; i < len(resp); i++ {
					_resp[i] = resp[i].id
				}

				require.Equal(t, tt.exp[i], _resp, i)
			}
		})
	}
}
