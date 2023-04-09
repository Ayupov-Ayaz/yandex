package main

import "testing"

func TestCheckSum(t *testing.T) {
	tests := []struct {
		n   int
		str string
		exp bool
	}{
		{
			n:   4,
			str: "1 5 7 1",
			exp: true,
		},
		{
			n:   3,
			str: "2 10 9",
		},
		{
			n:   6,
			str: "7 9 3 4 6 7",
			exp: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			data, sum := getPoints(tt.str)
			got := CheckSum(tt.n, sum, data)
			if got != tt.exp {
				t.Errorf("got %v, want %v", got, tt.exp)
			}
		})
	}
}
