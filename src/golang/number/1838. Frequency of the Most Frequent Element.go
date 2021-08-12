package number

import "sort"

//Input: nums = [1,2,4], k = 5
//Output: 3
func MaxFrequency(nums []int, k int) int {
	var l int = len(nums)
	sort.Ints(nums)
	var res int = 1
	var left int = 0
	var right int = left
	var sum int = 0
	for left < l{
		for right < l{
			sum += nums[right]
			cur_len := right - left + 1
			if nums[right] * cur_len - sum <= k{
				if cur_len > res{
					res = cur_len
				}
			}else{
				break
			}
			right++
		}
		sum -= nums[left]
		left++
		right++
	}
	return res
}