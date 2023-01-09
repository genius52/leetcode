package number

func maximumCount(nums []int) int {
	var neg_cnt int = 0
	var pos_cnt int = 0
	var l int = len(nums)
	var i int = 0
	for ;i < l;i++{
		if nums[i] >= 0{
			neg_cnt = i
			break
		}
	}
	for ;i < l;i++{
		if nums[i] > 0{
			pos_cnt = l - i
			break
		}
	}
	if neg_cnt > pos_cnt{
		return neg_cnt
	}
	return pos_cnt
}