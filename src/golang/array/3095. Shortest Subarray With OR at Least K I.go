package array

func MinimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}
	var l int = len(nums)
	var res int = -1
	for sub_len := 1; sub_len <= l; sub_len++ {
		for left := 0; left+sub_len <= l; left++ {
			var sum int = 0
			for visit := left; visit < left+sub_len; visit++ {
				sum |= nums[visit]
			}
			if sum >= k {
				return sub_len
			}
		}
	}
	return res
}
