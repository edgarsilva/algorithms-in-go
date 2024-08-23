package arrayduplicates

import "testing"

func TestHasDuplicates(t *testing.T) {
	testcases := []struct {
		input []int
		want  bool
	}{
		{[]int{3}, false},
		{[]int{3, 4, 5, 6}, false},
		{[]int{4, 5, 6, 4}, true},
		{[]int{5, 5}, true},
	}

	for _, tc := range testcases {
		got := HasDuplicates(tc.input)

		if got != tc.want {
			t.Errorf("got %#v want %#v", got, tc.want)
		}
	}
}
