package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	sumTests := []struct {
		nums   []int
		target int
		want   [2]int
	}{
		{[]int{3, 4, 5, 6}, 7, [2]int{0, 1}},
		{[]int{4, 5, 6}, 10, [2]int{0, 2}},
		{[]int{5, 5}, 10, [2]int{0, 1}},
	}

	for _, st := range sumTests {
		nums := st.nums
		target := st.target
		got := TwoSum(target, nums)

		if got != st.want {
			t.Errorf("got %#v want %#v", got, st.want)
		}
	}
}
