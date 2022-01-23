package array

func rearrangeArray(nums []int) []int {
	var l int = len(nums)
	var pos_idx int = 0
	var neg_idx int = 0
	var res []int = make([]int,l)
	var idx int = 0
	for pos_idx < l && neg_idx < l{
		for pos_idx < l && nums[pos_idx] < 0{
			pos_idx++
		}
		for neg_idx < l && nums[neg_idx] > 0{
			neg_idx++
		}
		res[idx] = nums[pos_idx]
		idx++
		pos_idx++
		res[idx] = nums[neg_idx]
		idx++
		neg_idx++
	}
	return res
}