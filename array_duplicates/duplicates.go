package arrayduplicates

type empty struct{}

func HasDuplicates(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	found := make(map[int]empty)
	for _, n := range nums {
		if _, exists := found[n]; exists {
			return true
		}

		found[n] = empty{}
	}

	return false
}
