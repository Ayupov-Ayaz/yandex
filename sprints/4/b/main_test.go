package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const size = 7

func TestHashTable_Add(t *testing.T) {
	ht := NewHashTable(size, SimpleHashFunc)

	for i := 0; i < size+1; i++ {
		ht.Add(i, i+1)

		var curr *node

		switch i {
		case size:
			curr = ht.table[0].next
		default:
			curr = ht.table[i]
		}

		require.NotNil(t, curr)
		require.Equal(t, i, curr.key)
		require.Equal(t, i+1, curr.val)
	}
}

func TestHashTable_Get(t *testing.T) {
	const (
		key    = 1235
		expVal = 4484848
	)

	ht := NewHashTable(size, SimpleHashFunc)
	val, found := ht.Get(key)
	require.False(t, found)
	require.Zero(t, val)

	ht.Add(key, expVal)
	val, found = ht.Get(key)
	require.True(t, found)
	require.Equal(t, expVal, val)
}

func TestHashTable_Remove(t *testing.T) {
	const (
		size  = 3
		steps = 3
		count = size * steps
	)

	checkBefore := func(t *testing.T, ht []*node) {
		curr := ht[1]
		require.NotNil(t, curr)
		k := 1

		for curr != nil {
			require.Equal(t, k, curr.key)
			curr = curr.next
			k += steps
		}

		require.Equal(t, 10, k)
	}

	tests := []struct {
		name       string
		key        int
		checkAfter func(t *testing.T, ht []*node)
	}{
		{
			name: "remove from head",
			key:  1,

			checkAfter: func(t *testing.T, ht []*node) {
				curr := ht[1]
				require.NotNil(t, curr)
				k := 4
				for curr != nil {
					require.Equal(t, k, curr.key)
					curr = curr.next
					k += steps
				}

				require.Equal(t, 10, k)
			},
		},
		{
			name: "remove from middle",
			key:  4,
			checkAfter: func(t *testing.T, ht []*node) {
				curr := ht[1]
				require.NotNil(t, curr)
				require.Equal(t, 1, curr.key)

				curr = curr.next
				require.NotNil(t, curr)
				require.Equal(t, 7, curr.key)
				require.Nil(t, curr.next)
			},
		},
		{
			name: "remove from tail",
			key:  7,
			checkAfter: func(t *testing.T, ht []*node) {
				curr := ht[1]
				require.NotNil(t, curr)
				require.Equal(t, 1, curr.key)

				curr = curr.next
				require.NotNil(t, curr)
				require.Equal(t, 4, curr.key)
				require.Nil(t, curr.next)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := NewHashTable(size, SimpleHashFunc)
			for i := 0; i < count; i++ {
				ht.Add(i, i+1)
			}

			k := 1
			curr := ht.table[k]
			for curr != nil {
				require.NotNil(t, curr)
				require.Equal(t, k, curr.key)
				curr = curr.next
				k += steps
			}
			require.Equal(t, 10, k)
			// remove

			checkBefore(t, ht.table)
			_, ok := ht.Remove(tt.key)
			require.True(t, ok)
			tt.checkAfter(t, ht.table)
		})
	}
}

func TestHashTable(t *testing.T) {
	tests := []struct {
		name string
		resp []string
	}{
		{
			name: "get 1\nput 1 10\nput 2 4\nget 1\nget 2\ndelete 2\nget 2\nput 1 5\nget 1\ndelete 2",
			resp: []string{"None", "10", "4", "4", "None", "5", "None"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := strings.Split(tt.name, "\n")
			ht := NewHashTable(len(cmd)/2, SimpleHashFunc)

			i := 0
			callback := func() string {
				defer func() {
					i++
				}()

				if i >= len(cmd) {
					return ""
				}

				return cmd[i]
			}

			iterator(ht, callback)
		})
	}
}
