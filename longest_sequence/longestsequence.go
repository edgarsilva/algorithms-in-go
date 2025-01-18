package longestsequence

// Longest Consecutive Sequence
// Given an array of integers nums, return the length of the longest consecutive sequence of elements.
//
// A consecutive sequence is a sequence of elements in which each element is exactly 1 greater than the previous element.
//
// You must write an algorithm that runs in O(n) time.
//
// Example 1:
// Input: nums = [2,20,4,10,3,4,5]
// Output: 4
// Explanation: The longest consecutive sequence is [2, 3, 4, 5].
//
// Example 2:
// Input: nums = [0,3,2,5,4,6,1,1]
// Output: 7
//
// Constraints:
// 0 <= nums.length <= 1000
// -10^9 <= nums[i] <= 10^9

type Empty struct{}

func LongestSequence(nums []int) int {
	numset := map[int]Empty{}

	for _, num := range nums {
		numset[num] = Empty{}
	}

	longest := 0
	for _, num := range nums {
		if _, ok := numset[num-1]; ok {
			continue
		}

		length := 0
		for {
			if _, nextExists := numset[num+length]; !nextExists {
				break
			}
			length++
		}
		longest = max(longest, length)
	}

	return longest
}
