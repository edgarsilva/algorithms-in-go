package main

import (
	"slices"
)

type Freq struct {
	key   int
	count int
}

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]Freq, len(nums)+1)

	for _, num := range nums {
		if f, ok := m[num]; !ok {
			m[num] = Freq{key: num, count: 1}
			continue
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
		}

		if a.count < b.count {
			return +1
		}

		return 0
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, sorted[i].key)
	}

	return res
}

// TopKFreqBucket using bucket sorted has bigger space but should have O time
// perf, basically O(n) since we are not acteally sorting the slice
func TopKFreqBucket(nums []int, k int) []int {
	counter := make(map[int]int, len(nums)+1)

	for _, num := range nums {
		f, exists := counter[num]
		if !exists {
			counter[num] = 1
			continue
		}

		f++
		counter[num] = f
	}

	bucket := make([][]int, len(nums)+1)
	for k, c := range counter {
		bucket[c] = append(bucket[c], k)
	}

	results := make([]int, 0, k)
	for i := len(bucket) - 1; i > 0; i-- {
		for _, n := range bucket[i] {
			results = append(results, n)
			if len(results) == k {
				return results
			}
		}
	}

	return results
}
