package main

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestMaxNumber(t *testing.T) {
	tests := []struct {
		arg []int
		exp string
	}{
		{
			exp: "78321",
			arg: []int{1, 783, 2},
		},
		{
			exp: "995928378776666564575515149433633181615151111111110",
			arg: []int{9, 18, 1, 65, 51, 16, 6, 43, 6, 36, 7, 11, 64, 5, 1, 76, 15, 11, 11, 15, 57, 95, 3, 49, 92, 78, 83, 51, 10, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.exp, func(t *testing.T) {
			str := make([]string, len(tt.arg))
			for i, n := range tt.arg {
				str[i] = strconv.Itoa(n)
			}

			sortByLen(str)

			require.Equal(t, tt.exp, build(str))
		})
	}
}
