package main

import "testing"

func TestCalculateLevenshteinDistance(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		exp int
	}{
		{
			s:   "abacaba",
			t:   "abaabc",
			exp: 2,
		},
		{
			s:   "dxqrpmratn",
			t:   "jdpmykgmaitn",
			exp: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.s+" "+test.t, func(t *testing.T) {
			got := LevenshteinDistance(test.s, test.t)
			if got != test.exp {
				t.Errorf("got %d, want %d", got, test.exp)
			}
		})
	}
}
