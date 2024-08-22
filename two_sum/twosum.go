package twosum

func TwoSum(target int, nums []int) [2]int {
	m := make(map[int]int)
	for i, v := range nums {
		if valIdx, exists := m[target-v]; exists {
			return [2]int{valIdx, i}
		}

		m[v] = i
	}

	return [2]int{}
}
