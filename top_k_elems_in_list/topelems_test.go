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

func TestTopKFreqBucket(t *testing.T) {
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
			nums: []int{7, 7, 7, 2, 3, 4, 4, 5, 6, 6, 6, 6},
			k:    3,
			want: []int{6, 7, 4},
		},
		{
			nums: []int{7, 7},
			k:    1,
			want: []int{7},
		},
	}

	for _, tc := range testcases {
		got := TopKFreqBucket(tc.nums, tc.k)
		want := tc.want

		if slices.Compare(got, want) != 0 {
			t.Errorf("got %#v want %#v", got, want)
		}
	}
}

func BenchmarkFrequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []int{7, 7, 7, 2, 3, 4, 4, 5, 6, 6, 6, 6}
		k := 3
		TopKFrequent(input, k)
	}
}

func BenchmarkTopKFreqBucket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []int{7, 7, 7, 2, 3, 4, 4, 5, 6, 6, 6, 6}
		k := 3
		TopKFrequent(input, k)
	}
}
