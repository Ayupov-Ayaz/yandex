package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDFS(t *testing.T) {
	arr := make([][]Vertex, 5)
	arr[1] = []Vertex{4, 2}
	arr[2] = []Vertex{3, 1}
	arr[3] = []Vertex{2, 4}
	arr[4] = []Vertex{3, 1}

	got := DFS(arr, 3)
	require.Equal(t, []int{3, 2, 1, 4}, got)
}
