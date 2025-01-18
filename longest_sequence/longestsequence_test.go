package longestsequence

import (
	"testing"
)

func TestLongestSequence(t *testing.T) {
	testcases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 20, 4, 10, 3, 4, 5}, 4},
		{[]int{0, 3, 2, 5, 4, 6, 1, 1}, 7},
	}

	for _, testcase := range testcases {
		nums := testcase.nums
		got := LongestSequence(nums)

		if got != testcase.want {
			t.Errorf("got %#v want %#v", got, testcase.want)
		}
	}
}
