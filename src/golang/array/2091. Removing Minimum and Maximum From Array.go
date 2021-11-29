package array

func minimumDeletions(nums []int) int {
	var l int = len(nums)
	var min_val int = 2147483647
	var max_val int = -2147483648
	var min_idx int = -1
	var max_idx int = -1
	for i := 0;i < l;i++{
		if nums[i] > max_val{
			max_val = nums[i]
			max_idx = i
		}
		if nums[i] < min_val{
			min_val = nums[i]
			min_idx = i
		}
	}
	res := min_int_number(max_int(min_idx,max_idx) + 1, l - min_int(min_idx,max_idx),
		min_int(min_idx,max_idx) + 1 + l - max_int(min_idx,max_idx))
	return res
}