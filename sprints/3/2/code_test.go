package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMember_IsBetter(t *testing.T) {
	this := NewMember("tommy", 32, 3)

	tests := []struct {
		login   string
		tasks   int
		forfeit int
		exp     bool
	}{
		{
			login:   "atom",
			tasks:   32,
			forfeit: 3,
			exp:     true,
		},
		{
			login:   "zed1",
			tasks:   32,
			forfeit: 2,
			exp:     true,
		},
		{
			login:   "zed2",
			tasks:   33,
			forfeit: 23232,
			exp:     true,
		},
		{
			login:   "zed3",
			tasks:   32,
			forfeit: 3,
			exp:     false,
		},
		{
			login:   "tommy2",
			tasks:   20,
			forfeit: 0,
			exp:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.login, func(t *testing.T) {
			that := NewMember(tt.login, tt.tasks, tt.forfeit)
			got := this.IsBetter(that)
			require.Equal(t, tt.exp, got)
		})
	}
}

func memberLogins(members []Member) []string {
	resp := make([]string, len(members))
	for i := range members {
		resp[i] = members[i].login
	}

	return resp
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name   string
		inputs []string
		exp    []string
	}{
		{
			name:   "1",
			inputs: []string{"alla 4 100", "gena 6 1000", "gosha 2 90", "rita 2 90", "timofey 4 80"},
			exp:    []string{"gena", "timofey", "alla", "gosha", "rita"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			members := make([]Member, len(tt.inputs))
			for i := range tt.inputs {
				m, err := parseMember(tt.inputs[i])
				require.NoError(t, err)
				members[i] = m
			}

			InPlaceQuickSort(members, 0, len(members)-1)
			require.Equal(t, tt.exp, memberLogins(members))
		})
	}
}
