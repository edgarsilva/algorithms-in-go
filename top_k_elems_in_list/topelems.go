package main

import (
	"slices"
)

type Freq struct {
	key   int
	count int
}

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]Freq)

	for _, num := range nums {
		if f, ok := m[num]; !ok {
			m[num] = Freq{key: num, count: 1}
		} else {
			f.count++
			m[num] = f
		}
	}

	sorted := make([]Freq, 0, len(m))
	for _, v := range m {
		sorted = append(sorted, v)
	}

	slices.SortFunc(sorted, func(a Freq, b Freq) int {
		if a.count > b.count {
			return -1
		} else if a.count < b.count {
			return +1
		} else {
			return 0
		}
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, sorted[i].key)
	}

	return res
}
