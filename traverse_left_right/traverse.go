package traverseleftright

import (
	"fmt"
	"slices"
)

func Assert() bool {
	t1 := NewTraverseTree()
	res1 := t1.Traverse(2, []int{4, 1, 1, 9, 0})
	fmt.Println("Traversing T1 ---->", res1)

	t2 := NewTraverseTree()
	res2 := t2.Traverse(4, []int{4, 1, 1, 9, 6, 0})
	fmt.Println("Traversing T2 ---->", res2)

	t3 := NewTraverseTree()
	res3 := t3.Traverse(1, []int{0, 9})
	fmt.Println("Traversing T3 ---->", res3)

	return res1
}

type TraverseTree struct {
	completed []int
}

func NewTraverseTree() TraverseTree {
	return TraverseTree{}
}

func (t *TraverseTree) Traverse(idx int, input []int) bool {
	if idx < 0 || idx >= len(input) {
		return false
	}

	if slices.Contains(t.completed, idx) {
		return false
	}

	valAtIdx := input[idx]
	if valAtIdx == 0 {
		return true
	}

	t.completed = append(t.completed, idx)
	lidx := idx - valAtIdx
	ridx := idx + valAtIdx

	return t.Traverse(lidx, input) || t.Traverse(ridx, input)
}
