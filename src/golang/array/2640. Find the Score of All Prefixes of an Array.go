package array

func FindPrefixScore(nums []int) []int64 {
	var l int = len(nums)
	var res []int64 = make([]int64, l)
	var pre_max int = nums[0]
	var pre_sum int64 = 0
	for i := 0; i < l; i++ {
		if nums[i] > pre_max {
			pre_max = nums[i]
		}
		cur_cover := int64(pre_max + nums[i])
		res[i] = cur_cover + pre_sum
		pre_sum = res[i]
	}
	return res
}
