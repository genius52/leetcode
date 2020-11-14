package number

import "sort"

func NumSubarrayProductLessThanK(nums []int, k int) int {
	var l int = len(nums)
	if l == 0{
		return 0
	}
	sort.Ints(nums)
	var res int = 0
	begin := 0
	end := 0
	product := nums[begin]
	for (begin < l && end < l) || (nums[begin] > k){
		if product < k{
			//res++
			end++
			if end >= l{
				break
			}
			product *= nums[end]
		}else{
			begin++
			end = begin
			product = nums[begin]
		}
	}
	return res
}