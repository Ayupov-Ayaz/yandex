package main

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		exp  []int
	}{
		{
			name: "100 23 41 123 53 52 2131 53",
			exp:  []int{23, 41, 52, 53, 53, 100, 123, 2131},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := strings.Split(tt.name, " ")
			input := make([]int, len(str))
			for i := 0; i < len(input); i++ {
				input[i], _ = strconv.Atoi(str[i])
			}

			got := QuickSort(input)
			require.Equal(t, tt.exp, got)
		})
	}
}

func TestInPlaceQuickSort(t *testing.T) {
	tests := []struct {
		name string
		exp  []int
	}{
		{
			name: "100 23 41 123 53 52 2131 53",
			exp:  []int{23, 41, 52, 53, 53, 100, 123, 2131},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := strings.Split(tt.name, " ")
			input := make([]int, len(str))
			for i := 0; i < len(input); i++ {
				input[i], _ = strconv.Atoi(str[i])
			}

			InPlaceQuickSort(input)
			require.Equal(t, tt.exp, input)
		})
	}
}
