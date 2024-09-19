package groupanagrams

import (
	"slices"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testcases := []struct {
		strs []string
		want [][]string
	}{
		{
			strs: []string{"x"},
			want: [][]string{
				{"x"},
			},
		},
		{
			strs: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			want: [][]string{
				{"act", "cat"},
				{"pots", "tops", "stop"},
				{"hat"},
			},
		},
	}

	for _, tc := range testcases {
		got := GroupAnagrams(tc.strs)
		want := tc.want

		if !slices.Equal(slices.Concat(got...), slices.Concat(want...)) {
			t.Errorf("got %#v want %#v", got, want)
		}
	}
}
