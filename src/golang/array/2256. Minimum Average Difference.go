package array

func MinimumAverageDifference(nums []int) int {
	var l int = len(nums)
	if l == 1{
		return 0
	}
	var prefix_sum []int = make([]int,l)
	prefix_sum[0] = nums[0]
	for i := 1;i < l;i++{
		prefix_sum[i] += prefix_sum[i - 1] + nums[i]
	}
	var suffix_sum []int = make([]int,l)
	suffix_sum[l - 1] = nums[l - 1]
	for i := l - 2;i >= 0;i--{
		suffix_sum[i] += suffix_sum[i + 1] + nums[i]
	}
	var idx int = -1
	var min_diff int = 2147483647
	for i := 0;i < l - 1;i++{
		cur_diff := abs_int(prefix_sum[i]/(i + 1) - suffix_sum[i + 1]/(l - 1 - i))
		if cur_diff < min_diff{
			min_diff = cur_diff
			idx = i
		}
	}
	if prefix_sum[l - 1]/l < min_diff{
		idx = l - 1
	}
	return idx
}