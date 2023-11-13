package array

func minimumSum(nums []int) int {
	var l int = len(nums)
	var res int = 1<<31 - 1
	for i := 1; i < l-1; i++ {
		var prefix_min int = nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] < prefix_min {
				prefix_min = nums[j]
			}
		}
		var suffix_min int = nums[i]
		for k := i + 1; k < l; k++ {
			if nums[k] < suffix_min {
				suffix_min = nums[k]
			}
		}
		if prefix_min < nums[i] && suffix_min < nums[i] {
			cur_sum := prefix_min + nums[i] + suffix_min
			if cur_sum < res {
				res = cur_sum
			}
		}
	}
	if res == 1<<31-1 {
		return -1
	}
	return res
}
