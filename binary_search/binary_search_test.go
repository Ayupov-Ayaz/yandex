package binary_search

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 4000, 5000}

	tests := []struct {
		need int
		exp  int
	}{
		{
			need: 6,
			exp:  5,
		},
		{
			need: 4000,
			exp:  12,
		},
		{
			need: 3232,
			exp:  -1,
		},
	}

	for _, tt := range tests {
		got := BinarySearch(arr, tt.need, 0, len(arr))
		require.Equal(t, tt.exp, got)
	}
}
