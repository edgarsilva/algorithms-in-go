package main

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	testcases := []struct {
		nums []int
		want []int
		k    int
	}{
		{
			nums: []int{1, 2, 2, 3, 3, 3},
			k:    2,
			want: []int{3, 2},
		},
		{
			nums: []int{7, 7},
			k:    1,
			want: []int{7},
		},
	}

	for _, tc := range testcases {
		got := TopKFrequent(tc.nums, tc.k)
		want := tc.want

		if slices.Compare(got, want) != 0 {
			t.Errorf("got %#v want %#v", got, want)
		}
	}
}
