package twosum

func TwoSum(target int, nums []int) [2]int {
	complementMap := make(map[int]int)
	for currentIdx, currentVal := range nums {
		if complementIdxVal, complementExists := complementMap[target-currentVal]; complementExists {
			return [2]int{complementIdxVal, currentIdx}
		}

		complementMap[currentVal] = currentIdx
	}

	return [2]int{}
}
