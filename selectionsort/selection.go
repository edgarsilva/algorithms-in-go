package selectionsort

import (
	"fmt"
	"slices"
)

func Assert() bool {
	list := List{
		items: []int{5, 4, 7, 9, 2, 1, 3, 8, 6},
	}
	fmt.Println("Selection Sort Algorithm")
	fmt.Println("Sorting -> ", list)

	Selectionsort(&list)

	fmt.Println("Sorted ->", list.items)

	return slices.Compare(list.items, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) == 0
}

type Sortable interface {
	Items() []interface{}
	Compare(i, j int) bool
	Swap(i, j int)
}

func Selectionsort(list Sortable) {
	items := list.Items()
	n := len(items)
	for i := 0; i < n-1; i++ {
		// the diff with Bubblesort is here
		// we compare the last sorted elem with all the following ones
		// always selecting/moving the smallest one to its final position
		for j := i + 1; j < n; j++ {
			if list.Compare(i, j) { // <- we alwayst comarse the last sorted one with all the remaining
				list.Swap(i, j)
				fmt.Println("Swap", list.Items()[i], "with", list.Items()[j])
				fmt.Printf("%v\033[F", list.Items())
				// time.Sleep(1000 * time.Millisecond)
			}
		}
	}
	fmt.Printf("\033[E\033[E\n\n")
}

type List struct {
	items []int
}

func (l *List) Items() []interface{} {
	items := make([]interface{}, len(l.items))
	for i, v := range l.items {
		items[i] = v
	}
	return items
}

func (l *List) Compare(a, b int) bool {
	return l.items[a] > l.items[b]
}

func (l *List) Swap(i, j int) {
	l.items[i], l.items[j] = l.items[j], l.items[i]
}
