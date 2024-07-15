package main

import (
	"fmt"
	"time"
)

func main() {
	list := List{
		items: []int{5, 4, 7, 9, 2, 1, 3, 8, 6},
	}
	fmt.Println("Bubblesort Algorithm")
	fmt.Println("Sorting -> ", list)

	Bubblesort(&list)

	fmt.Println("Sorted ->", list.items)
}

type Sortable interface {
	Items() []interface{}
	Compare(i, j int) bool
	Swap(i, j int)
}

func Bubblesort(list Sortable) {
	items := list.Items()
	n := len(items)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if list.Compare(j, j+1) {
				list.Swap(j, j+1)
				fmt.Println("Swap", list.Items()[j], "with", list.Items()[j+1])
				fmt.Printf("%v\033[F", list.Items())
				time.Sleep(500 * time.Millisecond)
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
