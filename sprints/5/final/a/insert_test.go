package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestInsert(t *testing.T) {
	arr := []string{
		"alla 4 100",
		"gena 6 1000",
		"gosha 2 90",
		"rita 2 90",
		"timofey 4 80",
	}

	var root *Node
	for _, v := range arr {
		val := strings.Split(v, " ")
		m := makeMember(val[0], val[1], val[2])

		if root != nil {
			Insert(root, m)
		} else {
			root = NewNode(m)
		}
	}

	got := root.RightToLeftOrder()
	const exp = "gena\ntimofey\nalla\ngosha\nrita"
	require.Equal(t, exp, got)
}
