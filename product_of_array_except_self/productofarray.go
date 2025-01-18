package productofarrayexceptself

// Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].
// Each product is guaranteed to fit in a 32-bit integer.
// Follow-up: Could you solve it in O(n) time without using the division operation?
//
// Example 1:
// Input: nums = [1,2,4,6]
// Output: [48,24,12,8]
//
// Example 2:
// Input: nums = [-1,0,1,2,3]
// Output: [0,-6,0,0,0]
//
// Constraints:
//
// 2 <= nums.length <= 1000
// -20 <= nums[i] <= 20
func Solution(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
