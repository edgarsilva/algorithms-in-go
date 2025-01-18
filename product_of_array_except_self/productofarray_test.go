package productofarrayexceptself

import (
	"slices"
	"testing"
)

func TestSolution(t *testing.T) {
	testcases := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{1, 2, 4, 6}, []int{48, 24, 12, 8}},
		{[]int{-1, 0, 1, 2, 3}, []int{0, -6, 0, 0, 0}},
	}

	for _, testcase := range testcases {
		nums := testcase.nums
		got := Solution(nums)

		if slices.Compare(got, testcase.want) != 0 {
			t.Errorf("got %#v want %#v", got, testcase.want)
		}
	}
}
