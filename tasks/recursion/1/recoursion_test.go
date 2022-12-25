package main

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
)

func Test_searchDaysWhenICanBuyTwoBikes(t *testing.T) {
	tests := []struct {
		name  string
		price int
		exp1  int
		exp2  int
	}{
		{
			name:  "1 2 4 4 4 4",
			price: 1,
			exp1:  1,
			exp2:  2,
		},
		{
			name:  "1 1 4 4 4 4",
			price: 1,
			exp1:  1,
			exp2:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := strings.Split(tt.name, " ")
			days := make([]int, len(str))
			for i, curr := range str {
				n, err := strconv.Atoi(curr)
				require.NoError(t, err)
				days[i] = n
			}
			got1, got2, err := searchDaysWhenICanBuyTwoBikes(days, tt.price)
			require.NoError(t, err)
			require.Equal(t, tt.exp1, got1)
			require.Equal(t, tt.exp2, got2)
		})
	}
}
