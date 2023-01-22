package selection_sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{10, 4, 2, 1, 2, 4, 6, 7, 9, 5}
	SelectSort(arr)
	require.Equal(t, []int{1, 2, 2, 4, 4, 5, 6, 7, 9, 10}, arr)
}
