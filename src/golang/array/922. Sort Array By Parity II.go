package array

//Input: nums = [4,2,5,7]
//Output: [4,5,2,7]
//Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.
func sortArrayByParityII(nums []int) []int {
	var l int = len(nums)
	var even_idx int = 0
	var odd_idx int = 1
	for even_idx < l && odd_idx < l{
		for even_idx < l && (nums[even_idx] | 1) != nums[even_idx]{
			even_idx += 2
		}
		for odd_idx < l && (nums[odd_idx] | 1) == nums[odd_idx]{
			odd_idx += 2
		}
		if even_idx < l && odd_idx < l{
			nums[even_idx],nums[odd_idx] = nums[odd_idx],nums[even_idx]
		}
		even_idx += 2
		odd_idx += 2
	}
	return nums
}