package selectionsort

import (
	"slices"
	"testing"
)

func TestAssert(t *testing.T) {
	passes := Assert()

	if !passes {
		t.Errorf("Assert punction failed")
	}
}

func TestSelectionsort(t *testing.T) {
	list := List{
		items: []int{2, 4, 5, 3, 1},
	}
	Selectionsort(&list)
	got := list.items
	want := []int{1, 2, 3, 4, 5}

	if slices.Compare(got, want) != 0 {
		t.Errorf("got %#v want %#v", got, want)
	}
}
