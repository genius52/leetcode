package number

//Input: nums = [10, 5, 2, 6], k = 100
//Output: 8
func NumSubarrayProductLessThanK(nums []int, k int) int {
	var l int = len(nums)
	if l == 0{
		return 0
	}
	var res int = 0
	var begin int = 0
	var end int = 0
	product := 1
	for begin < l && end < l{
		product *= nums[end]
		for begin <= end && product >= k{
			product /= nums[begin]
			begin++
		}
		res += end - begin + 1
		end++
	}
	return res
}