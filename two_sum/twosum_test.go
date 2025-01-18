package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	testcases := []struct {
		nums   []int
		target int
		want   [2]int
	}{
		{[]int{3, 4, 5, 6}, 7, [2]int{0, 1}},
		{[]int{4, 5, 6}, 10, [2]int{0, 2}},
		{[]int{5, 5}, 10, [2]int{0, 1}},
		{[]int{1}, 2, [2]int{0, 0}},
	}

	for _, st := range testcases {
		nums := st.nums
		target := st.target
		got := TwoSum(target, nums)

		if len(got) == 0 {
			t.Errorf("got %#v want %#v", got, st.want)
		}

		if got != st.want {
			t.Errorf("got %#v want %#v", got, st.want)
		}
	}
}
