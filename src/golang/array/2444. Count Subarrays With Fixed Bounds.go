package array

//Input: nums = [1,3,5,2,7,5], minK = 1, maxK = 5
//Output: 2
//Explanation: The fixed-bound subarrays are [1,3,5] and [1,3,5,2].
func CountSubarrays2444(nums []int, minK int, maxK int) int64 {
	var l int = len(nums)
	var res int64 = 0
	var max_val_idx int = -1
	var min_val_idx int = -1
	var bad_idx int = -1
	for i := 0;i < l;i++{
		if nums[i] == minK{
			min_val_idx = i
		}
		if nums[i] == maxK{
			max_val_idx = i
		}
		if nums[i] < minK || nums[i] > maxK{
			bad_idx = i
		}
		if min_val_idx != -1 && max_val_idx != -1{
			res += max_int64(0,int64(min_int(max_val_idx,min_val_idx) - bad_idx))
		}
	}
	return res
}